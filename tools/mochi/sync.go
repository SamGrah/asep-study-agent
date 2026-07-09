package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

// sync.go: live Mochi REST API integration (requires a Mochi Pro account).
//
// Unlike `build` (which writes an offline .mochi import bundle), `sync` talks
// to https://app.mochi.cards/api directly:
//
//	mochi sync     [-cards DIR] [-week N] [-key K] [-dry-run] [-prune]
//	mochi decks    list                  # list remote decks (id, name, parent)
//	mochi decks    archive <id> [<id>...] # set archived?: true on decks
//
// Mochi assigns its own card/deck ids server-side, so we persist a ledger
// (highlights/mochi-sync.json) mapping our local ids -> remote ids. That lets
// re-syncs UPDATE existing cards (preserving review history) instead of
// creating duplicates.

const (
	mochiAPIBase = "https://app.mochi.cards/api"
	keyFile      = "tools/mochi/.mochi-key"
	syncLedger   = "highlights/mochi-sync.json"
)

// ---------------------------------------------------------------------------
// API key resolution: -key flag > MOCHI_API_KEY env > tools/mochi/.mochi-key
// ---------------------------------------------------------------------------

func resolveKey(flagKey string) (string, error) {
	if flagKey != "" {
		return strings.TrimSpace(flagKey), nil
	}
	if k := strings.TrimSpace(os.Getenv("MOCHI_API_KEY")); k != "" {
		return k, nil
	}
	data, err := os.ReadFile(keyFile)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("no API key: pass -key, set MOCHI_API_KEY, or create %s", keyFile)
		}
		return "", err
	}
	k := strings.TrimSpace(string(data))
	if k == "" {
		return "", fmt.Errorf("%s is empty", keyFile)
	}
	return k, nil
}

// ---------------------------------------------------------------------------
// HTTP client (Basic auth: key as username, empty password; 1 concurrent
// request per account, so all calls are sequential with 429 backoff)
// ---------------------------------------------------------------------------

type mochiClient struct {
	key  string
	http *http.Client
}

func newClient(key string) *mochiClient {
	return &mochiClient{key: key, http: &http.Client{Timeout: 30 * time.Second}}
}

// do issues a request to path (e.g. "/cards/") with an optional JSON body and
// decodes a JSON response into out (may be nil). Retries on HTTP 429.
func (m *mochiClient) do(method, path string, body any, out any) error {
	var payload []byte
	if body != nil {
		var err error
		payload, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}
	const maxAttempts = 5
	for attempt := 1; ; attempt++ {
		var rdr io.Reader
		if payload != nil {
			rdr = bytes.NewReader(payload)
		}
		req, err := http.NewRequest(method, mochiAPIBase+path, rdr)
		if err != nil {
			return err
		}
		req.SetBasicAuth(m.key, "")
		if payload != nil {
			req.Header.Set("Content-Type", "application/json")
		}
		req.Header.Set("Accept", "application/json")

		resp, err := m.http.Do(req)
		if err != nil {
			return err
		}
		respBody, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode == http.StatusTooManyRequests && attempt < maxAttempts {
			wait := time.Duration(attempt) * 2 * time.Second
			fmt.Printf("  rate-limited (429); retrying in %s...\n", wait)
			time.Sleep(wait)
			continue
		}
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return fmt.Errorf("%s %s -> %d: %s", method, path, resp.StatusCode, strings.TrimSpace(string(respBody)))
		}
		if out != nil && len(respBody) > 0 {
			if err := json.Unmarshal(respBody, out); err != nil {
				return fmt.Errorf("decoding %s %s response: %w", method, path, err)
			}
		}
		return nil
	}
}

// ---------------------------------------------------------------------------
// sync ledger (local id -> remote id, per account; gitignored)
// ---------------------------------------------------------------------------

type ledger struct {
	Decks map[string]string `json:"decks"` // deck path  -> Mochi deck id
	Cards map[string]string `json:"cards"` // local card id -> Mochi card id
}

func loadLedger() (*ledger, error) {
	l := &ledger{Decks: map[string]string{}, Cards: map[string]string{}}
	data, err := os.ReadFile(syncLedger)
	if os.IsNotExist(err) {
		return l, nil
	}
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, l); err != nil {
		return nil, err
	}
	if l.Decks == nil {
		l.Decks = map[string]string{}
	}
	if l.Cards == nil {
		l.Cards = map[string]string{}
	}
	return l, nil
}

