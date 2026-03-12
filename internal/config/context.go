package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// ContextType identifies one of the 12 context dimensions.
type ContextType string

const (
	ContextBusiness      ContextType = "business"
	ContextProduct       ContextType = "product"
	ContextICP           ContextType = "icp"
	ContextCompetitive   ContextType = "competitive"
	ContextBrandVoice    ContextType = "brand-voice"
	ContextChannels      ContextType = "channels"
	ContextContentSEO    ContextType = "content-seo"
	ContextSales         ContextType = "sales"
	ContextGoals         ContextType = "goals"
	ContextLanguageBank  ContextType = "language-bank"
	ContextProofStack    ContextType = "proof-stack"
	ContextAntiPatterns  ContextType = "anti-patterns"
)

// AllContextTypes returns all 12 context types in recommended fill order.
var AllContextTypes = []ContextType{
	ContextBusiness, ContextProduct, ContextICP, ContextCompetitive,
	ContextBrandVoice, ContextChannels, ContextContentSEO, ContextSales,
	ContextGoals, ContextLanguageBank, ContextProofStack, ContextAntiPatterns,
}

// contextLabels maps types to human-readable names.
var contextLabels = map[ContextType]string{
	ContextBusiness:     "Business Foundation",
	ContextProduct:      "Product / Service",
	ContextICP:          "Ideal Customer Profile",
	ContextCompetitive:  "Competitive Landscape",
	ContextBrandVoice:   "Brand Voice & Identity",
	ContextChannels:     "Marketing Channels",
	ContextContentSEO:   "Content & SEO",
	ContextSales:        "Sales & Conversion",
	ContextGoals:        "Goals & Constraints",
	ContextLanguageBank: "Customer Language Bank",
	ContextProofStack:   "Proof Stack",
	ContextAntiPatterns: "Anti-Patterns & Constraints",
}

// Label returns the human-readable name for this context type.
func (ct ContextType) Label() string {
	if l, ok := contextLabels[ct]; ok {
		return l
	}
	return string(ct)
}

// Filename returns the YAML filename for this context type.
func (ct ContextType) Filename() string {
	return string(ct) + ".yaml"
}

// --- Per-type context structs ---

// BusinessContext holds company fundamentals.
type BusinessContext struct {
	CompanyName string `yaml:"company_name,omitempty" json:"company_name,omitempty"`
	Stage       string `yaml:"stage,omitempty" json:"stage,omitempty"` // pre-revenue, early, growth, scale
	Revenue     string `yaml:"revenue,omitempty" json:"revenue,omitempty"`
	TeamSize    int    `yaml:"team_size,omitempty" json:"team_size,omitempty"`
	Funding     string `yaml:"funding,omitempty" json:"funding,omitempty"`
	Founded     string `yaml:"founded,omitempty" json:"founded,omitempty"`
}

// ProductContext holds product/service details.
type ProductContext struct {
	Name       string   `yaml:"name,omitempty" json:"name,omitempty"`
	Category   string   `yaml:"category,omitempty" json:"category,omitempty"`
	ValueProp  string   `yaml:"value_prop,omitempty" json:"value_prop,omitempty"`
	Pricing    string   `yaml:"pricing,omitempty" json:"pricing,omitempty"`
	Features   []string `yaml:"features,omitempty" json:"features,omitempty"`
	UseCases   []string `yaml:"use_cases,omitempty" json:"use_cases,omitempty"`
}

// ICPContext holds Ideal Customer Profile data.
type ICPContext struct {
	Personas     []Persona `yaml:"personas,omitempty" json:"personas,omitempty"`
	PainPoints   []string  `yaml:"pain_points,omitempty" json:"pain_points,omitempty"`
	Goals        []string  `yaml:"goals,omitempty" json:"goals,omitempty"`
	WateringHoles []string `yaml:"watering_holes,omitempty" json:"watering_holes,omitempty"`
	BuyerStages  []string  `yaml:"buyer_stages,omitempty" json:"buyer_stages,omitempty"`
}

// Persona describes a single ICP persona.
type Persona struct {
	Title      string `yaml:"title,omitempty" json:"title,omitempty"`
	Role       string `yaml:"role,omitempty" json:"role,omitempty"`
	Challenge  string `yaml:"challenge,omitempty" json:"challenge,omitempty"`
	Motivation string `yaml:"motivation,omitempty" json:"motivation,omitempty"`
}

// CompetitiveContext holds competitive landscape data.
type CompetitiveContext struct {
	DirectCompetitors   []string `yaml:"direct_competitors,omitempty" json:"direct_competitors,omitempty"`
	IndirectCompetitors []string `yaml:"indirect_competitors,omitempty" json:"indirect_competitors,omitempty"`
	Differentiator      string   `yaml:"differentiator,omitempty" json:"differentiator,omitempty"`
	Moat                string   `yaml:"moat,omitempty" json:"moat,omitempty"`
}

