# AGENTS.md — ASEP Certification Study Workspace

This repo is a personal study workspace for the **INCOSE ASEP** (Associate
Systems Engineering Professional) exam, based on the **INCOSE Systems
Engineering Handbook, 5th Edition (2023)**. Goal: **pass ASEP by mid-to-late
August** to unlock the MBSE course at CSU (program starts Sep 1). It is not a software
project — it's a tutor + flashcard system built around the Handbook PDF.

## Session opener
At the start of a session, **ask which mode the user wants** — unless their
first message already makes it clear, in which case just proceed:

1. **Handbook search & learnings** — answer/explain SE questions, cite the
   Handbook. Skill: `se-tutor`.
2. **Weekly Q&A session** — you quiz the user, Socratic + adversarial. Skill:
   `se-qa`.
3. **Flashcard management** — create/organize/export Mochi cards. Skill:
   `mochi-cards`.

## Golden rule (tutor & Q&A modes)
Everything must trace to **SEHB v5**. Cite `§<section> (p.<printed page>)`.
Never invent definitions, process names, I/O, or page numbers. Use the
handbook's verbatim wording for definitions/acronyms (the exam tests precise
language). If it isn't in the handbook, say so.

## Output style (all modes)
- Citation-first: every claim ends with `§x.y (p.N)`; no cite → say "uncited".
- **Bold** = verbatim handbook wording; outside the bold is paraphrase.
- Any X-vs-Y comparison goes in a table (mirrors exam discrimination).
- Q&A mode: one question per turn, terse, no answer-leaking preamble.
- Slash commands exist for the common session types: `/week-start N`,
  `/tutor <topic>`, `/qa N`, `/cards N` (see `.opencode/commands/`).

## Key files
- `memory/sehb5-index.md` — **start here** to locate content: full section→page
  map + concept→section lookup + exam-weighting notes.
- `memory/sehb5-text/sehb5-full.txt` — grep-able full text (pdftotext `-layout`).
- `memory/summaries/ch{1..6}-summary.md` — exam-targeted **pre-reading summaries**
  (read before the dense chapter). Reading aid, NOT a citation source; see
  `memory/summaries/README.md`. Slash command: `/summary N`.
- `study-plan.md` — the 8-week schedule + progress checkboxes; defines the
  "current week" for Q&A mode.
- `cards/**/*.md` — flashcard source of truth (one card per file).
- `qa/coverage.md`, `qa/gaps.md` — Q&A coverage log and weak-area log.
- `INCOSE_SEHB5_..._v5.pdf` — the Handbook; the user highlights it in Preview.

## Chapter summaries & read-aloud
- **Summaries** live in `memory/summaries/`. Each: exam-relevance banner · concept
  walkthrough (verbatim-**bold** terms, every handbook claim cited `§x.y (p.N)`) ·
  term/acronym checklist · `> **Outside the handbook**` example callouts (external,
  **uncited** — never cite on the exam) · exam-trap pairs. Summaries never override
  the golden rule: only the handbook is citable.
- **Read-aloud of the handbook** (see summaries/README.md): macOS Siri "Speak
  Selection" is the best free path for live read-along. A `tools/tts/` pipeline can
  pre-generate `audio/chN.mp3` (`make-chapter.sh N`), but the local Piper voice was
  judged too robotic, so **no MP3s are kept** — a future session will wire the
  `openai`/`elevenlabs` backend (stubs in `tools/tts/tts.sh`) for a natural voice.
  `audio/` is gitignored.

## Subagents (pinned models — cost-tiered)
Defined in `.opencode/agent/`. Note: after editing an agent file, **restart
OpenCode** for the Task tool to pick it up.
- `summary-author` (Opus 4.8, max reasoning) — authors chapter summaries + finds
  external examples. Highest-judgment task.
- `summary-auditor` (Opus 4.8) — verifies citations + external-example labeling;
  read-only. Run on every summary; fix BLOCKERs before trusting it.
- `text-cleaner` (Haiku 4.5) — cleans raw `-layout` chapter text for TTS (mechanical;
  cheapest model for the high token volume).
