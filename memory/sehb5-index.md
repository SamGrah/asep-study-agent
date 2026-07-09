# SEHB v5 Index & Memory File

Authoritative navigation map for the **INCOSE Systems Engineering Handbook, 5th
Edition (2023)** — `INCOSE_SEHB5_Systems_Engineering_Handbook_v5.pdf`. Used by
**handbook-search** and **se-qa** modes to locate and cite content. Every tutor
answer and every Q&A question must trace to a section/page here.

## How to cite & search

- **Citations use PRINTED page numbers** (what the reader sees / what Preview &
  the TOC show), e.g. *SEHB v5 §2.3.5.9 Verification, p.138*.
- **Grep corpus:** `memory/sehb5-text/sehb5-full.txt` (pdftotext `-layout`,
  372 physical pages separated by form-feed `\f`).
- **Page offset:** `physical_page = printed_page + 26` (printed p.1 = physical
  p.27). Front matter (roman numerals i–xxi) occupies physical pages 1–26.
  To jump in the grep corpus to printed page _N_, go to the (N+26)th `\f` chunk.
- Source standard: **ISO/IEC/IEEE 15288:2023**. Published 5.0, Jul 2023.
  Editors: Walden, Shortell, Roedler, Delicado, Mornas, Yip, Endler.

## Exam relevance map (ASEP)

The ASEP knowledge exam is drawn from this handbook, weighted toward the
**life-cycle processes in Chapter 2** (esp. §2.3.5 Technical Processes and
§2.3.4 Technical Management Processes). Prioritize:

- **§2.3.5 Technical Processes** (the 14 processes) — highest yield.
- **§2.3.4 Technical Management Processes** (the 8 processes).
- **§2.1–2.2** life-cycle terms, stages, decision gates, Vee/iterative models.
- **§1.3–1.4** systems concepts & SE foundations (definitions, emergence).
- **Appendix C: Terms & Definitions (p.311)** and **Appendix B: Acronyms
  (p.305)** — memorize verbatim; exam tests precise wording.
- **Appendix E: Input/Output Descriptions (p.321)** — process I/O is testable.

## Chapter 1 — Systems Engineering Introduction (p.1)

| § | Title | p. |
|---|-------|----|
| 1.1 | What Is Systems Engineering? | 1 |
| 1.2 | Why Is Systems Engineering Important? | 4 |
| 1.3 | Systems Concepts | 8 |
| 1.3.1 | System Boundary and the System of Interest (SoI) | 8 |
| 1.3.2 | Emergence | 9 |
| 1.3.3 | Interfacing, Interoperating, and Enabling Systems | 10 |
| 1.3.4 | System Innovation Ecosystem | 11 |
| 1.3.5 | The Hierarchy within a System | 12 |
| 1.3.6 | Systems States and Modes | 14 |
| 1.3.7 | Complexity | 15 |
| 1.4 | Systems Engineering Foundations | 15 |
| 1.4.1 | Uncertainty | 15 |
| 1.4.2 | Cognitive Bias | 17 |
| 1.4.3 | Systems Engineering Principles | 17 |
| 1.4.4 | Systems Engineering Heuristics | 20 |
| 1.5 | System Science and Systems Thinking | 21 |

## Chapter 2 — System Life Cycle Concepts, Models, and Processes (p.25)

| § | Title | p. |
|---|-------|----|
| 2.1 | Life Cycle Terms and Concepts | 25 |
| 2.1.1 | Life Cycle Characteristics | 25 |
| 2.1.2 | Typical Life Cycle Stages | 26 |
| 2.1.3 | Decision Gates | 29 |
| 2.1.4 | Technical Reviews and Audits | 31 |
| 2.2 | Life Cycle Model Approaches | 33 |
| 2.2.1 | Sequential Methods | 35 |
| 2.2.2 | Incremental Methods | 36 |
| 2.2.3 | Evolutionary Methods | 38 |
| 2.3 | System Life Cycle Processes | 39 |
| 2.3.1 | Introduction to the System Life Cycle Processes | 39 |
| 2.3.1.1 | Format and Conventions | 40 |
| 2.3.1.2 | Concurrency, Iteration, and Recursion | 42 |

### §2.3.2 Agreement Processes (p.44)
| § | Process | p. |
|---|---------|----|
| 2.3.2.1 | Acquisition | 45 |
| 2.3.2.2 | Supply | 48 |

