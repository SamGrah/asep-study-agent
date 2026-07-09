---
name: mochi-cards
description: Use for "flashcard management" mode — creating, organizing, and exporting Mochi flashcards for ASEP study. Triggers when the user wants to make/edit flashcards, process PDF highlights into cards, build a Mochi import bundle, or mentions Mochi, decks, or cards. Covers the highlight→card pipeline and the Go tool in tools/mochi.
---

# Mochi Flashcards

Local markdown files under `cards/` are the **source of truth**. The Go tool in
`tools/mochi` pushes them to Mochi. The user is on a **Mochi Pro account**, so the
**primary path is `mochi sync`**, which talks to the live REST API. The offline
`.mochi` bundle (`mochi build`) is retained only as a backup/export.

## The Go tool
Build once from `tools/mochi/`. On this machine (Go 1.22 + Apple CLT) the default
internal linker omits `LC_UUID`, which makes the binary fail to launch with a
`dyld: missing LC_UUID load command` error. **Build with the external linker:**

```
cd tools/mochi && go build -ldflags=-linkmode=external -o mochi .
```

Run all commands from the **repo root**:

- `tools/mochi/mochi highlights` — runs `pdfannots` on the Handbook, classifies
  new highlights by color, dedupes against `highlights/processed.json`, writes
  draft card stubs to `cards/_inbox/`, and logs blue items to `cards/_discuss.md`.
- `tools/mochi/mochi build [-week N]` — compiles `cards/**/*.md` into
  `dist/sehb5-cards.mochi` (a zip containing `data.edn`). Skips files that are
  drafts (`status: draft` or containing `TODO front`) and any `_`-prefixed dir.
  **Build also pins each card's generated `id:` into its front matter** so later
  edits/moves keep the same id. Never hand-edit or copy `id:` lines; build errors
  on duplicate ids. **Always run `build` before `sync`** — sync requires every
  card to have a pinned id.
- `tools/mochi/mochi sync [-week N] [-dry-run] [-prune]` — **primary path.** Pushes
  cards to the Mochi API: creates decks (nested by path), then creates/updates each
  card. Records `local-id → Mochi-id` in `highlights/mochi-sync.json` (gitignored)
  so re-syncs **update in place** (preserving review history) instead of
  duplicating. `-dry-run` reports create/update counts without calling the API;
  `-prune` trashes remote cards whose local file is gone. Term `reverse: true`
  cards are sent as a single card with the native `review-reverse?` flag (no
  synthetic second card).
- `tools/mochi/mochi decks list|archive <id>...` — list or archive remote decks.
- `tools/mochi/mochi processed [count|list]` — inspect the dedupe ledger.

### API key & sync state
- The key lives in **`tools/mochi/.mochi-key`** (gitignored, mode 600). Resolution
  order: `-key K` flag → `MOCHI_API_KEY` env → that file. Auth is HTTP Basic (key
  as username, empty password).
- `highlights/mochi-sync.json` (gitignored) is the per-account id ledger. Don't
  commit it; it's machine/account state, not source.
- API allows **one concurrent request**; sync is sequential with 429 backoff.

Each study week: `mochi build -week N && mochi lint && mochi sync -week N`.

## Highlight → card pipeline
The user highlights passages in **Apple Preview** (in the main Handbook PDF) and
saves. Then:
1. Run `mochi highlights`. It reports new highlights grouped by category and
   drops a stub per highlight in `cards/_inbox/`.
2. For each stub, write a high-quality card: fill the front (question) and back
   (answer), keep the `source:`/`section:`/`highlight:` front matter, set
   `status: ready`, and move it to the correct `cards/<deck>/...` path.
3. `mochi build`, then the user imports.

### Color convention (set in Preview)
| Color | Category | Default deck |
|-------|----------|--------------|
| Yellow | term/definition | `terms` |
| Green | process | `processes` |
| Pink/Red | exam-trap / must-memorize | `exam-traps` |
| Purple | list/enumeration | `lists` |
| Blue | explain-deeper (NO card; logged to `cards/_discuss.md`) | — |

Classification is by hue, so exact Preview shades don't need to match.

## Card file format
One card per `.md` file. Front matter + body; the body **is** the Mochi card
content (front, then a `---` separator line, then back):

```markdown
---
deck: processes/technical-processes   # optional; else inferred from folder path
category: exam-trap
source: SEHB v5 §2.3.5.9 p.138
section: "Verification Process"        # optional
highlight: <hash>                      # optional; set by the stub generator
tags: [verification, technical-process]
status: ready                          # 'draft' (or a TODO body) excludes from build
---
Front: the question

---

Back: the answer, with a citation like (§2.3.5.9, p.138).
```

> **Do NOT start the front with a Markdown `#` heading** — Mochi renders H1 at a
> large size. Keep the question as plain text (bold inline emphasis is fine).

- **Deck** comes from the `deck:` field, else from the folder path under
  `cards/` (nested decks via `/`). Decks are auto-created in the bundle.
- **Card id** is derived deterministically if omitted, so re-imports update
  rather than duplicate.

## Foundation cards (per week) — standing workflow
When the user starts a new study week (see `study-plan.md`), proactively create
a **starter set of foundation flashcards** covering that week's core SEHB
sections, BEFORE the user's highlights come in. PDF-annotation cards then
*supplement* these.

- **Source:** the sections listed for that week in `study-plan.md`. Read them
  via `memory/sehb5-index.md` + `memory/sehb5-text/sehb5-full.txt`.
- **Coverage:** for each section hit the essentials — key term/definition
  (verbatim), each process's purpose + key inputs/outputs, and the classic
  exam distinctions (e.g. verification vs validation). Aim ~15–30 cards/week.
