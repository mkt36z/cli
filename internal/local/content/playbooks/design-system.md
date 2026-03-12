# Playbook: Design System / Brand Kit

## Sources
- Marty Neumeier — *The Brand Gap*, Brand as a System
- Jacob Cass — Brand Identity Design Process
- InVision / Brad Frost — Atomic Design Systems
- Marketing36z OS synthesis

---

## Part 1: Why You Need a Design System

> **Every inconsistent touchpoint erodes trust. A design system ensures your brand looks intentional across every pixel, every page, every platform.**

A design system is not just a logo file. It's the complete set of rules, components, and assets that ensure visual consistency — from your website to your social posts to your sales deck.

---

## Part 2: Brand Identity Foundation

### Logo System

| Asset | Specifications | Usage Rules |
|-------|---------------|-------------|
| **Primary logo** | Full logo (mark + wordmark) | Default usage on light backgrounds |
| **Secondary logo** | Simplified or stacked version | When space is limited |
| **Logo mark only** | Icon/symbol without wordmark | Favicons, app icons, social avatars |
| **Wordmark only** | Name in brand typeface | Editorial contexts, co-branding |
| **Reversed/white logo** | For dark backgrounds | Ensure contrast ratio ≥ 4.5:1 |

### Logo Clear Space and Minimum Size

```
┌──────────────────────────────┐
│                              │
│    ┌──┐                      │
│    │  │  ← Clear space =     │
│    └──┘    height of the     │
│            mark on all sides │
│                              │
│    Minimum size:             │
│    Digital: 24px height      │
│    Print: 0.5" / 12mm       │
│                              │
└──────────────────────────────┘
```

### Logo Don'ts

| Never Do This | Why |
|--------------|-----|
| Stretch or distort | Damages perceived quality |
| Change logo colors | Brand recognition requires consistency |
| Add effects (drop shadow, glow) | Dates the brand instantly |
| Place on busy backgrounds | Reduces readability |
| Rotate the logo | Undermines brand recognition |
| Crop or partially hide | Legal trademark issues |

---

## Part 3: Color System

### Brand Color Palette

| Role | Color Name | Hex | RGB | Usage |
|------|-----------|:---:|:---:|-------|
| **Primary** | | #______ | | Headlines, CTAs, primary actions |
| **Secondary** | | #______ | | Supporting elements, accents |
| **Accent** | | #______ | | Highlights, notifications, badges |
| **Dark / Text** | | #______ | | Body text, headings |
| **Light / Background** | | #______ | | Page backgrounds, cards |
| **Success** | | #______ | | Confirmations, positive states |
| **Warning** | | #______ | | Alerts, caution states |
| **Error** | | #______ | | Errors, destructive actions |
| **Neutral 100-900** | | #______ × 9 | | Borders, disabled states, dividers |

### Color Accessibility Rules

| Combination | Minimum Contrast | WCAG Level | Use Case |
|-------------|:----------------:|:----------:|----------|
| Text on background | 4.5:1 | AA | Body text |
| Large text on background | 3:1 | AA | Headings (18px+ or 14px+ bold) |
| UI components | 3:1 | AA | Buttons, form fields, icons |
| Enhanced readability | 7:1 | AAA | Critical content |

### Color Usage Ratios

```
Primary:    60% of branded surfaces (backgrounds, large areas)
Secondary:  30% of branded surfaces (supporting elements)
Accent:     10% of branded surfaces (CTAs, highlights, key actions)
```

---

## Part 4: Typography System

### Font Stack

| Role | Typeface | Weight | Size (Desktop) | Size (Mobile) | Usage |
|------|----------|:------:|:--------------:|:-------------:|-------|
| **Display / H1** | | Bold (700) | 48-64px | 32-40px | Hero headlines, landing pages |
| **H2** | | Semi-Bold (600) | 36-40px | 24-28px | Section headings |
| **H3** | | Semi-Bold (600) | 24-28px | 20-24px | Sub-section headings |
| **H4** | | Medium (500) | 20-24px | 18-20px | Card titles, labels |
| **Body** | | Regular (400) | 16-18px | 16px | Paragraphs, descriptions |
| **Small / Caption** | | Regular (400) | 12-14px | 12-14px | Footnotes, labels, metadata |
| **Code / Mono** | | Regular (400) | 14-16px | 14px | Code blocks, data |

### Typography Rules

| Rule | Specification |
|------|-------------|
| **Line height (body)** | 1.5-1.6× font size |
| **Line height (headings)** | 1.2-1.3× font size |
| **Paragraph spacing** | 1× font size between paragraphs |
| **Maximum line length** | 65-75 characters (for readability) |
| **Font pairings** | Maximum 2 typefaces (1 heading + 1 body) |
| **Fallback fonts** | Always specify system fallbacks: `-apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif` |

---

## Part 5: Spacing and Layout

### The Spacing Scale

Use a consistent base unit (recommended: 4px or 8px) with multipliers:

| Token | Size (8px base) | Usage |
|-------|:---------------:|-------|
| `space-xs` | 4px | Inline gaps, icon padding |
| `space-sm` | 8px | Tight groupings, form field gaps |
| `space-md` | 16px | Default spacing between elements |
| `space-lg` | 24px | Section separators within a group |
| `space-xl` | 32px | Between major sections |
| `space-2xl` | 48px | Between page sections |
| `space-3xl` | 64px | Major section dividers, hero padding |

