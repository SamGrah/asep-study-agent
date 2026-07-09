# Chapter pre-reading summaries + read-aloud

Concise, exam-targeted summaries of each SEHB v5 chapter, meant to be read (and/or
listened to) **before** the dense handbook chapter itself.

## The summaries (`ch1`…`ch6-summary.md`)
Written in a **conversational, plain-English** voice — meant to be a quick skim that
gives you the lay of the land *before* you read the dense handbook chapter, so you
don't get lost. They are deliberately NOT bibliographic: there are **no `§x.y (p.N)`
citations**. **Bold** is used sparingly, only for the key terms the exam cares about.
Where the handbook uses a formula, the summary spells it out in full plain English
and explains what it means. Each ends with a short **"Watch out for (exam traps)"**
section, and uses `> **Outside the handbook**` callouts for external analogies.

**Important:** summaries are a *reading aid, not a citation source*. Anything inside
a `> **Outside the handbook**` callout is my own illustration and is **not in SEHB**
— never repeat it on the exam as handbook content. The authoritative, citable record
is the handbook itself (and the `cards/`, which keep the verbatim wording + exact
section/page references for drilling).

How they're generated: the `summary-author` subagent (Opus 4.8, max reasoning)
drafts from the corpus; the `summary-auditor` subagent (Opus 4.8) verifies every
citation and external-example label before the summary is trusted. See `/summary N`.

## Read-aloud of the handbook

You have two free, local paths — both narrate the **handbook chapter text** (not the
summaries).

### Option A — macOS Siri "Speak Selection" (live read-along, best voice)
1. System Settings → Accessibility → **Spoken Content**.
2. **System Voice** → pick a voice → the menu → **Manage Voices…** → download an
   **English (Siri)** voice (e.g. *Siri Voice 4*). These are far more natural than the
   default robotic voices.
3. Enable **Speak Selection** and set a hotkey (default ⌥+Esc).
4. In Preview, select handbook text and press the hotkey to hear it while you read.
   Offline, free, no setup beyond the voice download.

### Option B — pre-generated MP3s (replayable audio files)
A TTS pipeline (`tools/tts/`) turns cleaned chapter text into MP3s.
- Generate a chapter: `tools/tts/make-chapter.sh N`
  → reads `tools/tts/clean/chN.txt`, writes `audio/chN.mp3`.
- The cleaned text for a chapter comes from `memory/sehb5-text/sehb5-full.txt`; for
  new chapters, clean the text first (the `text-cleaner` subagent does this) into
  `tools/tts/clean/chN.txt`.
- Engine is pluggable via the `TTS_BACKEND` env var: `piper` (local/free, **default**)
  plus drop-in stubs for `openai` and `elevenlabs` in `tools/tts/tts.sh`.

> **TODO (future session) — use a cloud voice.** Piper's quality was judged not good
> enough (too robotic for read-along), so **no chapter MP3s are kept**. When you have
> an API key, finish the `openai` or `elevenlabs` stub in `tools/tts/tts.sh` (the API
> call is already sketched there), then run e.g.
> `TTS_BACKEND=elevenlabs ELEVENLABS_API_KEY=… tools/tts/make-chapter.sh 1`.
> ElevenLabs/OpenAI TTS will sound far more natural. For live read-along in the
> meantime, **Option A (macOS Siri voices) is the best free path.**

`audio/` and `tools/tts/clean/` are gitignored (regenerable). The local Piper voice
model in `tools/tts/models/` can stay (it's the free fallback) or be deleted to
reclaim ~110 MB — it'll re-download if ever needed.
