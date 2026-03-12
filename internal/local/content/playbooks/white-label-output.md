# White-Label Output System — Agency & Reseller Guide

*Source: Roadmap item 40 — System enhancement for white-label deliverables (Phase 4, Score 3.3)*

> Remove all Marketing36z OS branding from client-facing deliverables. Replace with client branding or your agency's branding for professional, seamless delivery.

---

## Overview

White-label output enables agencies and consultants to deliver Marketing36z OS outputs as their own work product. This system covers branding removal, custom branding injection, export format configuration, and template customization.

---

## Part 1: Configuration

### Enabling White-Label Mode

```yaml
# config/policies.yaml
white_label:
  enabled: true
  remove_branding: true          # Strip all Marketing36z OS references
  custom_header: "Prepared by [Your Agency Name]"
  custom_footer: "Confidential — [Client Name] © 2025"
  export_formats:
    - markdown
    - pdf
    - docx
  client_logo_path: "assets/client/logo.png"
  agency_logo_path: "assets/agency/logo.png"
  custom_report_template: "templates/agency-report.md"

  # Branding replacement rules
  replacements:
    "Marketing36z OS": "[Your Agency Name] Platform"
    "Brain36z": "[Your Agency Name] Intelligence"
    "marketing36z": "[your-agency]"
```

### Per-Client Override

```yaml
white_label:
  clients:
    client-alpha:
      header: "Prepared by [Agency] for Client Alpha"
      footer: "Confidential — Client Alpha © 2025"
      logo: "assets/clients/alpha/logo.png"
      color_primary: "#2563EB"
      template: "templates/alpha-report.md"

    client-beta:
      header: "Prepared by [Agency] for Client Beta"
      footer: "Confidential — Client Beta © 2025"
      logo: "assets/clients/beta/logo.png"
      color_primary: "#EC4899"
      template: "templates/beta-report.md"
```

---

## Part 2: Branding Removal Checklist

Before delivering any output to a client, verify all internal branding has been removed:

| Check | What to Remove | Replace With |
|-------|---------------|-------------|
| **Document headers** | "Marketing36z OS" references | Agency name or blank |
| **Footers** | System-generated footers | Client/agency footer |
| **Agent references** | "Strategy Planner", "Content Engine", etc. | Generic: "Strategic analysis", "Content recommendations" |
| **Playbook citations** | "See playbook X" or "per Marketing36z framework" | "Per our methodology" or remove |
| **System metadata** | File paths, config references, version numbers | Remove entirely |
| **Watermarks** | Any visual branding elements | Agency branding or blank |
| **URLs** | Links to Marketing36z resources | Links to agency resources or remove |

### Automated Branding Scan

Before export, scan all output for these patterns and flag/replace:

```
Patterns to catch:
- "Marketing36z" (any capitalization)
- "Brain36z"
- "marketing36z-os"
- Agent names: "Orchestrator", "Strategy Planner", "Content Engine",
  "Channel Operator", "Quality Guardian", "LLM Council"
- Internal file paths: "/playbooks/", "/workflows/", "/templates/"
- Internal references: "See SYSTEM.md", "per config/policies.yaml"
```

---

## Part 3: Export Formats

### Markdown (Default)

Best for: Internal handoff, CMS import, developer-friendly clients.

```markdown
# [Report Title]

**Prepared by:** [Agency Name]
**For:** [Client Name]
**Date:** [Date]

---

[Content body — all branding stripped, agency branding applied]

---

*Confidential — [Client Name] © [Year]*
```

### PDF Export

Best for: Client presentations, QBR decks, formal deliverables.

**PDF Generation Pipeline:**
1. Start from white-labeled markdown
2. Apply custom CSS template with client colors and fonts
3. Inject client logo in header
4. Add page numbers and footer
5. Generate PDF via markdown-to-PDF tool (e.g., Pandoc, WeasyPrint, or Puppeteer)

**CSS Template Variables:**