### Grid System

| Context | Columns | Gutter | Margin |
|---------|:-------:|:------:|:------:|
| **Desktop (1200px+)** | 12 | 24px | 64px |
| **Tablet (768-1199px)** | 8 | 16px | 32px |
| **Mobile (320-767px)** | 4 | 16px | 16px |

---

## Part 6: Component Library

### Buttons

| Variant | Usage | Hierarchy |
|---------|-------|:---------:|
| **Primary** | Main CTA (1 per section) | Highest |
| **Secondary** | Supporting actions | Medium |
| **Tertiary / Ghost** | Navigational, low-priority | Low |
| **Destructive** | Delete, cancel, irreversible | Situational |
| **Disabled** | Action not yet available | N/A |

### Button Rules
1. One primary button per visible screen
2. Button text = action verb ("Get Started", not "Submit")
3. Minimum touch target: 44×44px (mobile accessibility)
4. Consistent border radius across all buttons

### Cards

| Element | Specification |
|---------|-------------|
| Border radius | Consistent (8px recommended) |
| Shadow | Subtle elevation (0 2px 4px rgba(0,0,0,0.1)) |
| Padding | `space-lg` (24px) internal |
| Spacing | `space-md` (16px) between cards |

### Form Elements

| Element | Specifications |
|---------|-------------|
| Input height | 44-48px (accessibility) |
| Label position | Above input (not inside) |
| Error state | Red border + error message below |
| Focus state | Visible focus ring (2px solid primary) |
| Placeholder text | Light gray, descriptive, disappears on focus |

---

## Part 7: Imagery and Iconography

### Photography Style Guide

| Dimension | Direction | Avoid |
|-----------|----------|-------|
| **Tone** | Authentic, natural, warm | Stock-photo staged poses |
| **Color treatment** | Consistent filter/preset | Over-saturated or inconsistent |
| **Subjects** | Real customers, real team, real product | Generic lifestyle imagery |
| **Composition** | Clean, purposeful, space for text overlay | Cluttered, busy |
| **Diversity** | Representative of actual customer base | Tokenistic or homogeneous |

### Icon System

| Property | Standard |
|----------|---------|
| **Style** | Outline / Filled / Duo-tone (pick one, be consistent) |
| **Size grid** | 16px, 20px, 24px, 32px |
| **Stroke weight** | Consistent (1.5px recommended) |
| **Corner radius** | Match button/card radius |
| **Color** | Use brand palette only — icons inherit text color by default |

---

## Part 8: Social and Marketing Templates

### Social Media Sizes

| Platform | Post | Story | Profile | Banner |
|----------|:----:|:-----:|:-------:|:------:|
| **LinkedIn** | 1200×628 | N/A | 400×400 | 1584×396 |
| **X/Twitter** | 1200×675 | N/A | 400×400 | 1500×500 |
| **Instagram** | 1080×1080 | 1080×1920 | 320×320 | N/A |
| **Facebook** | 1200×630 | 1080×1920 | 320×320 | 820×312 |
| **TikTok** | N/A | 1080×1920 | 200×200 | N/A |

### Marketing Collateral Templates

| Template | Specification | Primary Use |
|----------|-------------|------------|
| **Blog post header** | 1200×628px | Blog featured images |
| **Email header** | 600×200px | Newsletter, drip sequences |
| **Presentation deck** | 16:9 (1920×1080) | Sales decks, pitches, webinars |
| **One-pager** | Letter (8.5×11) / A4 | Sales collateral, leave-behinds |
| **Ad creative** | Platform-specific (above) | Paid campaigns |
| **Event banner** | 2000×500px | Conference, webinar branding |

---

## Part 9: Brand Kit Checklist

### Essential Assets (Must Have)

- [ ] Logo files (SVG, PNG, EPS) in all variants
- [ ] Brand color palette with hex/RGB codes
- [ ] Typography specification with font files
- [ ] Brand guidelines document (this playbook, customized)
- [ ] Social media templates (Figma/Canva)
- [ ] Email template
- [ ] Presentation template

### Advanced Assets (Growth Stage)

- [ ] Icon library
- [ ] Photography style guide with example images
- [ ] Motion/animation guidelines
- [ ] Component library (Figma/Storybook)
- [ ] Print templates (business cards, letterhead)
- [ ] Merchandise guidelines

### Enterprise Assets (Scale Stage)

- [ ] Full design system documentation (Storybook or similar)
- [ ] Multi-brand guidelines (sub-brands, acquired brands)
- [ ] Accessibility audit documentation
- [ ] Co-branding guidelines
- [ ] Localization adaptation rules
- [ ] Video production style guide

---

## Integration with Marketing36z OS

| System Component | How This Playbook Connects |
|-----------------|---------------------------|
| **Brand Foundation** | Design system is the visual expression of brand foundation |
| **Content Engine** | All content production follows design system templates |
| **Quality Guardian** | Checks visual consistency against design system standards |
| **Channel Operator** | Platform-specific templates ensure channel-fit formatting |
| **White-Label Output** | Agency tier uses design system for client brand management |
