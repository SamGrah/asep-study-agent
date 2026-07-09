# ASEP Study Agent

An AI-powered study system for the **INCOSE ASEP** (Associate Systems Engineering Professional) certification exam, based on the **INCOSE Systems Engineering Handbook, 5th Edition (2023)**.

**Goal:** Pass the ASEP knowledge exam by mid-to-late August 2026 to unlock the MBSE course at CSU (program starts Sep 1, 2026).

---

## What This Is

This is **not a software project** — it's an intelligent tutor and flashcard system built around the SE Handbook PDF. The system uses OpenCode (an AI coding agent) with custom skills and workflows to:

1. **Tutor** — Answer systems engineering questions with exact citations from the Handbook
2. **Quiz** — Conduct Socratic/adversarial Q&A sessions to test understanding
3. **Generate flashcards** — Convert PDF highlights into Mochi flashcards for spaced repetition

The exam is **120 multiple-choice questions in 120 minutes**, with content drawn 100% from the SE Handbook v5. This workspace treats the Handbook as the single source of truth.

---

## How It Works

### Three Operating Modes

When you start an OpenCode session in this directory, the agent asks which mode you want (or infers from your message):

1. **Handbook search & learnings** (`se-tutor` skill)  
   Ask SE questions, get explanations with citations: `§<section> (p.<page>)`.  
   Every claim traces to the Handbook. No invented definitions.

2. **Weekly Q&A session** (`se-qa` skill)  
   The agent quizzes you on the week's content using Socratic method + adversarial dialectic.  
   You answer; the agent pushes back until you demonstrate mastery.

3. **Flashcard management** (`mochi-cards` skill)  
   Create, organize, and sync Mochi flashcards. Process PDF highlights into cards, build decks, export to Mochi.

### The Study Workflow (Weekly Cadence)

See `study-plan.md` for the full 8-week schedule. Each week (~10–15 hrs):

1. **Foundation cards + cold pretest** (`/week-start N`)  
   Agent generates starter flashcards for the week, then quizzes you **before** reading (pretesting boosts retention).

2. **Pre-read the chapter summary** (`/summary N`)  
   Read `memory/summaries/chN-summary.md` (~20 min) to prime concepts before the dense Handbook text.

3. **Read + highlight**  
   Read the week's Handbook sections. **Highlight in Apple Preview** using color codes:
   - **Yellow** = term/definition
   - **Green** = process/method
   - **Pink** = exam trap (easy to confuse)
   - **Purple** = list/enumeration
   - **Blue** = "explain deeper" (discuss, no card)

4. **Convert highlights → flashcards**  
   Run `tools/mochi/mochi highlights` to extract highlights into card stubs.  
   Fill stubs into real cards, then `mochi build -week N && mochi sync -week N`.  
   **Review daily (15–20 min)** — spaced repetition is critical.

5. **Q&A session** (`/qa N`)  
   One Socratic drill session (~1–2 hrs) with the agent. Gaps logged to `qa/gaps.md` → new cards.

---

## Repository Structure

