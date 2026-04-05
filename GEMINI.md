# Gemini Context: SkillRx

`SkillRx` is a structured library of **AI Agent prompts**, **skills**, **system instructions**, and **configuration templates** optimized for Google Gemini. It includes a Go-based driver (`main.go`) that programmatically assembles and executes these resources using the latest Gemini SDK.

## Project Overview

*   **Purpose:** A comprehensive collection of modular AI resources organized by professional domains, designed to be consumed programmatically or manually.
*   **Core Driver:** `main.go` is the primary application that calls the Gemini API. it dynamically loads system instructions, prompts, and modular skills from the domain folders to execute tasks.
*   **Technologies:**
    *   **SDK:** Exclusively uses `google.golang.org/genai`. No legacy SDKs, Genkit, or Vertex AI packages are permitted.
    *   **Models:** Configured to use Gemini 3+ models (e.g., `gemini-3-flash`).
    *   **Resources:** Markdown (`.md`) files with `{{VARIABLE}}` placeholders located in domain-specific subdirectories.

## Building and Running

### Main Driver (Go)
The `main.go` program is the central point of execution for the library content.

*   **Prerequisites:** Go 1.25.1 or later.
*   **Configuration:** Set the `GEMINI_API_KEY` environment variable.
*   **Execution:**
    ```bash
    go run main.go
    ```
*   **Functionality:** `main.go` parses the `.md` files in the category folders (e.g., `software-engineering/prompts/code-review.md`), injects required variables, and sends the request to the Gemini API using the `google.golang.org/genai` client.

## Development Conventions

### Resource Organization
Every domain folder follows a strict layout:
*   `system-instructions/`: Persona and behavior definitions. Used as the `SystemInstruction` in the GenAI client configuration.
*   `prompts/`: Specific task templates. Used as the primary user message.
*   `skills/`: Modular capability blocks. These are appended to the system instruction or prompt to extend capabilities.

### Implementation Standards
*   **SDK Usage:** All logic MUST use `google.golang.org/genai`. Use the latest stable patterns for client creation and content generation.
*   **Variables:** Continue using `{{VARIABLE_NAME}}` for placeholders. `main.go` is responsible for replacing these before API calls.
*   **Modernization:** Avoid legacy vertex or generic AI Platform packages. Strictly use the dedicated Gemini `genai` package.

## Key Files

*   `main.go`: The core program that orchestrates prompts, skills, and system instructions using the Gemini SDK.
*   `go.mod`: Manages the `google.golang.org/genai` dependency.
*   `templates/`: Scaffolding for creating new, compatible resources.
