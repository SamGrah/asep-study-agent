# ASEP Study Plan — Jun 22 → mid-Aug 2026

**Goal:** Pass the INCOSE ASEP knowledge exam by **mid-to-late August 2026** so the
CSU MBSE course prerequisite is satisfied before the program starts **Sep 1, 2026**
(the hard deadline).
**Budget:** 10–15 hrs/week × 8 weeks ≈ 90–115 hrs.

## Exam facts (verified incose.org, 2026)
- INCOSE Knowledge Exam: **120 multiple-choice questions, 120 minutes**
  (100 scored + up to 20 unscored "beta"; 1 min/question). Single correct
  answer per question; you may review and change answers.
- **Content is 100% from the SE Handbook 5th Edition** (since 15 Mar 2025).
- Computer exam: take from home, **live remote video proctoring**, online
  scratch paper, **immediate pass/fail** on screen. Schedule at your convenience
  after registering.
- ASEP path = **exam only** (no SE work-experience requirement, unlike CSEP).
- "Most people who pass have read the handbook more than once."
- Official practice: 10 sample questions for $10 (format preview only).

## Admin checklist (do in Week 1)
- [ ] Join INCOSE (membership → free handbook download + exam discount).
- [ ] Submit ASEP application (auto-enrolls you for the computer exam).
- [ ] Register for the computer exam; confirm system/firewall requirements.
- [ ] **Schedule the exam for mid-to-late August** (the deadline anchors
      everything; rescheduling is allowed; stay before the Sep 1 hard deadline).
- [ ] Buy/start the two Udemy courses (concepts + practice exams).
- [ ] Take the $10 official practice (10 Q) once, just to see the format.

## Weekly cadence (each week, ~10–15 hr)
- **Foundation cards + cold pretest (start of week):** the agent generates a
  starter deck for the week's sections (tagged `foundation`, `week-N`), then
  **quizzes you on them cold, BEFORE you read** (`/week-start N`). Expect to
  miss most — pretesting measurably boosts encoding of the subsequent reading.
  Your PDF highlights then supplement the deck.
- **Pre-read the chapter summary (~20 min):** before the dense handbook text, read
  `memory/summaries/chN-summary.md` (`/summary N`) to prime the concepts/terms.
  Optionally listen along — macOS Siri "Speak Selection" on the PDF, or the
  Piper MP3 (`tools/tts/make-chapter.sh N` → `audio/chN.mp3`). See summaries/README.
- **Read + course (~5–6 hr):** read the week's SEHB sections; **highlight in
  Preview** (color = card type, see AGENTS.md); watch the paired Udemy concept
  lessons.
