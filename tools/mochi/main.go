// Command mochi manages the ASEP study flashcard pipeline.
//
// Subcommands:
//
//	highlights   Extract PDF highlights (via pdfannots), classify by color,
//	             dedupe against highlights/processed.json, and write draft card
//	             stubs into cards/_inbox/.
//	build        Compile cards/**/*.md into a Mochi-importable bundle
//	             (dist/sehb5-cards.mochi, a zip containing data.edn).
//	processed    Inspect/modify the processed-highlight ledger.
//
// The tool intentionally has zero external dependencies so it builds offline.
// Live Mochi API sync (requires a Mochi Pro account + MOCHI_API_KEY) is not
// implemented; the card model is identical either way, so adding it later
// requires no change to card files.
package main

import (
	"archive/zip"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	var err error
	switch os.Args[1] {
	case "highlights":
		err = cmdHighlights(os.Args[2:])
	case "build":
		err = cmdBuild(os.Args[2:])
	case "sync":
		err = cmdSync(os.Args[2:])
	case "decks":
		err = cmdDecks(os.Args[2:])
	case "glossary":
		err = cmdGlossary(os.Args[2:])
	case "stats":
		err = cmdStats(os.Args[2:])
	case "lint":
		err = cmdLint(os.Args[2:])
	case "processed":
		err = cmdProcessed(os.Args[2:])
	case "help", "-h", "--help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n\n", os.Args[1])
		usage()
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Print(`mochi — ASEP flashcard pipeline

USAGE
  mochi highlights [-pdf PATH] [-annots PATH] [-no-stubs] [-offset N]
  mochi glossary   [-text PATH]   # Appendix B acronyms + C terms -> cards
  mochi lint       [-cards DIR]   # citation/quality gate on ready cards
  mochi stats      [-cards DIR]   # progress + coverage dashboard
  mochi build      [-cards DIR] [-out FILE] [-week N]  # -week N: only weeks 1..N
                                  # build also pins generated card ids into files
                                  # (offline .mochi bundle; backup/export path)
  mochi sync       [-cards DIR] [-week N] [-key K] [-dry-run] [-prune]
                                  # PRIMARY path: push cards to the Mochi API
                                  # (Pro account). Tracks ids in highlights/
                                  # mochi-sync.json so re-syncs update in place.
  mochi decks      list | archive <id>...   # list/archive remote Mochi decks
  mochi processed  [list|count]

Run from the repository root. See .opencode/skills/mochi-cards/SKILL.md.
Run "mochi build" before "mochi sync" so every card has a pinned id.
`)
}

// ---------------------------------------------------------------------------
// repo paths
// ---------------------------------------------------------------------------

const (
	defaultPDF     = "INCOSE_SEHB5_Systems_Engineering_Handbook_v5.pdf"
	defaultAnnots  = ".venv/bin/pdfannots"
	highlightsJSON = "highlights/highlights.json"
	processedJSON  = "highlights/processed.json"
	inboxDir       = "cards/_inbox"
)

// pageOffset: physical = printed + pageOffset (fallback when pdfannots gives no
// page_label). Overridable via `mochi highlights -offset N`.
var pageOffset = 26

// ---------------------------------------------------------------------------
// pdfannots JSON model
// ---------------------------------------------------------------------------

type annot struct {
	Type      string `json:"type"`
	Page      int    `json:"page"`       // physical page (1-based)
	PageLabel string `json:"page_label"` // printed page label if present
	Outline   string `json:"prior_outline"`
	Text      string `json:"text"`     // highlighted text
	Contents  string `json:"contents"` // user comment
	Color     string `json:"color"`    // #rrggbb
}

// printedPage returns the citation page (printed) for an annotation.
func (a annot) printedPage() string {
	if a.PageLabel != "" {
		return a.PageLabel
	}
	return strconv.Itoa(a.Page - pageOffset)
}

// hash is the stable dedupe identity of a highlight.
func (a annot) hash() string {
	h := sha1.Sum([]byte(a.printedPage() + "\x00" + norm(a.Text) + "\x00" + norm(a.Contents)))
	return hex.EncodeToString(h[:])[:12]
}

