# Long Context Repository Analysis — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | Repository Analysis |
| **Skill Name** | Long Context Repository Analysis |
| **Compatible Roles** | Codebase Analyzer, Documentation Generator |
| **Gemini Feature Used** | Long context window (up to 1M tokens on gemini-2.5-pro) |

---

## Skill Description

Use Gemini's million-token context window to load entire repositories — or large portions of them — into a single prompt, enabling holistic codebase understanding without chunking or multiple passes.

## When to Activate

Activate this skill when:

* The codebase is too large to analyse manually file-by-file
* The user wants cross-file analysis (e.g. tracing a call chain across 10+ files)
* The task requires understanding the relationship between many modules simultaneously
* Generating comprehensive documentation that covers the whole codebase

## How to Prepare the Repository for Analysis

### Method 1: Flatten with repomix

```bash
# Install repomix
npm install -g repomix

# Generate a flat text file of the entire repo
repomix --output repo_context.txt

# Or exclude certain directories
repomix --ignore "node_modules,dist,.git" --output repo_context.txt
```

Then paste the contents of `repo_context.txt` into your prompt.

### Method 2: Manual selection with cat

```bash
# Combine key files
cat README.md package.json src/**/*.ts > repo_context.txt
```

### Method 3: Files API (for Gemini API users)

```python
import google.generativeai as genai

# Upload the repository file
repo_file = genai.upload_file("repo_context.txt", mime_type="text/plain")

model = genai.GenerativeModel("gemini-2.5-pro")

response = model.generate_content([
    repo_file,
    "Provide a comprehensive architectural overview of this codebase."
])
```

## Behaviour

When this skill is active:

1. Accept a large text block containing the full repository content.
2. Build a mental model of the complete codebase before answering specific questions.
3. Provide cross-file analysis — trace data flows, dependencies, and patterns across the entire codebase.
4. Reference specific file paths and line numbers in findings.
5. Explicitly note when a file was not included and how that affects the analysis.

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: Long Context Repository Analysis

When the user provides a full repository dump or large codebase:
1. Read the entire content before responding — do not answer based on partial reading.
2. Build a cross-file understanding: identify module boundaries, shared utilities, and architectural patterns.
3. Reference specific file paths in all findings (e.g. "In `src/auth/middleware.py`, line 42...").
4. Flag files that appear to be missing from the provided context.
5. When the codebase is very large, acknowledge any limitations in the analysis.
```

## Notes and Limitations

* `gemini-2.5-pro` supports up to 1 million tokens (approximately 750,000 words or ~3,000 code files).
* Very large repositories may still exceed context limits — use selective file inclusion to prioritise core logic.
* Latency increases significantly with large context sizes — expect 30–120 seconds for 500K+ token inputs.
* Costs are proportional to input tokens — consider the economic trade-off for large-scale analyses.