```
.
├── AGENTS.md                   # Agent instructions (how OpenCode behaves)
├── study-plan.md               # 8-week study schedule + progress tracking
├── opencode.json               # OpenCode configuration
│
├── .opencode/
│   ├── agent/                  # Subagent definitions (summary-author, etc.)
│   ├── commands/               # Slash commands (/tutor, /qa, /summary, /cards)
│   └── skills/                 # Mode-specific instructions (se-tutor, se-qa, mochi-cards)
│
├── memory/
│   ├── sehb5-index.md          # Section→page map, concept lookup (START HERE)
│   ├── sehb5-text/             # Full handbook text (grep-able, pdftotext -layout)
│   └── summaries/              # Exam-targeted pre-reading summaries (ch1–6)
│
├── cards/                      # Flashcard source of truth (one card per .md file)
│   ├── concepts/               # Conceptual cards (emergence, SE value, etc.)
│   ├── life-cycle/             # Life cycle models, stages, gates
│   ├── processes/              # Technical processes
│   └── reference/              # ~300 auto-generated glossary cards (acronyms + terms)
│
├── qa/
│   ├── coverage.md             # Q&A session log
│   └── gaps.md                 # Weak areas to address
│
├── highlights/
│   ├── highlights.json         # Extracted PDF highlights (dedupe ledger)
│   └── mochi-sync.json         # Local↔Mochi id mapping (gitignored)
│
├── tools/mochi/                # Go CLI for flashcard pipeline
│   └── mochi                   # Binary (build once, see below)
│
├── dist/                       # Build artifacts (gitignored, regenerable)
├── audio/                      # TTS output (gitignored, regenerable)
└── INCOSE_SEHB5_..._v5.pdf     # The Handbook (source of truth)
```

---

## Commands

### OpenCode Slash Commands

From within an OpenCode session:

- `/tutor <topic>` — Start tutor mode on a specific topic
- `/qa N` — Quiz session for week N
- `/summary N` — Display chapter N summary
- `/week-start N` — Generate foundation cards + cold pretest for week N
- `/cards N` — Flashcard management for week N

### Mochi Flashcard Tool

First, **build the tool once** (requires Go 1.22+):

```bash
cd tools/mochi
go build -ldflags=-linkmode=external -o mochi .
```

Then from the **repo root**:

#### Daily/Weekly Commands

```bash
# Extract PDF highlights → card stubs
tools/mochi/mochi highlights

# Build flashcard bundle for week N (excludes future weeks)
tools/mochi/mochi build -week N

# Lint cards before syncing (citation/quality checks)
tools/mochi/mochi lint

# Sync to Mochi API (create/update cards)
tools/mochi/mochi sync -week N

# Weekly workflow (run before daily review)
tools/mochi/mochi build -week N && tools/mochi/mochi lint && tools/mochi/mochi sync -week N
```

#### Utility Commands

```bash
# Progress dashboard (cards by deck, week, foundation vs. highlight-derived)
tools/mochi/mochi stats

# Generate ~300 glossary cards from Appendix B & C (run once per week)
tools/mochi/mochi glossary

# List/archive remote Mochi decks
tools/mochi/mochi decks list
tools/mochi/mochi decks archive <deck-id>...

# Dedupe ledger (how many highlights processed)
tools/mochi/mochi processed count
tools/mochi/mochi processed list
```

#### Advanced Options

```bash
# Dry-run sync (preview changes)
tools/mochi/mochi sync -week N -dry-run

# Prune deleted local cards from Mochi
tools/mochi/mochi sync -week N -prune
```

---

## Setup & Dependencies

### Required

