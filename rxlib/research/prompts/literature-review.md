# Literature Review Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Research |
| **Task** | Structured literature review synthesis |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `research/system-instructions/research-assistant.md` |
| **Output Format** | Structured Markdown |

---

## Prompt

Please write a structured literature review on the following topic, synthesising the provided sources and your knowledge of the field.

### Topic

**{{TOPIC}}**

### Research Question (optional)

{{RESEARCH_QUESTION}}

### Sources Provided

Paste any sources you want included (abstracts, full texts, citations):

```
{{SOURCES}}
```

### Literature Review Parameters

* **Scope:** {{SCOPE}} (e.g. "empirical studies published 2015–2025", "theoretical and empirical", "English language only")
* **Citation style:** {{CITATION_STYLE}} (APA 7th / MLA 9th / Chicago / Harvard / Vancouver)
* **Target length:** {{TARGET_LENGTH}} (e.g. 800–1200 words, 2000–3000 words)
* **Audience:** {{AUDIENCE}} (e.g. undergraduate, graduate, expert practitioner)

### Task

1. Organise the literature thematically (not chronologically or source-by-source).
2. For each theme, synthesise what multiple sources say — highlight agreements, contradictions, and gaps.
3. Critically evaluate the methodological quality of the research where relevant.
4. Identify the key debates, unresolved questions, or under-researched areas.
5. Write a conclusion that summarises the state of knowledge and points to future research directions.

### Output Requirements

* Formal academic register.
* Every factual claim attributed with an in-text citation.
* No quotations longer than two sentences — paraphrase instead.
* A complete reference list at the end in the specified citation style.
* Clearly flag any claims drawn from training knowledge rather than provided sources.