- **Tagging:** every foundation card gets `status: ready` and
  `tags: [foundation, week-N, <topic>...]`. The `foundation` + `week-N` tags
  let the user filter/review them as a set and distinguish them from
  highlight-derived cards.
- **Decks:** file by topic as usual (`terms`, `processes/...`, `exam-traps`,
  `lists`), NOT a per-week deck — weekly grouping is via the `week-N` tag.
- **Dedupe vs highlights:** foundation cards are authored, not hash-tracked.
  When later processing highlights that restate a foundation card, fold the
  nuance into the existing card or skip it rather than duplicating.
- After generating, run `mochi build` so they're ready to import.

## Special card formats
- **Cloze deletion (Mochi `{{...}}`):** wrap hidden text in double braces;
  group with `{{1::x}}` `{{2::y}}` for separate review variants. Ideal for
  **list recall** and definitions in context. Prefer clozes over plain
  front/back for enumerations and the "name the N…" lists (purple highlights).
  **Cloze cards are FRONT-ONLY:** no `---` separator and **no back side.** Put the
  SEHB citation inline on the front. Reviewing is single-tap (reveal the hidden
  text); a back side would force a pointless second click in rapid-fire review.
  A cloze card must **never** set `reverse: true` (reversing a cloze is nonsense).
  Lint enforces all three: clozes balanced/non-empty, no back content, no reverse.
  Pattern:

  ```markdown
  The six generic life-cycle stages:
  {{1::Concept}} → {{2::Development}} → … (§2.1.2, p.26)
  ```
- **Reverse cards (`reverse: true` front matter):** Use ONLY for symmetric
  term↔definition cards (so you drill both "define X" and "what term means …?").
  Never on question/answer cards (a swapped Q/A is nonsense) and never on cloze
  cards (lint rejects cloze+reverse).
  - **`mochi sync` (primary):** sends one card with the native Mochi
    `review-reverse?: true` flag — Mochi reviews it both directions; no second card.
  - **`mochi build` (offline bundle):** emits a second swapped card (id suffixed
    `r`) since the EDN format has no reverse flag.
  `mochi glossary` sets `reverse: true` on **term** cards automatically;
  acronym cards are deliberately not reversed (expansion→acronym recall is
  low-yield and would double 245+ cards' review load).
- **Process I/O cards:** author these from each process section's IPO
  description (§2.3.x) — purpose + key inputs/outputs. (Appendix E's two-column
  layout doesn't parse reliably, so I/O cards are authored by hand; the old
  unwired auto-generator was removed.)
- **Figures/images:** the build has no attachment support (deliberately
  deferred). For dual coding on key figures (V-model, Fig 1.4 cost-commitment
  curve, N² diagrams), paste the figure into the card **directly in the Mochi
  app** after import — realistic need is only ~5–10 figures.

## Reference deck generators
- `mochi glossary` builds **acronym** cards (Appendix B, p.305) and **term**
  cards (Appendix C, p.311) into `cards/reference/`. Re-runnable (idempotent by
  slug, and it preserves pinned `id:`s so review history survives regeneration).
  Distinct entries that slug identically (e.g. **Ai** vs **AI**) get a hash
  suffix instead of silently overwriting. The exam tests verbatim wording, so
  they're high-yield (~300 cards).
- **Each reference card is tagged with the study week in which its term first
  appears in the body** (`week-N`), via `firstWeek`/`weekOfPage` in
  `appendix.go`. Terms that never appear in the body fall to `week-7` (review).
  This prevents showing acronyms/terms from not-yet-studied weeks.
- After any generation, run `mochi lint`, then `mochi build`, then `mochi sync`.

## Week-scoped builds/syncs (don't surface future-week material)
`-week N` includes only cards whose highest `week-N` tag is ≤ N (cards with no
week tag are always included). Both `build` and `sync` accept it. Each study week,
run with the current week number so reference + foundation + highlight cards stay
limited to covered content:

```
mochi build -week 1      # pin ids; offline bundle for weeks 1..1
mochi sync  -week 1      # push weeks 1..1 to Mochi (create/update via API)
```
Sync is idempotent (the id ledger maps each local card to its Mochi card), so
running `-week N` each week simply creates that week's newly-unlocked cards and
updates any edited earlier ones.

## Card-writing guidance
- Prefer atomic cards (one fact/idea each); split compound highlights.
- Always include the SEHB citation: on the **back** for front/back cards, inline on
  the **front** for cloze cards (`mochi lint` enforces this).
- Use the handbook's verbatim wording for definitions/acronyms.
- Good types: term→definition, process→purpose/inputs/outputs, "compare X vs Y"
  (exam traps), and list-recall via cloze.
- **Contrast cards:** for every confusable pair (verification/validation,
  architecture/design, product/process…) add a "both do Z — what distinguishes
  them?" card under `exam-traps`. Discrimination is what MCQ distractors test.
- When generating cards from `qa/gaps.md`, target the specific misconception.

## What `mochi lint` enforces (gate before every build/sync)
- `status:` is exactly `ready` (anything else on a non-draft card fails).
- No `TODO`; cited pages within 1..339; cloze `{{…}}` balanced and non-empty.
- No two ready cards share a normalized front (case-sensitive) or a pinned id.
- **Front/back cards:** separator present; front and back both non-empty; SEHB
  citation (`§…` or `p.N`) **on the back**.
- **Cloze cards (front-only):** **no** back side after `---`; citation anywhere on
  the front; **`reverse: true` forbidden**.
