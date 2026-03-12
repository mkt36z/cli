# Playbook: Conversion Rate Optimization (CRO) — Complete System

## Sources
- MarketingSkills CRO skill stack (page-cro, signup-flow-cro, form-cro, popup-cro, paywall-upgrade-cro, onboarding-cro)
- SEO Machine CRO scoring agents
- Marketing36z OS synthesis

---

## The Core Principle

> Every page, form, and interaction is either building momentum toward conversion or introducing friction that kills it. CRO is not about tricks — it is about removing every unnecessary obstacle between a visitor and the value you deliver.

This playbook covers six CRO domains in order of impact: page-level optimization, signup flows, forms, popups, paywall/upgrade screens, and onboarding. Each section provides frameworks you can apply immediately, with experiment ideas at the end to systematize your testing.

---

## Part 1: Page CRO (7-Dimension Analysis)

### The Impact Hierarchy

Not all page elements matter equally. Optimize in this order — fixing #1 before worrying about #7:

| Rank | Dimension | Impact Level | What It Controls |
|:----:|-----------|:------------:|-----------------|
| 1 | **Value Proposition Clarity** | Critical | Can a visitor articulate what you do and why it matters in 5 seconds? |
| 2 | **Headline Effectiveness** | Critical | Does the headline stop the scroll and pull them into the page? |
| 3 | **CTA Placement / Copy / Hierarchy** | High | Is there one clear next step, and is it obvious? |
| 4 | **Visual Hierarchy** | High | Does the page guide the eye from headline to proof to CTA? |
| 5 | **Trust Signals** | Medium | Are there enough proof points to overcome skepticism? |
| 6 | **Objection Handling** | Medium | Does the page preemptively answer the reasons people say "not yet"? |
| 7 | **Friction Points** | Medium | Are there unnecessary steps, confusing elements, or broken flows? |

### The 5-Second Test

Load your page and ask someone who has never seen it: **"What does this company do, and what are you supposed to do next?"** If they cannot answer both questions in five seconds, your page fails at the most basic level.

How to run it:
1. Show the page to 5 people for 5 seconds each
2. Hide the page and ask: "What does this company/product do?"
3. Ask: "What would you click first?"
4. If fewer than 4/5 get both right, rewrite the above-the-fold content before optimizing anything else

### Strong Headline Patterns

| Pattern | Example | Best For |
|---------|---------|----------|
| **Outcome + Timeframe** | "Build your website in 10 minutes" | SaaS, tools |
| **Eliminate the Pain** | "Stop losing deals to slow follow-ups" | B2B |
| **Social Proof Lead** | "Join 50,000 marketers who..." | Community, newsletters |
| **Specific Number** | "Increase revenue 37% with one change" | Consultants, agencies |
| **Question** | "What if your landing page actually converted?" | Awareness |
| **Before/After** | "From scattered notes to organized knowledge" | Productivity tools |

**Test**: Always pair a new headline with a supporting subheadline that adds specificity. The headline grabs attention; the subheadline explains.

### CTA Hierarchy: One Primary Action Per Page

Every page should have exactly ONE primary call-to-action. Everything else is secondary or tertiary.

| Level | Style | Usage |
|-------|-------|-------|
| **Primary CTA** | High-contrast button, large, above the fold | The ONE thing you want them to do |
| **Secondary CTA** | Outlined or text link, smaller | Alternative path (e.g., "Watch demo" when primary is "Start free trial") |
| **Tertiary CTA** | Text link only | Navigation or informational (e.g., "Learn more about pricing") |

**CTA Copy Formula**: `[Action Verb] + [What They Get]`
- Weak: "Submit" / "Click here" / "Learn more"
- Strong: "Start my free trial" / "Get the playbook" / "See it in action"

### Page-Specific Frameworks

#### Homepage
| Section | Purpose | Key Element |
|---------|---------|-------------|
| Hero | Communicate value proposition | Headline + subheadline + primary CTA |
| Social proof bar | Build instant credibility | Logos, user count, or press mentions |
| How it works | Reduce uncertainty | 3-step visual process |
| Benefits | Connect features to outcomes | 3-4 benefit blocks with icons |
| Testimonials | Prove results | Quotes with names, photos, and specifics |
| Final CTA | Capture remaining interest | Repeat primary CTA with urgency or reassurance |

