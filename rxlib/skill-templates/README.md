# Templates

Blank scaffolding files for every resource type in this library.  
Copy the appropriate file, fill in every `{{PLACEHOLDER}}` section, and place the result in the correct domain folder.

| File | Use When |
|---|---|
| [`system-instruction-template.md`](system-instruction-template.md) | Defining a new agent persona or role |
| [`prompt-template.md`](prompt-template.md) | Creating a reusable user-turn task prompt |
| [`skill-template.md`](skill-template.md) | Describing a modular capability block |
| [`agent-config-template.md`](agent-config-template.md) | Specifying full agent configuration (model, tools, safety) |

## Naming Conventions

* Use `kebab-case` for all file names.
* System instructions describe **roles**: `senior-software-engineer.md`, `quantitative-analyst.md`.
* Prompts describe **tasks**: `code-review.md`, `market-analysis.md`.
* Skills describe **capabilities**: `tool-use.md`, `chain-of-thought.md`.
