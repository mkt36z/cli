package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mkt36z/cli/internal/config"
)

var initDeep bool

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Set up your brand context",
	Long: `Initialize mkt36z with your brand context for personalized output.

  Quick Start (5 questions, <2 min):
    mkt36z init

  Deep Init (all 12 dimensions):
    mkt36z init --deep

  The more context you provide, the better the generated content.
  Check your context score: mkt36z doctor`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if initDeep {
			return runDeepInit()
		}
		return runQuickInit()
	},
}

func init() {
	initCmd.Flags().BoolVar(&initDeep, "deep", false, "Fill all 12 context dimensions")
	rootCmd.AddCommand(initCmd)
}

// runQuickInit asks 5 essential questions to get started fast.
func runQuickInit() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("mkt36z init — Quick Start")
	fmt.Println("Answer 5 questions to personalize your output. (<2 min)")
	fmt.Println()

	// 1. Company name
	companyName := prompt(reader, "Company name")
	if companyName == "" {
		return fmt.Errorf("company name is required")
	}

	// 2. What do you do?
	valueProp := prompt(reader, "What does your company do? (one sentence)")

	// 3. Ideal customer
	idealCustomer := prompt(reader, "Who is your ideal customer?")

	// 4. Differentiator
	differentiator := prompt(reader, "What makes you different from competitors?")

	// 5. Brand voice
	fmt.Println("\nSelect your brand voice (comma-separated numbers):")
	voiceOptions := []string{
		"Confident", "Warm", "Direct", "Playful",
		"Professional", "Bold", "Empathetic", "Witty",
		"Authoritative", "Casual", "Inspirational", "Technical",
	}
	for i, v := range voiceOptions {
		fmt.Printf("  %2d. %s\n", i+1, v)
	}
	voiceInput := prompt(reader, "Voice (e.g., 1,3,5)")
	selectedVoice := parseVoiceSelection(voiceInput, voiceOptions)

	// Save context files
	biz := &config.BusinessContext{CompanyName: companyName}
	if err := config.SaveContext(config.ContextBusiness, biz); err != nil {
		return fmt.Errorf("saving business context: %w", err)
	}

	prod := &config.ProductContext{ValueProp: valueProp}
	if err := config.SaveContext(config.ContextProduct, prod); err != nil {
		return fmt.Errorf("saving product context: %w", err)
	}

	icp := &config.ICPContext{
		Personas: []config.Persona{{Title: idealCustomer}},
	}
	if err := config.SaveContext(config.ContextICP, icp); err != nil {
		return fmt.Errorf("saving ICP context: %w", err)
	}

	comp := &config.CompetitiveContext{Differentiator: differentiator}
	if err := config.SaveContext(config.ContextCompetitive, comp); err != nil {
		return fmt.Errorf("saving competitive context: %w", err)
	}

	brand := &config.BrandVoice{
		CompanyName: companyName,
		Tone:        selectedVoice,
	}
	if err := config.SaveContext(config.ContextBrandVoice, brand); err != nil {
		return fmt.Errorf("saving brand voice context: %w", err)
	}

	// Ask about analytics
	analyticsOpt := prompt(reader, "\nShare anonymous usage data to improve mkt36z? (y/N)")
	analytics := strings.ToLower(strings.TrimSpace(analyticsOpt)) == "y"

	// Save config
	cfg, err := config.Load(flagConfig)
	if err != nil {
		cfg = config.DefaultConfig()
	}
	cfg.Brand = brand
	cfg.Analytics = analytics
	if err := config.Save(cfg); err != nil {
		return fmt.Errorf("saving config: %w", err)
	}

	// Show result
	score := config.ContextScore()
	fmt.Println()
	fmt.Printf("Context saved! Score: %d/100\n", score)
	fmt.Println()

	if flagJSON {
		out, _ := json.MarshalIndent(map[string]interface{}{
			"company":       companyName,
			"context_score": score,
			"dimensions":    5,
			"analytics":     analytics,
		}, "", "  ")
		fmt.Println(string(out))
	}

	populated := len(config.PopulatedContextFiles())
	if populated < 12 {
		fmt.Printf("Filled %d/12 dimensions. Run `mkt36z init --deep` for the rest.\n", populated)
	}

	fmt.Println("\nNext steps:")
	fmt.Println("  mkt36z generate headline \"your topic\"   Generate headlines")
	fmt.Println("  mkt36z analyze positioning               Analyze your position")
	fmt.Println("  mkt36z doctor                            Check your setup")

	return nil
}

