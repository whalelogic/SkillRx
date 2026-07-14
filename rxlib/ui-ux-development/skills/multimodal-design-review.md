# Multimodal Design Review — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | UI/UX Development |
| **Skill Name** | Multimodal Design Review |
| **Compatible Roles** | UX Designer, Product Designer, UI Developer |
| **Gemini Feature Used** | Multimodal (vision) — image input (`gemini-2.0-flash` or `gemini-2.5-pro`) |

---

## Skill Description

Analyse UI screenshots, wireframes, mockups, or design exports using Gemini's vision capability to provide design feedback without requiring the user to describe the interface in text.

## When to Activate

Activate this skill when the user:

* Attaches a screenshot, mockup, or exported image of a UI
* Asks "what do you think of this design?"
* Wants an accessibility or usability review of a visual interface
* Wants to compare two design options (attach both images)

## Behaviour

When this skill is active:

1. **Identify the interface type** from the image (web page, mobile screen, dashboard, form, dialog, etc.).
2. **Describe what you see** briefly to confirm correct interpretation before giving feedback.
3. **Analyse visual hierarchy:** what element draws the eye first? Is that the correct primary action?
4. **Check layout and spacing:** consistency of padding, alignment, whitespace usage.
5. **Check typography:** heading/body contrast, readable font sizes (flag anything appearing < 16px for body text).
6. **Check colour:** flag obvious colour contrast issues and over-reliance on colour alone.
7. **Check interactive affordances:** are buttons obviously clickable? Are links distinguishable from body text?
8. **ARIA/accessibility flags** visible from the image (missing labels, small touch targets, missing focus indicators).

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: Multimodal Design Review

When the user provides an image of a UI:
1. Identify the interface type and briefly describe what you see.
2. Analyse visual hierarchy, layout consistency, typography, and colour.
3. Flag potential accessibility issues visible in the image.
4. Provide actionable feedback with specific references to elements in the image (e.g. "the blue CTA button in the top right corner").
5. List 3 strengths and 3 improvement opportunities.
```

## API Configuration

To enable image input with Gemini:

```python
import google.generativeai as genai
from PIL import Image

model = genai.GenerativeModel("gemini-2.0-flash")

image = Image.open("my_design_screenshot.png")

response = model.generate_content([
    image,
    "Please review this UI design for usability and accessibility issues."
])
print(response.text)
```

## Notes and Limitations

* Vision capability is available on `gemini-2.0-flash`, `gemini-2.5-flash`, and `gemini-2.5-pro`.
* The model cannot measure exact pixel dimensions or colour hex values from images — approximate assessments only.
* For precise colour contrast ratios, use a dedicated tool such as the WebAIM Contrast Checker.
* High-fidelity prototypes with complex interactions cannot be analysed from static screenshots — describe interaction behaviours in text alongside the image.
