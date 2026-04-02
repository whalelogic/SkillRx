# Code Execution — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | Software Engineering |
| **Skill Name** | Code Execution |
| **Compatible Roles** | Senior Software Engineer, Backend Developer, Data Scientist |
| **Gemini Feature Used** | Code Execution tool (requires `tools=[{"code_execution": {}}]` in API request) |

---

## Skill Description

Run Python code snippets inside the model's secure sandboxed environment to validate logic, compute results, and demonstrate correctness before returning the final answer.

## When to Activate

Activate this skill when the user:

* Asks for a numerical result that requires calculation (e.g. algorithm benchmarks, data transformations)
* Requests validation that a code snippet produces the expected output
* Asks "does this work?" or "what does this return?"
* Wants a demonstration of a library's API with real output

## Behaviour

When this skill is active:

1. Determine whether the answer requires running code to verify or compute.
2. If yes, write a minimal test harness and execute it using the Code Execution tool.
3. Include the actual output in your response (not simulated output).
4. If execution raises an error, debug and re-run before presenting the final answer.
5. After execution, present the validated code and its output clearly.

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: Code Execution

When you need to verify logic or compute a result, use the Code Execution tool to run Python code in the sandbox. Always:
- Run minimal, self-contained snippets.
- Include the real execution output in your response.
- Debug any errors and re-run before presenting the final answer.
- Clearly label executed code blocks with a comment: `# Executed — output below`.
```

## API Configuration

Enable this skill by adding `code_execution` to your tools list:

```python
import google.generativeai as genai

model = genai.GenerativeModel(
    model_name="gemini-2.0-flash",
    tools=[{"code_execution": {}}],
)
```

## Notes and Limitations

* Only Python is supported in the code execution sandbox.
* The sandbox is stateless — variables do not persist between separate tool calls.
* External network calls, file system access, and GPU operations are not available in the sandbox.
* Execution has a time limit; avoid long-running computations.
