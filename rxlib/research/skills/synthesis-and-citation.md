# Synthesis and Citation — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | Research |
| **Skill Name** | Synthesis and Citation |
| **Compatible Roles** | Research Assistant, Academic Researcher |
| **Gemini Feature Used** | Standard generation |

---

## Skill Description

Synthesise information from multiple sources into a coherent, well-cited narrative — moving beyond source-by-source summary to thematic integration with properly formatted citations.

## When to Activate

Activate this skill when the user:

* Provides multiple sources and asks for a unified summary or literature review
* Wants to identify agreements and contradictions across sources
* Needs to produce academic writing with in-text citations and a reference list
* Asks for an annotated bibliography

## Behaviour

When this skill is active:

1. **Read all sources** before writing — identify the overarching themes.
2. **Group sources thematically** — do not discuss sources one by one; weave them together.
3. **Signal synthesis explicitly:**
   - Agreement: "Consistent with [Author A, Year], [Author B, Year] also found that..."
   - Contradiction: "In contrast, [Author C, Year] argues that... possibly because..."
   - Gap: "Notably absent from the literature is..."
4. **Cite accurately:** use the requested format for every factual claim. If a citation cannot be verified, flag it.
5. **Build a reference list** at the end with full citations in the requested style.

## Citation Format Quick Reference

### APA 7th — In-text
Single author: (Smith, 2021)  
Two authors: (Smith & Jones, 2021)  
Three+ authors: (Smith et al., 2021)  
Direct quote: (Smith, 2021, p. 45)

### APA 7th — Reference List (Journal Article)
`Author, A. A., & Author, B. B. (Year). Title of article. *Journal Name*, *Volume*(Issue), pages. https://doi.org/xxxxx`

### MLA 9th — In-text
(Smith 45)

### Chicago 17th Author-Date — In-text
(Smith 2021, 45)

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: Synthesis and Citation

When writing research content based on multiple sources:
1. Organise content thematically, not source-by-source.
2. Explicitly signal synthesis using phrases like "Similarly...", "In contrast...", "Building on X, Y argues..."
3. Cite every factual claim with an in-text citation in the user's requested format.
4. End every research output with a complete reference list.
5. If a source cannot be verified, flag it: "[UNVERIFIED — recommend checking original source]".
```

## Notes and Limitations

* Always check that citations are real — the model can sometimes confuse author names, years, or journal titles. When accuracy is critical, verify against the original source.
* For large volumes of sources, instruct the user to provide the citations explicitly rather than relying on the model to recall them.
