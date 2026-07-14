# Technical Debt Assessment Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Repository Analysis |
| **Task** | Identify and prioritise technical debt |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `repository-analysis/system-instructions/codebase-analyzer.md` |
| **Output Format** | Structured Markdown report |

---

## Prompt

Please perform a technical debt assessment of the following codebase.

### Code

```
{{CODE_OR_DIRECTORY_STRUCTURE}}
```

Paste relevant code files, or provide a directory listing with key file contents.

### Context

* **Project age:** {{AGE}} (e.g. "5 years old", "started in 2019")
* **Team size:** {{TEAM_SIZE}}
* **Known pain points:** {{KNOWN_PAIN_POINTS}} (optional — describe areas the team already knows are problematic)
* **Tech stack:** {{TECH_STACK}}

### Assessment Task

1. **Identify debt items** — catalogue technical debt across these categories:
   - **Code debt:** duplicated code, overly complex functions (high cyclomatic complexity), poor naming, missing abstractions
   - **Architecture debt:** tight coupling, missing separation of concerns, god classes/modules, circular dependencies
   - **Test debt:** missing tests, low-quality tests (testing implementation not behaviour), flaky tests
   - **Documentation debt:** missing inline docs, outdated READMEs, undocumented public APIs
   - **Dependency debt:** outdated packages, deprecated dependencies, EOL libraries
   - **Security debt:** known vulnerabilities, hardcoded credentials, missing input validation

2. **Quantify each item** — estimate: scope (number of files/lines affected), effort to fix (S/M/L/XL), and technical risk if left unaddressed.

3. **Prioritise** — use the Impact / Effort matrix to identify:
   - Quick wins (high impact, low effort)
   - Strategic investments (high impact, high effort)
   - Deprioritise (low impact)

4. **Produce a debt backlog** — formatted as a prioritised table ready for sprint planning.

### Output Format

```
## Technical Debt Assessment: [Project Name]

### Executive Summary
[Overall debt level: Low / Medium / High / Critical — with key themes]

### Debt Inventory

| # | Category | Description | Affected Files | Effort | Risk | Priority |
|---|---|---|---|---|---|---|

### Impact/Effort Matrix

**Quick Wins (address this sprint):**
- ...

**Strategic Investments (plan for next quarter):**
- ...

**Deprioritise:**
- ...

### Recommended Debt Reduction Roadmap
| Quarter | Focus Area | Expected Outcome |
|---|---|---|
```