### §2.3.3 Organizational Project-Enabling Processes (p.50)
| § | Process | p. |
|---|---------|----|
| 2.3.3.1 | Life Cycle Model Management | 51 |
| 2.3.3.2 | Infrastructure Management | 54 |
| 2.3.3.3 | Portfolio Management | 57 |
| 2.3.3.4 | Human Resource Management | 60 |
| 2.3.3.5 | Quality Management | 63 |
| 2.3.3.6 | Knowledge Management | 67 |

### §2.3.4 Technical Management Processes (p.70)
| § | Process | p. |
|---|---------|----|
| 2.3.4.1 | Project Planning | 70 |
| 2.3.4.2 | Project Assessment and Control | 75 |
| 2.3.4.3 | Decision Management | 78 |
| 2.3.4.4 | Risk Management | 81 |
| 2.3.4.5 | Configuration Management | 87 |
| 2.3.4.6 | Information Management | 91 |
| 2.3.4.7 | Measurement | 93 |
| 2.3.4.8 | Quality Assurance | 98 |

### §2.3.5 Technical Processes (p.101) — highest exam yield
| § | Process | p. |
|---|---------|----|
| 2.3.5.1 | Business or Mission Analysis | 103 |
| 2.3.5.2 | Stakeholder Needs and Requirements Definition | 107 |
| 2.3.5.3 | System Requirements Definition | 112 |
| 2.3.5.4 | System Architecture Definition | 118 |
| 2.3.5.5 | Design Definition | 124 |
| 2.3.5.6 | System Analysis | 129 |
| 2.3.5.7 | Implementation | 132 |
| 2.3.5.8 | Integration | 134 |
| 2.3.5.9 | Verification | 138 |
| 2.3.5.10 | Transition | 143 |
| 2.3.5.11 | Validation | 146 |
| 2.3.5.12 | Operation | 152 |
| 2.3.5.13 | Maintenance | 154 |
| 2.3.5.14 | Disposal | 156 |

## Chapter 3 — Life Cycle Analyses and Methods (p.159)

| § | Title | p. |
|---|-------|----|
| 3.1 | Quality Characteristics and Approaches | 159 |
| 3.1.2 | Affordability Analysis | 160 |
| 3.1.3 | Agility Engineering | 165 |
| 3.1.4 | Human Systems Integration | 168 |
| 3.1.5 | Interoperability Analysis | 171 |
| 3.1.6 | Logistics Engineering | 172 |
| 3.1.7 | Manufacturability/Producibility Analysis | 175 |
| 3.1.8 | Reliability, Availability, Maintainability Engineering | 176 |
| 3.1.9 | Resilience Engineering | 180 |
| 3.1.10 | Sustainability Engineering | 184 |
| 3.1.11 | System Safety Engineering | 185 |
| 3.1.12 | System Security Engineering | 190 |
| 3.1.13 | Loss-Driven Systems Engineering | 191 |
| 3.2 | Systems Engineering Analyses and Methods | 192 |
| 3.2.1 | Modeling, Analysis, and Simulation | 192 |
| 3.2.2 | Prototyping | 200 |
| 3.2.3 | Traceability | 201 |
| 3.2.4 | Interface Management | 202 |
| 3.2.5 | Architecture Frameworks | 206 |
| 3.2.6 | Patterns | 208 |
| 3.2.7 | Design Thinking | 212 |
| 3.2.8 | Biomimicry | 213 |

## Chapter 4 — Tailoring and Application Considerations (p.215)

| § | Title | p. |
|---|-------|----|
| 4.1 | Tailoring Considerations | 215 |
| 4.2 | SE Methodology/Approach Considerations | 219 |
| 4.2.1 | Model-Based SE (MBSE) | 219 |
| 4.2.2 | Agile Systems Engineering | 221 |
| 4.2.3 | Lean Systems Engineering | 224 |
| 4.2.4 | Product Line Engineering (PLE) | 226 |
| 4.3 | System Types Considerations | 229 |
| 4.3.1–4.3.9 | Greenfield, Brownfield, COTS, Software-Intensive, Cyber-Physical, SoS, IoT/Big-Data, Service, Enterprise | 229–241 |
| 4.4 | Application by Product Sector/Domain | 244 |
| 4.4.1–4.4.10 | Automotive, Biomedical/Healthcare, Aerospace, Defense, Infrastructure, Oil & Gas, Power & Energy, Space, Telecom, Transportation | 245–258 |