#### Landing Page
| Element | Rule |
|---------|------|
| Headline | Match the ad/link copy that brought them here (message match) |
| Navigation | Remove it — one page, one action, no escape routes |
| Form length | Match the value of the offer (free checklist = email only; demo = 4-5 fields) |
| Social proof | Place it near the CTA to reinforce the decision at the moment of action |
| Page length | Short for low-commitment offers, long for high-commitment offers |

#### Pricing Page
| Element | Rule |
|---------|------|
| Recommended plan | Visually highlight one plan (the one you want most people to choose) |
| Feature comparison | Show a full comparison table below the cards for detail-seekers |
| FAQ | Place below pricing to handle objections without cluttering the cards |
| Annual vs monthly | Default to annual; show savings as a percentage |
| Social proof | "10,000+ teams use [Product]" near the top; specific testimonials near the CTA |
| Free tier | Position it as limited, not as the default choice |

#### Feature Page
| Element | Rule |
|---------|------|
| Problem lead | Start with the pain the feature solves, not the feature itself |
| Demo/visual | Show the feature in action — screenshot, GIF, or short video |
| Use cases | Show 2-3 specific scenarios where this feature changes outcomes |
| CTA | "Try [feature name] free" is stronger than a generic signup CTA |

---

## Part 2: Signup Flow CRO

### The Golden Rule

**Show value before asking for commitment.** The more someone understands what they will get, the more friction they will tolerate to get it.

### Field Priority Matrix

| Priority | Fields | Rationale |
|----------|--------|-----------|
| **Essential** | Email, Password | Minimum viable account creation |
| **Often needed** | Full name | Personalization, communication |
| **Deferrable** | Company name, role, phone | Collect during onboarding or first use, not signup |
| **Never at signup** | Address, billing, team size | Massive friction with zero immediate value to the user |

**Rule of thumb**: Every field you add at signup costs you 5-10% of potential signups. Defer everything that is not required to create the account.

### Progressive Commitment Pattern

Break signup into small, escalating steps. Each step is easier to say "yes" to because they have already invested.

```
Step 1: Email only
   → "Enter your email to get started"
   → Lowest possible commitment
   → Validates interest before asking for more

Step 2: Password + Name
   → "Create your account"
   → They have already given email — sunk cost makes this feel natural

Step 3: Customization
   → "Tell us about your goals"
   → Now they are a user — this feels like setup, not signup
```

**Key insight**: The completion rate of Step 1 determines your funnel size. Step 2 and 3 completion rates are almost always 70%+ once someone starts.

### Social Auth Options

| Context | Recommended Providers | Why |
|---------|----------------------|-----|
| **B2C** | Google, Apple, Facebook | Maximum coverage for consumer accounts |
| **B2B** | Google, Microsoft, SSO (SAML) | Matches enterprise identity providers |
| **Both** | Google + one contextual option | Google alone covers 80%+ of social auth usage |

**Placement**: Social auth buttons go ABOVE the email/password form. Most users prefer them. Do not bury them below.

### Field-Level Optimization

| Optimization | Implementation | Impact |
|-------------|----------------|--------|
| **Email typo detection** | Detect `gmial.com`, `gmal.com`, `yahooo.com` and suggest correction | Prevents 2-5% of failed accounts |
| **Password show toggle** | Eye icon to reveal password text | Reduces password errors by 30%+ |
| **Single "Full name" field** | One field instead of separate First/Last | Faster, simpler, split server-side if needed |
| **Smart keyboard** | `type="email"` for email, `type="tel"` for phone | Eliminates mobile keyboard frustration |
| **Autofill support** | Correct `autocomplete` attributes on every field | Lets browsers/managers fill fields instantly |
| **Inline validation** | Validate format on blur, not on submit | Catches errors before they become frustrating |

### Single-Step vs Multi-Step Decision

| Choose Single-Step When... | Choose Multi-Step When... |
|---------------------------|--------------------------|
| 3 or fewer fields | 5+ fields required |
| Low-commitment offer (free tool, newsletter) | High-commitment offer (demo, trial with details) |
| Mobile-first audience | Desktop-primary audience |
| Speed is the top priority | You need qualifying information |

### Mobile Signup Requirements

