# Chapter 5 — Systems Engineering in Practice (pre-reading summary)

This is the "people and neighbors" chapter — the skills SE practitioners need,
their professional ethics, how SE relates to nearby disciplines, and where the field
is heading with digital engineering. It's lower exam weight, so expect a handful of
recall and tell-them-apart questions rather than deep process detail. A short read.
**Bold** = key terms; `> Outside the handbook` boxes are my analogies.

---

## The big picture

Don't over-invest here. The bits most likely to be tested are the **digital
engineering vocabulary** (especially *digital thread* vs. *digital twin*) and the
**"which discipline does what?"** comparisons (SE vs. project management, software
engineering, industrial engineering, operations research). The competency and ethics
material is worth a skim for a stray question or two.

## Competencies and professionalism

First, a definition pair the handbook is careful about: **competence** is your
overall ability to do a job, while a **competency** is one specific skill in the set
— your competencies *add up to* your competence. The reference framework is INCOSE's
**SECF (Systems Engineering Competency Framework)**, which organizes competencies
into five themes: core, professional, technical, management, and integrating. A
companion guide (SECAG) grades proficiency across five levels.

SE practitioners are described as **"T-shaped"**: deep in one engineering discipline
(the vertical stroke) but broad across many (the horizontal stroke), so they can
lead a multidisciplinary team. The job also leans hard on **soft skills** (also
called interpersonal skills) — the subjective, hard-to-measure stuff like
influencing and motivating — alongside the measurable **hard skills**, because SE
practitioners rarely have the authority to *order* outcomes and must **lead through
influence**.

> **Outside the handbook —** Competencies are the *ingredients* (separately
> measurable skills); competence is the finished *dish*. T-shaped is a great film
> director: deep in one craft but conversant in all the others to lead the crew. And
> the hard/soft split is "passing the calculus exam" vs. "talking a panicked
> teammate off a ledge at 2 a.m. before a launch."

On **ethics**: there's always pressure to cut corners for speed or cost, and the
INCOSE Code of Ethics exists to push back. Just know it has three section types:
fundamental principles, fundamental duties to society, and rules of practice. The
chapter also covers **diversity, equity, and inclusion (DEI)**, where the one
distinction worth holding is *equality* vs. *equity*: equality gives everyone the
*same* resources (which may not produce equal outcomes), while equity *proactively
addresses the disparities* so everyone can reach a *similar outcome*.

> **Outside the handbook —** The fence picture: three people of different heights
> trying to see over a fence. Equality hands each the same-size box (the tall person
> already saw fine; the short one still can't). Equity sizes the boxes so everyone
> sees over.

## SE versus its neighbor disciplines

This whole section is comparison fuel. The pattern: each neighbor is defined, then
related to SE.

- **Software Engineering (SWE)** is engineering applied to software. The SE angle:
  software is *malleable* (easy to change), but a single wrong symbol can sink the
  whole system, so SE owns the partitioning and the clean interfaces between
  software and hardware.
- **Hardware Engineering (HWE)** covers the physical (mechanical and electrical)
  elements. Unlike software, hardware performance generally has to be built into the
  first deliverable rather than iterated in later. This is where the **make-or-buy**
  decision lives: *make* means decompose and build further; *buy* means the
  decomposition for that piece *stops*.
- **Project Management (PM)** applies knowledge and tools to meet project
  requirements. The clean split: **PM delivers the overall project's benefits; SE
  delivers the technical aspect.** PMs focus on *project* stakeholders, SE on
  *system* stakeholders. They overlap early (concept and development) and diverge
  later.
- **Industrial Engineering (IE)** optimizes the use of people, equipment, and
  materials to produce things efficiently — heavily overlapping with SE.
- **Operations Research (OR)**, sometimes called management science, turns data into
  better decisions using analytical methods (optimization, simulation, decision
  analysis). It's a *toolbox* SE uses or commissions, not a rival.

> **Outside the handbook —** On a film: the *producer* (PM) owns budget, schedule,
> and delivery; the *director* (SE) owns whether the artistic-technical whole works.
> SWE perfects each Lego brick's molding; SE makes sure the bricks snap together
> into the model on the box. OR is the casino/weather-forecast math — roll the dice
> 10,000 times in software to learn the odds rather than betting once for real.

## Digital engineering (the high-yield bit)

**Digital engineering (DE)** is an integrated, model-and-data-driven way of working
that runs as a continuous thread across disciplines and the whole life cycle. MBSE
is one of its core elements. The vocabulary cluster the exam cares about — get the
boundaries straight:

- A **Digital System Model** is the integrated digital representation of the system
  — a federated set of models acting as the authoritative source of truth.
- A **digital twin** is a digital *surrogate* that emulates the actual system (or
  parts of it) and evolves alongside it through life.
- A **digital thread** is the connective tissue — the interconnected, cross-
  discipline data that links everything across the life cycle.

The distinction to lock in: the **twin** is a live model of the *product*; the
**thread** is the data backbone that runs *across the lifecycle*.

> **Outside the handbook —** A digital twin is NASA's ground simulators during
> Apollo 13 — a live "twin" of the crippled craft engineers could rehearse on; or
> an F1 team running a live model of the car during a race. The digital thread is
> the "tracked-changes + revision history" of the whole system — one continuous
> chain from requirement to CAD to test data, so you can trace any decision end to
> end. Twin = the flight simulator; thread = the maintenance logbook that follows
> the airframe forever.

## Where SE is heading

The chapter wraps with **transformation** — there are three big shifts: from no SE
to full SE, from traditional to agile methods, and from document-based to
model-based work. And the **future of SE** points at AI: the field coined **AI4SE**
(using AI to help human engineers) and **SE4AI** (applying rigorous SE to AI
systems).

> **Outside the handbook —** AI4SE is a co-pilot that drafts your requirements (AI
> *helps* the engineer); SE4AI is applying disciplined SE *to* a self-driving car's
> AI so it's safe. Same letters, opposite directions.

---

## Watch out for (exam traps)

- **Digital twin vs. digital thread.** The twin is a digital *surrogate that
  emulates the product*; the thread is the *data backbone across the life cycle*.
  Both ride on MBSE. Don't swap them.
- **Competence vs. competency.** A competency is one skill for a job; competence is
  your total ability. Competencies *sum to* competence.
- **Equality vs. equity.** Equality = same resources; equity = proactively closing
  disparities so outcomes are similar. The exam favors equity for "similar
  outcomes."
- **SE vs. PM.** PM delivers the overall project; SE delivers the technical part.
  PMs focus on project stakeholders, SE on system stakeholders. They overlap early,
  diverge late.
- **Make-or-buy.** *Buy* stops the decomposition for that piece; *make* keeps it
  going.
- **OR is a toolbox, not a rival** — SE uses or commissions OR studies.
- **SECF vs. SECAG.** SECF is the *framework* of competencies (five themes); SECAG
  is the *assessment guide* (five proficiency levels). And the five themes are core,
  professional, technical, management, integrating — "leadership" is *not* one of
  them.
- **The three transformations** are no-SE→full-SE, traditional→agile, and
  document→model-based. "Manual→automated" is a distractor.
