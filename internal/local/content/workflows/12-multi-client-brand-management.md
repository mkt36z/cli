# Workflow: Multi-Client Brand Management

*Source: Roadmap item 31 — Multi-brand policies.yaml + voice switching (Phase 4, Score 4.2)*

> Manage multiple client brands from a single Marketing36z OS instance. Switch between brand voices, compliance rules, and content strategies without cross-contamination.

---

## When to Use This Workflow

- Agency managing 2+ client brands
- Multi-brand company (parent company with sub-brands)
- Freelancer with multiple retainer clients
- Any scenario where output must match a specific brand identity

---

## Part 1: Brand Profile Setup

### Creating a New Brand Profile

For each client or brand, create a dedicated configuration block in `config/policies.yaml`:

```yaml
multi_brand:
  enabled: true
  active_brand: "client-alpha"  # Currently active brand

  brands:
    client-alpha:
      name: "Client Alpha"
      industry: "B2B SaaS"
      icp: "Mid-market HR directors, 200-2000 employees"

      voice:
        adjectives: ["confident", "empathetic", "precise"]
        avoid: ["synergy", "disrupt", "leverage", "guru"]
        tone: "Professional but warm. Use contractions. Avoid jargon."
        reading_level: "8th grade"
        pov: "we"  # we | you | they

      positioning:
        category: "Employee engagement platform"
        differentiator: "Only platform with real-time pulse surveys + AI-driven action plans"
        competitors: ["Culture Amp", "Lattice", "15Five"]
        tagline: "Engagement insights that drive action"

      compliance:
        required_disclaimers: []
        restricted_claims: ["guaranteed", "proven to"]
        regulatory_framework: "none"  # none | finra | hipaa | gdpr

      content:
        primary_topics: ["employee engagement", "HR technology", "workplace culture"]
        content_pillars: ["Engagement Science", "Manager Enablement", "Culture Building"]
        seo_primary_keywords: ["employee engagement software", "pulse survey tool"]
        cta_primary: "Start your free trial"
        cta_secondary: "See it in action"

      assets:
        logo_path: "assets/client-alpha/logo.png"
        brand_colors: ["#2563EB", "#1E40AF", "#DBEAFE"]
        font_primary: "Inter"

    client-beta:
      name: "Client Beta"
      industry: "E-commerce DTC"
      icp: "Female consumers 25-40, health-conscious, $75K+ income"

      voice:
        adjectives: ["playful", "trustworthy", "bold"]
        avoid: ["cheap", "diet", "miracle", "hack"]
        tone: "Friendly, casual, like talking to a smart friend. Use humor when appropriate."
        reading_level: "6th grade"
        pov: "you"

      positioning:
        category: "Clean beauty supplements"
        differentiator: "Dermatologist-formulated, transparent ingredient sourcing"
        competitors: ["Hum Nutrition", "Ritual", "Care/of"]
        tagline: "Beauty that starts within"

      compliance:
        required_disclaimers: ["These statements have not been evaluated by the FDA"]
        restricted_claims: ["cure", "treat", "prevent", "diagnose"]
        regulatory_framework: "fda-supplement"

      content:
        primary_topics: ["clean beauty", "skin health", "supplement science"]
        content_pillars: ["Ingredient Science", "Real Results", "Clean Living"]
        seo_primary_keywords: ["clean beauty supplements", "skin vitamins"]
        cta_primary: "Shop now"
        cta_secondary: "Take the skin quiz"

      assets:
        logo_path: "assets/client-beta/logo.png"
        brand_colors: ["#EC4899", "#BE185D", "#FCE7F3"]
        font_primary: "Playfair Display"
```

### Brand Profile Checklist

Before activating a new brand, ensure all fields are complete:

| Section | Required Fields | Quality Check |
|---------|:--------------:|---------------|
| **Identity** | Name, industry, ICP | ICP is specific (not "everyone") |
| **Voice** | 3 adjectives, avoid list, tone description | Read a paragraph aloud — does it sound like the brand? |
| **Positioning** | Category, differentiator, competitors | Differentiator is falsifiable and specific |
| **Compliance** | Disclaimers, restricted claims, framework | Legal team has reviewed |
| **Content** | Topics, pillars, keywords, CTAs | Pillars are distinct and cover the value prop |
| **Assets** | Logo, colors, font | Files exist and are current |

---

## Part 2: Brand Switching Protocol

### How to Switch Brands

**Step 1: Explicit Switch Command**

Every brand switch must be deliberate — never automatic:

```
Switch active brand to: [brand-id]
```

**Step 2: Context Load**

