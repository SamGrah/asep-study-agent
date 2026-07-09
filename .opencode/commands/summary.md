---
description: Open/work with a chapter pre-reading summary
---
Chapter pre-reading summary for chapter $ARGUMENTS.

1. Read `memory/summaries/ch$ARGUMENTS-summary.md` and present its key points so I
   can read it before tackling the dense handbook chapter.
2. If the file is missing, generate it: use the `summary-author` subagent
   (Opus 4.8, max reasoning) to author `memory/summaries/ch$ARGUMENTS-summary.md`
   from `memory/sehb5-index.md` + `memory/sehb5-text/sehb5-full.txt` (printed page
   N = the (N+26)th form-feed chunk), then audit it with the `summary-auditor`
   subagent and apply any BLOCKER fixes.
3. Reminders: summaries are a **reading aid, never a citation source**; handbook
   claims cite `§x.y (p.N)`, external examples stay in `> **Outside the handbook**`
   callouts. For audio of the full chapter: `tools/tts/make-chapter.sh $ARGUMENTS`
   (Piper, local) or use macOS Siri "Speak Selection" (see memory/summaries/README).
