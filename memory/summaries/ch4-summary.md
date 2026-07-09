# Chapter 4 — Tailoring and Application Considerations (pre-reading summary)

This is the "it depends" chapter. Chapters 2 and 3 gave you the generic SE
machinery; Chapter 4 is about *adapting* it to real situations — different
methodologies, different kinds of systems, different industries. The exam tests it
as definitions and tell-them-apart questions. The high-value stuff is concentrated
in the first half (tailoring and the methodologies); the system-types and
industry-domain material is broad context. **Bold** = key terms; `> Outside the
handbook` boxes are my analogies.

---

## The big picture

Four parts: **tailoring** (how to adapt the standard processes), the **SE
methodologies** (MBSE, Agile, Lean, Product Line Engineering), the **system types**
(greenfield, brownfield, COTS, cyber-physical, systems of systems, and so on), and
a quick lap through **industry domains**. Spend your energy on tailoring and the
methodologies, plus the four types of systems-of-systems — those are the
likely-to-be-tested bits.

## Tailoring

The core idea: you can't just apply a standard straight out of the box. **Tailoring**
means adapting the processes to fit your organization or project, scaled to the
right level of rigor *based on risk*. Too little rigor and you under-engineer and
get burned; too much process and you're paying for paperwork you don't need —
there's a sweet spot in the middle, and it shifts as the project goes.

The single most exam-worthy line: **tailoring can include the deletion,
modification, *or addition* of activities.** People assume tailoring just means
cutting steps — but you can also *add* rigor. There are two levels: *organizational*
tailoring (adapting an external standard to your company) and *project* tailoring
(adapting your company's processes to one project's quirks). And two named traps:
reusing a tailored setup from another project without re-doing the tailoring, and
"using every process just to be safe."

> **Outside the handbook —** Tailoring is buying a bespoke suit off a standard
> pattern — let it out here, take it in there. Both failure modes are real: a
> wedding suit with no alterations (under-tailored, doesn't fit) and paying a tailor
> to hand-stitch every seam on a t-shirt (over-tailored, wasted money). The
> exam's favorite twist: a nuclear plant tailors the standard *up*, not just down.

## The four methodologies

**Model-Based Systems Engineering (MBSE)** is the formal use of *models* — rather
than scattered documents — as the primary product of SE work, beginning early and
continuing through the life cycle. The contrast that gets tested is **model-based
vs. document-based**: in a document-based world, information is spread across specs,
spreadsheets, and reports that are hard to keep in sync and hard to check for
correctness/completeness/consistency; in MBSE, one shared model is the main artifact
(often expressed in SysML). The representative method is OOSEM.

> **Outside the handbook —** Document-centric SE is a pile of Word docs and Excel
> sheets emailed around — the moment one changes, the others are silently wrong.
> MBSE is moving to a single shared model, like going from scattered paper ledgers
> to one live accounting database, or from a paper map to Google Maps where every
> view comes from the same data. It's the "single source of truth" idea again.

**Agile Systems Engineering** is for when requirements keep changing or the
environment is dynamic. The line to remember: **Agile SE is a *what*, not a *how*.**
It's a principle-based mindset (respond to change, learn continuously) — distinct
from the specific software recipes people borrow from (Scrum, XP, SAFe, DevOps).
Those are *hows* from the software world; they inform Agile SE but aren't the same
thing. Agile SE almost always uses incremental or evolutionary delivery underneath.
(And mind the name clash from Chapter 3: *agile systems-engineering* = the process
is agile; *agile-systems engineering* = the product is an agile system.)

> **Outside the handbook —** The mindset (be responsive, keep learning) is the
> *what*; Scrum/XP/SAFe (sprints, stand-ups, backlogs) are particular *hows*.
> Sequential vs. agile is building a house from a frozen blueprint — you can't move
> the kitchen once the slab is poured — versus how a startup builds an app: ship,
> get feedback, rebuild.

**Lean Systems Engineering** comes from Toyota's just-in-time, eliminate-waste
philosophy. The anti-myth line the exam loves: **lean does *not* mean *less* SE —
it means *better* SE.** It's about removing waste so the real engineering can
happen, not about cutting corners. Its three fundamentals are value, waste, and
creating value without waste; value in lean SE is defined as *mission assurance*.
**Waste** is any work that adds no value in the customer's eyes — and there's a
classic list of seven (overprocessing, waiting, unnecessary movement,
overproduction, transportation, inventory, defects), with an eighth often added:
wasted human potential.

> **Outside the handbook —** Lean is the Toyota Production System applied to SE.
> "Waste" is anything a customer wouldn't pay for: a half-finished car sitting in
> inventory, an engineer waiting three days for a sign-off, a report nobody reads.
> Crucial exam point: lean is *not* "faster-cheaper" budget-cutting — don't read
> "lean" as "skimp."

**Product Line Engineering (PLE)** is for building a *family* of similar systems
(think a car platform that yields a sedan, wagon, and SUV). The old bad way is
"clone-and-own" — copy last year's design and hack it, then maintain ten diverging
copies forever. Feature-based PLE instead treats the whole family as one system
with a single shared "factory" and a model of what varies between members.

> **Outside the handbook —** PLE is how Dell lets you configure a laptop — pick the
> CPU, RAM, screen — from one product line instead of designing each laptop from
> scratch. The "factory" is the configurator; clone-and-own is the bad old way.

## The system types

The kind of system you're building changes how SE applies. The pair to know best is
**greenfield vs. brownfield**: greenfield ("clean sheet") is a brand-new design with
little or no legacy baggage — and the handbook notes it's *almost theoretical*,
rarely seen in practice. Brownfield ("legacy") means modifying or replacing an
existing system, usually with explicit *continuity* requirements (it has to keep
working through the change).

> **Outside the handbook —** Greenfield is building a new house on an empty lot.
> Brownfield is renovating a 1920s house while the family still lives in it — you
> work around old wiring (technical debt), keep the water running (continuity), and
> might find asbestos (hazards to remediate). Pure greenfield is almost a myth;
> nearly every real project inherits *something*.

The rest are quicker hits: **COTS-based** systems hinge on the make-vs-buy decision
(and note the vendor, not you, controls the updates, and you use it *without*
modifying its source); **software-intensive** systems are ones where software is
the essential influence; **cyber-physical systems (CPS)** fuse software control with
physical processes via sensors and feedback. There's a nesting the exam likes: a
*networked* CPS connected over the internet is an **IoT** and is always a
**system of systems** — but a CPS connected by non-internet means isn't networked.

The big one here is **System of Systems (SoS)**: a set of systems that together
deliver a capability none of them could alone, where each piece is a useful system
in its own right. The two *principal* distinguishing features are **operational
independence** and **managerial independence** (the constituents can run, and are
run, on their own). And the four SoS types — sorted by how much central authority
there is — are worth memorizing:

- **Directed** — there's a central authority and the constituents are subordinated
  to the SoS's purpose.
- **Acknowledged** — there's a designated manager and resources for the SoS, *but*
  the constituents keep their own ownership, funding, and goals.
- **Collaborative** — no central manager; the constituents cooperate voluntarily
  toward shared purposes.
- **Virtual** — no central authority *and* no agreed central purpose; large-scale
  behavior just emerges.

> **Outside the handbook —** Think of an authority spectrum from tightest to
> loosest. **Directed** = a carrier strike group (ships can sail alone, but the
> admiral commands). **Acknowledged** = a city's emergency response (a designated
> coordinator, but police/fire/hospitals keep their own budgets). **Collaborative**
> = the internet (providers voluntarily honor the same protocols, no boss).
> **Virtual** = the global economy (no manager, behavior just emerges). Memory hook:
> **D**irected = **D**ictator, **A**cknowledged = **A**ppointed coordinator,
> **C**ollaborative = **C**onsensus, **V**irtual = **V**oid.

