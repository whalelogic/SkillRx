# Commit Analysis — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | Repository Analysis |
| **Skill Name** | Commit Analysis |
| **Compatible Roles** | Codebase Analyzer, Documentation Generator |
| **Gemini Feature Used** | Standard generation; Code Execution for git stats |

---

## Skill Description

Extract insights from git commit history and pull request metadata — including team patterns, change velocity, hotspot files, and quality trends — to inform technical decisions.

## When to Activate

Activate this skill when the user:

* Provides git log output and wants analysis
* Asks about code churn, change frequency, or file hotspots
* Wants to generate a changelog from commit history
* Wants to understand team contribution patterns or onboarding risk areas

## How to Extract Git Data

### Commit history (past 90 days)

```bash
git log --oneline --since="90 days ago" > git_log.txt
```

### Commit stats per file (change frequency)

```bash
git log --pretty=format: --name-only | sort | uniq -c | sort -rn | head -30
```

### Author contributions

```bash
git shortlog -sn --all
```

### Files changed per commit (for churn analysis)

```bash
git log --pretty=format:"%H %s" --stat --since="90 days ago"
```

## Behaviour

When this skill is active:

1. **Parse commit messages** — extract the type (feat, fix, chore, docs, refactor, test) and scope if Conventional Commits format is used.
2. **Identify hotspot files** — files changed most frequently are highest-risk for bugs and warrant extra review and testing.
3. **Detect patterns** — large commits, commit message quality, frequent reverts, late-night/weekend commits (process or urgency indicators).
4. **Knowledge silos** — if only 1–2 contributors touch certain modules, flag as bus factor / knowledge silo risk.
5. **Trend analysis** — compare feature vs. bug fix vs. churn commits over time to assess project health.

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: Commit Analysis

When analysing git history:
1. Categorise commits by type (feat/fix/chore/docs/refactor/test).
2. Identify the top 10 most frequently changed files (hotspots).
3. Assess commit message quality (descriptive vs. vague).
4. Flag knowledge silos — modules with only 1–2 contributors.
5. Note any concerning patterns: large commits, frequent reverts, no tests alongside features.
6. Summarise findings as: velocity, quality, risk hotspots, and process health.
```

## Notes and Limitations

* Git history analysis is observational — high commit frequency can indicate either high activity or instability.
* Commit message quality varies widely; analysis depends on the discipline of the team's commit practices.
* This skill works with raw git log text — the model cannot execute git commands directly.
