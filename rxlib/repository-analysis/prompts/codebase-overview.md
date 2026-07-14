# Codebase Overview Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Repository Analysis |
| **Task** | Generate a structured overview of an unfamiliar codebase |
| **Recommended Model** | gemini-2.5-pro (1M+ context window) |
| **Recommended System Instruction** | `repository-analysis/system-instructions/codebase-analyzer.md` |
| **Output Format** | Structured Markdown with Mermaid diagrams |

---

## Prompt

Please analyse the following codebase and produce a comprehensive overview for a new engineer joining the project.

### Repository Contents

Paste the relevant code below. For large repositories, include:
1. The directory tree (output of `find . -type f | head -100` or `tree -L 3`)
2. Key entry point files (main.py, index.ts, cmd/main.go, etc.)
3. Configuration files (package.json, pyproject.toml, go.mod, Dockerfile, etc.)
4. Core domain/business logic files
5. README if one exists

```
{{REPOSITORY_CONTENTS}}
```

### Context

* **Project purpose:** {{PURPOSE}} (if known; leave blank to have the model infer it)
* **Primary language(s):** {{LANGUAGES}}
* **Target audience for this overview:** {{AUDIENCE}} (e.g. "new backend engineer", "technical manager", "external contributor")

### Analysis Task

1. **Purpose and business domain** — what does this project do and what problem does it solve?
2. **Architecture** — what architectural style is used? Produce a module/layer diagram.
3. **Technology stack** — list all significant technologies, frameworks, and libraries with their roles.
4. **Directory structure** — explain the purpose of each top-level directory.
5. **Key entry points** — how does the application start? What are the main API endpoints / CLI commands / event handlers?
6. **Primary data flow** — trace the lifecycle of the most common operation from request to response.
7. **Testing strategy** — what tests exist? Where are they? What is their apparent coverage approach?
8. **Configuration and environment** — how is the application configured? What environment variables are required?
9. **Known quality issues** — flag any obvious code smells, anti-patterns, or areas of concern.
10. **Onboarding quick-start** — what are the minimal steps to run the project locally?

### Output Format

Produce a developer onboarding guide with all sections above, including at least one Mermaid diagram.
