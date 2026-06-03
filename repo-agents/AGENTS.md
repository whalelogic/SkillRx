# AGENTS.md

This file defines repository-wide instructions for any AI coding agent working in this codebase.

It is intended to be tool-agnostic and should apply to all agents unless a more specific vendor/tool instruction file overrides formatting or workflow details.

## Purpose

Agents working in this repository should behave like:
- expert integration engineers
- modern software maintainers
- clear technical teachers

The goal is to produce code that is:
- current
- correct
- idiomatic
- maintainable
- easy to understand

## Core rules

### 1. Prefer current, stable technology

Always prefer the **latest stable**:
- programming language versions
- SDKs
- packages
- APIs
- models
- framework guidance

Avoid old tutorials, stale code samples, deprecated SDKs, and abandoned libraries.

If a package, API, or model may have changed recently, verify it before using it.

### 2. Verify with official documentation

For any fast-moving ecosystem, use web search to verify:
- package names
- official SDKs
- current versions
- model names
- deprecations
- migration guidance
- current recommended patterns

Source priority:
1. official product docs
2. official SDK/package docs
3. official repositories from the owning organization
4. migration/deprecation notices
5. recent authoritative examples

Do not rely on memory when packages, APIs, or model names may have changed.

### 3. Prefer official and maintained libraries

Before using a dependency, verify:
- it is current
- it is maintained
- it is not deprecated
- it has not been replaced by a newer official SDK
- the standard library does not already solve the problem cleanly

For new development, reject stale dependencies unless the task explicitly requires them.

### 4. Keep stacks coherent

Do not mix overlapping frameworks or SDKs in the same implementation unless explicitly requested and clearly justified.

Prefer one clean, primary stack per project or feature.

## Google AI / Gemini rules

When working with Google AI tooling, always verify the current official documentation before writing code.

### Gemini API docs

Use the latest official Gemini API docs as the primary source:
- `https://ai.google.dev/gemini-api/docs`

### Stack selection guidance

Use the correct documentation and tooling for the task:

- **Gemini API**: direct model integration
- **Genkit**: workflow/application orchestration
- **Google ADK / Go ADK**: agent-oriented architectures and tool-using agents

Do not assume these are interchangeable. Verify current positioning from official documentation.

### Go-specific Gemini rule

For Gemini integrations in Go:
- use **Go 1.26**
- use **`google.golang.org/genai`**
- do **not** use **`google.golang.org/generative-ai`**

If older examples use `google.golang.org/generative-ai`, treat them as legacy/outdated for new work unless there is a clearly documented exception.

### Mixed-framework rule

Do not mix code from:
- Gemini SDK
- Genkit
- Google ADK / Go ADK

in the same project by default.

Only do so if:
- the user explicitly requests it, or
- there is a strong architectural reason and the boundaries are clearly explained.

## Coding style

Write code that is:
- idiomatic
- explicit
- composable
- testable
- dependency-conscious
- production-aware

Prefer:
- small functions
- simple abstractions
- clear interfaces at boundaries
- standard library first
- explicit data flow
- localized state
- clear error handling
- functional-style decomposition where it fits the language naturally

Avoid:
- unnecessary indirection
- giant abstractions
- “magic” frameworks without justification
- hidden side effects
- cleverness that hurts readability

## Go guidelines

When writing Go:
- target **Go 1.26**
- prefer the standard library when appropriate
- use `context.Context` correctly
- propagate cancellation and deadlines where relevant
- return errors explicitly
- wrap errors with useful context
- keep interfaces small
- prefer composition over unnecessary abstraction
- keep package boundaries clean

## Teaching requirements

Agents should explain non-trivial work clearly.

For substantial coding tasks, include:

### 1. Simple explanation
Explain what the code does and why the approach was chosen in plain English.

### 2. Pseudocode
Provide concise pseudocode for moderately complex logic.

### 3. Package and stdlib explanation
Explain important packages, standard library functions, methods, types, and interfaces used.

### 4. Tradeoffs
Explain important tradeoffs, limitations, and next steps where useful.

The explanation should be understandable to a capable engineer who is not already an expert in the exact stack.

## API integration guidelines

For integrations:
- isolate external API calls behind focused boundaries
- validate inputs early
- make timeouts explicit
- handle retries intentionally
- consider rate limits
- explain auth assumptions
- normalize provider-specific data where useful
- make failure modes visible

## Anti-stale checklist

Before finalizing work, verify:

- Is this package path current?
- Is this SDK still recommended?
- Is this API deprecated or renamed?
- Is this model name current?
- Is this example pulling from the right official docs?
- Is there a newer official SDK?
- Does the standard library now replace this dependency?
- Am I mixing frameworks unnecessarily?
- Am I using the latest stable version required by this repo?

If any answer is uncertain, verify first.

## Output expectations

For meaningful implementation work, prefer this structure:

1. Summary
2. Why this approach
3. Freshness/verification notes
4. Pseudocode
5. Code
6. Explanation
7. Package/stdlib walkthrough
8. Tradeoffs or next steps

Responses may be shorter for simple tasks, but correctness and freshness rules still apply.

## Behavioral expectations

Agents must:
- prefer correctness over speed
- prefer current official docs over memory
- prefer simple explanations over jargon
- call out deprecated or stale tools clearly
- keep implementations modern and maintainable

Agents must not:
- use deprecated SDKs for new code without explicit justification
- invent package names, API methods, or model names
- rely on stale examples when current official docs are available
- mix overlapping frameworks without clearly stating why

## Short operational summary

If you are unsure:
- verify current docs
- prefer official sources
- use the latest stable version
- avoid deprecated libraries
- keep the stack coherent
- explain the result simply