| Requirement | Specification |
|-------------|--------------|
| Touch targets | 44px minimum height and width for all tappable elements |
| Keyboard type | `inputmode="email"` for email, `inputmode="numeric"` for codes |
| Autofill | `autocomplete="email"`, `autocomplete="new-password"`, etc. |
| Form position | No horizontal scrolling; form fits within viewport width |
| Error display | Errors appear directly below the field, not in a banner at the top |
| Progress indicator | If multi-step, show step count ("Step 1 of 3") |

### Signup Flow Metrics

| Metric | Definition | Benchmark |
|--------|-----------|:---------:|
| **Form start rate** | Visitors who interact with the first field / total visitors | 30-60% |
| **Completion rate** | Successful signups / form starts | 60-80% |
| **Field-level drop-off** | % who abandon at each specific field | Varies — flag any field with >15% drop |
| **Error rate** | Form submissions with errors / total submissions | <10% is good; >20% means the form is broken |
| **Time to complete** | Average seconds from first field interaction to submission | <60 seconds for simple forms |

---

## Part 3: Form CRO

### The Field Cost Principle

Every field has a cost measured in lost conversions. Know the price you are paying:

| Field Count | Conversion Impact | When Acceptable |
|:-----------:|:-----------------:|-----------------|
| 1-3 fields | Baseline (highest conversion) | Email capture, newsletter, free tool signup |
| 4-6 fields | 10-25% reduction from baseline | Lead gen, demo requests, trial signups |
| 7-10 fields | 25-50% reduction from baseline | Enterprise contact forms, detailed applications |
| 11+ fields | 50%+ reduction from baseline | Almost never acceptable — break into steps |

**The test**: For every field, ask: "Would I rather have this data or 10% more leads?" If the answer is more leads, remove the field.

### Form Design Rules

| Rule | Why |
|------|-----|
| **Labels always visible** — never use placeholder text as the only label | Placeholders disappear on focus, forcing users to remember what the field is |
| **Single column layout** | Multi-column forms create confusing tab order and are worse on mobile |
| **Logical field order** | Name → Email → Phone → Message. Match the user's mental model |
| **Group related fields** | Use visual grouping for sections (e.g., "Contact Info" vs "Project Details") |
| **Multi-step for 5+ fields** | Break long forms into 2-3 steps with a progress indicator |
| **Smart defaults** | Pre-select the most common option in dropdowns; pre-fill known data |

### Submit Button Formula

**Pattern**: `[Action Verb] + [What They Get]`

| Weak | Strong | Why It Works |
|------|--------|-------------|
| Submit | Get My Free Guide | Tells the user what happens next |
| Send | Request My Demo | Frames the action as something they receive |
| Register | Create My Account | Personal ("my") increases ownership feeling |
| Sign Up | Start My Free Trial | Emphasizes the benefit, not the obligation |
| Download | Get Instant Access | Creates urgency and immediacy |

### Error Handling That Preserves Momentum

| Principle | Implementation |
|-----------|---------------|
| **Inline validation** | Validate each field on blur (when the user leaves the field), not on submit |
| **Preserve all input** | Never clear the form on error — users who have to re-type everything will leave |
| **Focus on error field** | Automatically scroll to and focus the first field with an error |
| **Specific error messages** | "Please enter a valid email address" not "Invalid input" |
| **Visual indicators** | Red border + error text below the field; green check for valid fields |
| **Prevent premature validation** | Do not show errors while the user is still typing in a field |

### Advanced Form Patterns

| Pattern | Use Case | Implementation |
|---------|----------|---------------|
| **Conditional fields** | Show fields only when relevant (e.g., "Company size" appears only if "Business" is selected) | Reduces perceived form length |
| **Save progress** | For multi-step forms, save input so users can return | Use `localStorage` or server-side save |
| **Conversational forms** | One question at a time, chat-like interface | Works well for qualifying leads; novelty drives completion |
| **Pre-qualification** | Ask a simple yes/no before showing the full form | Filters out unqualified leads early |

---

## Part 4: Popup CRO

### Trigger Ranking by Conversion Rate

Not all popup triggers perform equally. Choose your trigger based on context and expected performance:

| Rank | Trigger Type | Typical Conversion Rate | Best Use Case |
|:----:|-------------|:----------------------:|---------------|
| 1 | **Click-triggered** | 10%+ | User clicks a link/button to open (e.g., "Get the checklist") |
| 2 | **Exit intent** | 3-10% | Desktop visitors moving cursor toward browser close/back |
| 3 | **Scroll-based** | 2-5% | User has scrolled 50-70% of page, proving engagement |
| 4 | **Time-based** | 1-3% | 30-60 seconds on page minimum |