func norm(s string) string { return strings.Join(strings.Fields(s), " ") }

// ---------------------------------------------------------------------------
// color -> category routing (see SKILL.md color convention)
// ---------------------------------------------------------------------------

type category struct {
	Name string // human label
	Deck string // suggested deck path under cards/
	Card bool   // false => no card, flag for tutor discussion
}

func classify(hexColor string) category {
	switch hueBucket(hexColor) {
	case "yellow":
		return category{"term/definition", "terms", true}
	case "green":
		return category{"process", "processes", true}
	case "pink":
		return category{"exam-trap", "exam-traps", true}
	case "purple":
		return category{"list/enumeration", "lists", true}
	case "blue":
		return category{"explain-deeper", "_discuss", false}
	default:
		return category{"uncategorized", "_inbox", true}
	}
}

// hueBucket maps a #rrggbb color to a coarse hue name.
func hueBucket(hx string) string {
	hx = strings.TrimPrefix(hx, "#")
	if len(hx) != 6 {
		return ""
	}
	r, _ := strconv.ParseInt(hx[0:2], 16, 0)
	g, _ := strconv.ParseInt(hx[2:4], 16, 0)
	b, _ := strconv.ParseInt(hx[4:6], 16, 0)
	rf, gf, bf := float64(r)/255, float64(g)/255, float64(b)/255
	max := maxf(rf, gf, bf)
	min := minf(rf, gf, bf)
	d := max - min
	if d < 0.08 { // near-gray; treat as yellow-ish default highlight
		return "yellow"
	}
	var h float64
	switch max {
	case rf:
		h = 60 * (gf - bf) / d
	case gf:
		h = 60 * (2 + (bf-rf)/d)
	default:
		h = 60 * (4 + (rf-gf)/d)
	}
	if h < 0 {
		h += 360
	}
	switch {
	case h < 40 || h >= 310:
		return "pink" // red/pink
	case h < 70:
		return "yellow"
	case h < 170:
		return "green"
	case h < 260:
		return "blue"
	default:
		return "purple"
	}
}

func maxf(xs ...float64) float64 {
	m := xs[0]
	for _, x := range xs[1:] {
		if x > m {
			m = x
		}
	}
	return m
}
func minf(xs ...float64) float64 {
	m := xs[0]
	for _, x := range xs[1:] {
		if x < m {
			m = x
		}
	}
	return m
}

// ---------------------------------------------------------------------------
// highlights command
// ---------------------------------------------------------------------------

func cmdHighlights(args []string) error {
	pdf := defaultPDF
	annotsBin := defaultAnnots
	noStubs := false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-pdf":
			i++
			pdf = args[i]
		case "-annots":
			i++
			annotsBin = args[i]
		case "-no-stubs":
			noStubs = true
		case "-offset":
			i++
			n, err := strconv.Atoi(args[i])
			if err != nil {
				return fmt.Errorf("-offset needs a number: %w", err)
			}
			pageOffset = n
		default:
			return fmt.Errorf("unknown flag %q", args[i])
		}
	}

	out, err := exec.Command(annotsBin, "-f", "json", pdf).Output()
	if err != nil {
		return fmt.Errorf("running pdfannots (%s): %w", annotsBin, err)
	}
	if err := os.MkdirAll("highlights", 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(highlightsJSON, out, 0o644); err != nil {
		return err
	}

	var annots []annot
	if err := json.Unmarshal(out, &annots); err != nil {
		return fmt.Errorf("parsing pdfannots json: %w", err)
	}

	processed, err := loadProcessed()
	if err != nil {
		return err
	}

	var fresh []annot
	for _, a := range annots {
		if strings.EqualFold(a.Type, "Highlight") || a.Contents != "" {
			if a.Text == "" && a.Contents == "" {
				continue
			}
			if !processed[a.hash()] {
				fresh = append(fresh, a)
			}
		}
	}

	fmt.Printf("Total annotations: %d | already processed: %d | new: %d\n\n",
		len(annots), len(processed), len(fresh))
	if len(fresh) == 0 {
		fmt.Println("No new highlights. Highlight more in Preview, save, and re-run.")
		return nil
	}

	// group report by category
	byCat := map[string][]annot{}
	for _, a := range fresh {
		c := classify(a.Color)
		byCat[c.Name] = append(byCat[c.Name], a)
	}
	cats := make([]string, 0, len(byCat))
	for k := range byCat {
		cats = append(cats, k)
	}
	sort.Strings(cats)
	for _, c := range cats {
		fmt.Printf("== %s (%d) ==\n", c, len(byCat[c]))
		for _, a := range byCat[c] {
			fmt.Printf("  [p.%s] %s\n", a.printedPage(), truncate(norm(a.Text), 90))
			if a.Contents != "" {
				fmt.Printf("        note: %s\n", truncate(norm(a.Contents), 90))
			}
		}
		fmt.Println()
	}

	if noStubs {
		return nil
	}
	n, err := writeStubs(fresh, processed)
	if err != nil {
		return err
	}
	fmt.Printf("Wrote %d card stub(s) to %s/ and recorded them in %s.\n", n, inboxDir, processedJSON)
	fmt.Println("Next: an agent fills each stub's front/back, files it under cards/<deck>/, then `mochi build`.")
	return nil
}

