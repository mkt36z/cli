# ML-Based Topic Clustering — Advanced SEO Content Strategy

*Source: SEO Machine TF-IDF + K-means topic clustering system (Roadmap item C21)*

> Move beyond manual keyword research. Use machine learning to discover natural topic clusters, identify content gaps, and build topical authority systematically.

---

## Overview

Topic clustering uses unsupervised machine learning to group related keywords and content into natural clusters — the same way search engines organize topics. This gives you a structural advantage: instead of guessing at content silos, you build them from data.

```
┌─────────────────────────────────────────────────┐
│ TOPIC CLUSTERING PIPELINE                        │
│                                                   │
│  Keywords ──► TF-IDF ──► K-Means ──► Clusters    │
│  (raw)      (vectors)   (groups)    (content     │
│                                      strategy)   │
└─────────────────────────────────────────────────┘
```

---

## Part 1: Data Collection

### Input Data Sources

| Source | Data | Purpose |
|--------|------|---------|
| **Google Search Console** | All queries with impressions | What you already rank for |
| **Keyword research tools** | Seed keywords + related terms | Expansion opportunities |
| **Competitor content** | URLs + titles + H1s from competing sites | Gap identification |
| **Your existing content** | All published URLs + titles + meta | Current coverage map |
| **Search suggestions** | Google Autocomplete, People Also Ask | Long-tail discovery |

### Minimum Data Requirements

| Data Point | Minimum Volume | Ideal Volume |
|-----------|:--------------:|:------------:|
| Keywords | 500 | 2,000-5,000 |
| Search volume data | For 80%+ of keywords | For all keywords |
| Your existing URLs | All indexed pages | All pages + drafts |
| Competitor URLs | Top 3 competitors | Top 5-10 competitors |

---

## Part 2: TF-IDF Vectorization

### What TF-IDF Does

TF-IDF (Term Frequency–Inverse Document Frequency) converts text into numerical vectors that represent how important each word is relative to the entire corpus.

- **TF (Term Frequency):** How often a word appears in a document
- **IDF (Inverse Document Frequency):** How rare the word is across all documents
- **TF-IDF = TF × IDF** — Words that are frequent in one document but rare overall score highest

### Implementation

```python
from sklearn.feature_extraction.text import TfidfVectorizer

# Prepare keyword data
keywords = [
    "content marketing strategy",
    "content marketing plan",
    "content marketing examples",
    "SEO content writing",
    "SEO blog posts",
    "keyword research tools",
    "keyword research strategy",
    # ... hundreds more
]

# Configure TF-IDF
vectorizer = TfidfVectorizer(
    max_features=5000,        # Limit vocabulary size
    ngram_range=(1, 3),       # Capture 1-word, 2-word, and 3-word phrases
    min_df=2,                 # Ignore terms appearing in fewer than 2 keywords
    max_df=0.95,              # Ignore terms appearing in >95% of keywords
    stop_words='english'      # Remove common words
)

# Transform keywords into TF-IDF vectors
tfidf_matrix = vectorizer.fit_transform(keywords)
```

### Parameter Tuning Guide

| Parameter | Default | Adjust When |
|-----------|:-------:|-------------|
| `ngram_range` | (1, 3) | Increase to (1, 4) for long-tail keyword analysis |
| `min_df` | 2 | Increase for large datasets (>5K keywords) to reduce noise |
| `max_df` | 0.95 | Decrease for more specific clusters |
| `max_features` | 5000 | Increase for larger keyword sets |

---

## Part 3: K-Means Clustering

### What K-Means Does

K-Means groups keywords into K clusters by minimizing the distance between each keyword vector and its cluster center. Keywords that are semantically similar end up in the same cluster.

### Implementation

```python
from sklearn.cluster import KMeans
from sklearn.metrics import silhouette_score
import numpy as np

# Method 1: Elbow method to find optimal K
inertias = []
K_range = range(5, 50)

for k in K_range:
    kmeans = KMeans(n_clusters=k, random_state=42, n_init=10)
    kmeans.fit(tfidf_matrix)
    inertias.append(kmeans.inertia_)

# Method 2: Silhouette score (more reliable)
silhouette_scores = []

for k in K_range:
    kmeans = KMeans(n_clusters=k, random_state=42, n_init=10)
    labels = kmeans.fit_predict(tfidf_matrix)
    score = silhouette_score(tfidf_matrix, labels)
    silhouette_scores.append(score)

# Pick K with highest silhouette score
optimal_k = K_range[np.argmax(silhouette_scores)]

# Final clustering
kmeans = KMeans(n_clusters=optimal_k, random_state=42, n_init=10)
clusters = kmeans.fit_predict(tfidf_matrix)
```

### Choosing K (Number of Clusters)

| Keyword Count | Starting K Range | Expected Clusters |
|:-------------:|:---------------:|:-----------------:|
| 500-1,000 | 10-25 | 12-20 |
| 1,000-3,000 | 15-40 | 20-30 |
| 3,000-5,000 | 20-50 | 25-40 |
| 5,000+ | 30-60 | 30-50 |

**Rule of thumb:** Start with K = √(n/2) where n = number of keywords, then refine using silhouette score.

