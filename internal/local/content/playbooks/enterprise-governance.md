# Playbook: Enterprise Governance

## Sources
- SOC 2, GDPR, CCPA Compliance frameworks
- Brand governance best practices (Interbrand, Brandwatch)
- Marketing36z OS Governance Observer
- Marketing36z OS synthesis

---

## Part 1: Marketing Governance Framework

### What Marketing Governance Covers

| Domain | What It Governs | Risk of Non-Compliance |
|--------|----------------|----------------------|
| **Brand governance** | Logo usage, voice, messaging consistency | Brand dilution, confused market perception |
| **Compliance** | Claims, disclaimers, data privacy, advertising regulations | Legal action, fines, reputation damage |
| **Data governance** | Customer data handling, consent, storage | GDPR/CCPA fines, data breach liability |
| **Content governance** | Approval workflows, version control, archiving | Unauthorized content, stale information |
| **Financial governance** | Budget management, vendor payments, ROI tracking | Budget overruns, audit failures |
| **AI/LLM governance** | Model usage, output review, audit trails | Hallucinations, brand risk, compliance violations |

---

## Part 2: Brand Governance

### Brand Usage Policy

| Rule | Specification | Enforcement |
|------|-------------|------------|
| **Logo usage** | Only approved versions, minimum clear space | Design system playbook |
| **Color palette** | Only brand-approved colors in all materials | Design system playbook |
| **Voice and tone** | Matches approved adjectives, avoids prohibited language | Quality Guardian + policies.yaml |
| **Messaging** | Aligns with approved messaging hierarchy | Strategy Planner review |
| **Co-branding** | Requires written approval + co-branding guidelines | Brand team approval |
| **User-generated** | Must comply with brand guidelines when featuring | Moderation process |

### Brand Approval Matrix

| Content Type | Creator | Reviewer | Approver | SLA |
|-------------|---------|---------|----------|:---:|
| Social media post | Content team | Editor | Auto-approved (after onboarding) | 24h |
| Blog post | Writer | Editor | Content lead | 48h |
| Email campaign | Content team | Marketing manager | Marketing lead | 48h |
| Landing page | Content + Design | Marketing lead | Marketing director | 72h |
| Press release | Communications | Legal | CEO / CMO | 5 days |
| Advertising creative | Creative team | Marketing lead + Legal | CMO | 5 days |
| Product messaging change | Product marketing | CMO + Product | CEO | 1 week |

---

## Part 3: Compliance Governance

### Advertising Compliance Checklist

| Regulation | Applies To | Key Requirement | How We Comply |
|-----------|-----------|----------------|-------------|
| **FTC Act** | All US advertising | Truthful, non-deceptive, evidence-based claims | Quality Guardian forbidden claims check |
| **CAN-SPAM** | Email marketing | Opt-out mechanism, physical address, honest subject lines | ESP compliance features |
| **GDPR** | EU data subjects | Consent for data processing, right to deletion, DPO | Consent management platform |
| **CCPA/CPRA** | California consumers | Opt-out of sale, data access rights | Privacy policy + opt-out mechanism |
| **ADA / WCAG** | Digital content | Accessible to people with disabilities | Accessibility checks in QA |
| **Industry-specific** | Varies (fintech, health, etc.) | Varies | Industry vertical compliance rules |

### Claims and Evidence Standards

| Claim Type | Evidence Required | Example |
|-----------|-----------------|---------|
| **Performance claims** | Data from controlled study or customer results | "Increases conversions by 30%" → needs source |
| **Comparison claims** | Verified, apples-to-apples comparison | "Faster than [competitor]" → needs benchmark |
| **Testimonials** | Real customer, typical results or disclaimer | "We grew 10x" → needs "results not typical" if atypical |
| **Awards / rankings** | Verifiable third-party source | "#1 rated" → needs source citation |
| **Superlatives** | Verifiable data | "Best-in-class" → must have evidence or add qualifier |

### The Compliance Review Process

```
Content created → Self-review (creator checklist) → Peer review (editor) →
Compliance review (if required by type) → Legal review (if high-risk) →
Approval → Publish → Archive
```

---

## Part 4: Data Governance

### Customer Data Handling Policy

| Data Type | Classification | Collection Rules | Storage Rules | Retention |
|-----------|:---:|-------------|------------|:-:|
| **Public** (company name, website) | Low | No restrictions | Standard | Indefinite |
| **Contact** (email, phone) | Medium | Consent required | Encrypted at rest | Until opt-out + 30 days |
| **Behavioral** (page views, clicks) | Medium | Privacy policy disclosure | Anonymized for analytics | 24 months |
| **Financial** (payment, billing) | High | PCI compliance required | Encrypted, restricted access | Per regulations |
| **Sensitive** (SSN, health data) | Critical | Explicit consent only | Encrypted, audited access | Minimum required |