func writeStubs(fresh []annot, processed map[string]bool) (int, error) {
	if err := os.MkdirAll(inboxDir, 0o755); err != nil {
		return 0, err
	}
	discussN := 0
	written := 0
	var discuss []annot
	for _, a := range fresh {
		c := classify(a.Color)
		h := a.hash()
		if !c.Card {
			discuss = append(discuss, a)
			discussN++
			processed[h] = true // tracked so it isn't re-reported
			continue
		}
		path := filepath.Join(inboxDir, h+".md")
		body := stubContent(a, c)
		if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
			return written, err
		}
		processed[h] = true
		written++
	}
	if len(discuss) > 0 {
		if err := appendDiscuss(discuss); err != nil {
			return written, err
		}
		fmt.Printf("(%d blue 'explain-deeper' highlight(s) logged to cards/_discuss.md, no card)\n", discussN)
	}
	if err := saveProcessed(processed); err != nil {
		return written, err
	}
	return written, nil
}

func stubContent(a annot, c category) string {
	var b strings.Builder
	fmt.Fprintf(&b, "---\n")
	fmt.Fprintf(&b, "deck: %s\n", c.Deck)
	fmt.Fprintf(&b, "category: %s\n", c.Name)
	fmt.Fprintf(&b, "source: SEHB v5 p.%s\n", a.printedPage())
	if a.Outline != "" {
		fmt.Fprintf(&b, "section: %q\n", norm(a.Outline))
	}
	fmt.Fprintf(&b, "highlight: %s\n", a.hash())
	fmt.Fprintf(&b, "status: draft\n")
	fmt.Fprintf(&b, "---\n")
	fmt.Fprintf(&b, "<!-- HIGHLIGHTED: %s -->\n", norm(a.Text))
	if a.Contents != "" {
		fmt.Fprintf(&b, "<!-- NOTE: %s -->\n", norm(a.Contents))
	}
	fmt.Fprintf(&b, "\nTODO front (question)\n\n---\n\nTODO back (answer)\n")
	return b.String()
}

