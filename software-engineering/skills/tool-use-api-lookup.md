# Tool Use: API Lookup — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | Software Engineering |
| **Skill Name** | Tool Use: API Lookup |
| **Compatible Roles** | All software engineering roles |
| **Gemini Feature Used** | Google Search Grounding / Function Calling |

---

## Skill Description

Query up-to-date library documentation, API references, and package release notes via Google Search grounding or a custom function-calling tool, instead of relying solely on training data.

## When to Activate

Activate this skill when the user:

* Asks about the current API of a library (version-sensitive questions)
* References a library or framework released after the model's knowledge cut-off
* Asks about the latest version, changelog, or deprecation status of a package
* Needs a confirmed method signature, not a recalled one

## Behaviour

When this skill is active:

1. Identify whether the question is version-sensitive or time-sensitive.
2. If yes, trigger a grounded search before answering.
3. Cite the source URL and the retrieved documentation version in your answer.
4. Clearly distinguish between information from search results and information from training data.
5. If search returns no relevant result, state that you are answering from training data and note the knowledge cut-off date.

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: API Lookup via Grounding

When answering questions about specific library APIs, method signatures, or package versions:
1. Use Google Search grounding to retrieve up-to-date documentation.
2. Cite the source URL and document version in your answer.
3. Clearly label answers drawn from search: "According to [Source, retrieved DATE]".
4. If grounding is unavailable, state: "I'm answering from training data (cut-off: [DATE]). Verify against the official documentation."
```

## API Configuration

Enable grounding in the Gemini API:

```python
import google.generativeai as genai
from google.generativeai import types

model = genai.GenerativeModel(model_name="gemini-2.0-flash")

response = model.generate_content(
    "What is the signature of sklearn.preprocessing.StandardScaler.fit_transform in the latest version?",
    tools=[types.Tool(google_search=types.GoogleSearch())],
)
```

## Notes and Limitations

* Google Search grounding is available on `gemini-2.0-flash` and `gemini-2.5-pro`.
* Grounding adds latency (~1–3 s). Only activate for version-sensitive or time-sensitive questions.
* The model may still produce inaccurate results even with grounding — always verify critical API usage against official documentation.