- `tts-engineer` (Sonnet 4.6) — maintains the `tools/tts/` Piper pipeline.

## Page-number gotcha
Citations use **printed** page numbers (what Preview/TOC show). In the grep
corpus, `physical_page = printed_page + 26` (printed p.1 = physical p.27); pages
are split on form-feed `\f`. The PDF has page labels, so `pdfannots` reports
printed pages directly.

## Commands
Build the flashcard tool once (zero deps). On this machine use the **external
linker** or the binary fails with `dyld: missing LC_UUID`:
`cd tools/mochi && go build -ldflags=-linkmode=external -o mochi .`
Then from the **repo root**:
- `tools/mochi/mochi highlights` — PDF highlights → classified, deduped draft
  card stubs in `cards/_inbox/` (blue → `cards/_discuss.md`).
- `tools/mochi/mochi glossary` — Appendix B acronyms + C terms → ~300 reference
  cards in `cards/reference/`, each tagged `week-N` by where the term first
  appears in the body (so future-week terms can be excluded). Term cards get
  `reverse: true` (definition→term drilled too); regeneration preserves ids.
- `tools/mochi/mochi lint` — citation/quality gate on `status: ready` cards
  (run before build/sync). Cloze cards are front-only (no back side, citation on
  the front, never `reverse: true`); front/back cards need a back-side citation.
- `tools/mochi/mochi stats` — progress dashboard (cards/deck, foundation &
  week-N counts, processed highlights, open Q&A gaps).
- `tools/mochi/mochi build [-week N]` — `cards/**/*.md` → `dist/sehb5-cards.mochi`
  (offline backup/export). `-week N` excludes cards tagged for weeks after N.
  Drafts/`TODO front`/`_`-dirs skipped. `reverse: true` emits a swapped card in
  the bundle. **Build pins each card's `id:`** (stable ids; required by sync);
  duplicate ids are a hard error.
- `tools/mochi/mochi sync [-week N] [-dry-run] [-prune]` — **primary path.** Push
  cards to the Mochi API (create decks + create/update cards). Maps each local id
  to its Mochi id in `highlights/mochi-sync.json` so re-syncs update in place.
  `reverse: true` → native `review-reverse?` (no second card). Run `build` first.
- `tools/mochi/mochi decks list|archive <id>...` — list/archive remote decks.
- `tools/mochi/mochi processed [count|list]` — dedupe ledger.

Weekly: `mochi build -week N && mochi lint && mochi sync -week N`.

## Toolchain / environment
- `pdftotext` (poppler, via Homebrew) for text extraction.
- `pdfannots` in a project venv: `./.venv/bin/pdfannots` (the Go tool calls it).
- Go 1.22+ (build with `-ldflags=-linkmode=external`; see Commands). Mochi is on
  **Pro** → live API sync via `mochi sync`. API key in `tools/mochi/.mochi-key`
  (gitignored, mode 600; resolution: `-key` → `MOCHI_API_KEY` → file). The
  per-account id ledger `highlights/mochi-sync.json` is also gitignored.

## Highlight → card workflow (color convention)
User highlights in **Apple Preview**; color picks the card type:
yellow=term, green=process, pink=exam-trap, purple=list, blue=explain-deeper
(no card). Run `mochi highlights`, fill each stub into a real card under
`cards/<deck>/`, set `status: ready`, then `mochi build && mochi sync`. Details:
`.opencode/skills/mochi-cards/SKILL.md`.

## Conventions
- **Weekly foundation cards:** at the start of each study week, proactively
  generate a starter set of flashcards for that week's `study-plan.md` sections
  (tagged `foundation`, `week-N`); highlight-derived cards supplement them. See
  `mochi-cards` skill.
- Keep cards atomic; always cite SEHB (back for front/back cards, front for cloze).
- Don't commit secrets; the Mochi key lives in `tools/mochi/.mochi-key`
  (gitignored) or `MOCHI_API_KEY`.
- `dist/` is a regenerable build artifact.
