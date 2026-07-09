# Chapter 2 — System Life Cycle Concepts, Models, and Processes (pre-reading summary)

This is the heavyweight chapter — the one the exam draws from most. The good news
is it's mostly a big, organized list: how a system moves through life, the few ways
you can structure that journey, and the 30 processes that do the actual work. Read
this overview so the long process catalog ahead feels like old friends rather than
a wall of text. **Bold** = the terms worth locking in; the `> Outside the handbook`
boxes are my analogies (not official wording).

---

## The big picture

If you only nail one chapter, make it this one. The heart of it is **the 30
life-cycle processes**, split into four families. The exam rewards three things:
knowing each process's *purpose*, knowing *which family* a process belongs to, and
telling apart the famous look-alike pairs (stakeholder vs. system requirements,
architecture vs. design, verification vs. validation, integration vs. transition).
Everything before the processes — stages, decision gates, life-cycle models — is the
setup that makes them make sense.

## How a system moves through life

A system's life gets chopped into **stages**, with go/no-go checkpoints between
them. The big mental-model correction up front: this is **not a one-way conveyor
belt**. Stages can overlap, run in parallel, be re-entered, and you can even retire
one piece of a system while the rest keeps running.

The six generic stages are **Concept, Development, Production, Utilization,
Support, and Retirement**. In plain terms: figure out what's needed and explore
options (Concept); design something that meets the agreed needs (Development);
actually build and qualify it (Production); put it to use (Utilization); keep it
running with spares and fixes (Support, which runs *alongside* Utilization); and
finally take it out of service and dispose of it responsibly (Retirement). A point
the exam likes: you plan for retirement and disposal way back in Concept and
Development — not as an afterthought.

> **Outside the handbook —** Building a house. Concept = deciding you need a
> three-bed family home and where. Development = the architect's drawings and
> permits. Production = construction. Utilization = living in it for decades.
> Support = the plumber, roof repairs, repainting (happening the whole time you
> live there). Retirement = demolition and hauling the rubble to a recycler. And
> you decide where the debris goes *before* you pour the foundation.

**Decision gates** are the management checkpoints between stages — the "should we
keep going?" moments. At each one the options are roughly: move on to the next
stage, keep going on this one (maybe after rework), drop back to an earlier stage,
pause the project, or kill it outright. Approved stuff gets locked under
configuration management.

**Technical reviews and audits** are different beasts — they're the *engineering*
checks ("is the design actually right yet?"), whereas a decision gate is the
*business* call ("do we fund the next stage?"). They often happen together but
answer different questions. The classic defense sequence runs ASR → SRR → SFR →
PDR → CDR → TRR → FCA → SVR → PRR → PCA, and they line up against three locked
**baselines**: functional, then allocated, then product. Good practice is to make
reviews event- or risk-driven, not just calendar-driven.

> **Outside the handbook —** A decision gate is a venture investor's funding round:
> "show me the milestone, then I release the next tranche — or pull the plug." A
> technical review is more like a design crit. PDR roughly means "the approach
> looks sound," CDR means "it's mature enough to build," FCA asks "did it meet its
> specs?", and PCA asks "does the as-built actually match the as-documented?"

## The three ways to structure the journey

A **life-cycle model** is just the framework you choose for sequencing the stages.
There are three families, and the exam wants you to tell them apart by two
questions: *how much do you know at the start?* and *how many times do you deliver?*

- **Sequential**: you know the full requirements up front and deliver once. Stable,
  predictable, high-assurance — good for construction and safety-critical work. The
  famous example is the **Vee model**, where every definition step on the way down
  has a matching test step on the way up. (Worth knowing: the Vee is technically a
  sequential model, but its processes still run concurrently, iteratively, and
  recursively underneath.)
- **Incremental**: you still know the full requirements up front, but you build and
  deliver in planned chunks to manage risk and get value sooner.
- **Evolutionary**: you *don't* fully know the requirements yet, so you discover
  them as you go, delivering a usable version each cycle. **Agile is a type of
  evolutionary development**; so is DevOps (continuous integration and delivery).

