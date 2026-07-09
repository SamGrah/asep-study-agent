# Chapter 6 — Case Studies (pre-reading summary)

Five real-world stories — some failures, one success — that put the abstract ideas
from earlier chapters into concrete situations. This is the lowest-weight chapter,
and it's not where definitions get tested. The whole value is being able to say
"which case is the poster child for which concept?" Read the stories once, attach
each to its lesson, and move on. `> Outside the handbook` boxes are my own extra
examples.

---

## The big picture

If Chapter 6 shows up on the exam at all, it'll be a matching question — "which case
illustrates X?" So the move is to lock in the one-line lesson for each, and to
remember that only the Øresund Bridge is told as a *success*; the rest are failures
(or, for the incubators, a failed approach fixed by a smarter one). The definitions
themselves live back in Chapters 2 and 3 — cite those concepts, not the cases.

## The five cases

**Therac-25 (radiation therapy).** A medical radiation machine whose software bugs
delivered massive overdoses, killing patients in the mid-1980s. The SE lesson is the
big one: the designers *removed the hardware interlocks* that earlier models had,
trusting the software to handle safety instead — so when the software failed, there
was no backstop. It was also built by evolving old, poorly documented code (a
brownfield reuse), and the root problem traces back to *concept-and-early-development
decisions*, not a stray typo. Poster child for: **safety as an emergent property,
software V&V, and the danger of removing redundancy.**

> **Outside the handbook —** Its cousin is the Boeing 737 MAX — a single sensor's
> software authority with no easy pilot override, the same "trust the software, drop
> the redundancy" failure. Mnemonic: *they traded a steel lock for a line of code.*

**Øresund Bridge (joining Denmark and Sweden).** The success story — a huge
road-and-rail link across a strait, delivered on time and budget. Why it worked: the
team did a disciplined **concept stage** and refused to rush into development before
the problem was well understood. It also nailed **interoperability** — the trains
drive on different sides and use different power in the two countries, so the system
had to switch on the fly — and managed stakeholders cleverly (they put the head of a
key environmental group on the board). Poster child for: **the payoff of solid
concept-stage work, plus interoperability and stakeholder management.**

> **Outside the handbook —** "Go slow to go fast" — finishing the concept stage is
> "measure twice, cut once." The different-sides-of-the-track detail is a vivid
> standards/interoperability picture, like the old "break of gauge" problem where
> two countries' railways used different track widths and everything had to be
> unloaded at the border.

**Stuxnet (cyber-physical security).** A cyber-attack that secretly took over the
industrial controllers running Iran's nuclear centrifuges and spun them to
destruction. Even though the network was *air-gapped* (physically isolated from the
internet), attackers still crossed it via USB drives, supply chains, and
maintenance staff — proving threats come from *both inside and outside* the system
boundary. It used a **zero-day** (an exploit of a flaw nobody had patched yet).
Poster child for: **system security as a first-class SE concern in cyber-physical
systems.**

> **Outside the handbook —** "Jumping the air gap" is like robbing a vault with no
> internet by having a trusted employee carry the loot out in their pocket — security
> is about people and supply chains, not just firewalls. A zero-day is a burglar with
> a key to a lock the locksmith doesn't even know is broken.

**Incubators (design for maintainability).** Donated high-tech infant incubators
kept breaking in the developing world — and most stayed broken, because no one local
could fix them or get parts. The smart redesign (called NeoNurture) deliberately
built the incubator out of *car parts* — headlights for warmth, dashboard fans for
air — so a local auto mechanic could repair it. The lesson: the highest-tech
solution isn't always the best, and **maintainability and logistics have to be
designed in from the start** — it's too late to bolt them on later. Poster child
for: **design for maintainability / matching technology to its real context.**

> **Outside the handbook —** This is "appropriate technology," the same spirit as
> the $25 prosthetic foot or the hand-crank laptop. A system you can't fix in the
> field is a system that's broken most of the time — design for the village mechanic,
> not the factory.

**Autonomous vehicles (AI in SE).** The fatal 2018 Uber self-driving car crash. A
pile of compounding failures: the perception software kept reclassifying the
pedestrian and threw away its tracking each time; a deliberate one-second "action
suppression" delay (meant to avoid false alarms) stopped the brakes from engaging in
time; Uber had *disabled the Volvo's own built-in safety system*; and the safety
culture and operator oversight were weak. Poster child for: **AI validation, the
limits of autonomy, human-machine teaming, and (again) the danger of removing a
redundant safety layer.**

> **Outside the handbook —** The classic AI-validation horror story is the image
> classifier that "learned" to spot wolves by detecting *snow in the background* —
> exactly the kind of brittle perception that failed here. And note the deliberate
> echo of Therac-25: both disasters removed an existing redundant safety layer.

---

## Watch out for (exam traps)

- **Match each case to its lesson.** Therac-25 → software safety / removed
  redundancy / brownfield reuse. Øresund → concept-stage rigor + interoperability.
  Stuxnet → security of a cyber-physical system. Incubators → maintainability /
  appropriate technology. Autonomous vehicles → AI validation + human-machine
  teaming. Don't swap Therac (medical software) with the Uber case (AI autonomy) —
  both are safety failures, but the named lesson differs.
- **Only Øresund is the success.** The others are failures (or a failed approach
  later fixed).
- **Therac-25's root cause was concept/early-development decisions** — reusing
  undocumented software and dropping hardware interlocks — *not* "a programming
  typo."
- **Air gap ≠ safe.** Stuxnet's whole point is that an isolated network was still
  breached, because threats come from inside and outside the boundary. A "zero-day"
  is an exploit of an unpatched, unknown flaw.
- **Incubators are about maintainability, not price.** Cost is the smaller hurdle;
  the real lesson is designing for repair in context, from the start.
- **The Uber crash had many causes**, not one villain — and it deliberately mirrors
  Therac-25 in removing an existing safety layer.
- **Don't cite Chapter 6 for definitions.** The actual definitions of safety,
  security, maintainability, and verification live in Chapters 2 and 3 — that's
  where the real exam content is.
