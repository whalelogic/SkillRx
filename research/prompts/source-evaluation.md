# Source Evaluation Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Research |
| **Task** | Evaluate the credibility and relevance of a source |
| **Recommended Model** | gemini-2.0-flash |
| **Recommended System Instruction** | `research/system-instructions/research-assistant.md` |
| **Output Format** | Structured Markdown |

---

## Prompt

Please evaluate the following source for credibility, quality, and relevance to my research.

### Source

```
{{SOURCE_CONTENT_OR_CITATION}}
```

Provide the full citation, abstract, URL, or paste the text of the source.

### Research Context

* **My research topic:** {{RESEARCH_TOPIC}}
* **My research question:** {{RESEARCH_QUESTION}}
* **Source type:** {{SOURCE_TYPE}} (academic journal article / book / news article / government report / blog / preprint / other)

### Evaluation Task

Evaluate this source using the **SIFT** method (Stop, Investigate the source, Find better coverage, Trace claims) and the **CRAAP** criteria (Currency, Relevance, Authority, Accuracy, Purpose):

1. **Currency:** When was it published? Is it recent enough for the topic? Have key claims been superseded?
2. **Relevance:** How directly does it address the research question? Is it appropriate for the academic level?
3. **Authority:** Who are the authors? What are their credentials? Is the publisher/journal reputable? Is it peer-reviewed?
4. **Accuracy:** Is the methodology described? Are claims supported by evidence? Are there citations? Can claims be verified?
5. **Purpose:** Why was this written? Is there a detectable bias, conflict of interest, or commercial motive?
6. **Overall verdict:** Recommended for use / Use with caveats (describe) / Not recommended (explain why)

### Output Format

```
## Source Evaluation

**Citation:** ...
**Source type:** ...

### Currency
...

### Relevance
Score: High / Medium / Low
Reason: ...

### Authority
...

### Accuracy
...

### Purpose / Bias
...

### Overall Verdict
**Use:** Recommended / Use with caveats / Not recommended
**Summary:** [2–3 sentences on why and how to use this source]
```
