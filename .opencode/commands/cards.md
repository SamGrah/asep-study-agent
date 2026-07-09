---
description: Process PDF highlights into ready cards and build the bundle
---
Flashcard mode (mochi-cards skill). Current study week: $ARGUMENTS

1. Run `tools/mochi/mochi highlights` and report new highlights by category.
2. For each stub in cards/_inbox/: write an atomic card (front = one question,
   back = answer + §x.y (p.N) citation, verbatim handbook wording for definitions),
   file it under the right deck, set status: ready. Use cloze ({{1::...}}) for
   lists/enumerations — clozes go on the FRONT. Fold near-duplicates into existing
   foundation cards instead of duplicating.
3. Summarize anything new in cards/_discuss.md (blue highlights) and offer to explain.
4. Run `tools/mochi/mochi lint`, fix any failures, then `build -week $ARGUMENTS`,
   then `stats`. Remind me to import dist/sehb5-cards.mochi (File → Import).
