# API Reference Generation Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Repository Analysis |
| **Task** | Generate a complete API reference from code or specs |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `repository-analysis/system-instructions/documentation-generator.md` |
| **Output Format** | Markdown API reference |

---

## Prompt

Please generate a complete API reference for the following code or specification.

### Source

Select one:

**Option A — Source code:**
```{{LANGUAGE}}
{{SOURCE_CODE}}
```

**Option B — OpenAPI / Swagger spec:**
```yaml
{{OPENAPI_SPEC}}
```

**Option C — GraphQL schema:**
```graphql
{{GRAPHQL_SCHEMA}}
```

### Configuration

* **Documentation style:** {{DOC_STYLE}} (e.g. Google Python docstring / JSDoc / OpenAPI Markdown / custom)
* **Target audience:** {{AUDIENCE}} (e.g. "external API consumers", "internal SDK users", "TypeScript developers")
* **Include examples:** {{INCLUDE_EXAMPLES}} (yes / no)
* **Output format:** {{OUTPUT_FORMAT}} (Markdown / OpenAPI YAML / JSON)

### Generation Task

1. Extract every public API surface: functions, classes, methods, REST endpoints, GraphQL queries/mutations.
2. For each item, document: name/path, description, parameters/arguments, return value/response schema, error/exception cases.
3. Generate a usage example for each major API element.
4. Create a top-level index/table of contents.
5. Flag any undocumented parameters, missing types, or behaviours that cannot be inferred from the code alone.

### Output Requirements

* Complete reference — do not skip any public symbol.
* Do not invent parameter descriptions — only document what is inferable from names, types, and code logic.
* Use consistent formatting throughout.
* For REST APIs: include HTTP method, path, path/query parameters, request body schema, response schemas for 2xx and common error codes.

---

## Example — Python Function

**Input:**
```python
def paginate(items: list, page: int, page_size: int = 20) -> dict:
    start = (page - 1) * page_size
    end = start + page_size
    return {
        "items": items[start:end],
        "page": page,
        "page_size": page_size,
        "total": len(items),
        "total_pages": -(-len(items) // page_size),
    }
```

**Expected output excerpt:**
```
## `paginate(items, page, page_size=20) -> dict`

Returns a paginated slice of a list.

### Parameters
| Name | Type | Default | Required | Description |
|---|---|---|---|---|
| `items` | `list` | — | Yes | The full list of items to paginate |
| `page` | `int` | — | Yes | The page number to return (1-indexed) |
| `page_size` | `int` | `20` | No | Number of items per page |

### Returns
`dict` with keys: `items`, `page`, `page_size`, `total`, `total_pages`
```