The remaining types — IoT/big-data systems, service systems, and enterprise systems
— are context. One distinction worth holding: an **enterprise is not the same as an
organization** — organizations *participate in* an enterprise but aren't necessarily
*part* of it, and you shape an enterprise more than you command it.

## The industry domains

The final section walks through ten sectors (automotive, biomedical, aerospace,
defense, infrastructure, oil & gas, power, space, telecom, transportation), listed
alphabetically. This is low-yield — don't memorize each one. Just know the theme is
"same SE grammar, different dialect," and if a question hands you a domain-specific
standard, match it to its sector.

> **Outside the handbook —** A pacemaker engineer and a fighter-jet engineer follow
> the same Vee but speak different regulatory languages. Rough matching: ISO 26262
> → cars, DO-178C / ARP4754A → aircraft, IEC 60601 → medical devices.

---

## Watch out for (exam traps)

- **Tailoring adds *and* deletes.** It can delete, modify, *or add* activities. An
  answer saying "tailoring only removes process" is wrong — and *too much* process
  is a tailoring failure just like too little.
- **MBSE vs. document-based.** MBSE makes the *model* the primary artifact;
  document-based scatters info across files that fall out of sync. Also keep three
  things separate: MBSE is the *approach*, SysML is the *language*, OOSEM is a
  *method*.
- **Agile SE is "a what, not a how."** The process being agile is different from the
  product being an agile system. Scrum/XP/SAFe are software *hows* that inform —
  but aren't identical to — Agile SE.
- **Lean ≠ less SE.** Lean means *better* SE by removing waste. A choice framing it
  as cost-cutting is wrong.
- **Greenfield vs. brownfield.** Clean-sheet/no legacy (and rarely seen for real)
  vs. modifying a legacy system with continuity requirements.
- **The four SoS types.** Discriminate by *authority*: Directed (subordinated),
  Acknowledged (a manager, but independent constituents), Collaborative (voluntary,
  no manager), Virtual (no authority and no agreed purpose). The classic swap to
  avoid is Acknowledged (has a manager) vs. Collaborative (no manager).
- **SoS's two principal features** are operational and managerial independence.
  Geographic spread, emergence, and evolutionary development alone don't make
  something an SoS.
- **CPS vs. IoT.** A CPS always has embedded software; a *networked* CPS over the
  internet is an IoT and always an SoS — but not every CPS is networked. Don't
  assume "every CPS is an IoT."
- **Enterprise ≠ organization.** Organizations participate in an enterprise; they
  aren't necessarily part of it. You influence an enterprise, not command it.
- **COTS.** The *vendor* controls the updates, and you use COTS *without* modifying
  its source code. "The buyer customizes the COTS source" contradicts the
  definition.