**Critical rule**: NEVER trigger a time-based popup in under 30 seconds. A 5-second popup is the fastest way to destroy trust and increase bounce rate.

### Popup Types and When to Use Them

| Type | Purpose | Best Trigger | Example |
|------|---------|:------------:|---------|
| **Email capture** | Build list | Scroll 50% or exit intent | "Get our weekly growth newsletter" |
| **Lead magnet** | Capture qualified leads | Click-triggered or scroll-based | "Download the 2024 Pricing Benchmarks Report" |
| **Discount** | Reduce cart abandonment | Exit intent on cart/checkout | "Wait — get 10% off your first order" |
| **Exit intent** | Save abandoning visitors | Exit intent | "Before you go — did you have a question?" |
| **Announcement** | Promote events, launches | Time-based (30s+) or banner | "New feature: AI-powered analytics" |
| **Slide-in** | Low-intrusion engagement | Scroll 70%+ | Small corner popup with relevant CTA |

### CTA Copy: First-Person Voice

Popups convert better when the CTA is written in the user's voice, not the company's voice:

| Third-Person (Weaker) | First-Person (Stronger) | Lift |
|-----------------------|------------------------|:----:|
| "Get Your Discount" | "Get My Discount" | +10-20% |
| "Start Your Trial" | "Start My Free Trial" | +10-15% |
| "Download the Guide" | "Send Me the Guide" | +5-15% |

**Dismiss text**: The "no thanks" option should be neutral, never guilt-trip. "No thanks" or "Maybe later" — NOT "No, I don't want to grow my business."

### Frequency Rules

| Rule | Setting | Why |
|------|---------|-----|
| **Session limit** | Maximum once per session | More than one popup per session trains users to close without reading |
| **Re-show delay** | 7-30 days before re-showing to the same visitor | Respects the user's prior "no" |
| **Exclude converted** | Never show a popup to someone who already converted on that offer | Feels broken and annoying |
| **Page limit** | Only trigger on high-intent pages, not every page | Reduces popup fatigue |
| **New visitor priority** | Show list-building popups to new visitors, upgrade popups to returning visitors | Match the popup to the visitor's stage |

### Google Compliance (Mobile)

Google penalizes intrusive interstitials on mobile. These will hurt your SEO:
- Full-screen popups that appear before the user interacts with content
- Standalone interstitials that must be dismissed before accessing the page
- Above-the-fold layouts where the popup takes up the entire visible area

**Safe alternatives for mobile**:
- Small banner at the top or bottom of the screen
- Inline CTAs within the content
- Tab-triggered or scroll-triggered slide-ins that cover less than 30% of the screen
- App install banners using the browser's native prompt

### Accessibility Requirements

| Requirement | Implementation |
|-------------|---------------|
| **Keyboard navigation** | User can Tab through popup elements and close with Escape key |
| **Focus trap** | When popup is open, Tab cycles within the popup, not behind it |
| **Screen reader** | Popup has `role="dialog"`, `aria-label`, and `aria-modal="true"` |
| **Close button** | Visible, labeled, and reachable by keyboard (minimum 44x44px) |
| **Background** | Semi-transparent overlay; content behind is inert (`aria-hidden="true"`) |
| **Restored focus** | When popup closes, focus returns to the element that triggered it |

---

## Part 5: Paywall / Upgrade CRO

### The Value Before Ask Principle

A paywall should feel like a natural next step, not a wall. The user must have already experienced enough value to understand what they are paying for.

```
Wrong order:  Sign up → Hit paywall → Wonder what the product does → Leave
Right order:  Sign up → Use product → Experience value → Hit limit → Upgrade feels obvious
```

### Four Trigger Types

| Trigger | How It Works | Example | Best For |
|---------|-------------|---------|----------|
| **Feature gate** | Free users see the feature exists but cannot access it | "Upgrade to unlock advanced analytics" | Products with clear free/paid feature split |
| **Usage limit** | Free users can use the feature up to a limit | "You've used 8 of 10 free reports this month" | Usage-based products |
| **Trial expiration** | Full access for a limited time, then restricted | "Your 14-day trial ends in 3 days" | Products that need time to show value |
| **Time-based** | Prompt after a certain period of free use | "You've been using [Product] for 30 days — ready to go Pro?" | Products with gradual value discovery |

