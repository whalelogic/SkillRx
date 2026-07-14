# Test Generation Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Software Engineering |
| **Task** | Generate unit or integration tests |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `software-engineering/system-instructions/senior-software-engineer.md` |
| **Output Format** | Markdown with code blocks |

---

## Prompt

Generate comprehensive `{{TEST_TYPE}}` tests for the following `{{LANGUAGE}}` code using `{{TEST_FRAMEWORK}}`.

### Test Configuration

* **Test type:** {{TEST_TYPE}} (unit / integration / contract / end-to-end)
* **Test framework:** {{TEST_FRAMEWORK}} (e.g. pytest, Jest, JUnit 5, Go testing, RSpec)
* **Mocking library:** {{MOCK_LIBRARY}} (e.g. pytest-mock, jest.mock, Mockito)
* **Coverage target:** {{COVERAGE_TARGET}} (e.g. "all branches", "happy path and primary error paths")

### Code Under Test

```{{LANGUAGE}}
{{CODE}}
```

### Additional Context

* **External dependencies to mock:** {{DEPENDENCIES_TO_MOCK}} (e.g. database, external HTTP API, file system)
* **Known edge cases to cover:** {{EDGE_CASES}}

### Task

1. Identify the distinct behaviours (happy paths, error paths, edge cases, boundary conditions) that need testing.
2. Generate a complete, runnable test file covering those behaviours.
3. For each test, briefly explain what scenario it covers.

### Output Requirements

* A single, complete, runnable test file — not fragments.
* Descriptive test names that read as sentences (e.g. `test_returns_empty_list_when_input_is_none`).
* All dependencies properly mocked or stubbed.
* No unnecessary `print` statements in test code.
* Add an inline comment above each test group explaining the scenario.

---

## Example Usage

**Code under test (Python):**
```python
def divide(a: float, b: float) -> float:
    if b == 0:
        raise ValueError("Cannot divide by zero")
    return a / b
```

**Framework:** pytest

**Expected output excerpt:**
```python
import pytest
from mymodule import divide

class TestDivide:
    # Happy path
    def test_divides_two_positive_numbers(self):
        assert divide(10.0, 2.0) == 5.0

    def test_divides_negative_by_positive(self):
        assert divide(-10.0, 2.0) == -5.0

    # Edge cases
    def test_divide_returns_zero_when_numerator_is_zero(self):
        assert divide(0.0, 5.0) == 0.0

    # Error path
    def test_raises_value_error_when_divisor_is_zero(self):
        with pytest.raises(ValueError, match="Cannot divide by zero"):
            divide(10.0, 0.0)
```
