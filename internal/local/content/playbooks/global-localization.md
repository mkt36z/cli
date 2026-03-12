# Playbook: Global Marketing & Localization

## Sources
- Nataly Kelly — *Take Your Company Global*
- Common Sense Advisory — Localization Best Practices
- W3C Internationalization Standards
- Marketing36z OS synthesis

---

## Part 1: Localization Strategy Framework

### Localization vs. Translation

| Dimension | Translation | Localization |
|-----------|-----------|-------------|
| **Scope** | Convert text from one language to another | Adapt the entire experience for a market |
| **Depth** | Words | Words + visuals + UX + cultural context + legal |
| **Output** | Same content in different language | Market-appropriate content that resonates locally |
| **Cost** | Lower | Higher (but dramatically higher ROI) |
| **Risk of errors** | Linguistic errors | Cultural missteps |

### Market Prioritization Matrix

| Market | Market Size | Competitive Intensity | Localization Effort | Expected ROI | Priority |
|--------|:---:|:---:|:---:|:---:|:---:|
| | High/Med/Low | High/Med/Low | High/Med/Low | High/Med/Low | 1-5 |
| | | | | | |
| | | | | | |

### Localization Readiness Checklist

- [ ] Product supports multi-language (i18n infrastructure)
- [ ] CMS supports localized content
- [ ] Payment processing supports local currencies
- [ ] Legal/compliance reviewed for target markets
- [ ] Customer support available in target language (or planned)
- [ ] SEO strategy defined per market (local domains, hreflang)
- [ ] Brand voice adaptation guidelines created
- [ ] Local market research completed (ICP, competitors, channels)

---

## Part 2: Content Localization Process

### The Localization Workflow

```
Source Content → Market Research → Cultural Adaptation → Translation →
Local QA Review → Legal/Compliance Check → Publish → Monitor Performance
```

### Content Prioritization for Localization

| Priority | Content Type | Why |
|:--------:|-------------|-----|
| 1 | **Website (homepage, pricing, signup)** | First impression, conversion-critical |
| 2 | **Product UI** | User experience, retention |
| 3 | **Email sequences (welcome, onboarding)** | Activation and engagement |
| 4 | **Help center / documentation** | Support deflection, user success |
| 5 | **Sales collateral** | Revenue enablement |
| 6 | **Blog / SEO content** | Organic discovery in local market |
| 7 | **Social media** | Brand awareness and community |

### Cultural Adaptation Guide

| Element | What to Adapt | Example |
|---------|-------------|---------|
| **Imagery** | People, settings, cultural references | Stock photos with local demographics |
| **Colors** | Color meaning varies by culture | White = purity (West), mourning (East Asia) |
| **Humor** | Rarely translates | What's funny in US may offend in Japan |
| **Formality** | Tu vs. Vous, Du vs. Sie | German marketing tends more formal than US |
| **Date/time** | MM/DD/YYYY vs. DD/MM/YYYY | Always localize |
| **Currency** | Local currency display | €, £, ¥ — not just USD |
| **Units** | Metric vs. Imperial | International = metric |
| **Social proof** | Local references and testimonials | "Used by 500 companies in [Country]" |
| **Legal requirements** | Privacy, advertising regulations | GDPR (EU), LGPD (Brazil), PIPL (China) |
| **Payment methods** | Local payment preferences | iDEAL (NL), PIX (Brazil), UPI (India) |

---

## Part 3: Multi-Language SEO

### International SEO Architecture

| Approach | Structure | Pros | Cons |
|----------|---------|------|------|
| **ccTLD** (country code) | example.de, example.fr | Strongest local signal | Expensive, separate domains |
| **Subdirectory** | example.com/de/, example.com/fr/ | Single domain authority | Less local signal |
| **Subdomain** | de.example.com, fr.example.com | Flexible | Google may treat as separate sites |

**Recommendation:** Subdirectory for most companies. ccTLDs for major market commitment.

### Hreflang Implementation

Every localized page needs hreflang tags:

```html
<link rel="alternate" hreflang="en" href="https://example.com/" />
<link rel="alternate" hreflang="de" href="https://example.com/de/" />
<link rel="alternate" hreflang="fr" href="https://example.com/fr/" />
<link rel="alternate" hreflang="x-default" href="https://example.com/" />
```

### Local Keyword Strategy

| Step | Action |
|------|--------|
| 1 | Research local search volume (don't just translate English keywords) |
| 2 | Identify local competitors ranking for target terms |
| 3 | Adapt content to local search intent (may differ from English) |
| 4 | Build local backlinks from market-relevant sites |
| 5 | Monitor local rankings separately from global |

---

## Part 4: Multi-Market Channel Strategy

### Channel Preferences by Region

| Channel | US/Canada | Western Europe | LATAM | APAC | MENA |
|---------|:---:|:---:|:---:|:---:|:---:|
| **LinkedIn** | Very High | High | Medium | Medium | Medium |
| **X/Twitter** | High | Medium | Medium | High (Japan) | Low |
| **Instagram** | High | High | Very High | High | High |
| **TikTok** | High | High | High | Very High | High |
| **WhatsApp** | Medium | High | Very High | High | Very High |
| **WeChat** | Very Low | Low | Low | Very High (China) | Low |
| **LINE** | Very Low | Low | Low | Very High (Japan, Thailand) | Low |
| **Email** | Very High | High | Medium | Medium | Medium |
| **Google Ads** | Very High | Very High | High | High | High |
| **Local search engines** | N/A | N/A | N/A | Baidu (CN), Naver (KR), Yandex (RU) | N/A |

---

## Part 5: Localization Quality Assurance

### QA Checklist Per Market

| Check | Status |
|-------|:------:|
| Translation accuracy (native speaker review) | |
| Cultural appropriateness (no offensive imagery/language) | |
| Legal/compliance (local regulations met) | |
| Date, time, currency formatting | |
| Links and CTAs work correctly | |
| Payment flow functional in local currency | |
| Images and visuals culturally appropriate | |
| SEO elements localized (title, meta, alt text) | |
| Customer support contact information updated | |
| Privacy policy / ToS localized and legally compliant | |

---

## Part 6: Measuring Global Marketing Performance

### Per-Market Dashboard

| Metric | Market A | Market B | Market C | Global |
|--------|:--------:|:--------:|:--------:|:------:|
| Website visitors | | | | |
| Leads | | | | |
| Conversion rate | % | % | % | % |
| Customers | | | | |
| Revenue | $ | $ | $ | $ |
| CAC | $ | $ | $ | $ |
| LTV:CAC | x | x | x | x |
| Market share est. | % | % | % | % |

---

## Integration with Marketing36z OS

| System Component | How This Playbook Connects |
|-----------------|---------------------------|
| **Content Engine** | Produces source content that enters localization workflow |
| **Quality Guardian** | Adds cultural and legal checks for each target market |
| **Channel Operator** | Adapts channel strategy per market preferences |
| **Design System** | Provides market-specific visual adaptation guidelines |
| **Multi-Brand Management** | Each market may operate as a separate "brand" in the system |
| **Enterprise Governance** | Compliance frameworks vary by market |
