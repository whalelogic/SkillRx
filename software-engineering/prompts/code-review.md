# Code Review Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Software Engineering |
| **Task** | Structured code review |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `software-engineering/system-instructions/code-reviewer.md` |
| **Output Format** | Markdown (table + prose) |

---

## Prompt

Please perform a thorough code review of the following `{{LANGUAGE}}` code.

### Context

* **Project type:** {{PROJECT_TYPE}} (e.g. REST API, CLI tool, data pipeline)
* **Review focus:** {{FOCUS}} (e.g. "all categories", "security only", "performance only")
* **Framework / version:** {{FRAMEWORK_VERSION}} (e.g. FastAPI 0.111, React 18.3)
* **Pull request description:** {{PR_DESCRIPTION}} (optional — describe what this change is meant to do)

### Code

```{{LANGUAGE}}
{{CODE}}
```

### Review Task

1. Identify all issues grouped by severity: 🔴 Critical, 🟠 High, 🟡 Medium, 🔵 Low, ℹ️ Info.
2. For each finding, cite the specific line or function name, describe the problem, and provide a corrected code snippet.
3. Note any positive patterns or good practices observed in the code.
4. Summarise the top 3 actions the author should take before merging.

### Output Requirements

* Present findings in a Markdown table with columns: `#`, `Severity`, `Location`, `Category`, `Description`, `Suggested Fix`.
* Follow the table with a "Positive Observations" section (2–5 bullet points).
* End with a numbered "Recommended Next Steps" section.

---

## Example Usage

**Input context:**
- Language: Python
- Project type: REST API (FastAPI)
- Focus: All categories

**Input code:**
```python
@app.post("/login")
def login(username: str, password: str, db: Session = Depends(get_db)):
    user = db.execute(f"SELECT * FROM users WHERE username='{username}'").fetchone()
    if user and user.password == password:
        return {"token": create_token(user.id)}
    raise HTTPException(status_code=401)
```

**Expected output excerpt:**

| # | Severity | Location | Category | Description | Suggested Fix |
|---|---|---|---|---|---|
| 1 | 🔴 Critical | `login`, line 3 | Security — SQL Injection | String interpolation in raw SQL allows an attacker to manipulate the query | Use parameterised queries: `db.execute("SELECT * FROM users WHERE username=:u", {"u": username})` |
| 2 | 🔴 Critical | `login`, line 4 | Security — Plain-text password comparison | Passwords are compared in plain text, indicating they are stored unhashed | Hash passwords with `bcrypt` or `argon2` and compare with `pwd_context.verify()` |