func appendDiscuss(as []annot) error {
	f, err := os.OpenFile("cards/_discuss.md", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, a := range as {
		fmt.Fprintf(f, "- [ ] p.%s — %s\n", a.printedPage(), norm(a.Text))
		if a.Contents != "" {
			fmt.Fprintf(f, "      note: %s\n", norm(a.Contents))
		}
	}
	return nil
}

// ---------------------------------------------------------------------------
// processed ledger
// ---------------------------------------------------------------------------

func loadProcessed() (map[string]bool, error) {
	m := map[string]bool{}
	data, err := os.ReadFile(processedJSON)
	if os.IsNotExist(err) {
		return m, nil
	}
	if err != nil {
		return nil, err
	}
	var ids []string
	if err := json.Unmarshal(data, &ids); err != nil {
		return nil, err
	}
	for _, id := range ids {
		m[id] = true
	}
	return m, nil
}

func saveProcessed(m map[string]bool) error {
	ids := make([]string, 0, len(m))
	for id := range m {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	data, _ := json.MarshalIndent(ids, "", "  ")
	if err := os.MkdirAll("highlights", 0o755); err != nil {
		return err
	}
	return os.WriteFile(processedJSON, data, 0o644)
}

func cmdProcessed(args []string) error {
	sub := "count"
	if len(args) > 0 {
		sub = args[0]
	}
	m, err := loadProcessed()
	if err != nil {
		return err
	}
	switch sub {
	case "count":
		fmt.Printf("%d processed highlight(s)\n", len(m))
	case "list":
		ids := make([]string, 0, len(m))
		for id := range m {
			ids = append(ids, id)
		}
		sort.Strings(ids)
		for _, id := range ids {
			fmt.Println(id)
		}
	default:
		return fmt.Errorf("unknown processed subcommand %q", sub)
	}
	return nil
}

// ---------------------------------------------------------------------------
// build command: cards/**/*.md -> dist/sehb5-cards.mochi
// ---------------------------------------------------------------------------

type card struct {
	id      string
	path    string // source file (for error reporting)
	deck    string // deck path, e.g. "technical-processes/verification"
	content string // mochi markdown (front --- back)
	tags    []string
	reverse bool // also emit a swapped front<->back card (term<->definition)
}

func cmdBuild(args []string) error {
	cardsDir := "cards"
	outFile := "dist/sehb5-cards.mochi"
	weekLimit := 0 // 0 = no limit
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-cards":
			i++
			cardsDir = args[i]
		case "-out":
			i++
			outFile = args[i]
		case "-week":
			i++
			n, err := strconv.Atoi(args[i])
			if err != nil {
				return fmt.Errorf("-week needs a number: %w", err)
			}
			weekLimit = n
		default:
			return fmt.Errorf("unknown flag %q", args[i])
		}
	}

	var cards []card
	pinned := 0
	untagged := 0
	err := filepath.WalkDir(cardsDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if strings.HasPrefix(d.Name(), "_") && path != cardsDir {
				return filepath.SkipDir // skip _inbox, _discuss, etc.
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
			fmt.Fprintf(os.Stderr, "skipping draft/incomplete card: %s\n", path)
			return nil
		}
		// Pin the computed id into the file so later edits/moves keep the same
		// Mochi card id (preserving review history across re-imports).
		didPin, err := pinCardID(path, c.id)
		if err != nil {
			return fmt.Errorf("%s: pinning id: %w", path, err)
		}
		if didPin {
			pinned++
		}
		if maxWeek(c.tags) == 0 {
			untagged++
		}
		if weekLimit > 0 {
			if w := maxWeek(c.tags); w > weekLimit {
				return nil // not yet covered by the current study week
			}
		}
		cards = append(cards, c)
		return nil
	})
	if err != nil {
		return err
	}
	if len(cards) == 0 {
		return fmt.Errorf("no complete cards found under %s/", cardsDir)
	}
	if pinned > 0 {
		fmt.Printf("(pinned id: into %d card file(s))\n", pinned)
	}
	if weekLimit > 0 && untagged > 0 {
		fmt.Printf("note: %d card(s) have no week-N tag and are included in every -week build\n", untagged)
	}

	// Expand reverse cards into a second, swapped card.
	reversed := 0
	for _, c := range cards {
		if !c.reverse {
			continue
		}
		if rc, ok := reversedContent(c.content); ok {
			cards = append(cards, card{id: c.id + "r", path: c.path, deck: c.deck, content: rc, tags: c.tags})
			reversed++
		}
	}

	// Card ids must be globally unique or Mochi will merge/clobber on import.
	seen := map[string]string{}
	for _, c := range cards {
		if first, dup := seen[c.id]; dup {
			return fmt.Errorf("duplicate card id %s:\n  %s\n  %s", c.id, first, c.path)
		}
		seen[c.id] = c.path
	}

	edn := buildEDN(cards)
	if reversed > 0 {
		fmt.Printf("(+%d reverse card(s))\n", reversed)
	}
	if err := os.MkdirAll(filepath.Dir(outFile), 0o755); err != nil {
		return err
	}
	if err := writeMochiZip(outFile, edn); err != nil {
		return err
	}
	scope := "all weeks"
	if weekLimit > 0 {
		scope = fmt.Sprintf("weeks 1..%d", weekLimit)
	}
	fmt.Printf("Built %s with %d card(s) [%s]. Import it via Mochi: File > Import.\n", outFile, len(cards), scope)
	return nil
}

