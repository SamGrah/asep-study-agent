package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// Generators that turn the Handbook's reference appendices into flashcards:
//
//	glossary  Appendix B (Acronyms, p.305) + Appendix C (Terms, p.311)
//
// Both appendices are two-column lists. We extract entries with a gap-split
// heuristic and emit one card file per entry under cards/reference/.
// (Appendix E process I/O cards are deliberately authored by hand — its layout
// doesn't parse reliably; see the mochi-cards skill.)

const defaultText = "memory/sehb5-text/sehb5-full.txt"

var gapRe = regexp.MustCompile(`\s{2,}`)

type entry struct {
	name string
	def  string
}

// splitGap splits s at its first run of 2+ spaces.
func splitGap(s string) (left, right string, found bool) {
	loc := gapRe.FindStringIndex(s)
	if loc == nil {
		return s, "", false
	}
	return strings.TrimSpace(s[:loc[0]]), strings.TrimSpace(s[loc[1]:]), true
}

func isAppendixNoise(line string) bool {
	t := strings.TrimSpace(line)
	if t == "" {
		return true
	}
	for _, sub := range []string{
		"INCOSE Systems Engineering Handbook", "John Wiley",
		"© 2023", "This PDF copy", "publicly accessible", "whether for a fee",
		"systema", "Sources of", "Definitions of the typical", "(Continued)",
	} {
		if strings.Contains(line, sub) {
			return true
		}
	}
	// case-insensitive: appendix titles and running page headers
	lc := strings.ToLower(line)
	for _, sub := range []string{
		"appendix", "terms and definitions", "input/output descriptions", "acronyms",
	} {
		if strings.Contains(lc, sub) {
			return true
		}
	}
	for _, pre := range []string{"Note:", "described in Appendix", "Edited by", "initions can be found"} {
		if strings.HasPrefix(t, pre) {
			return true
		}
	}
	// column headers
	if (strings.HasPrefix(t, "Term") && strings.HasSuffix(t, "Definition")) ||
		t == "Typical Input/Output Description" {
		return true
	}
	// page header/footer: pure number, or "<n> Appendix X" / "Appendix X <n>"
	if regexp.MustCompile(`^\d+$`).MatchString(t) ||
		regexp.MustCompile(`^\d+\s+Appendix`).MatchString(t) ||
		regexp.MustCompile(`Appendix\s+[A-Z]\s+\d+$`).MatchString(t) {
		return true
	}
	return false
}

// findSpan returns the lines of the corpus between the start marker (inclusive
// of its page, after the marker line) and the end marker (exclusive).
func findSpan(corpus, start, end string) []string {
	pages := strings.Split(corpus, "\f")
	var s, e = -1, len(pages)
	for i, p := range pages {
		if s < 0 && strings.Contains(p, start) {
			s = i
		} else if s >= 0 && strings.Contains(p, end) {
			e = i + 1 // include the end page's pre-marker lines too
			break
		}
	}
	if s < 0 {
		return nil
	}
	var lines []string
	for i := s; i < e; i++ {
		lines = append(lines, strings.Split(pages[i], "\n")...)
	}
	return lines
}

// parseAppendix walks lines and accumulates entries (gap-split mode, for the
// two-column Appendix B/C layouts): a new entry begins in column 0; the
// name/def boundary is the first run of 2+ spaces; indented lines are
// definition continuations.
func parseAppendix(lines []string) []entry {
	var out []entry
	var cur *entry
	flush := func() {
		if cur != nil && cur.name != "" && cur.def != "" {
			out = append(out, *cur)
		}
		cur = nil
	}
	for _, line := range lines {
		if isAppendixNoise(line) {
			continue
		}
		if line[0] != ' ' { // new entry
			flush()
			name, def, _ := splitGap(strings.TrimSpace(line))
			cur = &entry{name: name, def: def}
			continue
		}
		if cur == nil {
			continue
		}
		nf, df, found := splitGap(strings.TrimSpace(line))
		if found {
			if nf != "" {
				cur.name += " " + nf
			}
			cur.def += " " + df
		} else {
			cur.def += " " + nf
		}
	}
	flush()
	for i := range out {
		out[i].name = strings.Join(strings.Fields(out[i].name), " ")
		out[i].def = strings.Join(strings.Fields(out[i].def), " ")
	}
	return out
}

var slugRe = regexp.MustCompile(`[^a-z0-9]+`)

func slug(s string) string {
	s = strings.ToLower(s)
	s = slugRe.ReplaceAllString(s, "-")
	return strings.Trim(s, "-")
}

// writtenPaths guards against two different entries slugging to the same file
// within one generation run (previously the later entry silently overwrote the
// earlier one).
var writtenPaths = map[string]string{}

func writeGenerated(deck, name, front, back, source, srcTag string, tags []string, reverse bool) (string, error) {
	dir := filepath.Join("cards", filepath.FromSlash(deck))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	path := filepath.Join(dir, slug(name)+".md")
	if prev, clash := writtenPaths[path]; clash && prev != name {
		path = filepath.Join(dir, slug(name)+"-"+shortHash(name)[:6]+".md")
		fmt.Fprintf(os.Stderr, "warning: slug collision: %q vs %q -> %s\n", prev, name, path)
	}
	writtenPaths[path] = name

	// Preserve a previously pinned card id so regeneration never re-keys a card
	// (which would orphan its Mochi review history on re-import).
	id := ""
	if old, err := os.ReadFile(path); err == nil {
		fm, _ := splitFrontMatter(string(old))
		id = fm["id"]
	}

	var b strings.Builder
	b.WriteString("---\n")
	if id != "" {
		fmt.Fprintf(&b, "id: %s\n", id)
	}
	fmt.Fprintf(&b, "source: %s\ntags: [%s]\nstatus: ready\ngenerated: %s\n",
		source, strings.Join(tags, ", "), srcTag)
	if reverse {
		b.WriteString("reverse: true\n")
	}
	b.WriteString("---\n")
	fmt.Fprintf(&b, "%s\n\n---\n\n%s\n", front, back)
	return path, os.WriteFile(path, []byte(b.String()), 0o644)
}