// runDeepInit walks through all 12 context dimensions.
func runDeepInit() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("mkt36z init --deep — Full Context Setup")
	fmt.Println("Walk through all 12 context dimensions. Press Enter to skip any.")
	fmt.Println()

	dimensions := []struct {
		ct    config.ContextType
		label string
		fn    func(*bufio.Reader) error
	}{
		{config.ContextBusiness, "Business Foundation", initBusiness},
		{config.ContextProduct, "Product / Service", initProduct},
		{config.ContextICP, "Ideal Customer Profile", initICP},
		{config.ContextCompetitive, "Competitive Landscape", initCompetitive},
		{config.ContextBrandVoice, "Brand Voice & Identity", initBrandVoice},
		{config.ContextChannels, "Marketing Channels", initChannels},
		{config.ContextContentSEO, "Content & SEO", initContentSEO},
		{config.ContextSales, "Sales & Conversion", initSales},
		{config.ContextGoals, "Goals & Constraints", initGoals},
		{config.ContextLanguageBank, "Customer Language Bank", initLanguageBank},
		{config.ContextProofStack, "Proof Stack", initProofStack},
		{config.ContextAntiPatterns, "Anti-Patterns & Constraints", initAntiPatterns},
	}

	for i, dim := range dimensions {
		exists := config.ContextFileExists(dim.ct)
		status := ""
		if exists {
			status = " (already populated)"
		}
		fmt.Printf("\n--- %d/%d: %s%s ---\n", i+1, 12, dim.label, status)

		if exists {
			skip := prompt(reader, "Already filled. Overwrite? (y/N)")
			if strings.ToLower(strings.TrimSpace(skip)) != "y" {
				fmt.Println("  Skipped.")
				continue
			}
		}

		if err := dim.fn(reader); err != nil {
			fmt.Printf("  Error: %v\n", err)
		}
	}

	score := config.ContextScore()
	populated := len(config.PopulatedContextFiles())
	fmt.Printf("\nDone! Context score: %d/100 (%d/12 dimensions filled)\n", score, populated)

	if flagJSON {
		out, _ := json.MarshalIndent(map[string]interface{}{
			"context_score": score,
			"populated":     populated,
			"total":         12,
		}, "", "  ")
		fmt.Println(string(out))
	}

	return nil
}

// Deep init helpers for each dimension.

func initBusiness(r *bufio.Reader) error {
	ctx := &config.BusinessContext{}
	ctx.CompanyName = prompt(r, "Company name")
	ctx.Stage = prompt(r, "Stage (pre-revenue/early/growth/scale)")
	ctx.Revenue = prompt(r, "Annual revenue range")
	ctx.Founded = prompt(r, "Founded year")
	if ctx.CompanyName == "" {
		return nil // skip if no input
	}
	return config.SaveContext(config.ContextBusiness, ctx)
}

func initProduct(r *bufio.Reader) error {
	ctx := &config.ProductContext{}
	ctx.Name = prompt(r, "Product/service name")
	ctx.Category = prompt(r, "Category (e.g., SaaS, consulting, e-commerce)")
	ctx.ValueProp = prompt(r, "Value proposition (one sentence)")
	ctx.Pricing = prompt(r, "Pricing model (e.g., freemium, subscription, one-time)")
	features := prompt(r, "Key features (comma-separated)")
	if features != "" {
		ctx.Features = splitCSV(features)
	}
	if ctx.Name == "" && ctx.ValueProp == "" {
		return nil
	}
	return config.SaveContext(config.ContextProduct, ctx)
}

