# Debugging Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Software Engineering |
| **Task** | Diagnose and fix a bug |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `software-engineering/system-instructions/senior-software-engineer.md` |
| **Output Format** | Markdown with code blocks |

---

## Prompt

I am debugging a `{{LANGUAGE}}` issue and need your help diagnosing the root cause and applying a fix.

### Environment

* **Language / runtime version:** {{LANGUAGE_VERSION}} (e.g. Python 3.12, Node 20.x)
* **Framework:** {{FRAMEWORK}} (e.g. Django 5.0, Express 4.x)
* **Operating system:** {{OS}} (optional)

### Error / Unexpected Behaviour

```
{{ERROR_OUTPUT}}
```

### Relevant Code

```{{LANGUAGE}}
{{CODE}}
```

### What I Have Already Tried

{{ATTEMPTED_FIXES}}

### Debugging Task

1. Identify the root cause of the error or unexpected behaviour.
2. Explain clearly *why* this is happening — not just what the fix is.
3. Provide the corrected code with a brief diff summary of what changed.
4. Suggest any preventive measures to avoid this class of bug in the future (e.g. linting rules, type annotations, test coverage).

### Output Requirements

* Start with a "Root Cause" section (2–4 sentences).
* Provide the fixed code in a fenced block.
* End with a "Prevention" bullet list.

---

## Example Usage

**Error:**
```
TypeError: unsupported operand type(s) for +: 'int' and 'str'
  File "app.py", line 14, in calculate_total
    total = base_price + tax_rate
```

**Code:**
```python
def calculate_total(base_price: int, tax_rate):
    return base_price + tax_rate
```

**Expected root cause explanation:**
`tax_rate` is passed as a string (e.g. `"0.2"`) because it was read directly from user input without conversion. Adding an `int` and a `str` raises a `TypeError` in Python.

**Expected fix:**
```python
def calculate_total(base_price: int, tax_rate: float) -> float:
    return base_price + float(tax_rate)
```