On switch, the system loads:
1. Voice settings (adjectives, tone, avoid list)
2. Positioning (category, differentiator, competitors)
3. Compliance rules (disclaimers, restricted claims)
4. Content strategy (pillars, keywords, CTAs)
5. Visual assets (logo, colors, fonts)

**Step 3: Confirmation**

Before generating any content, confirm:
```
Active brand: [Brand Name]
Voice: [3 adjectives]
Compliance: [Framework]
Content pillar: [Relevant pillar for this task]
```

### Brand Switching Rules

| Rule | Rationale |
|------|-----------|
| **Never generate content without confirming active brand** | Prevents cross-contamination |
| **Log every brand switch** | Audit trail for client accountability |
| **Clear context between switches** | No bleed between brand voices |
| **Validate output against brand voice before delivery** | Quality Guardian checks alignment |
| **Never reference one client's data in another's content** | Confidentiality |

---

## Part 3: Cross-Brand Content Production

### Daily Multi-Brand Workflow

```
Morning Block (Brand A):
  1. Switch to Brand A
  2. Review Brand A content calendar
  3. Produce Brand A deliverables
  4. Quality Guardian review (Brand A voice check)
  5. Save/export Brand A outputs

Afternoon Block (Brand B):
  1. Switch to Brand B
  2. Review Brand B content calendar
  3. Produce Brand B deliverables
  4. Quality Guardian review (Brand B voice check)
  5. Save/export Brand B outputs
```

### Quality Guardian Multi-Brand Checks

For each piece of content, the Quality Guardian validates:

| Check | What It Catches |
|-------|----------------|
| **Voice alignment** | Content using wrong brand's adjectives or tone |
| **Competitor mentions** | Accidentally mentioning wrong brand's competitors |
| **CTA consistency** | Using wrong brand's CTAs |
| **Compliance** | Missing required disclaimers or using restricted claims |
| **Topic relevance** | Content outside the brand's pillar structure |
| **Visual identity** | Wrong colors, fonts, or logo in templates |

### Cross-Brand Analytics

Track performance separately for each brand:

| Metric | Track Per Brand | Review Cadence |
|--------|:--------------:|:--------------:|
| Content output volume | Weekly | Weekly standup |
| Content quality score (0-100) | Per piece | Every review cycle |
| Channel performance | Weekly | Monthly review |
| SEO ranking progress | Monthly | Monthly QBR |
| Conversion metrics | Weekly | Monthly QBR |

---

## Part 4: Client Onboarding into Multi-Brand System

### New Client Setup Process

| Step | Action | Time |
|:----:|--------|:----:|
| 1 | Collect brand assets (logo, colors, fonts, style guide) | Day 1 |
| 2 | Interview for voice definition (3 adjectives, tone, avoid list) | Day 1 |
| 3 | Document positioning (category, differentiator, competitors) | Day 1-2 |
| 4 | Identify compliance requirements | Day 2 |
| 5 | Define content pillars and primary keywords | Day 2-3 |
| 6 | Create brand profile in `policies.yaml` | Day 3 |
| 7 | Generate 3 test pieces and validate voice alignment | Day 3-4 |
| 8 | Client review and voice refinement | Day 4-5 |
| 9 | Activate brand for production | Day 5 |

### Voice Calibration Test

Before going live with a new brand, generate 3 content samples:
1. A social media post
2. A blog introduction paragraph
3. An email subject line + opening sentence

Have the client rate each on a 1-5 scale for voice accuracy. Minimum score: 4/5 on all three before activating.

---

## Part 5: Scaling to 10+ Brands

### Organization Strategies

| Brand Count | Organization Strategy |
|:-----------:|----------------------|
| 2-5 | Single `policies.yaml` with all brands |
| 6-15 | One config file per brand: `config/brands/[brand-id].yaml` |
| 15+ | Brand database with API-driven config loading |

### Team Assignment Matrix

| Role | Brands per Person | Rationale |
|------|:-----------------:|-----------|
| Content strategist | 3-5 | Needs deep brand knowledge |
| Content writer | 2-3 | Voice consistency degrades beyond 3 |
| Designer | 5-8 | Visual assets are more systematic |
| Account manager | 5-10 | Relationship management scales better |

### Common Multi-Brand Pitfalls

| Pitfall | Prevention |
|---------|-----------|
| Voice bleed between brands | Batch by brand, never interleave |
| Using wrong brand's examples | Quality Guardian cross-brand check |
| Forgetting compliance requirements | Mandatory compliance check before publish |
| Uneven attention across brands | Weekly hours allocation by brand in calendar |
| Sharing strategy insights across clients | Strict information barriers — never cross-pollinate |