func parseCard(path, root string) (card, bool, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return card{}, false, err
	}
	fm, body := splitFrontMatter(string(raw))
	body = strings.TrimSpace(stripComments(body))
	if fm["status"] == "draft" || strings.Contains(body, "TODO front") || body == "" {
		return card{}, false, nil
	}
	if st := fm["status"]; st != "" && st != "ready" {
		fmt.Fprintf(os.Stderr, "warning: %s: unknown status %q (treating as ready; use 'ready' or 'draft')\n", path, st)
	}
	deck := fm["deck"]
	if deck == "" {
		rel, _ := filepath.Rel(root, filepath.Dir(path))
		deck = filepath.ToSlash(rel)
		if deck == "." {
			deck = "general"
		}
	}
	id := fm["id"]
	if id == "" {
		id = "c" + shortHash(deck+"\x00"+body)
	}
	var tags []string
	if t := fm["tags"]; t != "" {
		tags = parseList(t)
	}
	rev := fm["reverse"] == "true"
	return card{id: id, path: path, deck: deck, content: body, tags: tags, reverse: rev}, true, nil
}

// pinCardID writes `id: <id>` into a card's front matter if it doesn't already
// have one, so the Mochi card id stays stable across content edits and deck
// moves. Returns true if the file was modified.
func pinCardID(path, id string) (bool, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}
	s := string(raw)
	fm, _ := splitFrontMatter(s)
	if fm["id"] != "" {
		return false, nil // already pinned (or explicitly set)
	}
	if !strings.HasPrefix(s, "---") {
		// No front matter at all: create one holding just the id.
		return true, os.WriteFile(path, []byte("---\nid: "+id+"\n---\n"+s), 0o644)
	}
	rest := strings.TrimPrefix(s, "---")
	trimmed := strings.TrimLeft(rest, "\r\n")
	end := strings.Index(trimmed, "\n---")
	if end < 0 {
		return false, fmt.Errorf("unterminated front matter")
	}
	head := s[:len(s)-len(trimmed)+end] // up to (not incl.) the "\n---" closer
	tail := trimmed[end:]
	out := head + "\nid: " + id + tail
	return true, os.WriteFile(path, []byte(out), 0o644)
}

// reversedContent swaps the front and back halves of a card around the first
// `---` separator. Returns ("", false) if there is no separator.
func reversedContent(content string) (string, bool) {
	front, back, found := strings.Cut(content, "\n---")
	if !found {
		return "", false
	}
	back = strings.TrimLeft(back, "-\r\n")
	front = strings.TrimSpace(front)
	back = strings.TrimSpace(back)
	if front == "" || back == "" {
		return "", false
	}
	return back + "\n\n---\n\n" + front, true
}

func stripComments(s string) string {
	for {
		i := strings.Index(s, "<!--")
		if i < 0 {
			break
		}
		j := strings.Index(s[i:], "-->")
		if j < 0 {
			break
		}
		s = s[:i] + s[i+j+3:]
	}
	return s
}

