# ⚕️pskillRx ⚕️

A structured, ever-expanding library of **AI Agent prompts**, **skills**, **system instructions**, and **configuration templates** optimised for use with Google Gemini (and compatible large language models).

---

## What's Inside

| Domain | Description |
|---|---|
| [`software-engineering/`](software-engineering/) | Code review, debugging, architecture, DevOps, security |
| [`financial-markets/`](financial-markets/) | Market analysis, portfolio review, risk, quant strategies |
| [`research/`](research/) | Academic research, literature review, hypothesis generation |
| [`writing/`](writing/) | Content, technical, creative writing and editing |
| [`ui-ux-development/`](ui-ux-development/) | Design review, component generation, accessibility |
| [`repository-analysis/`](repository-analysis/) | Codebase analysis, dependency audits, changelog generation |
| [`data-science/`](data-science/) | EDA, ML pipeline design, model evaluation, visualisation |
| [`productivity/`](productivity/) | Task management, meeting summaries, planning, GTD |
| [`templates/`](templates/) | Blank scaffolding for prompts, skills, system instructions |

---

## How to Use This Library

### Quick Start with Gemini

1. **Pick a domain** folder that matches your use case.
2. **Open a `system-instructions/` file** — paste the full content into the *System Instructions* field of your Gemini agent or Gemini API request.
3. **Browse `prompts/`** for ready-to-use user-turn messages.
4. **Combine `skills/`** entries to extend an agent's capabilities beyond its base system instruction.

### Gemini API Example

```python
import google.generativeai as genai

# Load a system instruction
with open("software-engineering/system-instructions/senior-software-engineer.md") as f:
    system_instruction = f.read()

model = genai.GenerativeModel(
    model_name="gemini-2.0-flash",
    system_instruction=system_instruction,
)

# Load a prompt template and fill it in
with open("software-engineering/prompts/code-review.md") as f:
    prompt_template = f.read()

response = model.generate_content(
    prompt_template.replace("{{CODE}}", my_code_string)
)
print(response.text)
```

### Google AI Studio

1. Open [Google AI Studio](https://aistudio.google.com).
2. Create a new prompt or agent.
3. Copy the contents of any `system-instructions/*.md` file into the **System Instructions** box.
4. Use the corresponding `prompts/*.md` files as starter messages.

---

## Folder Conventions

Every domain folder follows the same layout:

```
<domain>/
├── README.md                  # Domain overview and usage guide
├── system-instructions/       # Full system prompts that define an agent's persona and behaviour
│   └── *.md
├── prompts/                   # Standalone user-turn prompts for specific tasks
│   └── *.md
└── skills/                    # Modular capability blocks that can be mixed into any agent
    └── *.md
```

### Template Variables

Prompts and skills use `{{VARIABLE_NAME}}` placeholders. Replace them with real values before sending to the model. Example:

```
Analyse the following {{LANGUAGE}} code for security vulnerabilities:

{{CODE}}
```

Becomes:

```
Analyse the following Python code for security vulnerabilities:

def login(user, pwd):
    ...
```

---

## Contributing

1. Copy the appropriate blank template from [`templates/`](templates/).
2. Fill in every section — do not leave placeholder text.
3. Place the file in the correct domain subfolder.
4. Update the domain `README.md` with a one-line description of your addition.
5. Open a pull request with a short summary of what you added and why it is useful.

---

## Licence

[MIT](LICENSE)

