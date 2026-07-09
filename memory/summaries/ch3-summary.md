# Chapter 3 — Life Cycle Analyses and Methods (pre-reading summary)

This is the "-ilities" chapter — all the quality characteristics a system is judged
on (reliability, safety, security, and friends), plus the cross-cutting methods SE
uses (modeling, traceability, interface management, and so on). It's lighter weight
than Chapter 2, and the exam mostly tests it as *definitions* and *telling
near-twins apart* (reliability vs. availability vs. maintainability; safety vs.
security; resilience vs. robustness). Read this so those distinctions are already
sitting in your head. **Bold** = key terms; `> Outside the handbook` boxes are my
analogies.

---

## The big picture

Two halves. First, the **quality characteristics** (QCs) — the "-ilities," formerly
called specialty engineering. These are how stakeholders judge whether a system is
any good, beyond just "does it function." Second, the **SE methods** that cut across
the whole life cycle — modeling, prototyping, traceability, interfaces, frameworks,
and a few more. The single highest-yield cluster is **RAM** (reliability,
availability, maintainability), followed by the safety-vs-security and
resilience-vs-robustness distinctions.

## Quality characteristics (the "-ilities")

A **quality characteristic** is an inherent property of a system tied to a
requirement — basically *how well* it does things, not *what* it does. The handbook
lists them alphabetically on purpose, so no one thinks one outranks another. One
nuance the exam likes: QCs usually generate *non-functional* requirements, but
**safety, security, and resilience can also generate functional requirements**.

> **Outside the handbook —** They're nicknamed "-ilities" because most end in
> *-ility*: reliab-ility, maintainab-ility, availab-ility, scalab-ility. Think of
> them as the *adverbs* of a system — not what it does, but how well and how
> gracefully it does it.

A quick tour of the ones worth recognizing:

**Affordability** is about maximizing value over the whole life, not just the
sticker price. Two ideas live here. First, **life cycle cost (LCC)** is the *total*
cost across the system's life — and the handbook is careful that LCC isn't the same
as "total cost of ownership," which often only counts costs after purchase. Second,
there's a small formula for **cost-effectiveness**:

> **CE = SE / (IC + SC)** — in plain English, *cost-effectiveness equals system
> effectiveness divided by the sum of initial cost plus sustainment cost.* In other
> words, you're measuring how much capability you get per dollar, where the dollars
> include both what you pay up front (initial cost) *and* what you pay to keep it
> running (sustainment cost). A bigger number means more bang for the total buck. The
> point of writing it this way is to stop people from judging a system on purchase
> price alone — the running costs belong in the denominator too.

> **Outside the handbook —** This is the printer trap: the printer is cheap, the ink
> bankrupts you. A pricier Toyota can be cheaper to *own* than a cheap car that
> guzzles fuel and breaks down. And remember the Chapter 1 lesson — about 80% of
> that lifetime cost is locked in by early design choices.

**Agility engineering** is about being able to change quickly and cheaply when the
world shifts. The enabling trick is plug-and-play modularity — loosely coupled
pieces you can swap. (Heads up for a name clash: *agile systems-engineering* means
the engineering *process* is agile, while *agile-systems engineering* means the
*product* is an agile system. Same words, different thing.)

**Human Systems Integration (HSI)** makes sure technology, organizations, and people
actually work together — designing for the human under stress.

> **Outside the handbook —** HSI is why a cockpit or an ATM is laid out for a real
> person. The iPhone is the success story; the confusing Three Mile Island control
> panel is the cautionary tale.

**Interoperability** ensures your system plays nicely with other systems, usually by
agreeing on shared standards. **Logistics engineering** keeps the system supportable
— the spare-parts-and-trained-mechanics web that keeps a fleet flying.
**Manufacturability/producibility** makes sure the thing can actually be mass-
produced affordably.

**RAM — reliability, availability, maintainability** — is the big one, and the three
are completely intertwined. Take them one at a time:

- **Reliability** is the ability to perform without failure for a stated time. Its
  metric is **MTBF — mean time between failure**. A trap to pre-load: *MTBF is not
  the same as expected lifespan.* A batch of bulbs rated 10,000-hour MTBF doesn't
  mean your particular bulb lasts 10,000 hours.
- **Maintainability** is how easily the system can be repaired and restored when it
  does break. Its metric is **MTTR — mean time to repair**. Important distinction:
  *maintainability* is a design property you build in, whereas *maintenance* is the
  actual repair activity that results from that design.
