# Changelog Generation Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Repository Analysis |
| **Task** | Generate a human-readable changelog from git history or PR descriptions |
| **Recommended Model** | gemini-2.0-flash |
| **Recommended System Instruction** | `repository-analysis/system-instructions/documentation-generator.md` |
| **Output Format** | Keep a Changelog Markdown format |

---

## Prompt

Please generate a formatted changelog from the following git history / PR descriptions.

### Source Material

Paste one of the following:

**Option A — Git log output:**
```
{{GIT_LOG}}
```

Generate with: `git log --oneline --no-merges v1.0.0..HEAD` or `git log --pretty=format:"%h %s (%an, %ad)" --date=short v1.0.0..HEAD`

**Option B — PR/MR descriptions:**
```
{{PR_DESCRIPTIONS}}
```

**Option C — Raw commit messages:**
```
{{COMMIT_MESSAGES}}
```

### Configuration

* **Version number:** {{VERSION}} (e.g. `2.3.0` — leave blank to use "Unreleased")
* **Release date:** {{RELEASE_DATE}} (YYYY-MM-DD — leave blank for today)
* **Changelog format:** {{FORMAT}} (Keep a Changelog / Conventional Commits auto / plain grouped list)
* **Audience:** {{AUDIENCE}} (developers / end users / both)
* **Group by:** {{GROUP_BY}} (type: Added/Changed/Fixed/etc. / component: auth, UI, API / both)

### Generation Task

1. Parse the raw input and extract meaningful changes (filter out noise like "fix typo", "bump version", merge commits).
2. Group changes into categories: **Added**, **Changed**, **Deprecated**, **Removed**, **Fixed**, **Security**.
3. For each entry: write a clear, concise description in 1 sentence; include the PR/commit reference in parentheses.
4. If audience is "end users", translate technical commit messages into user-facing language.
5. Flag any commits that are ambiguous or need clarification.

### Output Format

```markdown
## [{{VERSION}}] — {{RELEASE_DATE}}

### Added
- Description of new feature ([#PR_NUMBER](link))

### Changed
- Description of change

### Fixed
- Description of bug fix

### Security
- Description of security fix — see advisory [GHSA-XXXX](link)
```
