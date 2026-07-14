# Dependency Analysis Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Repository Analysis |
| **Task** | Analyse project dependencies for risk, outdatedness, and security |
| **Recommended Model** | gemini-2.0-flash |
| **Recommended System Instruction** | `repository-analysis/system-instructions/dependency-auditor.md` |
| **Output Format** | Structured Markdown audit report |

---

## Prompt

Please analyse the following project dependencies for security vulnerabilities, outdated packages, licence issues, and supply chain risks.

### Dependency Files

Paste the dependency manifest file(s):

```
{{DEPENDENCY_FILES}}
```

Examples of what to paste:
* `package.json` + `package-lock.json` (or `yarn.lock`)
* `requirements.txt` + `pip freeze` output
* `pyproject.toml`
* `go.mod` + `go.sum`
* `Cargo.toml` + `Cargo.lock`
* `pom.xml`
* `Gemfile` + `Gemfile.lock`

### Audit Configuration

* **Environment:** {{ENV}} (production / development / both)
* **Language/ecosystem:** {{ECOSYSTEM}}
* **Licence policy:** {{LICENCE_POLICY}} (e.g. "only MIT, Apache 2.0, and BSD permitted", "no copyleft in production", "no restrictions")

### Analysis Task

1. **Vulnerability scan** — identify known CVEs for all listed packages (direct and transitive where inferable). Prioritise Critical and High severity.
2. **Outdated packages** — flag packages where there are major version updates available. Note breaking change risk.
3. **EOL/unmaintained packages** — identify packages that are end-of-life, abandoned, or have no recent releases.
4. **Licence compliance** — flag any packages whose licences conflict with the stated licence policy.
5. **Dependency health** — assess the overall dependency graph size, depth, and any circular dependency risks.
6. **Supply chain risks** — flag any packages with known supply chain incidents, unusual publish patterns, or very low download counts.
7. **Remediation plan** — provide a prioritised table of recommended updates with effort estimates.

### Output Requirements

* Prioritise by severity: Critical > High > Medium > Low.
* For each vulnerability, provide: Package, Current version, CVE/Advisory ID, CVSS score, Description, Fixed version.
* Include a "run this command" section for automated scanning.
* Note the knowledge cut-off date and recommend running live scanning tools.