- **OpenCode** — The AI agent framework. Install from [opencode.ai](https://opencode.ai)
- **Go 1.22+** — For building `tools/mochi/mochi`
- **Mochi Pro account** — For flashcard sync. API key goes in:
  - `tools/mochi/.mochi-key` (gitignored, mode 600), OR
  - Environment variable `MOCHI_API_KEY`

### Optional

- **pdftotext** (poppler) — For extracting text: `brew install poppler`
- **pdfannots** — For highlight extraction (Python): install in `.venv/bin/pdfannots`
- **Python venv** — If using pdfannots: `python3 -m venv .venv && source .venv/bin/activate && pip install pdfannots`

---

## Key Concepts

### The Golden Rule

Everything traces to **SEHB v5**. Citations use the format `§<section> (p.<printed page>)`.

- **Bold text** = verbatim Handbook wording
- Regular text = paraphrase
- If it's not in the Handbook, the agent says "uncited"
- The exam tests **precise language** — definitions and acronyms must match the Handbook exactly

### Page Numbers

Citations use **printed** page numbers (what you see in Preview/TOC), not PDF physical pages.

- Printed p.1 = physical PDF page 27 (offset of 26)
- The grep corpus (`memory/sehb5-text/sehb5-full.txt`) splits pages on form-feed `\f`
- `pdfannots` reports printed pages directly

### Flashcard Conventions

- **One card per file** in `cards/<deck>/<card-name>.md`
- Each card has YAML frontmatter: `id`, `deck`, `tags`, `status`, optional `reverse`
- `status: ready` = included in builds; `status: draft` = excluded
- `reverse: true` = Mochi drills both directions (term→def AND def→term)
- **Cloze cards:** citation on the `front` only (no back side, never `reverse: true`)
- **Front/back cards:** citation on the `back` (or both sides)
- Weekly tags: `week-1`, `week-2`, ..., `week-8` + `foundation` for agent-generated starter cards

### Subagents (Specialized AI Agents)

Defined in `.opencode/agent/`, pinned to cost-tiered models:

- **summary-author** (Opus 4.8, max reasoning) — Writes chapter summaries + external examples
- **summary-auditor** (Opus 4.8) — Verifies citations and labels (read-only correctness checker)
- **text-cleaner** (Haiku 4.5) — Cleans raw text for TTS (mechanical, cheap)
- **tts-engineer** (Sonnet 4.6) — Maintains TTS pipeline

After editing an agent file, **restart OpenCode** for changes to take effect.

---

## Typical Session

1. Open terminal in this directory
2. Run `opencode` (or your preferred CLI to start the agent)
3. Agent asks which mode you want (tutor / Q&A / flashcards)
4. Use slash commands or natural language to work

Example interactions:

```
# Tutor mode
You: /tutor stakeholder needs

Agent: [Explains stakeholder needs with citations from §2.3.1 (p.45), etc.]

# Q&A mode (week 3)
You: /qa 3

Agent: Question 1: What are the three fundamental life cycle stages?
You: Concept, development, production...
Agent: Close, but you're conflating stages with phases. Let's clarify...

# Flashcard mode
You: I highlighted 20 new terms in chapter 4. Process them.

Agent: [Loads mochi-cards skill, runs `mochi highlights`, shows you the stubs]
```

---

## Git Conventions

- **Don't commit secrets:** `tools/mochi/.mochi-key` and `highlights/mochi-sync.json` are gitignored
- `dist/` and `audio/` are regenerable build artifacts (gitignored)
- `tools/tts/` is excluded (large model files, ~114MB)
- Commit flashcards (`cards/**/*.md`) and progress logs (`qa/*.md`, `study-plan.md`)

---

## Weekly Workflow Summary

```bash
# Monday: Week N starts
opencode
> /week-start N
# Agent generates foundation cards, quizzes you cold, syncs cards

# Tue–Thu: Read + highlight
# (In Preview, highlight the Handbook; colors = card types)

# Friday: Process highlights
tools/mochi/mochi highlights
# Edit stubs in cards/_inbox/, move to appropriate decks, set status: ready

# Saturday: Build + sync
tools/mochi/mochi build -week N && tools/mochi/mochi lint && tools/mochi/mochi sync -week N

# Daily: Review in Mochi app (15–20 min)

# Sunday: Q&A session
opencode
> /qa N
# Agent drills you, logs gaps
```

---

## Questions?

- **Agent behavior:** See `AGENTS.md` (instructions for OpenCode)
- **Study schedule:** See `study-plan.md` (8-week breakdown)
- **Flashcard workflow:** See `.opencode/skills/mochi-cards/SKILL.md`
- **Tutor/Q&A details:** See `.opencode/skills/se-tutor/SKILL.md` and `.opencode/skills/se-qa/SKILL.md`
- **Handbook index:** Start with `memory/sehb5-index.md` to locate content

---

## License

Personal study workspace. The INCOSE SE Handbook is copyrighted by INCOSE and not included in this public repo (users must obtain their own copy via INCOSE membership or purchase).
