# Research Question Refinement Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Research |
| **Task** | Refine a broad topic into a well-scoped research question |
| **Recommended Model** | gemini-2.0-flash |
| **Recommended System Instruction** | `research/system-instructions/academic-researcher.md` |
| **Output Format** | Structured Markdown |

---

## Prompt

I want to conduct research on the following broad topic. Please help me refine it into a precise, answerable research question.

### Broad Topic

{{BROAD_TOPIC}}

### Context

* **Field / discipline:** {{FIELD}} (e.g. computer science, sociology, economics, medicine)
* **Research purpose:** {{PURPOSE}} (e.g. thesis, journal article, policy brief, industry report)
* **Available resources:** {{RESOURCES}} (e.g. "I have access to X dataset", "I can conduct interviews", "literature review only")
* **Time horizon:** {{TIME_HORIZON}} (e.g. 3 months, 1 year)
* **Geographic or demographic scope:** {{SCOPE}} (e.g. US adults, global, EU policy only)

### Refinement Task

1. **Gap analysis** — briefly describe the current state of research on this topic and identify where knowledge gaps exist.
2. **Scoped question options** — propose 3 distinct, well-scoped research questions derived from the broad topic. Each should be:
   - Specific (not "why is X important?" but "what is the relationship between X and Y in context Z?")
   - Researchable with the available resources
   - Significant — the answer should matter to the field
3. **Question evaluation** — for each proposed question, note: feasibility, scope, potential contribution to knowledge.
4. **Recommended question** — recommend the strongest question with a brief justification.
5. **Hypothesis (if applicable)** — if the recommended question is empirical, propose a falsifiable null and alternative hypothesis.
6. **Key search terms** — provide 5–10 Boolean search strings for Google Scholar/PubMed to find relevant prior literature.

### Output Format

```
## Current State of Research
...

## Proposed Research Questions
### Option 1: [Question]
- Feasibility: ...
- Scope: ...
- Contribution: ...

### Option 2: [Question]
...

### Option 3: [Question]
...

## Recommended Question
...

## Hypothesis (if applicable)
- H0: ...
- H1: ...

## Key Search Terms
1. ...
```