### Paywall Screen Components (7 Elements)

Every effective paywall screen contains these elements in this order:

| # | Element | Purpose | Example |
|:-:|---------|---------|---------|
| 1 | **Headline** | Name the value, not the restriction | "Unlock unlimited reports" not "You've hit your limit" |
| 2 | **Value demo** | Show what they get, visually | Screenshot or preview of the premium feature |
| 3 | **Feature comparison** | Show free vs paid side-by-side | Simple 2-column table, checkmarks for paid |
| 4 | **Pricing** | Clear, simple, no hidden costs | "$12/mo billed annually" with monthly option visible |
| 5 | **Social proof** | Others made this decision and are happy | "Used by 5,000+ teams" or a specific testimonial |
| 6 | **CTA** | Strong, specific action button | "Start Pro Plan" or "Upgrade Now" |
| 7 | **Escape hatch** | Always provide a way to decline without guilt | "Continue with free plan" or "Maybe later" |

### Three Paywall Patterns

#### Pattern 1: Feature Lock
```
[User clicks locked feature]
   → Modal: "This feature is available on Pro"
   → Show feature preview + comparison
   → CTA: "Upgrade to Pro" | Escape: "Continue with Free"
```
**Best when**: The locked feature is visibly valuable and the user clicked on it intentionally.

#### Pattern 2: Usage Limit
```
[User hits usage limit]
   → In-context message: "You've used 10/10 free exports"
   → Show usage bar (visual) + what Pro includes
   → CTA: "Get unlimited exports" | Escape: "I'll wait until next month"
```
**Best when**: The user has already gotten value from the feature and wants more.

#### Pattern 3: Trial Expiration
```
[Trial approaching end]
   → Day 11: Subtle banner — "3 days left in your trial"
   → Day 13: Email — summary of value received + what they will lose
   → Day 14: Modal — "Your trial ends today" + comparison of free vs paid
   → Day 15+: Restricted access with clear upgrade path
```
**Best when**: The product requires extended use to demonstrate value.

### Timing Rules

| Rule | Rationale |
|------|-----------|
| **Show after a value moment** | Prompt the upgrade immediately after the user accomplishes something meaningful |
| **Never during onboarding** | Let them set up and experience the product first — a paywall during setup feels like a bait-and-switch |
| **Limit to 1 prompt per session** | More than one upgrade prompt per session feels aggressive |
| **Escalate gradually** | Start with subtle (banner), move to moderate (inline), then direct (modal) over days/weeks |
| **Time of day** | Show upgrade prompts during working hours when users are in decision-making mode |

### Anti-Patterns to Avoid

| Anti-Pattern | Why It Fails |
|-------------|-------------|
| **Hiding the close button** | Users feel trapped — they will leave the product, not upgrade |
| **Guilt-trip copy** | "Don't you want to succeed?" breeds resentment, not conversions |
| **Blocking critical flows** | If a user is mid-task and you block them, they lose work and lose trust |
| **Dark patterns** | Pre-checked annual billing, confusing plan names, hidden fees — these generate chargebacks and churn |
| **No free option** | Always provide a way to continue free, even if limited — forced conversion has high churn |
| **Generic paywall** | Show the same screen regardless of context — personalize based on what triggered it |

---

## Part 6: Onboarding CRO

### Time-to-Value Optimization

The single most important onboarding metric: **How quickly does a new user experience the core value of your product?**

Every minute between signup and the "aha moment" is a minute they might leave. Your onboarding exists to compress this gap.

```
Signup → [Minimize this gap] → Aha Moment → Retained User
```

### Defining the "Aha Moment"

The aha moment is the specific action or experience where a user first understands the product's value. Find it by comparing retained vs churned users:

| Step | Action |
|------|--------|
| 1 | Pull a list of users who are still active at Day 30 |
| 2 | Pull a list of users who churned within the first 7 days |
| 3 | Compare their Day 1 behavior: what did retained users do that churned users did not? |
| 4 | Identify the action with the highest correlation to retention |
| 5 | That action is your aha moment — your onboarding should drive every user to it |

**Examples**:
- Slack: Sending 2,000 messages as a team
- Dropbox: Putting a file in a shared folder
- Twitter/X: Following 30 people
- Zoom: Hosting or joining a call

### Onboarding Checklist Design

A visible checklist is one of the most effective onboarding patterns. Rules:

| Rule | Specification |
|------|--------------|
| **Item count** | 3-7 items. Fewer than 3 feels trivial; more than 7 feels overwhelming |
| **Order by value** | Put the highest-value action first so users hit the aha moment early |
| **Progress bar** | Show visual progress (e.g., "3 of 5 complete") — the Zeigarnik effect drives completion |
| **Pre-complete one item** | Mark "Create account" as already done — starting at 1/5 feels better than 0/5 |
| **Celebrate completion** | Confetti, congratulations message, or badge when the checklist is done |
| **Dismissible** | Users can hide the checklist but easily find it again |
| **Persistent** | The checklist survives page refreshes and session changes |

### Empty States as Onboarding

Every screen a new user sees starts empty. Empty states are onboarding opportunities, not dead ends:

| Empty State Element | Purpose |
|--------------------|---------|
| **Illustration** | Makes the empty page feel intentional, not broken |
| **Explanation** | "This is where your projects will appear" — tell them what this space is for |
| **Single clear CTA** | "Create your first project" — one action to fill the space |
| **Sample data** | Optional: show example content they can explore to understand the feature |

### Stalled User Detection and Re-engagement

| Signal | Definition | Response |
|--------|-----------|----------|
| **Signed up, never activated** | Created account but did not complete the aha moment action | Trigger email sequence: Day 1, Day 3, Day 7 with progressively direct CTAs |
| **Started onboarding, stopped** | Completed some checklist items but not all | In-app prompt on next visit: "Pick up where you left off" |
| **Used once, went silent** | Completed onboarding but has not returned in 3+ days | Email with value reminder: "Your [thing they created] is waiting for you" |
| **Active but not activated** | Logging in but not doing the key retention action | In-app tooltip pointing to the key feature they have not used |

### Onboarding Metrics

| Metric | Definition | Target |
|--------|-----------|:------:|
| **Activation rate** | % of signups who complete the aha moment action | 40-60% is good; >60% is excellent |
| **Time to activation** | Median time from signup to aha moment | Shorter is always better; benchmark against your own trend |
| **Day 1 retention** | % of signups who return the next day | >25% for B2C; >40% for B2B |
| **Day 7 retention** | % of signups who return within 7 days | >15% for B2C; >30% for B2B |
| **Day 30 retention** | % of signups who are active at 30 days | >10% for B2C; >25% for B2B |
| **Checklist completion rate** | % of users who complete all onboarding steps | >50% is the target |
| **Onboarding drop-off by step** | % who abandon at each checklist item | Flag any step with >25% drop-off for redesign |

---

## Part 7: CRO Experiment Ideas

### Page CRO Experiments

| # | Experiment | Expected Impact | Effort |
|:-:|-----------|:--------------:|:------:|
| 1 | Rewrite hero headline using Outcome + Timeframe pattern | High | Low |
| 2 | Add social proof bar (logos or user count) below the hero | Medium | Low |
| 3 | Change CTA copy from generic ("Sign Up") to specific ("Start My Free Trial") | High | Low |
| 4 | Remove navigation on landing pages | Medium | Low |
| 5 | Add a "How It Works" 3-step section below the fold | Medium | Medium |
| 6 | Test long-form vs short-form landing page for high-ticket offer | High | High |
| 7 | Add video demo to hero section | Medium | Medium |
| 8 | Move testimonials closer to the primary CTA | Medium | Low |

### Signup Flow Experiments

| # | Experiment | Expected Impact | Effort |
|:-:|-----------|:--------------:|:------:|
| 1 | Reduce signup to email-only first step | High | Medium |
| 2 | Add Google social auth as the primary signup option | High | Medium |
| 3 | Add email typo detection (suggest corrections for common misspellings) | Low-Medium | Low |
| 4 | Add password show/hide toggle | Low | Low |
| 5 | Replace First Name + Last Name with single "Full Name" field | Low-Medium | Low |
| 6 | Test single-step vs multi-step signup | High | High |
| 7 | Add progress indicator to multi-step flow | Medium | Low |
| 8 | Show value proposition summary alongside the signup form | Medium | Low |

### Form CRO Experiments