func saveLedger(l *ledger) error {
	if err := os.MkdirAll(filepath.Dir(syncLedger), 0o755); err != nil {
		return err
	}
	data, _ := json.MarshalIndent(l, "", "  ")
	return os.WriteFile(syncLedger, data, 0o644)
}

// ---------------------------------------------------------------------------
// mochi sync
// ---------------------------------------------------------------------------

func cmdSync(args []string) error {
	cardsDir := "cards"
	weekLimit := 0
	flagKey := ""
	dryRun := false
	prune := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-cards":
			i++
			cardsDir = args[i]
		case "-week":
			i++
			n, err := strconv.Atoi(args[i])
			if err != nil {
				return fmt.Errorf("-week needs a number: %w", err)
			}
			weekLimit = n
		case "-key":
			i++
			flagKey = args[i]
		case "-dry-run":
			dryRun = true
		case "-prune":
			prune = true
		default:
			return fmt.Errorf("unknown flag %q", args[i])
		}
	}

	cards, err := collectCards(cardsDir, weekLimit)
	if err != nil {
		return err
	}
	if len(cards) == 0 {
		return fmt.Errorf("no complete cards found under %s/", cardsDir)
	}

	led, err := loadLedger()
	if err != nil {
		return err
	}

	scope := "all weeks"
	if weekLimit > 0 {
		scope = fmt.Sprintf("weeks 1..%d", weekLimit)
	}
	if dryRun {
		creates, updates := 0, 0
		for _, c := range cards {
			if led.Cards[c.id] == "" {
				creates++
			} else {
				updates++
			}
		}
		fmt.Printf("DRY RUN [%s]: %d card(s) — %d create, %d update; %d deck(s) tracked.\n",
			scope, len(cards), creates, updates, len(led.Decks))
		return nil
	}

	key, err := resolveKey(flagKey)
	if err != nil {
		return err
	}
	cli := newClient(key)

	// 1. Ensure all decks (and ancestors) exist; record their remote ids.
	deckPaths := neededDeckPaths(cards)
	for _, p := range deckPaths {
		if led.Decks[p] != "" {
			continue
		}
		parts := strings.Split(p, "/")
		name := titleize(parts[len(parts)-1])
		req := map[string]any{"name": name}
		if len(parts) > 1 {
			parent := strings.Join(parts[:len(parts)-1], "/")
			if pid := led.Decks[parent]; pid != "" {
				req["parent-id"] = pid
			}
		}
		var created struct {
			ID string `json:"id"`
		}
		if err := cli.do("POST", "/decks/", req, &created); err != nil {
			return fmt.Errorf("creating deck %q: %w", p, err)
		}
		led.Decks[p] = created.ID
		_ = saveLedger(led) // persist incrementally so a crash doesn't orphan decks
		fmt.Printf("  + deck %-40s -> %s\n", p, created.ID)
	}

	// 2. Create or update each card.
	created, updated := 0, 0
	for _, c := range cards {
		deckID := led.Decks[c.deck]
		if deckID == "" {
			return fmt.Errorf("internal: no deck id for %q (card %s)", c.deck, c.path)
		}
		isCloze := strings.Contains(c.content, "{{")
		req := map[string]any{
			"content":         c.content,
			"deck-id":         deckID,
			"review-reverse?": c.reverse && !isCloze,
		}
		if len(c.tags) > 0 {
			req["manual-tags"] = c.tags
		}

		if remoteID := led.Cards[c.id]; remoteID != "" {
			if err := cli.do("POST", "/cards/"+remoteID, req, nil); err != nil {
				return fmt.Errorf("updating card %s (%s): %w", c.id, c.path, err)
			}
			updated++
		} else {
			var res struct {
				ID string `json:"id"`
			}
			if err := cli.do("POST", "/cards/", req, &res); err != nil {
				return fmt.Errorf("creating card %s (%s): %w", c.id, c.path, err)
			}
			led.Cards[c.id] = res.ID
			created++
			_ = saveLedger(led)
		}
	}
	if err := saveLedger(led); err != nil {
		return err
	}

	fmt.Printf("Synced [%s]: %d created, %d updated (%d deck(s)).\n", scope, created, updated, len(led.Decks))

	// 3. Optional prune: trash remote cards whose local id is gone.
	if prune {
		present := map[string]bool{}
		for _, c := range cards {
			present[c.id] = true
		}
		trashed := 0
		now := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
		for localID, remoteID := range led.Cards {
			if present[localID] {
				continue
			}
			if err := cli.do("POST", "/cards/"+remoteID, map[string]any{"trashed?": now}, nil); err != nil {
				return fmt.Errorf("pruning card %s: %w", localID, err)
			}
			delete(led.Cards, localID)
			trashed++
		}
		if err := saveLedger(led); err != nil {
			return err
		}
		fmt.Printf("Pruned %d card(s) no longer present locally.\n", trashed)
	}
	return nil
}