// splitFrontMatter parses a minimal `key: value` YAML-ish front matter block.
func splitFrontMatter(s string) (map[string]string, string) {
	fm := map[string]string{}
	if !strings.HasPrefix(s, "---") {
		return fm, s
	}
	rest := strings.TrimPrefix(s, "---")
	rest = strings.TrimLeft(rest, "\r\n")
	end := strings.Index(rest, "\n---")
	if end < 0 {
		return fm, s
	}
	block := rest[:end]
	body := rest[end+4:]
	body = strings.TrimLeft(body, "\r\n")
	for _, line := range strings.Split(block, "\n") {
		line = strings.TrimRight(line, "\r")
		if strings.TrimSpace(line) == "" {
			continue
		}
		k, v, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}
		fm[strings.TrimSpace(k)] = strings.Trim(strings.TrimSpace(v), `"`)
	}
	return fm, body
}

// maxWeek returns the highest N among `week-N` tags, or 0 if none are present.
func maxWeek(tags []string) int {
	max := 0
	for _, t := range tags {
		if n, err := strconv.Atoi(strings.TrimPrefix(t, "week-")); err == nil && strings.HasPrefix(t, "week-") && n > max {
			max = n
		}
	}
	return max
}

func parseList(s string) []string {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	var out []string
	for _, p := range strings.Split(s, ",") {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, strings.Trim(p, `"`))
		}
	}
	return out
}

// buildEDN emits a Mochi import data.edn (format reference: version 2).
// Decks are created hierarchically from slash-separated deck paths.
func buildEDN(cards []card) string {
	// Collect every deck path and its ancestors.
	deckIDs := map[string]string{}
	parentOf := map[string]string{}
	nameOf := map[string]string{}
	var order []string
	ensure := func(p string) {
		if p == "" {
			return
		}
		parts := strings.Split(p, "/")
		for i := range parts {
			sub := strings.Join(parts[:i+1], "/")
			if _, ok := deckIDs[sub]; ok {
				continue
			}
			deckIDs[sub] = "d" + shortHash(sub)
			nameOf[sub] = titleize(parts[i])
			if i > 0 {
				parentOf[sub] = strings.Join(parts[:i], "/")
			}
			order = append(order, sub)
		}
	}
	for _, c := range cards {
		ensure(c.deck)
	}

	cardsByDeck := map[string][]card{}
	for _, c := range cards {
		cardsByDeck[c.deck] = append(cardsByDeck[c.deck], c)
	}

	var b strings.Builder
	b.WriteString("{:version 2\n :decks [")
	for di, dpath := range order {
		if di > 0 {
			b.WriteString("\n         ")
		}
		fmt.Fprintf(&b, "{:id :%s :name %s", deckIDs[dpath], ednStr(nameOf[dpath]))
		if p, ok := parentOf[dpath]; ok {
			fmt.Fprintf(&b, " :parent-id :%s", deckIDs[p])
		}
		dc := cardsByDeck[dpath]
		if len(dc) > 0 {
			b.WriteString("\n          :cards [")
			for ci, c := range dc {
				if ci > 0 {
					b.WriteString("\n                  ")
				}
				fmt.Fprintf(&b, "{:id :%s :content %s", c.id, ednStr(c.content))
				if len(c.tags) > 0 {
					b.WriteString(" :tags [")
					for ti, t := range c.tags {
						if ti > 0 {
							b.WriteByte(' ')
						}
						b.WriteString(ednStr(t))
					}
					b.WriteString("]")
				}
				b.WriteString("}")
			}
			b.WriteString("]")
		}
		b.WriteString("}")
	}
	b.WriteString("]}\n")
	return b.String()
}

func writeMochiZip(outFile, edn string) error {
	f, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer f.Close()
	zw := zip.NewWriter(f)
	w, err := zw.Create("data.edn")
	if err != nil {
		return err
	}
	if _, err := w.Write([]byte(edn)); err != nil {
		return err
	}
	return zw.Close()
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func ednStr(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	return "\"" + s + "\""
}

func shortHash(s string) string {
	h := sha1.Sum([]byte(s))
	return hex.EncodeToString(h[:])[:11]
}

func titleize(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	words := strings.Fields(s)
	for i, w := range words {
		words[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	return strings.Join(words, " ")
}

func truncate(s string, n int) string {
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[:n-1]) + "…"
}
