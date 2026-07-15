package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/whalelogic/SkillRx/app/config"

	"google.golang.org/genai"
)

// Repo path
// 	"github.com/whalelogic/SkillRx/app/<module>"

// ---------------------------------------------------------------------
// Data model
// ---------------------------------------------------------------------

type LibraryEntry struct {
	ID    string
	Label string
}

type Category struct {
	Name    string
	Code    string
	Entries []LibraryEntry
}

type PageData struct {
	Categories               []Category
	ActiveCategory           string
	ActiveEntry              string
	ActiveLabel              string
	ModelName                string
	SourcePath               string
	SourceContent            string
	InstructionOptions       []LibraryEntry
	PromptOptions            []LibraryEntry
	SkillOptions             []LibraryEntry
	SelectedInstruction      string
	SelectedInstructionLabel string
	SelectedPrompt           string
	SelectedPromptLabel      string
	SelectedSkills           []string
	SelectedSkillLabels      []string
	CompositionSummary       string
}

var (
	categories  []Category
	tmpl        *template.Template
	appRoot     string
	staticRoot  string
	libraryRoot string
)

const defaultModel = "gemini-3-flash-preview"
const defaultCategory = "financial-markets"
const defaultEntry = "system-instructions/market-analyst.md"

// loadCategories walks rxlib/<category> recursively and exposes the actual
// library documents on disk.
func loadCategories(root string) []Category {
	entries, err := os.ReadDir(root)
	if err != nil {
		log.Printf("rxlib not found at %q, using demo data (%v)", root, err)
		return demoCategories()
	}

	var cats []Category
	i := 0
	for _, e := range entries {
		if !e.IsDir() || strings.HasPrefix(e.Name(), ".") {
			continue
		}
		i++
		docs, err := loadLibraryEntries(filepath.Join(root, e.Name()))
		if err != nil {
			continue
		}
		if len(docs) == 0 {
			continue
		}
		cats = append(cats, Category{
			Name:    e.Name(),
			Code:    fmt.Sprintf("C%02d", i),
			Entries: docs,
		})
	}
	if len(cats) == 0 {
		return demoCategories()
	}
	return cats
}