- **Availability** ties the other two together. The simplest way to think about it:

  > **Availability = uptime / total time** — in plain English, *the fraction of the
  > time the system is actually ready to use.* If something is up 90 hours out of
  > every 100, it's 90% available. Because "uptime" depends on how rarely it breaks
  > (reliability) and how fast you fix it (maintainability) — plus how long you wait
  > for parts and paperwork — availability is really a *combination* of reliability,
  > maintainability, and logistics.

  Availability comes in three flavors, and the only difference between them is
  *which delays you count*. **Inherent availability (Ai)** counts only the built-in
  reliability and repair time — no scheduled maintenance, no waiting for parts, no
  admin. **Achieved availability (Aa)** adds scheduled (preventive) maintenance.
  **Operational availability (Ao)** is the real-world number — it adds the time
  spent waiting for parts (logistics delay) and paperwork (administrative delay).
  So Ai is the best-case "perfect garage" figure, and Ao is the honest "real world"
  figure, with Aa in between.

> **Outside the handbook —** Use a car. Reliability = how rarely it breaks down.
> Maintainability = how fast and cheap it is to fix when it does. Availability =
> what fraction of the time it's actually ready to drive — which needs *both*. A
> Formula-1 car is high-performance but low-availability (constant maintenance); a
> Toyota Hilux is the availability champion. Inherent vs. operational availability
> is "in a perfect garage with every part on the shelf" versus "in the real world
> waiting two weeks for a part to ship from overseas."

**Resilience** is about still delivering capability when something bad happens. Its
three fundamental objectives are **avoid, withstand, recover**. The subtle, tested
point: resilience cares about keeping the *capability*, not necessarily keeping the
*architecture* — a resilient system is allowed to adapt and change shape to keep
working. **Robustness** is the narrower idea of just *withstanding* damage without
breaking. So robustness is one slice of resilience, not the whole thing. (Resilience
also absorbs the older idea of survivability.)

> **Outside the handbook —** Robustness is a fortress wall — strong, resists damage
> (withstand only). Resilience is the bigger story: a starfish regrowing a lost arm
> (recover), a power grid rerouting around a downed line (adapt), an immune system
> avoiding infection (avoid). The internet was *designed* for resilience — packets
> reroute around damage.

**Safety** reduces the likelihood of harm to people, property, and the environment.
Two things to remember: safety is an **emergent property** (it lives in the whole
system in its real environment, not in any one part), and engineered systems are
never 100% safe — the goal is acceptable risk, often phrased as **ALARP** ("as low
as reasonably practicable"). **Security** protects the system from *malicious or
disruptive* events. The discriminator between them is *intent*: safety is mostly
about accidents and chance; security is about an intelligent adversary. Security
leans on the **CIA triad** — confidentiality, integrity, availability.

> **Outside the handbook —** Safe *from accidents*, secure *from adversaries*.
> Therac-25 (a radiation machine that overdosed patients) and the Boeing 737 MAX
> are the textbook "safety is emergent" disasters. Stuxnet (a worm that physically
> wrecked Iranian centrifuges) is the textbook security case. CIA triad:
> confidentiality (no one reads it), integrity (no one alters it), availability
> (you can still use it).

Finally, **loss-driven systems engineering** is the umbrella insight that
resilience, safety, security, sustainability, and availability are all really
asking the same question — "what could we lose, and how do we cope?" — so they
should be handled together instead of in separate silos.

## The cross-cutting SE methods

**Modeling, analysis, and simulation** are three *distinct* activities people often
blur. Modeling is *creating* a model; analysis is *examining* something to gain
insight; simulation is *running* a model to predict behavior. A clean rule the exam
uses: simulation always involves a *digital* model, and any examination using a
*physical* model is a *test*. Models are described by their breadth (how much they
cover), granularity (how detailed), and fidelity (how accurately they match
reality). The payoff of going model-based is a single "authoritative source of
truth" instead of a pile of contradicting documents.

> **Outside the handbook —** A wind-tunnel scale model is *physical*, so running it
> is a *test*; a CAD or flight-sim model is *digital*, so running it is a
> *simulation*. Fidelity is the giveaway: a paper map (low fidelity, wide breadth)
> vs. Google Street View (high fidelity).

**Prototyping** is for *learning*, not for shipping — a prototype is meant to be set
aside once it's taught you what you needed. Don't confuse it with "version one."

**Traceability** is the ability to follow the thread of any item through its history
and relationships. Three kinds: **bidirectional** (a link that automatically creates
the reverse link back), **vertical** (up and down the levels — parent/child), and
**horizontal** (across peers at one level and across the life cycle). It's enabled
by configuration management and is the backbone of the "digital thread."

> **Outside the handbook —** Traceability is the audit trail: every requirement
> links up to *why* it exists (a stakeholder need) and down to *how* it's proven (a
> test). Vertical is up and down the family tree; horizontal is sideways across
> peers and stages. It's how you answer "if I change this bolt, what requirements
> and tests does it touch?" — the same idea as tracing a contaminated lettuce back
> to one farm.

