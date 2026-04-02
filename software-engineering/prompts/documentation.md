# Documentation Generation Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Software Engineering |
| **Task** | Generate or improve code documentation |
| **Recommended Model** | gemini-2.0-flash |
| **Recommended System Instruction** | `software-engineering/system-instructions/senior-software-engineer.md` |
| **Output Format** | Markdown or inline docstrings |

---

## Prompt

Generate {{DOC_TYPE}} documentation for the following `{{LANGUAGE}}` code.

### Documentation Type

Select one:
- [ ] **Inline docstrings / JSDoc** — add doc comments to every public function, class, and method
- [ ] **README** — project overview, installation, usage, configuration, contributing
- [ ] **API reference** — OpenAPI/Swagger YAML or a Markdown reference table
- [ ] **Architecture doc** — describe how the system is structured and why
- [ ] **Runbook** — step-by-step operational guide

### Documentation Style Guide

* **Format:** {{FORMAT}} (e.g. Google Python docstrings, NumPy docstrings, JSDoc, TSDoc)
* **Target audience:** {{AUDIENCE}} (e.g. internal engineers, external API consumers, ops team)
* **Tone:** {{TONE}} (e.g. technical, beginner-friendly)
* **Include examples:** {{INCLUDE_EXAMPLES}} (yes / no)

### Code

```{{LANGUAGE}}
{{CODE}}
```

### Task

1. Generate complete, accurate documentation for every public symbol.
2. Describe parameters (name, type, description), return values, and exceptions raised.
3. Include at least one usage example per public function if examples are requested.
4. Flag any code that is undocumentable without more context and explain what additional information is needed.

### Output Requirements

* Return the fully documented code (not just the doc comments in isolation).
* Do not alter any logic — documentation only.
* Use the specified docstring format consistently throughout.

---

## Example Usage

**Language:** Python  
**Doc type:** Inline docstrings (Google style)  
**Original code:**
```python
def parse_date(date_str, fmt="%Y-%m-%d"):
    from datetime import datetime
    return datetime.strptime(date_str, fmt)
```

**Expected output:**
```python
def parse_date(date_str: str, fmt: str = "%Y-%m-%d") -> datetime:
    """Parse a date string into a datetime object.

    Args:
        date_str: The date string to parse (e.g. "2024-01-15").
        fmt: The expected format of the date string. Defaults to ISO 8601 ("%Y-%m-%d").

    Returns:
        A ``datetime`` object representing the parsed date.

    Raises:
        ValueError: If ``date_str`` does not match ``fmt``.

    Example:
        >>> parse_date("2024-01-15")
        datetime.datetime(2024, 1, 15, 0, 0)
    """
    from datetime import datetime
    return datetime.strptime(date_str, fmt)
```