### Data Subject Rights (GDPR/CCPA)

| Right | Response SLA | Process |
|-------|:---:|---------|
| **Right to access** | 30 days | Export customer data on request |
| **Right to deletion** | 30 days | Delete from all systems, confirm completion |
| **Right to correction** | 30 days | Update records as requested |
| **Right to portability** | 30 days | Provide data in machine-readable format |
| **Right to opt-out** | Immediate | Suppress from all marketing communications |

### Consent Management

| Consent Type | How Captured | How Stored | How Enforced |
|-------------|-------------|-----------|-------------|
| **Email marketing** | Opt-in checkbox (no pre-checked) | CRM consent field | ESP suppression lists |
| **Cookie tracking** | Cookie consent banner | Consent management platform | Tag manager conditional loading |
| **Data processing** | Privacy policy acceptance | CRM consent field | Processing only with consent |
| **Third-party sharing** | Explicit opt-in | CRM consent field | Data sharing only with consent |

---

## Part 5: AI/LLM Governance

### LLM Usage Policy for Marketing

| Rule | Specification | Enforcement |
|------|-------------|------------|
| **All LLM outputs reviewed by human** | No auto-publish | Governance Observer audit trail |
| **No PII in prompts** | Redact customer data before sending to models | PII Shield (Governance Observer) |
| **Claims fact-checked** | LLM-generated claims verified before use | Quality Guardian review |
| **Audit trail maintained** | Every LLM call logged with hash-chain | Governance Observer |
| **Cost controls enforced** | Per-request, per-session, per-day caps | Policy Engine |
| **Brand voice validated** | Output checked against voice adjectives | Brand voice scanner |

### AI Content Disclosure

| Context | Disclosure Required? | How |
|---------|:---:|-----|
| Blog post (AI-assisted) | Best practice | "This article was written with AI assistance and reviewed by [Name]" |
| Email marketing | Not required by most regulations | Internal policy decision |
| Paid advertising | Emerging regulation | Monitor jurisdiction-specific requirements |
| Customer-facing chatbot | Yes (in most jurisdictions) | "You're chatting with an AI assistant" |
| Internal documents | Not required | Optional |

---

## Part 6: Content Governance

### Content Lifecycle Management

| Stage | Activity | Owner |
|-------|---------|-------|
| **Creation** | Produced following brief and guidelines | Content team |
| **Review** | Quality, compliance, and brand checks | Editor + compliance |
| **Approval** | Authorized for publication | Approver (per matrix) |
| **Publication** | Published to designated channels | Distribution |
| **Monitoring** | Performance tracked, feedback collected | Analytics |
| **Update** | Refreshed when stale or inaccurate (annual minimum) | Content team |
| **Archival** | Retired content removed from active distribution | Content ops |
| **Deletion** | Permanently removed (if required) | Content ops + legal |

### Version Control Standards

| Standard | Specification |
|----------|-------------|
| **Naming convention** | `[document-name]-v[X.Y]-[YYYY-MM-DD]` |
| **Major version (X)** | Significant content change or strategic shift |
| **Minor version (Y)** | Updates, corrections, refreshes |
| **Changelog** | Every version change documented with who, what, when, why |
| **Single source of truth** | One authoritative location per document |

---

## Part 7: Audit Trail and Reporting

### Governance Audit Report (Quarterly)

| Section | Content |
|---------|---------|
| **Brand compliance** | % of content passing brand voice checks |
| **Claims compliance** | % of claims with supporting evidence |
| **Data privacy** | Data subject requests fulfilled, consent rates |
| **Content freshness** | % of content reviewed/updated in last 12 months |
| **LLM governance** | Audit trail integrity, PII incidents, cost compliance |
| **Incidents** | Any compliance violations, how they were resolved |
| **Recommendations** | Process improvements for next quarter |

### The Governance Scorecard

| Dimension | Score (1-5) | Evidence | Action Needed |
|-----------|:-----------:|----------|-------------|
| Brand consistency | | | |
| Claims compliance | | | |
| Data privacy compliance | | | |
| Content governance maturity | | | |
| AI/LLM governance | | | |
| Audit trail completeness | | | |
| **Overall governance health** | | | |

---

## Integration with Marketing36z OS

| System Component | How This Playbook Connects |
|-----------------|---------------------------|
| **Governance Observer** | Operationalizes AI/LLM governance with audit trails and policy enforcement |
| **Quality Guardian** | Enforces brand and compliance rules on every deliverable |
| **Policies.yaml** | Configuration file for all governance rules |
| **Design System** | Visual brand governance standards |
| **Content Operations** | Content lifecycle management at scale |
| **Crisis Communications** | Governance failures may trigger crisis response |
