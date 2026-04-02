# Chain-of-Thought Coding — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | Software Engineering |
| **Skill Name** | Chain-of-Thought Coding |
| **Compatible Roles** | All software engineering roles |
| **Gemini Feature Used** | Extended thinking / reasoning tokens (gemini-2.5-pro) |

---

## Skill Description

Before producing any code, reason step-by-step about the problem in a visible scratchpad section, then deliver the final implementation.

## When to Activate

Activate this skill when the user:

* Asks for a solution to a non-trivial algorithmic or architectural problem
* Provides incomplete requirements and the agent must make assumptions
* Requests an explanation of *how* a solution works, not just *what* it is
* Asks to compare multiple implementation approaches

## Behaviour

When this skill is active:

1. Open a `### Reasoning` section and think through: (a) what the problem is really asking, (b) constraints, (c) candidate approaches with trade-offs.
2. State the chosen approach and why before writing any code.
3. Write the implementation with inline comments explaining non-obvious choices.
4. Close with a `### Verification` section that traces through one example manually to confirm correctness.

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: Chain-of-Thought Coding

For any non-trivial coding task, before writing code:
1. Open a "### Reasoning" section and think step-by-step through the problem, constraints, and candidate approaches.
2. State your chosen approach with a one-sentence justification.
3. Write the implementation with inline comments on non-obvious decisions.
4. Close with a "### Verification" section tracing one concrete example manually through your solution.
```

## Notes and Limitations

* Adds latency and output length — best suited for `gemini-2.5-pro` with a high `max_output_tokens` budget.
* For trivial or purely stylistic tasks (e.g. rename a variable), skip the reasoning section and answer directly.