---

## Part 4: Cluster Analysis & Labeling

### Extracting Cluster Insights

```python
import pandas as pd

# Create results dataframe
results = pd.DataFrame({
    'keyword': keywords,
    'cluster': clusters,
    'search_volume': search_volumes  # from your keyword data
})

# Analyze each cluster
for cluster_id in range(optimal_k):
    cluster_data = results[results['cluster'] == cluster_id]

    # Get top terms for this cluster (cluster center)
    center = kmeans.cluster_centers_[cluster_id]
    feature_names = vectorizer.get_feature_names_out()
    top_terms = [feature_names[i] for i in center.argsort()[-5:][::-1]]

    print(f"\nCluster {cluster_id}: {', '.join(top_terms)}")
    print(f"  Keywords: {len(cluster_data)}")
    print(f"  Total search volume: {cluster_data['search_volume'].sum()}")
    print(f"  Top keywords: {cluster_data.nlargest(5, 'search_volume')['keyword'].tolist()}")
```

### Cluster Labeling Convention

Each cluster becomes a **content pillar** with this naming:

```
[Topic] — [Primary Intent]

Examples:
- "Content Marketing — Strategy & Planning"
- "SEO Tools — Comparison & Selection"
- "Email Marketing — Automation & Sequences"
```

### Quality Checks

| Check | Pass Criteria | Fail Action |
|-------|:------------:|-------------|
| Cluster size | 10-100 keywords per cluster | Split large clusters, merge tiny ones |
| Semantic coherence | Top 5 terms clearly related | Increase K or adjust TF-IDF params |
| Silhouette score | > 0.3 | Re-run with different K |
| No orphan keywords | < 5% of keywords in wrong cluster | Manually reassign outliers |

---

## Part 5: Content Strategy Output

### From Clusters to Content Plans

Each cluster maps to a content pillar:

```
Cluster: "Content Marketing — Strategy & Planning"
├── Pillar Page: "The Complete Guide to Content Marketing Strategy"
├── Supporting Content:
│   ├── "How to Build a Content Marketing Plan (Template)"
│   ├── "Content Marketing Examples from 50 SaaS Companies"
│   ├── "Content Marketing vs. Content Strategy: What's the Difference"
│   ├── "Content Marketing ROI: How to Measure What Matters"
│   └── "Content Marketing for Startups: The $0 Budget Playbook"
└── Internal Linking: All supporting pages link to pillar page
```

### Cluster-to-Content Mapping Table

| Cluster Attribute | Content Decision |
|-------------------|-----------------|
| **Highest search volume keyword** | Pillar page target keyword |
| **Second-tier keywords (100-1K volume)** | Individual supporting articles |
| **Long-tail keywords (<100 volume)** | FAQ sections, subsections within articles |
| **Commercial intent keywords** | Bottom-of-funnel comparison/alternative pages |
| **Informational keywords** | Top-of-funnel educational content |

### Prioritization Matrix

Score each cluster for content prioritization:

| Factor | Weight | Scoring |
|--------|:------:|---------|
| Total cluster search volume | 25% | Relative to other clusters |
| Business relevance | 25% | How close to your product value prop |
| Current coverage gap | 20% | % of cluster keywords with no content |
| Competition level | 15% | Average KD of cluster keywords |
| Conversion potential | 15% | % of commercial-intent keywords in cluster |

---

## Part 6: Ongoing Maintenance

### Monthly Refresh Cycle

1. **Add new keywords** from GSC (queries gaining impressions)
2. **Re-run clustering** with updated keyword set
3. **Compare clusters** to previous month — identify emerging topics
4. **Update content calendar** based on shifting cluster priorities
5. **Track cluster coverage** — % of cluster keywords with published content

### Cluster Health Metrics

| Metric | Target | Measurement |
|--------|:------:|-------------|
| Cluster coverage | >70% of keywords have content | Published URLs ÷ cluster keywords |
| Pillar page ranking | Top 3 for primary keyword | GSC position data |
| Internal link density | 5+ supporting links per pillar | Crawl data |
| Cluster traffic growth | +10% MoM | GA4 filtered by cluster URLs |
| Topical authority score | Rising over time | Average position improvement across cluster |

### When to Re-Cluster

| Trigger | Action |
|---------|--------|
| Added 500+ new keywords | Full re-cluster |
| Major industry shift (new topic emergence) | Full re-cluster |
| Monthly refresh | Incremental update |
| Silhouette score dropping | Re-tune parameters |
| New competitor entering space | Add competitor content to corpus and re-cluster |

---

## Tools & Requirements

| Need | Open Source | Commercial |
|------|-----------|------------|
| **TF-IDF + K-Means** | scikit-learn (Python) | — |
| **Keyword data** | Google Search Console (free) | Ahrefs, SEMrush, Moz |
| **Visualization** | matplotlib, seaborn | Tableau, Looker |
| **Data processing** | pandas, numpy | — |
| **Alternative clustering** | DBSCAN, hierarchical (scipy) | — |

**Minimum technical requirement:** Python 3.8+ with scikit-learn, pandas, and numpy installed.