// ChannelsContext holds marketing channel data.
type ChannelsContext struct {
	ActiveChannels []Channel `yaml:"active_channels,omitempty" json:"active_channels,omitempty"`
	CAC            string    `yaml:"cac,omitempty" json:"cac,omitempty"`
}

// Channel represents a single marketing channel.
type Channel struct {
	Name        string `yaml:"name,omitempty" json:"name,omitempty"`
	Performance string `yaml:"performance,omitempty" json:"performance,omitempty"`
	Spend       string `yaml:"spend,omitempty" json:"spend,omitempty"`
}

// ContentSEOContext holds content and SEO data.
type ContentSEOContext struct {
	BlogURL    string   `yaml:"blog_url,omitempty" json:"blog_url,omitempty"`
	Frequency  string   `yaml:"frequency,omitempty" json:"frequency,omitempty"`
	Pillars    []string `yaml:"pillars,omitempty" json:"pillars,omitempty"`
	Keywords   []string `yaml:"keywords,omitempty" json:"keywords,omitempty"`
}

// SalesContext holds sales and conversion data.
type SalesContext struct {
	Model           string `yaml:"model,omitempty" json:"model,omitempty"` // self-serve, sales-led, hybrid
	AvgDealSize     string `yaml:"avg_deal_size,omitempty" json:"avg_deal_size,omitempty"`
	ConversionRates string `yaml:"conversion_rates,omitempty" json:"conversion_rates,omitempty"`
	Churn           string `yaml:"churn,omitempty" json:"churn,omitempty"`
}

// GoalsContext holds goals and constraints.
type GoalsContext struct {
	PrimaryGoal    string   `yaml:"primary_goal,omitempty" json:"primary_goal,omitempty"`
	Budget         string   `yaml:"budget,omitempty" json:"budget,omitempty"`
	Timeline       string   `yaml:"timeline,omitempty" json:"timeline,omitempty"`
	SuccessMetrics []string `yaml:"success_metrics,omitempty" json:"success_metrics,omitempty"`
	KeyInsight     string   `yaml:"key_insight,omitempty" json:"key_insight,omitempty"`
}

// LanguageBankContext holds customer language data.
type LanguageBankContext struct {
	PainPhrases    []string `yaml:"pain_phrases,omitempty" json:"pain_phrases,omitempty"`
	DesirePhrases  []string `yaml:"desire_phrases,omitempty" json:"desire_phrases,omitempty"`
	ObjectionPhrases []string `yaml:"objection_phrases,omitempty" json:"objection_phrases,omitempty"`
	SuccessPhrases []string `yaml:"success_phrases,omitempty" json:"success_phrases,omitempty"`
	Vocabulary     []string `yaml:"vocabulary,omitempty" json:"vocabulary,omitempty"`
	Sources        []string `yaml:"sources,omitempty" json:"sources,omitempty"`
}

// ProofStackContext holds proof and social evidence.
type ProofStackContext struct {
	Testimonials   []string `yaml:"testimonials,omitempty" json:"testimonials,omitempty"`
	CaseStudies    []string `yaml:"case_studies,omitempty" json:"case_studies,omitempty"`
	Metrics        []string `yaml:"metrics,omitempty" json:"metrics,omitempty"`
	Certifications []string `yaml:"certifications,omitempty" json:"certifications,omitempty"`
	ReasonsToBelie []string `yaml:"reasons_to_believe,omitempty" json:"reasons_to_believe,omitempty"`
}

// AntiPatternsContext holds negative constraints.
type AntiPatternsContext struct {
	NegativePersonas []string            `yaml:"negative_personas,omitempty" json:"negative_personas,omitempty"`
	ForbiddenClaims  []string            `yaml:"forbidden_claims,omitempty" json:"forbidden_claims,omitempty"`
	Compliance       []string            `yaml:"compliance,omitempty" json:"compliance,omitempty"`
	ContentRules     map[string][]string `yaml:"content_rules,omitempty" json:"content_rules,omitempty"`
}

// LoadContext loads a single context file by type.
func LoadContext(ct ContextType, target interface{}) error {
	path := filepath.Join(ContextDir(), ct.Filename())
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, target)
}

// SaveContext writes a single context file by type.
func SaveContext(ct ContextType, source interface{}) error {
	dir := ContextDir()
	if err := EnsureDir(dir); err != nil {
		return err
	}
	data, err := yaml.Marshal(source)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, ct.Filename()), data, 0600)
}

// ContextFileExists returns true if the context file exists and is non-empty.
func ContextFileExists(ct ContextType) bool {
	path := filepath.Join(ContextDir(), ct.Filename())
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Size() > 0
}

// PopulatedContextFiles returns the list of context types that have files on disk.
func PopulatedContextFiles() []ContextType {
	var populated []ContextType
	for _, ct := range AllContextTypes {
		if ContextFileExists(ct) {
			populated = append(populated, ct)
		}
	}
	return populated
}

// ContextSummary returns a human-readable summary of context completeness.
func ContextSummary() string {
	populated := PopulatedContextFiles()
	return fmt.Sprintf("%d/%d context files populated", len(populated), len(AllContextTypes))
}
