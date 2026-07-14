# Hypothesis Generation Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Research |
| **Task** | Generate testable hypotheses from observations or prior research |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `research/system-instructions/academic-researcher.md` |
| **Output Format** | Structured Markdown |

---

## Prompt

Based on the following observations or prior research findings, generate a set of testable hypotheses.

### Observations / Prior Findings

```
{{OBSERVATIONS_OR_FINDINGS}}
```

### Field / Domain

{{FIELD}}

### Research Paradigm

* **Approach:** {{APPROACH}} (quantitative / qualitative / mixed methods)
* **Causal inference goal:** {{CAUSAL_GOAL}} (yes — establish causation / no — describe associations only)
* **Study type envisaged:** {{STUDY_TYPE}} (e.g. RCT, observational cohort, survey, experiment, computational model)

### Hypothesis Generation Task

1. **Theory identification** — which existing theoretical frameworks from {{FIELD}} are most relevant to the observations?
2. **Generate hypotheses** — propose 3–5 falsifiable hypotheses. Each must include:
   - H0 (null hypothesis)
   - H1 (alternative hypothesis)
   - Directional prediction (if applicable)
   - The key variables (independent, dependent, moderating, mediating)
3. **Operationalisation** — for each hypothesis, suggest how the key variables could be measured or operationalised.
4. **Ranking** — rank the hypotheses by: theoretical grounding, novelty, and testability.
5. **Confounds** — for each top-ranked hypothesis, list the main confounding variables that would need to be controlled.

### Output Format

```
## Relevant Theoretical Frameworks
...

## Hypotheses

### Hypothesis 1: [Short title]
- H0: ...
- H1: ...
- Key variables: IV = ..., DV = ..., Moderator = ... (if applicable)
- Operationalisation: ...
- Confounds: ...

### Hypothesis 2: [Short title]
...

## Priority Ranking
| Rank | Hypothesis | Theoretical Grounding | Novelty | Testability |
|---|---|---|---|---|
```