// weekOfPage maps a printed Handbook page to the study week that covers it,
// per study-plan.md. Pages not covered by a specific week (Ch.5/6, anything
// past p.259) fall to week 7 (the review weeks).
func weekOfPage(p int) int {
	switch {
	case p <= 43: // Ch.1 + §2.1–2.3.1 (foundations)
		return 1
	case p <= 69: // §2.3.2–2.3.3 agreement & org-enabling
		return 6
	case p <= 100: // §2.3.4 technical management
		return 5
	case p <= 117: // §2.3.5.1–3 requirements
		return 2
	case p <= 133: // §2.3.5.4–7 architecture/design
		return 3
	case p <= 158: // §2.3.5.8–14 integration→disposal
		return 4
	case p <= 259: // Ch.3 analyses + Ch.4 tailoring
		return 6
	default:
		return 7
	}
}

var parenRe = regexp.MustCompile(`\s*\([^)]*\)`)

// firstWeek returns the study week in which `name` first appears in the body of
// the Handbook (printed pages 1..304, before the appendices). Returns 7 if it
// never appears in the body. `acronym` selects case-sensitive matching.
func firstWeek(pages []string, name string, acronym bool) int {
	key := name
	if !acronym {
		key = strings.TrimSpace(parenRe.ReplaceAllString(name, ""))
	}
	if key == "" {
		return 7
	}
	pat := `\b` + regexp.QuoteMeta(key) + `\b`
	if !acronym {
		pat = `(?i)` + pat
	}
	re, err := regexp.Compile(pat)
	if err != nil {
		return 7
	}
	for idx := 26; idx < len(pages); idx++ { // idx 26 == printed p.1
		printed := idx - 25
		if printed < 1 || printed > 304 {
			continue
		}
		if re.MatchString(pages[idx]) {
			return weekOfPage(printed)
		}
	}
	return 7
}

// cmdGlossary generates acronym (Appendix B) and term (Appendix C) cards, each
// tagged with the study week in which the term is first introduced so that
// `mochi build -week N` can exclude not-yet-covered material.
func cmdGlossary(args []string) error {
	textPath := defaultText
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-text":
			i++
			textPath = args[i]
		default:
			return fmt.Errorf("unknown flag %q", args[i])
		}
	}
	corpus, err := os.ReadFile(textPath)
	if err != nil {
		return err
	}
	c := string(corpus)
	pages := strings.Split(c, "\f")

	// --- Appendix B: acronyms (merge duplicate keys) ---
	acr := parseAppendix(findSpan(c, "APPENDIX B: ACRONYMS", "APPENDIX C: TERMS"))
	merged := map[string][]string{}
	var order []string
	for _, e := range acr {
		if len(e.name) > 12 || strings.ContainsAny(e.name, " ") {
			continue // acronyms are short single tokens; drop stray prose
		}
		if _, ok := merged[e.name]; !ok {
			order = append(order, e.name)
		}
		merged[e.name] = append(merged[e.name], e.def)
	}
	nB := 0
	for _, k := range order {
		defs := merged[k]
		var back strings.Builder
		if len(defs) == 1 {
			back.WriteString(defs[0])
		} else {
			for _, d := range defs {
				fmt.Fprintf(&back, "- %s\n", d)
			}
		}
		fmt.Fprintf(&back, "\n(§Appendix B, p.305)")
		front := fmt.Sprintf("What does **%s** stand for?", k)
		wk := firstWeek(pages, k, true)
		// Acronyms are NOT reversed: expansion->acronym recall is low-yield for
		// the exam and would double 245 cards' review load.
		if _, err := writeGenerated("reference/acronyms", k, front, back.String(),
			"SEHB v5 Appendix B p.305", "appendix-b",
			[]string{"acronym", "reference", fmt.Sprintf("week-%d", wk)}, false); err != nil {
			return err
		}
		nB++
	}

	// --- Appendix C: terms (reverse: also drill definition -> term) ---
	terms := parseAppendix(findSpan(c, "APPENDIX C: TERMS", "APPENDIX D"))
	nC := 0
	for _, e := range terms {
		if e.name == "" || len(e.name) > 60 || strings.HasPrefix(e.def, "See ") {
			continue
		}
		front := fmt.Sprintf("Define: **%s**", e.name)
		back := e.def + "\n\n(§Appendix C, p.311)"
		wk := firstWeek(pages, e.name, false)
		if _, err := writeGenerated("reference/terms", e.name, front, back,
			"SEHB v5 Appendix C p.311", "appendix-c",
			[]string{"term", "reference", fmt.Sprintf("week-%d", wk)}, true); err != nil {
			return err
		}
		nC++
	}

	fmt.Printf("Generated %d acronym card(s) -> cards/reference/acronyms/\n", nB)
	fmt.Printf("Generated %d term card(s)    -> cards/reference/terms/\n", nC)
	fmt.Println("Run `mochi lint` then `mochi build`.")
	return nil
}

// sortedKeys is a small helper used by stats.
func sortedKeys[T any](m map[string]T) []string {
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	return ks
}