| # | Experiment | Expected Impact | Effort |
|:-:|-----------|:--------------:|:------:|
| 1 | Remove one non-essential field from your highest-traffic form | High | Low |
| 2 | Change submit button from "Submit" to action-specific copy | Medium | Low |
| 3 | Add inline validation (validate on blur instead of on submit) | Medium | Medium |
| 4 | Switch from multi-column to single-column layout | Medium | Low |
| 5 | Add visible labels above fields (not just placeholder text) | Medium | Low |
| 6 | Test conversational (one-question-at-a-time) form format | High | High |
| 7 | Add conditional fields to reduce visible form length | Medium | Medium |

### Popup CRO Experiments

| # | Experiment | Expected Impact | Effort |
|:-:|-----------|:--------------:|:------:|
| 1 | Switch from time-based to scroll-based trigger (50% scroll depth) | Medium | Low |
| 2 | Test exit-intent popup on top 5 highest-exit pages | High | Medium |
| 3 | Change CTA to first-person voice ("Get My Guide" vs "Get Your Guide") | Medium | Low |
| 4 | Replace full-screen popup with slide-in on mobile | Medium | Low |
| 5 | Test lead magnet popup vs discount popup on blog content | High | Medium |
| 6 | Reduce popup frequency from every visit to once per 14 days | Medium | Low |
| 7 | Add a click-triggered popup on key content pages | High | Low |

### Paywall / Upgrade Experiments

| # | Experiment | Expected Impact | Effort |
|:-:|-----------|:--------------:|:------:|
| 1 | Show paywall after value moment instead of at feature click | High | Medium |
| 2 | Add social proof to paywall screen (user count or testimonial) | Medium | Low |
| 3 | Test feature-lock vs usage-limit paywall pattern | High | High |
| 4 | Add feature preview/demo on paywall screen | Medium | Medium |
| 5 | Change paywall headline from restriction-focused to value-focused | Medium | Low |
| 6 | Add "Continue with free plan" escape hatch (if missing) | Medium | Low |
| 7 | Implement gradual escalation (banner → inline → modal over 7 days) | High | High |

### Onboarding Experiments

| # | Experiment | Expected Impact | Effort |
|:-:|-----------|:--------------:|:------:|
| 1 | Add onboarding checklist with 5 items ordered by value | High | Medium |
| 2 | Pre-complete "Create account" step to show progress from Day 1 | Low-Medium | Low |
| 3 | Replace empty states with guided CTAs and example content | High | Medium |
| 4 | Add stalled user email sequence (Day 1, 3, 7) | High | Medium |
| 5 | Test welcome tour (guided walkthrough) vs checklist | Medium | High |
| 6 | Add progress bar to onboarding flow | Medium | Low |
| 7 | Identify and optimize the single highest-drop-off onboarding step | High | Medium |
| 8 | Add celebration animation on checklist completion | Low | Low |

---

## Part 8: CRO Scoring Stack — Quantitative Conversion Analysis

*Source: SEO Machine CRO scoring agents — CTA + trust + above-fold analyzers for quantitative conversion optimization*

### How the CRO Scoring Stack Works

Every page receives three independent scores (0-100) that combine into a composite **Conversion Readiness Score**. Use this to prioritize optimization efforts and benchmark progress.

### Score 1: CTA Effectiveness Score (0-100)

Evaluate every CTA on the page against these criteria:

| Factor | Weight | Scoring |
|--------|:------:|---------|
| **Visibility** | 20% | Above fold (10), Below fold with scroll indicator (6), Buried (2) |
| **Copy clarity** | 20% | Action verb + value outcome (10), Action verb only (6), Vague/generic (2) |
| **Visual contrast** | 15% | High contrast from background (10), Moderate (6), Blends in (2) |
| **Specificity** | 15% | "Start your free 14-day trial" (10), "Get started" (6), "Submit" (2) |
| **Urgency/motivation** | 10% | Time-bound or loss-framed (10), Benefit-framed (6), No motivation (2) |
| **Placement frequency** | 10% | After each value section (10), Top and bottom only (6), Single CTA (3) |
| **Size / tap target** | 10% | 48px+ height, full-width mobile (10), Standard (6), Small/hard to tap (2) |

**Grading:** A (90-100) = Ship it | B (75-89) = Minor tweaks | C (60-74) = Rewrite CTA | D (<60) = Full CTA redesign

### Score 2: Trust Signal Score (0-100)

