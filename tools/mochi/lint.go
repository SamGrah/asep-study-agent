package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// cmdLint enforces card quality on ready cards:
//   - status is explicitly 'ready' (no missing/typo'd statuses build silently)
//   - body is free of TODO placeholders
//   - any cited printed page is within the Handbook's range (1..339)
//   - no two ready cards share the same (normalized) front
//
// Non-cloze cards additionally must:
//   - have a front/back separator; front and back both non-empty
//   - carry a SEHB citation (a § ref or a p.<n>) on the BACK
//
// Cloze cards (body contains a {{…}} deletion) are reviewed FRONT-ONLY, so:
//   - they must NOT have a back side (no content after a '---' separator)
//   - their SEHB citation may live anywhere in the (front) body
//   - clozes must be balanced and non-empty
//   - a cloze card may never set `reverse: true` (reversing a cloze is nonsense)
//
// Exits non-zero if any problems are found, so it can gate a build.
func cmdLint(args []string) error {
	cardsDir := "cards"
	if len(args) == 2 && args[0] == "-cards" {
		cardsDir = args[1]
	}
	pageRe := regexp.MustCompile(`p\.\s?(\d{1,3})`)
	citeRe := regexp.MustCompile(`§|p\.\s?\d`)

	var problems []string
	checked := 0
	firstSeen := map[string]string{} // normalized front -> first path
	idSeen := map[string]string{}    // pinned id -> first path
	_ = eachCardFile(cardsDir, func(c cardFile) error {
		if !isReady(c) {
			return nil
		}
		checked++
		add := func(msg string) { problems = append(problems, c.path+": "+msg) }

		if st := c.fm["status"]; st != "ready" {
			add(fmt.Sprintf("status %q (ready cards must say 'status: ready')", st))
		}
		if strings.Contains(c.body, "TODO") {
			add("contains TODO placeholder")
		}

		isCloze := strings.Contains(c.body, "{{")
		front, back, found := strings.Cut(c.body, "\n---")
		if !found {
			front, back = c.body, ""
		}
		front = strings.TrimSpace(front)
		back = strings.TrimSpace(strings.TrimLeft(back, "-\n "))

		if isCloze {
			// Cloze cards are front-only: no back side allowed.
			if c.fm["reverse"] == "true" {
				add("cloze card must not set reverse: true (clozes are front-only)")
			}
			if back != "" {
				add("cloze card has a back side (clozes are front-only; remove the '---' and put the citation on the front)")
			}
			if !citeRe.MatchString(front) {
				add("no SEHB citation on the cloze front (need a § or p.<page>)")
			}
			// cloze well-formedness
			if o, cl := strings.Count(c.body, "{{"), strings.Count(c.body, "}}"); o != cl {
				add(fmt.Sprintf("unbalanced cloze braces ({{=%d, }}=%d)", o, cl))
			}
			if strings.Contains(c.body, "{{}}") {
				add("empty cloze deletion {{}}")
			}
		} else {
			// Front/back card: separator + non-empty sides + citation on back.
			if !found {
				add("missing front/back separator '---'")
			}
			if front == "" {
				add("empty front")
			}
			if back == "" {
				add("empty back")
			}
			if !citeRe.MatchString(back) {
				add("no SEHB citation on the back (need a § or p.<page>)")
			}
		}

		for _, m := range pageRe.FindAllStringSubmatch(c.body, -1) {
			if n, _ := strconv.Atoi(m[1]); n < 1 || n > 339 {
				add(fmt.Sprintf("cited page p.%d out of range (1..339)", n))
			}
		}
		// duplicate fronts (normalized: collapsed whitespace; case-SENSITIVE,
		// since acronym case is meaningful — e.g. Ai vs AI in Appendix B)
		if front != "" {
			key := strings.Join(strings.Fields(front), " ")
			if first, dup := firstSeen[key]; dup {
				add("duplicate front (same as " + first + ")")
			} else {
				firstSeen[key] = c.path
			}
		}
		// duplicate pinned ids (would merge/clobber on Mochi import)
		if id := c.fm["id"]; id != "" {
			if first, dup := idSeen[id]; dup {
				add("duplicate id " + id + " (same as " + first + ")")
			} else {
				idSeen[id] = c.path
			}
		}
		return nil
	})

	if len(problems) == 0 {
		fmt.Printf("lint OK: %d ready card(s), no problems.\n", checked)
		return nil
	}
	for _, p := range problems {
		fmt.Println("FAIL " + p)
	}
	return fmt.Errorf("%d problem(s) across %d ready card(s)", len(problems), checked)
}