func initICP(r *bufio.Reader) error {
	ctx := &config.ICPContext{}
	persona := prompt(r, "Primary persona title (e.g., VP of Marketing)")
	if persona != "" {
		role := prompt(r, "Their role/responsibility")
		challenge := prompt(r, "Their biggest challenge")
		ctx.Personas = []config.Persona{{Title: persona, Role: role, Challenge: challenge}}
	}
	pains := prompt(r, "Pain points (comma-separated)")
	if pains != "" {
		ctx.PainPoints = splitCSV(pains)
	}
	goals := prompt(r, "Their goals (comma-separated)")
	if goals != "" {
		ctx.Goals = splitCSV(goals)
	}
	if persona == "" && pains == "" {
		return nil
	}
	return config.SaveContext(config.ContextICP, ctx)
}

func initCompetitive(r *bufio.Reader) error {
	ctx := &config.CompetitiveContext{}
	direct := prompt(r, "Direct competitors (comma-separated)")
	if direct != "" {
		ctx.DirectCompetitors = splitCSV(direct)
	}
	ctx.Differentiator = prompt(r, "Your key differentiator")
	ctx.Moat = prompt(r, "Your moat (what's hard to copy)")
	if ctx.Differentiator == "" && direct == "" {
		return nil
	}
	return config.SaveContext(config.ContextCompetitive, ctx)
}

func initBrandVoice(r *bufio.Reader) error {
	ctx := &config.BrandVoice{}
	ctx.CompanyName = prompt(r, "Company name")
	ctx.Industry = prompt(r, "Industry")
	tone := prompt(r, "Brand tone adjectives (comma-separated, e.g., confident,warm,direct)")
	if tone != "" {
		ctx.Tone = splitCSV(tone)
	}
	forbidden := prompt(r, "Forbidden words (comma-separated)")
	if forbidden != "" {
		ctx.ForbiddenWords = splitCSV(forbidden)
	}
	if ctx.CompanyName == "" && tone == "" {
		return nil
	}
	return config.SaveContext(config.ContextBrandVoice, ctx)
}

func initChannels(r *bufio.Reader) error {
	ctx := &config.ChannelsContext{}
	channels := prompt(r, "Active marketing channels (comma-separated, e.g., LinkedIn,Email,SEO)")
	if channels == "" {
		return nil
	}
	for _, ch := range splitCSV(channels) {
		ctx.ActiveChannels = append(ctx.ActiveChannels, config.Channel{Name: ch})
	}
	ctx.CAC = prompt(r, "Customer acquisition cost (approximate)")
	return config.SaveContext(config.ContextChannels, ctx)
}

func initContentSEO(r *bufio.Reader) error {
	ctx := &config.ContentSEOContext{}
	ctx.BlogURL = prompt(r, "Blog URL")
	ctx.Frequency = prompt(r, "Publishing frequency (e.g., 2x/week)")
	pillars := prompt(r, "Content pillars (comma-separated)")
	if pillars != "" {
		ctx.Pillars = splitCSV(pillars)
	}
	keywords := prompt(r, "Target keywords (comma-separated)")
	if keywords != "" {
		ctx.Keywords = splitCSV(keywords)
	}
	if ctx.BlogURL == "" && pillars == "" {
		return nil
	}
	return config.SaveContext(config.ContextContentSEO, ctx)
}

func initSales(r *bufio.Reader) error {
	ctx := &config.SalesContext{}
	ctx.Model = prompt(r, "Sales model (self-serve/sales-led/hybrid)")
	ctx.AvgDealSize = prompt(r, "Average deal size")
	ctx.ConversionRates = prompt(r, "Key conversion rate (e.g., trial→paid: 12%)")
	ctx.Churn = prompt(r, "Monthly churn rate")
	if ctx.Model == "" && ctx.AvgDealSize == "" {
		return nil
	}
	return config.SaveContext(config.ContextSales, ctx)
}