**Interface management** is about defining and controlling what crosses the
boundaries between pieces — and it matters because most system failures happen *at
the interfaces*, not inside the boxes. The key framing: it defines *what* crosses
the interface (the characteristics), not *how* that interaction is physically
realized — the "how" (the actual data bus, wiring, Wi-Fi) is a design decision. Two
diagram tools to recognize: an **N2 diagram** is a grid where interactions flow
*clockwise* and a blank cell means no interaction; a **DSM (design structure
matrix)** is the same idea but uses the *opposite* (counterclockwise) convention.

> **Outside the handbook —** The \$125M Mars Climate Orbiter was lost because one
> team used metric units and another used imperial across an interface — a textbook
> interface-management failure. An interface control document is the "contract"
> between two teams about exactly what crosses the boundary — it's what keeps a
> modern car's 150-plus computer modules from fighting.

The last few methods are lower-yield but worth recognizing: **architecture
frameworks** (standard sets of viewpoints like DoDAF, TOGAF, Zachman — the
"required blueprints" of a domain, distinct from a modeling *language* like SysML);
**patterns** (reusable solutions to recurring problems); **design thinking** (the
human-needs-first, diverge-then-converge approach); and **biomimicry** (borrowing
strategies from nature).

> **Outside the handbook —** Architecture frameworks are like building codes that
> say "every house plan needs a floor plan, an electrical plan, and a plumbing
> plan." Biomimicry gave us the kingfisher-beak nose on Japan's bullet train and
> Velcro (copied from burrs on a dog's fur).

---

## Watch out for (exam traps)

- **Safety vs. security.** The discriminator is *intent*: safety guards against
  accidents and hazards; security guards against malice and disruption. Memory
  hook: safe *from accidents*, secure *from adversaries*. Both can be emergent.
- **Reliability vs. availability vs. maintainability.** Reliability = how rarely it
  fails (metric MTBF). Maintainability = how easily it's fixed (metric MTTR).
  Availability = is it ready when needed (a function of both, plus logistics).
  Don't pick "reliability" when the question stresses *being ready when needed* —
  that's availability.
- **Maintainability vs. maintenance.** Maintainability is a *design property* built
  in; maintenance is the *activity* of repairing. One is a noun about the design,
  the other is the act.
- **The three availabilities.** Inherent (Ai) counts only reliability and repair;
  Achieved (Aa) adds preventive maintenance; Operational (Ao) adds logistics and
  admin delays. Ai ≥ Aa ≥ Ao. The discriminator is *which delays you count*.
- **Resilience vs. robustness.** Resilience is the superset (avoid, withstand,
  recover) and may *adapt* the system's shape to keep delivering capability.
  Robustness is just the "withstand" piece. Resilience also subsumes survivability.
- **MTBF is not lifespan.** A distractor equating "mean time between failure" with
  "how long it lasts" is wrong.
- **Modeling vs. analysis vs. simulation.** Three distinct activities — create,
  examine, run-to-predict. And remember: simulation = digital model; examining a
  physical model = a test.
- **Vertical vs. horizontal traceability.** Vertical = up/down levels
  (parent/child); horizontal = across peers and across the life cycle (the
  stakeholder-req → system-req → architecture → product → V&V chain is horizontal).
- **Interface "what" vs. "how."** Interface management defines *what* crosses the
  boundary; the *how* (the physical medium) is a design decision.
- **N2 vs. DSM.** N2 flows clockwise, DSM counterclockwise; a blank cell means no
  interaction.
- **QCs and requirement types.** They usually create non-functional requirements,
  but safety, security, and resilience can create functional ones too — "QCs only
  make non-functional requirements" is a trap.
- **Architecture framework vs. modeling language.** A framework (DoDAF, TOGAF,
  Zachman) is the set of required viewpoints; SysML is the notation. Don't conflate.
