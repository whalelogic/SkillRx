package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// ---------------------------------------------------------------------
// Data models
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
	Categories     []Category
	ActiveCategory string
	ActiveEntry    string
	ActiveLabel    string
	ModelName      string
	SourcePath     string
	SourceContent  string
}

// Default values
const defaultModel = "gemini-3-flash-preview"
const defaultCategory = "financial-markets"
const defaultEntry = "system-instructions/market-analyst.md"


// API key loading function
func LoadApiKey() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv("GEMINI_API_KEY")
}