func initGoals(r *bufio.Reader) error {
	ctx := &config.GoalsContext{}
	ctx.PrimaryGoal = prompt(r, "Primary marketing goal")
	ctx.Budget = prompt(r, "Monthly marketing budget")
	ctx.Timeline = prompt(r, "Timeline (e.g., next 90 days)")
	metrics := prompt(r, "Success metrics (comma-separated)")
	if metrics != "" {
		ctx.SuccessMetrics = splitCSV(metrics)
	}
	if ctx.PrimaryGoal == "" {
		return nil
	}
	return config.SaveContext(config.ContextGoals, ctx)
}

func initLanguageBank(r *bufio.Reader) error {
	ctx := &config.LanguageBankContext{}
	pains := prompt(r, "Customer pain phrases (comma-separated, exact words they use)")
	if pains != "" {
		ctx.PainPhrases = splitCSV(pains)
	}
	desires := prompt(r, "Customer desire phrases (comma-separated)")
	if desires != "" {
		ctx.DesirePhrases = splitCSV(desires)
	}
	objections := prompt(r, "Common objections (comma-separated)")
	if objections != "" {
		ctx.ObjectionPhrases = splitCSV(objections)
	}
	if pains == "" && desires == "" {
		return nil
	}
	return config.SaveContext(config.ContextLanguageBank, ctx)
}

func initProofStack(r *bufio.Reader) error {
	ctx := &config.ProofStackContext{}
	testimonials := prompt(r, "Key testimonial quotes (comma-separated)")
	if testimonials != "" {
		ctx.Testimonials = splitCSV(testimonials)
	}
	metrics := prompt(r, "Key metrics/stats (comma-separated, e.g., 10x faster, 50% cost reduction)")
	if metrics != "" {
		ctx.Metrics = splitCSV(metrics)
	}
	certs := prompt(r, "Certifications/awards (comma-separated)")
	if certs != "" {
		ctx.Certifications = splitCSV(certs)
	}
	if testimonials == "" && metrics == "" {
		return nil
	}
	return config.SaveContext(config.ContextProofStack, ctx)
}

func initAntiPatterns(r *bufio.Reader) error {
	ctx := &config.AntiPatternsContext{}
	negative := prompt(r, "Who is NOT your customer? (comma-separated)")
	if negative != "" {
		ctx.NegativePersonas = splitCSV(negative)
	}
	forbidden := prompt(r, "Forbidden claims (comma-separated)")
	if forbidden != "" {
		ctx.ForbiddenClaims = splitCSV(forbidden)
	}
	compliance := prompt(r, "Compliance requirements (comma-separated, e.g., GDPR, HIPAA)")
	if compliance != "" {
		ctx.Compliance = splitCSV(compliance)
	}
	if negative == "" && forbidden == "" && compliance == "" {
		return nil
	}
	return config.SaveContext(config.ContextAntiPatterns, ctx)
}

// prompt prints a question and reads a line of input.
func prompt(r *bufio.Reader, question string) string {
	fmt.Printf("  %s: ", question)
	line, err := r.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(line)
}

// splitCSV splits a comma-separated string into trimmed parts.
func splitCSV(s string) []string {
	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			result = append(result, p)
		}
	}
	return result
}

// parseVoiceSelection converts "1,3,5" into selected voice labels.
func parseVoiceSelection(input string, options []string) []string {
	var selected []string
	for _, part := range splitCSV(input) {
		var idx int
		if _, err := fmt.Sscanf(part, "%d", &idx); err == nil && idx >= 1 && idx <= len(options) {
			selected = append(selected, options[idx-1])
		}
	}
	if len(selected) == 0 {
		// Default voice
		selected = []string{"Confident", "Direct", "Professional"}
	}
	return selected
}
