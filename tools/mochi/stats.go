package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type cardFile struct {
	path string
	fm   map[string]string
	body string
}

// eachCardFile visits every card .md under root, skipping `_`-prefixed dirs and
// files. body is comment-stripped and trimmed.
func eachCardFile(root string, fn func(cardFile) error) error {
	return filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if strings.HasPrefix(d.Name(), "_") && path != root {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".md") || strings.HasPrefix(d.Name(), "_") {
			return nil
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		fm, body := splitFrontMatter(string(raw))
		return fn(cardFile{path: path, fm: fm, body: strings.TrimSpace(stripComments(body))})
	})
}

func isReady(c cardFile) bool {
	return c.fm["status"] != "draft" && !strings.Contains(c.body, "TODO front") && c.body != ""
}

func cmdStats(args []string) error {
	cardsDir := "cards"
	if len(args) == 2 && args[0] == "-cards" {
		cardsDir = args[1]
	}

	var total, ready, draft int
	byDeck := map[string]int{}
	byWeek := map[string]int{}
	byGen := map[string]int{}
	foundation := 0

	_ = eachCardFile(cardsDir, func(c cardFile) error {
		total++
		if isReady(c) {
			ready++
		} else {
			draft++
		}
		rel, _ := filepath.Rel(cardsDir, c.path)
		top := strings.SplitN(filepath.ToSlash(rel), "/", 2)[0]
		byDeck[top]++
		tags := parseList(c.fm["tags"])
		for _, t := range tags {
			switch {
			case t == "foundation":
				foundation++
			case strings.HasPrefix(t, "week-"):
				byWeek[t]++
			}
		}
		if g := c.fm["generated"]; g != "" {
			byGen[g]++
		}
		return nil
	})

	fmt.Printf("CARDS  total=%d  ready=%d  draft=%d  foundation=%d\n\n", total, ready, draft, foundation)

	fmt.Println("By deck (top level):")
	for _, k := range sortedKeys(byDeck) {
		fmt.Printf("  %-16s %d\n", k, byDeck[k])
	}

	if len(byWeek) > 0 {
		fmt.Println("\nBy study week:")
		for _, k := range sortedKeys(byWeek) {
			fmt.Printf("  %-16s %d\n", k, byWeek[k])
		}
	}
	if len(byGen) > 0 {
		fmt.Println("\nGenerated from appendices:")
		for _, k := range sortedKeys(byGen) {
			fmt.Printf("  %-16s %d\n", k, byGen[k])
		}
	}

	// pipeline state
	if proc, err := loadProcessed(); err == nil {
		fmt.Printf("\nHighlights processed: %d\n", len(proc))
	}
	if stubs := countGlob(filepath.Join(cardsDir, "_inbox", "*.md")); stubs > 0 {
		fmt.Printf("Inbox stubs awaiting authoring: %d\n", stubs)
	}
	if open := countOpenGaps("qa/gaps.md"); open >= 0 {
		fmt.Printf("Open Q&A gaps: %d\n", open)
	}
	return nil
}

func countGlob(pattern string) int {
	m, _ := filepath.Glob(pattern)
	return len(m)
}

func countOpenGaps(path string) int {
	data, err := os.ReadFile(path)
	if err != nil {
		return -1
	}
	n := 0
	for _, ln := range strings.Split(string(data), "\n") {
		if strings.Contains(ln, "[ ]") && !strings.Contains(strings.ToLower(ln), "example") {
			n++
		}
	}
	return n
}

var _ = sort.Strings