```css
:root {
  --color-primary: #2563EB;     /* Client brand color */
  --color-secondary: #1E40AF;   /* Client secondary */
  --color-background: #FFFFFF;
  --font-heading: 'Inter', sans-serif;
  --font-body: 'Inter', sans-serif;
  --logo-url: url('client-logo.png');
}
```

### DOCX Export

Best for: Clients who need to edit deliverables, collaborative workflows.

**DOCX Template Requirements:**
1. Pre-built `.docx` template with client branding in headers/footers
2. Style mapping: H1-H6, body, table, code block styles
3. Cover page with client logo and report metadata
4. Generate via Pandoc: `pandoc input.md -o output.docx --reference-doc=template.docx`

---

## Part 4: Report Templates

### Strategy Report Template

```markdown
# [Client Name] — Marketing Strategy Report

**Prepared by:** [Agency Name]
**Date:** [Quarter] [Year]
**Version:** [X.X]

---

## Executive Summary

[2-3 paragraph overview of findings and recommendations]

## Current State Analysis

### Market Position
[Positioning analysis — no internal framework references]

### Competitive Landscape
[Competitor analysis output]

### Performance Review
[Data from GA4/GSC integration — presented as agency analysis]

## Recommendations

### Priority 1: [Recommendation]
- **Impact:** [Expected outcome]
- **Timeline:** [Implementation window]
- **Investment:** [Resource/budget requirement]

### Priority 2: [Recommendation]
[...]

### Priority 3: [Recommendation]
[...]

## 90-Day Action Plan

| Week | Action | Owner | Deliverable |
|:----:|--------|-------|-------------|
| 1-2 | [Action] | [Team member] | [Output] |
| 3-4 | [Action] | [Team member] | [Output] |
[...]

## Appendix

### Methodology
[Generic description of approach — no Marketing36z references]

### Data Sources
[List data sources used — GA4, GSC, etc.]

---

*Confidential — Prepared exclusively for [Client Name]*
*© [Year] [Agency Name]. All rights reserved.*
```

### Content Calendar Template (White-Labeled)

```markdown
# [Client Name] — Content Calendar

**Period:** [Month/Quarter]
**Prepared by:** [Agency Name]

| Week | Topic | Format | Channel | Target Keyword | Status |
|:----:|-------|--------|---------|---------------|:------:|
| 1 | [Topic] | Blog | SEO | [keyword] | Draft |
| 1 | [Topic] | Social | LinkedIn | — | Scheduled |
[...]
```

### QBR Deck Template (White-Labeled)

```markdown
# Quarterly Business Review
## [Client Name] — [Quarter] [Year]

### Prepared by [Agency Name]

---

## Slide 1: Quarter in Review
- Key metrics vs. targets
- Wins and highlights
- Challenges encountered

## Slide 2: Channel Performance
[GA4/GSC data presented in agency format]

## Slide 3: Content Performance
[Top performing content with engagement metrics]

## Slide 4: Next Quarter Plan
[Strategic priorities and tactical calendar]

## Slide 5: Investment & ROI
[Budget allocation and return metrics]
```

---

## Part 5: Quality Assurance for White-Label

### Pre-Delivery Checklist

Run this checklist before every client delivery:

| # | Check | Pass/Fail |
|---|-------|:---------:|
| 1 | All "Marketing36z" references removed | |
| 2 | All agent name references removed | |
| 3 | Client logo present in header | |
| 4 | Client name correct throughout | |
| 5 | Agency branding applied consistently | |
| 6 | Footer includes confidentiality notice | |
| 7 | No internal file paths or system references | |
| 8 | Export format matches client preference | |
| 9 | Date and version number current | |
| 10 | Methodology section uses generic language | |

### Common White-Label Mistakes

| Mistake | How It Happens | Prevention |
|---------|---------------|-----------|
| Internal branding leak | Copy-paste from system output | Automated branding scan before export |
| Wrong client's branding | Forgot to switch brand profile | Mandatory brand confirmation before generating |
| Inconsistent formatting | Manual template application | Use pre-built templates, not ad-hoc formatting |
| Missing confidentiality notice | Template not updated | Include in all templates by default |
| Referring to "the AI" | Output mentions AI generation | Review and humanize all deliverables |