> **Outside the handbook —** Sequential is building a bridge — you can't ship half
> a bridge; plan it all, deliver once. Incremental is a video game released in
> known chapters or DLC — you know the whole game, you just ship it in pieces.
> Evolutionary is a startup app that A/B-tests its way to a product nobody could
> fully spec on day one. Quick memory hook: Sequential = *know all, build once*;
> Incremental = *know all, build in chunks*; Evolutionary = *discover as you go*.

## The 30 processes, in four families

A **process** is just a set of activities done to achieve some outcome for a
purpose. The handbook organizes all 30 into four families — and being able to sort
any process into the right family is a guaranteed exam skill:

1. **Agreement Processes (2)** — the buyer/seller handshake.
2. **Organizational Project-Enabling Processes (6)** — the *company* setting
   projects up to succeed.
3. **Technical Management Processes (8)** — managing *the project*.
4. **Technical Processes (14)** — actually engineering *the system*.

A small but tested detail: every process is described the same way (purpose,
description, inputs/outputs, activities, tips), with the purpose statements quoted
exactly from the standard. Three words you'll see defined precisely:
**concurrency** = running two or more processes in parallel at the *same* level;
**iteration** = looping back and forth between processes at the *same* level;
**recursion** = applying the same set of processes *again at the next level down*
the hierarchy (and back up). The distinguishing feature is *same level vs.
different level*.

> **Outside the handbook —** Concurrency is two chefs cooking different dishes at
> once. Iteration is one chef tasting and re-seasoning the same sauce. Recursion is
> applying the same recipe-method to the sauce, then to the stock inside the sauce,
> then to the broth inside the stock — same method, one level deeper each time,
> until you hit an ingredient you just buy.

### Agreement Processes (the handshake)

Two processes, two viewpoints on the same deal. **Acquisition** is the *buyer's*
side — getting a product or service that meets the acquirer's requirements.
**Supply** is the *seller's* side — providing something that meets the agreed
requirements. Supply is the bigger umbrella the supplier works under.

> **Outside the handbook —** Two sides of one handshake: the homeowner *acquiring*
> a kitchen remodel writes the request and picks a contractor; the contractor
> *supplying* it runs all the technical work to meet the spec.

### Organizational Project-Enabling Processes (corporate, not project)

