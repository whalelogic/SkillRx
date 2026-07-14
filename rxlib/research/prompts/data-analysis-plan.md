# Data Analysis Plan Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Research |
| **Task** | Design a data collection and analysis plan |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `research/system-instructions/academic-researcher.md` |
| **Output Format** | Structured Markdown |

---

## Prompt

Please help me design a rigorous data collection and analysis plan for the following research question.

### Research Question

{{RESEARCH_QUESTION}}

### Context

* **Field:** {{FIELD}}
* **Study type:** {{STUDY_TYPE}} (e.g. survey, experiment, secondary data analysis, qualitative interviews)
* **Available resources:** {{RESOURCES}} (e.g. access to a specific dataset, ability to recruit N participants, budget)
* **Ethical constraints:** {{ETHICS}} (e.g. IRB approval required, sensitive population, anonymisation needed)
* **Target output:** {{OUTPUT}} (e.g. peer-reviewed paper, thesis chapter, internal report)

### Plan Design Task

1. **Research design** — recommend the most appropriate study design given the question and resources. Justify why.
2. **Sampling strategy** — who or what will be sampled? What is the target sample size? How will participants be recruited/data sourced? Include a power analysis rationale for quantitative studies.
3. **Data collection instruments** — what instruments, surveys, interview guides, or data extraction forms are needed? Provide a draft structure.
4. **Variables and measures** — list all key variables with their operationalisation and measurement scale (nominal, ordinal, interval, ratio).
5. **Analysis approach** — specify the statistical or qualitative analysis methods. For quantitative studies, state the primary statistical test and its assumptions.
6. **Validity and reliability** — how will internal and external validity be maximised? What are the main threats?
7. **Ethical considerations** — what ethical issues arise and how will they be addressed?
8. **Limitations** — what are the inherent limitations of this design?
9. **Timeline** — provide a high-level Gantt-style table with major milestones.
