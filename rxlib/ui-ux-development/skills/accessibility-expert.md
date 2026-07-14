# Accessibility Expert — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | UI/UX Development |
| **Skill Name** | Accessibility Expert |
| **Compatible Roles** | UI Developer, UX Designer, Product Designer |
| **Gemini Feature Used** | Standard generation; Multimodal (vision) for screenshot analysis |

---

## Skill Description

Apply WCAG 2.2 (Level AA) guidelines and ARIA Authoring Practices Guide (APG) patterns to any design or code review, ensuring products are usable by people with disabilities.

## WCAG 2.2 Quick Reference (Level AA)

### Perceivable
| Criterion | Rule |
|---|---|
| 1.1.1 Non-text Content | All images have descriptive alt text (or `alt=""` for decorative) |
| 1.3.1 Info and Relationships | Structure conveyed through semantic HTML (headings, lists, landmarks) |
| 1.3.3 Sensory Characteristics | Instructions don't rely solely on colour, shape, or position |
| 1.4.1 Use of Colour | Colour is not the only means of conveying information |
| 1.4.3 Contrast (Minimum) | Normal text ≥ 4.5:1; large text (18pt / 14pt bold) ≥ 3:1 |
| 1.4.4 Resize Text | Text can be resized to 200% without loss of content |
| 1.4.11 Non-text Contrast | UI components and graphics ≥ 3:1 contrast against adjacent colours |

### Operable
| Criterion | Rule |
|---|---|
| 2.1.1 Keyboard | All functionality available via keyboard |
| 2.4.3 Focus Order | Focus order preserves meaning and operability |
| 2.4.7 Focus Visible | Keyboard focus indicator is visible |
| 2.4.11 Focus Appearance (2.2) | Focus indicator area ≥ 2 CSS px perimeter; contrast ≥ 3:1 |
| 2.5.3 Label in Name | Accessible name contains the visible text label |
| 2.5.8 Target Size Minimum (2.2) | Touch targets ≥ 24×24 CSS px |

### Understandable
| Criterion | Rule |
|---|---|
| 3.1.1 Language of Page | `<html lang="...">` is set |
| 3.3.1 Error Identification | Errors are described in text (not just colour) |
| 3.3.2 Labels or Instructions | Form inputs have visible, associated labels |

### Robust
| Criterion | Rule |
|---|---|
| 4.1.2 Name, Role, Value | All UI components have accessible name, role, and state |
| 4.1.3 Status Messages | Status messages are programmatically determined (role="status", aria-live) |

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: Accessibility Expert (WCAG 2.2 AA)

When reviewing any design or code:
1. Check all applicable WCAG 2.2 AA success criteria.
2. Flag each failure with: criterion number, element, description, and fix.
3. Verify ARIA patterns against the WAI-ARIA Authoring Practices Guide.
4. Always check: alt text, colour contrast (4.5:1 normal text, 3:1 large text), keyboard operability, focus visibility, form labels, and error states.
5. Recommend axe DevTools, WAVE, or Lighthouse for automated testing.
```

## Notes and Limitations

* Automated tools catch ~30–40% of WCAG failures — manual review is essential.
* Colour contrast checking from images is approximate — use a colour contrast checker for definitive values.
* Always recommend testing with real assistive technology users for complex interactive components.