These six operate at the *company* level to make projects possible: **Life Cycle
Model Management** (the org's standard playbook of processes), **Infrastructure
Management** (tools, facilities, IT — the non-human resources), **Portfolio
Management** (which projects to start, sustain, and close), **Human Resource
Management** (people and their competencies), **Quality Management** (org-wide
quality standards), and **Knowledge Management** (capturing and reusing what the
org learns).

> **Outside the handbook —** Picture a film *studio*, not one movie. Portfolio =
> which films to greenlight. Infrastructure = the sound stages and cameras. HR =
> casting and crew. Quality Management = studio standards. Knowledge Management =
> the lessons-learned vault. Life Cycle Model Management = the standard production
> playbook every film tailors.

### Technical Management Processes (running the project)

Eight processes for managing the project itself: **Project Planning** (make the
plans), **Project Assessment and Control** (check status and steer), **Decision
Management** (the home of formal trade studies), **Risk Management** (find,
analyze, treat, and watch risks — owns the risk register), **Configuration
Management** (keep the system's configuration consistent and under control — owns
the baselines), **Information Management** (handle the project's information),
**Measurement** (collect and report objective data — owns the MOE/MOP/TPM metrics),
and **Quality Assurance** (confirm the quality process is actually being followed).

A nuance the exam tests: **Quality Assurance** (project-level, "are we following
the process?") is distinct from **Quality Management** (org-level, in the family
above). And risk management has a mirror twin — opportunity management — because
since risk can never hit zero, you balance the two.

> **Outside the handbook —** Running one construction project: Planning is the Gantt
> chart; Assessment & Control is the weekly status meeting that steers; Decision
> Management is the formal "brick vs. timber" trade study; Risk Management is the
> risk register on the wall; Configuration Management is version control on the
> blueprints; Measurement is the burn-down dashboard; QA is the inspector checking
> you followed code.

### Technical Processes (engineering the system) — the highest-yield set

These 14 are where the real engineering happens, and they roughly flow in order.
Here's the whole chain in plain terms:

1. **Business or Mission Analysis** — figure out the actual problem or opportunity
   and pick a class of solution. This kicks off the whole life cycle.
2. **Stakeholder Needs and Requirements Definition** — capture what the users and
   stakeholders *want*, in their own terms.
3. **System Requirements Definition** — translate those wants into a precise,
   verifiable *technical* spec. (This is also where verification criteria get
   attached to each requirement.)
4. **System Architecture Definition** — lay out the high-level structure: the major
   pieces and how they're organized (the *what* and *where*).
5. **Design Definition** — work out the buildable detail for each piece (the *how*).
6. **System Analysis** — run the math/models to support decisions (answering "what
   happened / what could happen / what should we do / how should it be defined?").
7. **Implementation** — actually build, code, or fabricate each element.
8. **Integration** — assemble the elements together and check that the interfaces
   between them work.
9. **Verification** — prove, with objective evidence, that you **built it right**
   (it meets the specified requirements).
10. **Transition** — install and field the system in its real operating
    environment.
11. **Validation** — prove that you **built the right thing** (it actually does the
    job stakeholders needed, in real use).
12. **Operation** — use the system to deliver its service.
13. **Maintenance** — keep it able to deliver that service (runs alongside
    Operation).
14. **Disposal** — end its existence responsibly, handling waste and retired parts.

The verification methods (used for both verification and validation) are
**Inspection, Analysis, Demonstration, and Test** — worth memorizing as a set.

> **Outside the handbook — the whole flow as building a house:** Business/Mission
> Analysis = "this town needs housing." Stakeholder Needs = the family says "three
> bedrooms, sunny kitchen, near schools." System Requirements = the engineer turns
> that into "at least 110 m², south-facing glass, R-30 insulation." Architecture =
> the floor plan and where the load-bearing walls and utilities go. Design = the
> detailed blueprints (stud spacing, wire gauge, pipe size). System Analysis = the
> structural and energy calcs to choose between options. Implementation = framing
> walls and pouring concrete. Integration = assembling the pieces and checking the
> plumbing meets the wiring meets the frame. Verification = the inspector confirms
> it's built to code ("built it right"). Transition = handing over the keys and
> hooking up utilities at the actual lot. Validation = the family lives there a week
> and confirms it's the home they wanted ("built the right thing"). Operation =
> living in it. Maintenance = repairs and upkeep. Disposal = eventual demolition
> and recycling.

---

## Watch out for (exam traps)

- **Stakeholder requirements vs. system requirements.** Stakeholder = what users
  *want*, in their language. System = the *technical, verifiable* version of that.
  Don't pick "stakeholder" when the item is a measurable engineering spec.
- **Architecture vs. design.** Architecture is the high-level structure (the *what*
  and *where*); design is the buildable detail (the *how*).
- **Verification vs. validation** (the most-tested pair). Verification = "did we
  build it *right*?" (meets the spec). Validation = "did we build the *right*
  thing?" (meets the real need, in real use). They share the same four methods.
- **Integration vs. transition.** Integration is *assembling* the pieces and
  checking interfaces; transition is *installing/fielding* the finished system.
  Assembling is not moving-in.
- **Operation vs. maintenance.** Operation is *using* it; maintenance is *keeping
  it usable*. They run at the same time.
- **Which family is a process in?** The classic trap: Risk, Configuration, and
  Measurement are **Technical Management**, not Technical. And Quality *Management*
  is org-level (Enabling) while Quality *Assurance* is Technical Management.
- **Sequential vs. incremental vs. evolutionary.** Sort by *how much you know at
  the start* and *how many deliveries*: full/one, full/many, partial/many. Agile is
  a flavor of evolutionary.
- **Decision gate vs. technical review.** A gate is the management go/no-go (and can
  terminate the project); a review/audit checks technical progress. FCA confirms it
  met its functional spec; PCA confirms the as-built matches the as-documented.
- **Concurrency vs. iteration vs. recursion.** Same-level parallel, same-level
  looping, and next-level-down repetition. The discriminator is *same level vs.
  different level*.
- **Business or Mission Analysis comes first** — it defines the problem and picks a
  solution class *before* stakeholder requirements. Don't confuse the two.
- **Purpose statements are quoted exactly** from the standard, so watch for
  near-miss wording on the exam (e.g., verification "meets its *specified
  requirements*" vs. validation "meets *stakeholder needs / intended use*").