## Chapter 5 — Systems Engineering in Practice (p.261)

| § | Title | p. |
|---|-------|----|
| 5.1 | Systems Engineering Competencies | 261 |
| 5.1.4 | Ethics | 264 |
| 5.2 | Diversity, Equity, and Inclusion | 265 |
| 5.3 | SE Relationships to Other Disciplines (SWE, HWE, PM, IE, OR) | 266 |
| 5.4 | Digital Engineering | 273 |
| 5.5 | Systems Engineering Transformation | 274 |
| 5.6 | Future of SE | 275 |

## Chapter 6 — Case Studies (p.277)

| § | Case | p. |
|---|------|----|
| 6.1 | Therac-25 (radiation therapy) | 277 |
| 6.2 | Øresund Bridge | 278 |
| 6.3 | Stuxnet (cyber-physical security) | 280 |
| 6.4 | Incubators (design for maintainability) | 282 |
| 6.5 | Autonomous Vehicles (AI in SE) | 283 |

## Appendices

| App. | Title | p. |
|------|-------|----|
| A | References | 287 |
| B | Acronyms | 305 |
| C | Terms and Definitions | 311 |
| D | N2 Diagram of SE Processes | 317 |
| E | Input/Output Descriptions | 321 |
| F | Acknowledgments | 335 |
| G | Comment Form | 337 |
| — | Index | 339 |

## Concept → section quick lookup

- **Definition of SE** → §1.1 (p.1)
- **Emergence** → §1.3.2 (p.9)
- **System of Interest (SoI), system boundary** → §1.3.1 (p.8)
- **Enabling systems** → §1.3.3 (p.10)
- **States and modes** → §1.3.6 (p.14)
- **SE principles / heuristics** → §1.4.3–1.4.4 (p.17, 20)
- **Systems thinking** → §1.5 (p.21)
- **Life cycle stages** → §2.1.2 (p.26)
- **Decision gates / milestones** → §2.1.3 (p.29)
- **Technical reviews & audits (SRR, PDR, CDR, etc.)** → §2.1.4 (p.31)
- **Vee model** → §2.2 + Fig 2.6 (p.33+)
- **Incremental / evolutionary / iterative / spiral / DevSecOps** → §2.2.1–2.2.3 (p.35–38)
- **Concurrency, iteration, recursion** → §2.3.1.2 (p.42)
- **The 4 process groups** → §2.3.2 Agreement, §2.3.3 Org Project-Enabling, §2.3.4 Technical Management, §2.3.5 Technical
- **Requirements (stakeholder vs system)** → §2.3.5.2 (p.107) vs §2.3.5.3 (p.112)
- **Architecture vs Design** → §2.3.5.4 (p.118) vs §2.3.5.5 (p.124)
- **Verification vs Validation** → §2.3.5.9 (p.138) vs §2.3.5.11 (p.146)
- **Integration** → §2.3.5.8 (p.134)
- **Transition** → §2.3.5.10 (p.143)
- **Risk management** → §2.3.4.4 (p.81)
- **Configuration management** → §2.3.4.5 (p.87)
- **Measurement / metrics** → §2.3.4.7 (p.93)
- **Decision management / trade studies** → §2.3.4.3 (p.78)
- **RAM (reliability/availability/maintainability)** → §3.1.8 (p.176)
- **Resilience** → §3.1.9 (p.180); **Safety** → §3.1.11 (p.185); **Security** → §3.1.12 (p.190)
- **Traceability** → §3.2.3 (p.201); **Interface management** → §3.2.4 (p.202)
- **Architecture frameworks** → §3.2.5 (p.206)
- **Tailoring** → §4.1 (p.215)
- **MBSE** → §4.2.1 (p.219); **Agile SE** → §4.2.2 (p.221); **Lean** → §4.2.3 (p.224)
- **Systems of Systems (SoS)** → §4.3.6 (p.235)
- **Digital Engineering** → §5.4 (p.273)
- **Ethics** → §5.1.4 (p.264)
- **Terms & definitions (verbatim)** → Appendix C (p.311)
- **Acronyms** → Appendix B (p.305)
- **Process inputs/outputs** → Appendix E (p.321)
</content>
</invoke>