- **Flashcards (~3–4 hr):** run `mochi highlights`, turn stubs into cards, then
  `tools/mochi/mochi build -week N` (N = current week, so acronyms/terms from
  not-yet-studied weeks are excluded) and import. **Review daily 15–20 min**
  across all decks (spaced repetition — don't skip days). Grade honestly: a
  hesitant answer is a **miss** — each card should be recalled cleanly in 3+
  separate sessions before you trust it (successive relearning).
- **Q&A session (~1–2 hr):** one Socratic/adversarial drill on the week's
  content, with ~20% interleaved review of earlier weeks (logs to
  `qa/coverage.md` / `qa/gaps.md`; gaps → new cards). `/qa N`.
- **Practice questions (~2–4 hr, ramps Week 5+):** light early, primary in 7–8.

---

## Schedule

### Week 1 · Jun 22–28 — Foundations + setup
- Read **Ch.1 Systems Engineering Introduction (p.1)** and **§2.1–2.2 life-cycle
  terms, stages, decision gates, model approaches (p.25–38)**.
- Do the full **Admin checklist** above.
- [ ] Ch.1 read & highlighted
- [ ] §2.1–2.2 read & highlighted
- [ ] First deck built (terms); daily review habit started
- [ ] Q&A #1 (foundations)
- **Milestone:** tooling + study loop running; exam scheduled for mid-to-late Aug.

### Week 2 · Jun 29–Jul 5 — Technical Processes I (requirements)
- **§2.3.5.1 Business/Mission Analysis (p.103)**, **§2.3.5.2 Stakeholder Needs &
  Requirements (p.107)**, **§2.3.5.3 System Requirements (p.112)**.
- [ ] Sections read & highlighted
- [ ] Cards created + daily review
- [ ] Q&A #2 (requirements) — *high exam weight*
- **Milestone:** requirements process group solid.

### Week 3 · Jul 6–12 — Technical Processes II (architecture/design)
- **§2.3.5.4 Architecture (p.118)**, **§2.3.5.5 Design (p.124)**,
  **§2.3.5.6 System Analysis (p.129)**, **§2.3.5.7 Implementation (p.132)**.
- [ ] Sections read & highlighted
- [ ] Cards created + daily review
- [ ] Q&A #3 (architecture vs design)
- **Milestone:** architecture/design distinction nailed.

### Week 4 · Jul 13–19 — Technical Processes III (integration→disposal)
- **§2.3.5.8 Integration (p.134)**, **§2.3.5.9 Verification (p.138)**,
  **§2.3.5.10 Transition (p.143)**, **§2.3.5.11 Validation (p.146)**,
  **§2.3.5.12–14 Operation/Maintenance/Disposal (p.152–156)**.
- [ ] Sections read & highlighted
- [ ] Cards created + daily review
- [ ] Q&A #4 (V&V) — *highest exam weight; spend extra time here*
- **Milestone:** all 14 technical processes covered.

### Week 5 · Jul 20–26 — Technical Management Processes
- **§2.3.4.1–2.3.4.8 (p.70–100):** Planning, Assessment & Control, Decision,
  Risk, Configuration, Information, Measurement, Quality Assurance.
- [ ] Sections read & highlighted
- [ ] Cards created + daily review
- [ ] Q&A #5 (risk, CM, decision mgmt)
- [ ] **Practice exam #1 (calibration)** — score: ___. Expect it to be rough;
      its job is to expose weak sections early and make Weeks 6–8 targeted.
      Log every miss per handbook section in `qa/gaps.md`.
- **Milestone:** technical management group solid; baseline score known.

### Week 6 · Jul 27–Aug 2 — Agreement + Org-Enabling + cross-cutting
- **§2.3.2 Agreement (p.44)**, **§2.3.3 Organizational Project-Enabling (p.50)**;
  skim **Ch.3 analyses/methods (RAM, safety, security, traceability,
  interfaces)** and **Ch.4 tailoring + MBSE (§4.1, §4.2.1)**.
- [ ] Sections read & highlighted
- [ ] Cards created + daily review
- [ ] Q&A #6 (process groups overview + tailoring)
- [ ] **Cap Mochi review intervals** (deck settings → max interval ~7–10 days)
      so well-known cards still resurface before exam day.
- **Milestone:** **full handbook coverage complete; all decks built.**

### Week 7 · Aug 3–9 — Practice exams + targeted review
- **Practice exams #2–3** (Udemy). After each: review every miss, make cards
  from gaps, re-read the cited SEHB sections. Log misses **per handbook
  section** in `qa/gaps.md` — Week 8 time is allocated by section miss rate.
- Take one of them as a **full exam-day rehearsal**: 120 min timed, one
  sitting, webcam on, online-only scratch paper (proctoring conditions).
- Memorize **Appendix B Acronyms (p.305)** and **Appendix C Terms (p.311)**.
- [ ] Practice exam #2 (score: ___)
- [ ] Practice exam #3 (score: ___) — exam-conditions rehearsal
- [ ] Gap cards created; weak areas logged in `qa/gaps.md`
- **Milestone:** trend vs Week 5 baseline known; gap list drives Week 8.

### Week 8 · Aug 10–16 — Cram + final pass
- **Practice exam #4** (+ #5 if time allows); re-read flagged (blue/`_discuss`)
  sections and Appendices D (N2) & E (I/O, p.321). No new material late —
  consolidate. Allocate review hours by per-section miss rate from `qa/gaps.md`.
- [ ] Practice exam #4 (score: ___)
- [ ] Full deck review pass
- **Go/no-go:** if practice consistently **~80%+**, keep the exam date; if not,
  push into the August buffer rather than risk a real attempt.

### Exam week · Aug 17–21
- Light review only; rest the day before. **Sit the ASEP exam.**
- [ ] Exam taken — result: ___

### Buffer · Aug 22–31
- Retake if needed (review the candidate retake policy/wait time). Certified
  before **Sep 1**.

---

## Risk notes
- **Daily spaced review is non-negotiable** — skipped days collapse retention
  and the Week 7–8 practice gains.
- **Grade reviews honestly** (hesitation = miss; 3+ clean recalls before a card
  counts as known) and **track practice-exam misses per handbook section** —
  unfocused Week-8 review is the classic way to waste the final week.
- Highest-yield topics: **§2.3.5 Technical Processes** (esp. V&V and
  requirements) and **§2.3.4 Technical Management Processes**. Protect Weeks 2,
  4, 5.
- If you fall behind, compress Ch.3/Ch.4 breadth in Week 6 — **do not** sacrifice
  the practice-exam weeks (7–8).
- Read the handbook **at least twice**; the schedule's daily cards + Q&A act as
  the second pass.
