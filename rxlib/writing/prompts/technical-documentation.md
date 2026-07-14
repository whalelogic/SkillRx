# Technical Documentation Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Writing |
| **Task** | Write or improve technical documentation |
| **Recommended Model** | gemini-2.0-flash |
| **Recommended System Instruction** | `writing/system-instructions/technical-writer.md` |
| **Output Format** | Markdown |

---

## Prompt

Please {{WRITE_OR_IMPROVE}} the following technical documentation.

### Documentation Brief

* **Documentation type:** {{DOC_TYPE}} (quickstart guide / API reference / user manual / tutorial / runbook / architecture doc / README / release notes)
* **Subject:** {{SUBJECT}} (e.g. "the authentication module of our REST API", "deploying the application to Kubernetes")
* **Target audience:** {{AUDIENCE}} (e.g. "developers with basic Python knowledge", "non-technical end users", "DevOps engineers")
* **Style guide:** {{STYLE_GUIDE}} (Google / Microsoft / Write the Docs / internal — describe key rules if internal)

### Existing Content (for improvement tasks)

```
{{EXISTING_CONTENT}}
```

### Technical Specifications / Context

```
{{TECH_CONTEXT}}
```

(Paste API specs, code snippets, architecture details, or feature descriptions the documentation should cover.)

### Task

1. Structure the documentation according to the type selected above.
2. Write task-oriented content — lead with what the user needs to do, not a product feature description.
3. Use imperative mood for steps: "Run the following command:" not "The user should run".
4. Include prerequisites, steps, expected outcomes, and troubleshooting tips.
5. Add code examples for any technical procedure.
6. If improving existing content, produce a change summary after the revised document.

### Output Requirements

* Use the Diátaxis documentation framework: separate **tutorials** (learning-oriented) from **how-to guides** (task-oriented) from **reference** (information-oriented) from **explanation** (understanding-oriented).
* Flag any gaps where you need more information to write accurately.
* Do not add marketing language or vague promises.
