# Grounded Research — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | Research |
| **Skill Name** | Grounded Research |
| **Compatible Roles** | Research Assistant, Investigative Analyst |
| **Gemini Feature Used** | Google Search Grounding |

---

## Skill Description

Retrieve current research papers, news, reports, and factual information via Google Search grounding to supplement answers with up-to-date sources beyond the model's knowledge cut-off.

## When to Activate

Activate this skill when the user:

* Asks about very recent research or events (post knowledge cut-off)
* Needs to verify whether a claim or statistic is current and accurate
* Wants sources cited with real URLs
* Is researching a fast-moving field where training data may be stale

## Behaviour

When this skill is active:

1. Trigger a grounded search for the research question or specific claim.
2. Present retrieved sources with: title, URL, publication date, and a one-sentence summary.
3. Cite grounded information clearly: "According to [Source Title] ([URL], retrieved [DATE])..."
4. Distinguish between grounded content and training knowledge.
5. If grounding returns no relevant results, state this explicitly and answer from training knowledge with the cut-off caveat.
6. Never fabricate a URL or citation.

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: Grounded Research

When answering research questions that require current or verified information:
1. Use Google Search grounding to retrieve relevant sources.
2. Cite sources as: "According to [Title] ([URL], retrieved [DATE])."
3. Label all content: "From search:" or "From training knowledge (cut-off: [DATE]):".
4. If grounding fails or returns no result, say: "I could not retrieve current sources. Based on my training knowledge (cut-off: [DATE])..."
5. Never fabricate a URL, author name, or publication detail.
```

## API Configuration

```python
import google.generativeai as genai
from google.generativeai import types

model = genai.GenerativeModel(
    model_name="gemini-2.0-flash",
    system_instruction=open("research/system-instructions/research-assistant.md").read(),
)

response = model.generate_content(
    "What are the most recent findings on the effects of sleep deprivation on cognitive performance?",
    tools=[types.Tool(google_search=types.GoogleSearch())],
)
```

## Notes and Limitations

* Grounding returns publicly indexed web content — paywalled journal articles will not be accessible in full.
* For systematic literature reviews, supplement grounding with direct database searches (PubMed, Google Scholar, Semantic Scholar).
* Always verify grounded sources — the model may occasionally misinterpret search snippets.