func loadLibraryEntries(categoryRoot string) ([]LibraryEntry, error) {
	var docs []LibraryEntry
	err := filepath.WalkDir(categoryRoot, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if strings.HasPrefix(d.Name(), ".") {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if d.IsDir() {
			return nil
		}
		if !isLibraryDoc(d.Name()) {
			return nil
		}

		rel, err := filepath.Rel(categoryRoot, path)
		if err != nil {
			return nil
		}
		rel = filepath.ToSlash(rel)
		label := strings.TrimSuffix(rel, filepath.Ext(rel))
		label = strings.ReplaceAll(label, "/", " / ")
		docs = append(docs, LibraryEntry{
			ID:    rel,
			Label: label,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(docs, func(i, j int) bool { return docs[i].ID < docs[j].ID })
	return docs, nil
}

func isLibraryDoc(name string) bool {
	switch strings.ToLower(filepath.Ext(name)) {
	case ".md", ".prompt":
		return true
	default:
		return false
	}
}

func demoCategories() []Category {
	raw := map[string][]LibraryEntry{
		"claude":               {{ID: "skill.md", Label: "skill"}},
		"data-science":         {{ID: "prompts/exploratory-data-analysis.md", Label: "prompts / exploratory-data-analysis"}, {ID: "skills/statistical-reasoning.md", Label: "skills / statistical-reasoning"}},
		"financial-markets":    {{ID: "prompts/earnings-summary.md", Label: "prompts / earnings-summary"}, {ID: "system-instructions/market-analyst.md", Label: "system-instructions / market-analyst"}},
		"productivity":         {{ID: "prompts/task-breakdown.md", Label: "prompts / task-breakdown"}, {ID: "skills/prioritisation.md", Label: "skills / prioritisation"}},
		"repo-agents":          {{ID: "AGENTS.md", Label: "AGENTS"}},
		"research":             {{ID: "prompts/literature-review.md", Label: "prompts / literature-review"}, {ID: "skills/source-evaluation.md", Label: "skills / source-evaluation"}},
		"repository-analysis":  {{ID: "prompts/codebase-overview.md", Label: "prompts / codebase-overview"}, {ID: "skills/commit-analysis.md", Label: "skills / commit-analysis"}},
		"skill-templates":      {{ID: "prompt-template.md", Label: "prompt-template"}, {ID: "skill-template.md", Label: "skill-template"}},
		"software-engineering": {{ID: "prompts/code-review.md", Label: "prompts / code-review"}, {ID: "skills/code-execution.md", Label: "skills / code-execution"}},
		"ui-ux-development":    {{ID: "prompts/design-review.md", Label: "prompts / design-review"}, {ID: "skills/accessibility-expert.md", Label: "skills / accessibility-expert"}},
		"writing":              {{ID: "prompts/blog-post.md", Label: "prompts / blog-post"}, {ID: "skills/tone-adjustment.md", Label: "skills / tone-adjustment"}},
		"dotprompt":            {{ID: "codeReview.prompt", Label: "codeReview"}, {ID: "mentorChat.prompt", Label: "mentorChat"}},
	}
	order := []string{"claude", "data-science", "financial-markets", "productivity",
		"repo-agents", "research", "repository-analysis", "skill-templates",
		"software-engineering", "ui-ux-development", "writing", "dotprompt"}

	var cats []Category
	for i, name := range order {
		cats = append(cats, Category{Name: name, Code: fmt.Sprintf("C%02d", i+1), Entries: raw[name]})
	}
	return cats
}

func sanitizeEntryID(entry string) (string, bool) {
	entry = filepath.Clean(filepath.FromSlash(entry))
	if entry == "." || entry == "" || filepath.IsAbs(entry) || entry == ".." || strings.HasPrefix(entry, ".."+string(filepath.Separator)) {
		return "", false
	}
	return filepath.ToSlash(entry), true
}

// readSkillSource loads the real library document if present, else a placeholder
// so the Source tab always has something to show.
func readSkillSource(root, category, entry string) (path, content string) {
	cleanEntry, ok := sanitizeEntryID(entry)
	if !ok {
		return "rxlib/invalid-selection", "Invalid library selection."
	}

	path = filepath.ToSlash(filepath.Join("rxlib", category, cleanEntry))
	full := filepath.Join(root, category, filepath.FromSlash(cleanEntry))
	if b, err := os.ReadFile(full); err == nil {
		return path, string(b)
	}
	content = fmt.Sprintf(`---
name: %s
category: %s
---

# %s

(No source file found on disk at %s — showing placeholder.
Point RXLIB_PATH at your real rxlib/ directory to load actual content.)`,
		cleanEntry, category, cleanEntry, full)
	return path, content
}

func readLibraryEntry(root, category, entry string) (string, error) {
	cleanEntry, ok := sanitizeEntryID(entry)
	if !ok {
		return "", fmt.Errorf("invalid library selection")
	}

	full := filepath.Join(root, category, filepath.FromSlash(cleanEntry))
	b, err := os.ReadFile(full)
	if err != nil {
		return "", fmt.Errorf("read library entry %q: %w", full, err)
	}
	return string(b), nil
}

// ---------------------------------------------------------------------
// Handlers
// ---------------------------------------------------------------------

func projectRoot() string {
	candidates := []string{}
	if cwd, err := os.Getwd(); err == nil {
		candidates = append(candidates, cwd)
	}
	if _, file, _, ok := runtime.Caller(0); ok {
		candidates = append(candidates, filepath.Dir(file))
	}
	if exe, err := os.Executable(); err == nil {
		candidates = append(candidates, filepath.Dir(exe))
	}
	for _, start := range candidates {
		if root := findRepoRoot(start); root != "" {
			return root
		}
	}
	return ""
}

func findRepoRoot(start string) string {
	dir := start
	for {
		if hasDir(filepath.Join(dir, "rxlib")) && hasDir(filepath.Join(dir, "app", "templates")) && hasDir(filepath.Join(dir, "app", "static")) {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}

func hasDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func rxlibRoot() string {
	if p := os.Getenv("RXLIB_PATH"); p != "" {
		return p
	}
	if libraryRoot != "" {
		return libraryRoot
	}
	return filepath.Join("..", "rxlib")
}

func currentSelection(categoryName, entryID string) (Category, LibraryEntry, bool) {
	for _, category := range categories {
		if category.Name != categoryName {
			continue
		}
		for _, entry := range category.Entries {
			if entry.ID == entryID {
				return category, entry, true
			}
		}
		return Category{}, LibraryEntry{}, false
	}
	return Category{}, LibraryEntry{}, false
}

func categoryByName(categoryName string) (Category, bool) {
	for _, category := range categories {
		if category.Name == categoryName {
			return category, true
		}
	}
	return Category{}, false
}

func entryKind(entryID string) string {
	parts := strings.SplitN(filepath.ToSlash(entryID), "/", 2)
	if len(parts) == 0 {
		return ""
	}
	return parts[0]
}

func entriesForKind(category Category, kind string) []LibraryEntry {
	if kind == "" {
		return append([]LibraryEntry(nil), category.Entries...)
	}

	var filtered []LibraryEntry
	for _, entry := range category.Entries {
		if entryKind(entry.ID) == kind {
			filtered = append(filtered, entry)
		}
	}
	if kind == "system-instructions" && len(filtered) == 0 {
		return append([]LibraryEntry(nil), category.Entries...)
	}
	return filtered
}

func findEntry(entries []LibraryEntry, entryID string) (LibraryEntry, bool) {
	for _, entry := range entries {
		if entry.ID == entryID {
			return entry, true
		}
	}
	return LibraryEntry{}, false
}

func defaultInstructionID(category Category) string {
	if category.Name == defaultCategory {
		if _, entry, ok := currentSelection(category.Name, defaultEntry); ok {
			return entry.ID
		}
	}

	options := entriesForKind(category, "system-instructions")
	if len(options) > 0 {
		return options[0].ID
	}
	if len(category.Entries) > 0 {
		return category.Entries[0].ID
	}
	return ""
}

func normalizeEntryID(options []LibraryEntry, requested, fallback string, allowEmpty bool) string {
	if requested == "" {
		if allowEmpty {
			return ""
		}
		return fallback
	}
	if _, ok := findEntry(options, requested); ok {
		return requested
	}
	if allowEmpty {
		return ""
	}
	return fallback
}

func splitSkillValues(values []string) []string {
	var out []string
	for _, value := range values {
		for _, part := range strings.Split(value, ",") {
			part = strings.TrimSpace(part)
			if part != "" {
				out = append(out, part)
			}
		}
	}
	return out
}

func normalizeSkillIDs(options []LibraryEntry, requested []string) []string {
	if len(requested) == 0 || len(options) == 0 {
		return nil
	}

	allowed := make(map[string]bool, len(requested))
	for _, skillID := range requested {
		allowed[skillID] = true
	}

	var normalized []string
	for _, option := range options {
		if allowed[option.ID] {
			normalized = append(normalized, option.ID)
		}
	}
	return normalized
}

func entryLabels(category Category, entryIDs []string) []string {
	var labels []string
	for _, entryID := range entryIDs {
		if entry, ok := findEntry(category.Entries, entryID); ok {
			labels = append(labels, entry.Label)
		}
	}
	return labels
}

func compositionSummary(instructionLabel, promptLabel string, skillLabels []string) string {
	parts := []string{}
	if instructionLabel != "" {
		parts = append(parts, "instruction `"+instructionLabel+"`")
	}
	if promptLabel != "" {
		parts = append(parts, "prompt `"+promptLabel+"`")
	}
	if len(skillLabels) > 0 {
		parts = append(parts, fmt.Sprintf("%d skill%s", len(skillLabels), map[bool]string{true: "", false: "s"}[len(skillLabels) == 1]))
	}
	if len(parts) == 0 {
		return "base library document"
	}
	return strings.Join(parts, " + ")
}

func defaultSelection() (Category, LibraryEntry) {
	if category, entry, ok := currentSelection(defaultCategory, defaultEntry); ok {
		return category, entry
	}
	for _, category := range categories {
		if len(category.Entries) > 0 {
			return category, category.Entries[0]
		}
	}
	return Category{}, LibraryEntry{}
}

func buildPageData(categoryName, sourceID, instructionID, promptID string, requestedSkills []string) (PageData, error) {
	category, ok := categoryByName(categoryName)
	if !ok {
		category, _ = defaultSelection()
	}

	instructionOptions := entriesForKind(category, "system-instructions")
	promptOptions := entriesForKind(category, "prompts")
	skillOptions := entriesForKind(category, "skills")

	selectedInstruction := normalizeEntryID(instructionOptions, instructionID, defaultInstructionID(category), false)
	selectedPrompt := normalizeEntryID(promptOptions, promptID, "", true)
	selectedSkills := normalizeSkillIDs(skillOptions, requestedSkills)

	if sourceID == "" {
		sourceID = selectedInstruction
	}
	if _, ok := findEntry(category.Entries, sourceID); !ok {
		sourceID = selectedInstruction
	}
	if sourceID == "" && len(category.Entries) > 0 {
		sourceID = category.Entries[0].ID
	}

	sourceEntry, ok := findEntry(category.Entries, sourceID)
	if !ok {
		return PageData{}, fmt.Errorf("library entry not found")
	}

	path, content := readSkillSource(rxlibRoot(), category.Name, sourceEntry.ID)
	instructionEntry, _ := findEntry(category.Entries, selectedInstruction)
	promptEntry, _ := findEntry(category.Entries, selectedPrompt)
	skillLabels := entryLabels(category, selectedSkills)

	return PageData{
		Categories:               categories,
		ActiveCategory:           category.Name,
		ActiveEntry:              sourceEntry.ID,
		ActiveLabel:              sourceEntry.Label,
		ModelName:                modelName(),
		SourcePath:               path,
		SourceContent:            content,
		InstructionOptions:       instructionOptions,
		PromptOptions:            promptOptions,
		SkillOptions:             skillOptions,
		SelectedInstruction:      selectedInstruction,
		SelectedInstructionLabel: instructionEntry.Label,
		SelectedPrompt:           selectedPrompt,
		SelectedPromptLabel:      promptEntry.Label,
		SelectedSkills:           selectedSkills,
		SelectedSkillLabels:      skillLabels,
		CompositionSummary:       compositionSummary(instructionEntry.Label, promptEntry.Label, skillLabels),
	}, nil
}

func modelName() string {
	if name := strings.TrimSpace(os.Getenv("GEMINI_MODEL")); name != "" {
		return name
	}
	return defaultModel
}

func geminiAPIKey() string {
	return config.LoadApiKey()
}

func shouldUseGrounding(category string) bool {
	return category == "financial-markets"
}

func composeSystemInstruction(category, instructionID string, skillIDs []string) (string, error) {
	baseInstruction, err := readLibraryEntry(rxlibRoot(), category, instructionID)
	if err != nil {
		return "", err
	}

	sections := []string{strings.TrimSpace(baseInstruction)}
	for _, skillID := range skillIDs {
		skillText, err := readLibraryEntry(rxlibRoot(), category, skillID)
		if err != nil {
			return "", err
		}
		_, skillEntry, _ := currentSelection(category, skillID)
		sections = append(sections, fmt.Sprintf("## Attached Skill: %s\n\n%s", skillEntry.Label, strings.TrimSpace(skillText)))
	}

	if shouldUseGrounding(category) {
		sections = append(sections,
			"## Runtime Grounding Requirement\n\n"+
				"You must ground any current financial, macro, yield, spread, valuation, performance, or news claim in live web data retrieved for this response. "+
				"If current data cannot be grounded, say so explicitly and do not substitute stale training-data figures.",
		)
	}

	return strings.Join(sections, "\n\n---\n\n"), nil
}

func groundedPrompt(category, prompt string) string {
	if !shouldUseGrounding(category) {
		return prompt
	}
	return fmt.Sprintf(
		"Current date: %s.\nUse grounded web data retrieved today for any current financial facts, prices, yields, spreads, macro releases, performance claims, or market news. If a figure cannot be grounded with current sources, say that explicitly instead of answering from memory.\n\nUser request:\n%s",
		time.Now().Format("2006-01-02"),
		prompt,
	)
}

func applyPromptTemplate(templateText, userInput string) string {
	replacements := map[string]string{
		"{{USER_INPUT}}":  userInput,
		"{{USER_PROMPT}}": userInput,
		"{{PROMPT}}":      userInput,
		"{{QUESTION}}":    userInput,
		"{{REQUEST}}":     userInput,
		"{{TODAY}}":       time.Now().Format("2006-01-02"),
	}

	usedPlaceholder := false
	composed := templateText
	for placeholder, value := range replacements {
		if strings.Contains(composed, placeholder) {
			composed = strings.ReplaceAll(composed, placeholder, value)
			usedPlaceholder = true
		}
	}
	if usedPlaceholder {
		return composed
	}

	return strings.TrimSpace(templateText) + "\n\n## Live User Input\n\n" + userInput
}

func composeUserPrompt(category, promptTemplateID, userInput string) (string, error) {
	composed := userInput
	if promptTemplateID != "" {
		templateText, err := readLibraryEntry(rxlibRoot(), category, promptTemplateID)
		if err != nil {
			return "", err
		}
		composed = applyPromptTemplate(templateText, userInput)
	}
	return groundedPrompt(category, composed), nil
}

func groundingTools(category string) []*genai.Tool {
	if !shouldUseGrounding(category) {
		return nil
	}
	return []*genai.Tool{
		{GoogleSearchRetrieval: &genai.GoogleSearchRetrieval{}},
	}
}

func groundingSources(result *genai.GenerateContentResponse) []string {
	if result == nil || len(result.Candidates) == 0 || result.Candidates[0].GroundingMetadata == nil {
		return nil
	}

	seen := map[string]bool{}
	var sources []string
	for _, chunk := range result.Candidates[0].GroundingMetadata.GroundingChunks {
		if chunk == nil || chunk.Web == nil || chunk.Web.URI == "" || seen[chunk.Web.URI] {
			continue
		}
		seen[chunk.Web.URI] = true

		if chunk.Web.Title != "" {
			sources = append(sources, fmt.Sprintf("- [%s](%s)", chunk.Web.Title, chunk.Web.URI))
			continue
		}
		sources = append(sources, fmt.Sprintf("- %s", chunk.Web.URI))
	}
	return sources
}

func appendGroundingSources(reply string, result *genai.GenerateContentResponse) string {
	sources := groundingSources(result)
	if len(sources) == 0 {
		return reply
	}

	if len(sources) > 6 {
		sources = sources[:6]
	}
	return reply + "\n\n**Sources**\n" + strings.Join(sources, "\n")
}

func generateReply(ctx context.Context, category, instructionID, promptTemplateID string, skillIDs []string, prompt string) (string, error) {
	_, entry, ok := currentSelection(category, instructionID)
	if !ok {
		return "", fmt.Errorf("instruction not found")
	}

	systemInstruction, err := composeSystemInstruction(category, entry.ID, skillIDs)
	if err != nil {
		return "", err
	}
	userPrompt, err := composeUserPrompt(category, promptTemplateID, prompt)
	if err != nil {
		return "", err
	}

	apiKey := geminiAPIKey()
	if apiKey == "" {
		return "", fmt.Errorf("set GEMINI_API_KEY or GOOGLE_API_KEY to enable chat")
	}

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return "", fmt.Errorf("create Gemini client: %w", err)
	}

	result, err := client.Models.GenerateContent(ctx, modelName(), genai.Text(userPrompt), &genai.GenerateContentConfig{
		SystemInstruction: &genai.Content{
			Parts: []*genai.Part{{Text: systemInstruction}},
		},
		Tools: groundingTools(category),
	})
	if err != nil {
		return "", fmt.Errorf("generate model response: %w", err)
	}
	if shouldUseGrounding(category) && len(groundingSources(result)) == 0 {
		return "", fmt.Errorf("Gemini returned an ungrounded response. Please retry with a more specific request.")
	}

	reply := strings.TrimSpace(result.Text())
	if reply == "" {
		return "", fmt.Errorf("model returned an empty response")
	}
	reply = appendGroundingSources(reply, result)
	return reply, nil
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(recorder, r)
		log.Printf("%s %s status=%d duration=%s", r.Method, r.URL.RequestURI(), recorder.status, time.Since(start).Round(time.Millisecond))
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	activeCategory, activeEntry := defaultSelection()
	data, err := buildPageData(activeCategory.Name, activeEntry.ID, "", "", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GET /select?category=<category>&entry=<path>
// Returns a fragment containing several out-of-band swaps: the topbar
// label, the sidebar active state, the source panel, the hidden form
// inputs, and a reset chat pane.
func selectHandler(w http.ResponseWriter, r *http.Request) {
	cat := r.URL.Query().Get("category")
	if cat == "" {
		cat = r.PathValue("category")
	}
	entryID := r.URL.Query().Get("entry")
	if entryID == "" {
		entryID = r.PathValue("skill")
	}
	data, err := buildPageData(
		cat,
		entryID,
		r.URL.Query().Get("instruction"),
		r.URL.Query().Get("prompt_template"),
		splitSkillValues([]string{r.URL.Query().Get("skills")}),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "select.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// POST /chat
// Appends a user message + a simulated assistant reply. Swap this out
// for a real call into your agent/skill runner — stream via SSE
// (htmx's sse extension) instead of returning a full fragment if you
// want token-by-token output.
func chatHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	prompt := strings.TrimSpace(r.FormValue("prompt"))
	category := r.FormValue("category")
	instructionID := r.FormValue("instruction")
	promptTemplateID := r.FormValue("prompt_template")
	skillIDs := splitSkillValues(r.Form["skills"])
	if prompt == "" {
		return
	}
	pageData, err := buildPageData(category, instructionID, instructionID, promptTemplateID, skillIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_, _, ok := currentSelection(category, pageData.SelectedInstruction)
	if !ok {
		http.Error(w, "instruction not found", http.StatusNotFound)
		return
	}
	log.Printf(
		"chat request category=%q instruction=%q prompt_template=%q skills=%q model=%q prompt_chars=%d",
		category,
		pageData.SelectedInstruction,
		pageData.SelectedPrompt,
		strings.Join(pageData.SelectedSkills, ","),
		modelName(),
		len(prompt),
	)

	ctx, cancel := context.WithTimeout(r.Context(), 45*time.Second)
	defer cancel()

	reply, err := generateReply(ctx, category, pageData.SelectedInstruction, pageData.SelectedPrompt, pageData.SelectedSkills, prompt)
	if err != nil {
		log.Printf(
			"chat request failed category=%q instruction=%q prompt_template=%q skills=%q model=%q err=%v",
			category,
			pageData.SelectedInstruction,
			pageData.SelectedPrompt,
			strings.Join(pageData.SelectedSkills, ","),
			modelName(),
			err,
		)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	log.Printf(
		"chat request completed category=%q instruction=%q prompt_template=%q skills=%q model=%q reply_chars=%d",
		category,
		pageData.SelectedInstruction,
		pageData.SelectedPrompt,
		strings.Join(pageData.SelectedSkills, ","),
		modelName(),
		len(reply),
	)

	data := struct {
		Role  string
		Reply string
	}{
		Role:  pageData.CompositionSummary,
		Reply: reply,
	}

	if err := tmpl.ExecuteTemplate(w, "chat_reply.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ---------------------------------------------------------------------
// main
// ---------------------------------------------------------------------

func main() {
	if root := projectRoot(); root != "" {
		appRoot = filepath.Join(root, "app")
		staticRoot = filepath.Join(appRoot, "static")
		libraryRoot = filepath.Join(root, "rxlib")
	}
	if appRoot == "" {
		appRoot = "."
	}
	if staticRoot == "" {
		staticRoot = filepath.Join(appRoot, "static")
	}

	categories = loadCategories(rxlibRoot())

	funcs := template.FuncMap{
		"eq2": func(a, b string) bool { return a == b },
	}
	tmpl = template.Must(template.New("").Funcs(funcs).ParseGlob(filepath.Join(appRoot, "templates", "*.html")))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticRoot))))
	mux.HandleFunc("GET /{$}", indexHandler)
	mux.HandleFunc("GET /select", selectHandler)
	mux.HandleFunc("POST /chat", chatHandler)

	addr := ":8080"
	log.Printf("SkillRx listening on http://localhost%s (rxlib: %s)", addr, rxlibRoot())
	log.Fatal(http.ListenAndServe(addr, requestLogger(mux)))
}