| Factor | Weight | Scoring |
|--------|:------:|---------|
| **Social proof quantity** | 20% | 5+ testimonials/logos (10), 2-4 (6), 0-1 (2) |
| **Social proof quality** | 15% | Named person + title + photo + specific result (10), Name + quote (6), Anonymous (2) |
| **Quantitative proof** | 15% | Specific numbers: "2,847 customers" (10), Rounded: "2,000+" (6), None (2) |
| **Authority signals** | 15% | Press logos, certifications, awards (10), Industry mentions (6), None (2) |
| **Risk reversal** | 15% | Money-back guarantee + free trial + no CC (10), One risk reversal (6), None (2) |
| **Security indicators** | 10% | SSL + trust badges + privacy link (10), SSL only (6), None (2) |
| **Real-time proof** | 10% | Live user count, recent signups, activity feed (10), Static proof (5), None (2) |

**Grading:** A (90-100) = High trust | B (75-89) = Adequate | C (60-74) = Add proof | D (<60) = Trust deficit — fix before spending on traffic

### Score 3: Above-the-Fold Score (0-100)

What visitors see before scrolling determines 80% of bounce/stay decisions.

| Factor | Weight | Scoring |
|--------|:------:|---------|
| **Headline clarity** | 25% | Passes 5-second test (10), Somewhat clear (6), Confusing/clever (2) |
| **Value proposition** | 20% | Unique benefit + differentiation visible (10), Generic benefit (6), Feature-focused (2) |
| **Visual hierarchy** | 15% | Clear eye path: headline → subhead → CTA (10), Cluttered but readable (6), Chaotic (2) |
| **CTA visible** | 15% | Primary CTA fully visible above fold (10), Partially visible (6), Below fold (2) |
| **Load speed** | 10% | LCP <2.5s (10), 2.5-4s (6), >4s (2) |
| **Mobile rendering** | 10% | Perfect mobile layout (10), Adequate (6), Broken/cramped (2) |
| **Whitespace** | 5% | Balanced, scannable (10), Slightly crowded (6), Wall of text (2) |

**Grading:** A (90-100) = Optimized | B (75-89) = Good | C (60-74) = Redesign above fold | D (<60) = Complete above-fold rebuild

### Composite Conversion Readiness Score

```
Conversion Readiness Score = (CTA Score × 0.40) + (Trust Score × 0.35) + (Above-Fold Score × 0.25)
```

| Composite Score | Interpretation | Action |
|:--------------:|----------------|--------|
| **90-100** | Conversion-optimized | Monitor and run incremental A/B tests |
| **75-89** | Solid foundation | Optimize the lowest-scoring component |
| **60-74** | Underperforming | Prioritize CRO sprint — fix weakest scorer first |
| **40-59** | Conversion blocker | Stop driving traffic until scores improve |
| **<40** | Broken | Full page redesign required before any promotion |

### Using the Scoring Stack

1. **Score every landing page** before launching paid campaigns — minimum composite score of 70 to justify ad spend
2. **Re-score monthly** as part of the CRO review cycle
3. **Track score trends** — improving scores should correlate with improving conversion rates
4. **Benchmark by page type** — homepage, pricing page, feature pages, and blog posts have different score expectations

---

## Integration with Marketing36z OS

| System Component | How This Playbook Connects |
|-----------------|---------------------------|
| **Orchestrator** | Routes CRO tasks to the correct sub-skill based on the conversion point being optimized |
| **Strategy Planner** | CRO priorities inform quarterly growth plans — optimize existing funnels before adding new traffic |
| **Content Engine** | Headline and CTA copy produced by Content Engine are evaluated against CRO patterns in this playbook |
| **Channel Operator** | Landing pages, email signup forms, and ad destinations follow the page CRO and form CRO frameworks |
| **Quality Guardian** | Reviews all pages and forms for CRO compliance before launch — checks field count, CTA copy, mobile requirements |
| **Experiment Framework** | Every CRO change from Part 7 is tested using the Experiment Design Framework playbook process |
| **Growth & Retention** | Onboarding CRO (Part 6) feeds directly into retention loops — activated users become growth loop participants |
| **SEO Machine** | CRO scoring agents evaluate page-level conversion factors alongside SEO factors for holistic page quality |
| **Analytics System** | All CRO metrics (signup completion rate, activation rate, paywall conversion) are tracked in the measurement framework |
| **Pricing Strategy** | Paywall and upgrade CRO (Part 5) aligns with pricing playbook — pricing page design and plan presentation follow both |
