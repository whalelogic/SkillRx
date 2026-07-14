package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// ---------------------------------------------------------------------
// Data model
// ---------------------------------------------------------------------

type Category struct {
	Name   string
	Code   string
	Skills []string
}

type PageData struct {
	Categories     []Category
	ActiveCategory string
	ActiveSkill    string
	SourcePath     string
	SourceContent  string
}

var (
	categories []Category
	tmpl       *template.Template
)

// loadCategories walks rxlib/<category>/<skill>/SKILL.md on disk.
// Falls back to demo data if the path isn't found, so this runs standalone.
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
		skillDirs, err := os.ReadDir(filepath.Join(root, e.Name()))
		if err != nil {
			continue
		}
		var skills []string
		for _, s := range skillDirs {
			if s.IsDir() {
				skills = append(skills, s.Name())
			} else if strings.EqualFold(filepath.Ext(s.Name()), ".md") {
				skills = append(skills, strings.TrimSuffix(s.Name(), filepath.Ext(s.Name())))
			}
		}
		sort.Strings(skills)
		if len(skills) == 0 {
			continue
		}
		cats = append(cats, Category{
			Name:   e.Name(),
			Code:   fmt.Sprintf("C%02d", i),
			Skills: skills,
		})
	}
	if len(cats) == 0 {
		return demoCategories()
	}
	return cats
}

func demoCategories() []Category {
	raw := map[string][]string{
		"claude":               {"prompt-tuning", "tool-use-patterns", "context-window-mgmt"},
		"data-science":         {"eda-report", "feature-engineering", "model-eval"},
		"financial-markets":    {"earnings-summary", "portfolio-analysis", "risk-model"},
		"productivity":         {"meeting-notes", "task-breakdown", "weekly-review"},
		"repo-agents":          {"issue-triage", "pr-reviewer", "changelog-writer"},
		"research":             {"lit-review", "hypothesis-gen", "source-synthesis"},
		"repository-analysis":  {"dependency-audit", "codebase-map", "hotspot-detect"},
		"skill-templates":      {"base-template", "eval-harness", "readme-scaffold"},
		"software-engineering": {"code-review", "refactor-plan", "test-generation"},
		"ui-ux-development":    {"component-spec", "design-critique", "accessibility-pass"},
		"writing":              {"blog-post", "technical-doc", "editorial-pass"},
		"dotprompt":            {"schema-validate", "prompt-lint", "variant-compare"},
	}
	order := []string{"claude", "data-science", "financial-markets", "productivity",
		"repo-agents", "research", "repository-analysis", "skill-templates",
		"software-engineering", "ui-ux-development", "writing", "dotprompt"}

	var cats []Category
	for i, name := range order {
		cats = append(cats, Category{Name: name, Code: fmt.Sprintf("C%02d", i+1), Skills: raw[name]})
	}
	return cats
}

// readSkillSource loads the real SKILL.md if present, else a placeholder
// so the Source tab always has something to show.
func readSkillSource(root, category, skill string) (path, content string) {
	path = fmt.Sprintf("rxlib/%s/%s/SKILL.md", category, skill)
	full := filepath.Join(root, category, skill, "SKILL.md")
	if b, err := os.ReadFile(full); err == nil {
		return path, string(b)
	}
	content = fmt.Sprintf(`---
name: %s
category: %s
---

# %s

(No SKILL.md found on disk at %s — showing placeholder.
Point RXLIB_PATH at your real rxlib/ directory to load actual content.)`,
		skill, category, skill, full)
	return path, content
}

// ---------------------------------------------------------------------
// Handlers
// ---------------------------------------------------------------------

func rxlibRoot() string {
	if p := os.Getenv("RXLIB_PATH"); p != "" {
		return p
	}
	return "../rxlib"
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	active := categories[0]
	path, content := readSkillSource(rxlibRoot(), active.Name, active.Skills[0])

	data := PageData{
		Categories:     categories,
		ActiveCategory: active.Name,
		ActiveSkill:    active.Skills[0],
		SourcePath:     path,
		SourceContent:  content,
	}
	if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GET /select/{category}/{skill}
// Returns a fragment containing several out-of-band swaps: the topbar
// label, the sidebar active state, the source panel, the hidden form
// inputs, and a reset chat pane.
func selectHandler(w http.ResponseWriter, r *http.Request) {
	cat := r.PathValue("category")
	skill := r.PathValue("skill")

	path, content := readSkillSource(rxlibRoot(), cat, skill)
	data := PageData{
		Categories:     categories,
		ActiveCategory: cat,
		ActiveSkill:    skill,
		SourcePath:     path,
		SourceContent:  content,
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
	skill := r.FormValue("skill")
	if prompt == "" {
		return
	}

	// Simulate think-time; replace with your real agent call.
	time.Sleep(400 * time.Millisecond)

	reply := fmt.Sprintf(
		"**Simulated response** from `%s`.\n\nWire `chatHandler` in `main.go` up to your real "+
			"agent runner — load `rxlib/%s/%s/SKILL.md`, call the model, and return the reply "+
			"here (or stream it via SSE for token-by-token output).",
		skill, category, skill)

	data := struct {
		Skill  string
		Prompt string
		Reply  string
	}{Skill: skill, Prompt: prompt, Reply: reply}

	if err := tmpl.ExecuteTemplate(w, "chat_reply.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ---------------------------------------------------------------------
// main
// ---------------------------------------------------------------------

func main() {
	categories = loadCategories(rxlibRoot())

	funcs := template.FuncMap{
		"eq2": func(a, b string) bool { return a == b },
	}
	tmpl = template.Must(template.New("").Funcs(funcs).ParseGlob("templates/*.html"))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("GET /{$}", indexHandler)
	mux.HandleFunc("GET /select/{category}/{skill}", selectHandler)
	mux.HandleFunc("POST /chat", chatHandler)

	addr := ":8080"
	log.Printf("SkillRx listening on http://localhost%s (rxlib: %s)", addr, rxlibRoot())
	log.Fatal(http.ListenAndServe(addr, mux))
}
