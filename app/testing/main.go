package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


func LoadApiKey() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		fmt.Println("GEMINI_API_KEY is not set")
	} else {
		fmt.Println("GEMINI_API_KEY loaded successfully")
	}
	fmt.Println("GEMINI_API_KEY:", apiKey)
}


func main() {
	LoadApiKey()
}