// collectCards walks cardsDir applying the same rules as build (skip _-dirs and
// drafts; honor -week N), returning ready cards. It does NOT pin ids or expand
// reverse cards — reverse is handled natively by the API via review-reverse?.
func collectCards(cardsDir string, weekLimit int) ([]card, error) {
	var cards []card
	err := filepath.WalkDir(cardsDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if strings.HasPrefix(d.Name(), "_") && path != cardsDir {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".md") || strings.HasPrefix(d.Name(), "_") {
			return nil
		}
		c, ok, err := parseCard(path, cardsDir)
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}
		if !ok {
			return nil
		}
		if c.id == "" || !strings.HasPrefix(c.id, "c") || len(c.id) < 5 {
			// Sync relies on a stable pinned id (set by `mochi build`). Refuse
			// to sync unpinned cards rather than create orphans we can't track.
			return fmt.Errorf("%s: card has no pinned id; run `mochi build` first", path)
		}
		if weekLimit > 0 && maxWeek(c.tags) > weekLimit {
			return nil
		}
		cards = append(cards, c)
		return nil
	})
	return cards, err
}

// neededDeckPaths returns every deck path used by cards plus all ancestor
// paths, ordered parents-before-children so parent ids exist when nesting.
func neededDeckPaths(cards []card) []string {
	seen := map[string]bool{}
	for _, c := range cards {
		parts := strings.Split(c.deck, "/")
		for i := range parts {
			seen[strings.Join(parts[:i+1], "/")] = true
		}
	}
	paths := make([]string, 0, len(seen))
	for p := range seen {
		paths = append(paths, p)
	}
	// shorter paths (fewer segments) first => parents before children
	sort.Slice(paths, func(i, j int) bool {
		di, dj := strings.Count(paths[i], "/"), strings.Count(paths[j], "/")
		if di != dj {
			return di < dj
		}
		return paths[i] < paths[j]
	})
	return paths
}

// ---------------------------------------------------------------------------
// mochi decks  (list / archive) — used for first-run migration
// ---------------------------------------------------------------------------

type remoteDeck struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ParentID string `json:"parent-id"`
}

func (m *mochiClient) listDecks() ([]remoteDeck, error) {
	var all []remoteDeck
	path := "/decks/"
	for {
		var page struct {
			Bookmark string       `json:"bookmark"`
			Docs     []remoteDeck `json:"docs"`
		}
		if err := m.do("GET", path, nil, &page); err != nil {
			return nil, err
		}
		all = append(all, page.Docs...)
		if page.Bookmark == "" || len(page.Docs) == 0 {
			break
		}
		path = "/decks?bookmark=" + page.Bookmark
	}
	return all, nil
}

func cmdDecks(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: mochi decks list | mochi decks archive <id>...")
	}
	flagKey := ""
	// allow a trailing -key K
	var rest []string
	for i := 0; i < len(args); i++ {
		if args[i] == "-key" && i+1 < len(args) {
			flagKey = args[i+1]
			i++
			continue
		}
		rest = append(rest, args[i])
	}
	key, err := resolveKey(flagKey)
	if err != nil {
		return err
	}
	cli := newClient(key)

	switch rest[0] {
	case "list":
		decks, err := cli.listDecks()
		if err != nil {
			return err
		}
		byID := map[string]string{}
		for _, d := range decks {
			byID[d.ID] = d.Name
		}
		fmt.Printf("%d deck(s):\n", len(decks))
		for _, d := range decks {
			parent := ""
			if d.ParentID != "" {
				parent = "  (parent: " + byID[d.ParentID] + " / " + d.ParentID + ")"
			}
			fmt.Printf("  %-10s %s%s\n", d.ID, d.Name, parent)
		}
		return nil
	case "archive":
		ids := rest[1:]
		if len(ids) == 0 {
			return fmt.Errorf("archive needs at least one deck id")
		}
		for _, id := range ids {
			if err := cli.do("POST", "/decks/"+id, map[string]any{"archived?": true}, nil); err != nil {
				return fmt.Errorf("archiving deck %s: %w", id, err)
			}
			fmt.Printf("  archived deck %s\n", id)
		}
		return nil
	default:
		return fmt.Errorf("unknown decks subcommand %q (use list | archive)", rest[0])
	}
}
