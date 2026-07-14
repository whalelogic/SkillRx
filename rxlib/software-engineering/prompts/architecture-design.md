# Architecture Design Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Software Engineering |
| **Task** | System architecture design or evaluation |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `software-engineering/system-instructions/senior-software-engineer.md` |
| **Output Format** | Markdown with Mermaid diagrams |

---

## Prompt

{{DESIGN_OR_EVALUATE}}: Please {{design a new / evaluate the existing}} architecture for the following system.

### System Overview

* **System name:** {{SYSTEM_NAME}}
* **Purpose:** {{PURPOSE}}
* **Stage:** {{STAGE}} (greenfield / brownfield / migration)

### Functional Requirements

{{FUNCTIONAL_REQUIREMENTS}}

### Non-Functional Requirements

| Attribute | Target |
|---|---|
| Availability | {{AVAILABILITY_TARGET}} (e.g. 99.9% monthly) |
| Latency (p99) | {{LATENCY_TARGET}} (e.g. < 200 ms) |
| Throughput | {{THROUGHPUT_TARGET}} (e.g. 10,000 RPS peak) |
| Data volume | {{DATA_VOLUME}} |
| Compliance | {{COMPLIANCE}} (e.g. GDPR, HIPAA, PCI-DSS, or "none") |

### Constraints

* **Budget / cost ceiling:** {{BUDGET}}
* **Team size / expertise:** {{TEAM}}
* **Existing technology stack:** {{EXISTING_STACK}}
* **Deployment target:** {{DEPLOYMENT_TARGET}} (e.g. Google Cloud, on-premise, hybrid)

### Task

1. Propose or evaluate a high-level architecture that satisfies the requirements.
2. Describe the major components and their responsibilities.
3. Illustrate the architecture with a Mermaid diagram.
4. Identify the top 3 architectural risks and mitigation strategies.
5. Recommend the technology choices for each component with brief justifications.
6. Describe how the architecture meets each non-functional requirement.

### Output Format

```
## Architecture Overview
[2–3 paragraph summary]

## Component Diagram
```mermaid
...
```

## Component Descriptions
[Table: Component | Responsibility | Technology Choice | Justification]

## Request Flow
[Mermaid sequence diagram for the primary user journey]

## Architectural Risks
[Numbered list: Risk | Impact | Likelihood | Mitigation]

## NFR Mapping
[Table: NFR | How It Is Met]
```
