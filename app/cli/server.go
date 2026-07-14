package main

import (
        "context"
		"os"
        "fmt"
        "google.golang.org/genai"
        "log"
)

type Temp *float32

func main() {

		f, err := os.Open("/home/whaler/github/SkillRx/financial-markets/system-instructions/market_pulse.md")
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		stringData := make([]byte, 1000)
		_, err = f.Read(stringData)
		if err != nil {
			log.Fatal(err)
		}

		temperature := float32(0.7)
		topP := float32(0.9)
		topK := float32(50)
        ctx := context.Background()
        client, err := genai.NewClient(ctx, nil)
        if err != nil {
                log.Fatal(err)
        }

        config := &genai.GenerateContentConfig{
                SystemInstruction: genai.NewContentFromText(string(stringData), genai.RoleUser),
                Temperature:       &temperature,
				TopP:             &topP,
				TopK:             &topK,
        }

        result, _ := client.Models.GenerateContent(
                ctx,
                "gemini-3-flash-preview",
                genai.Text("Hello, cat. What's your favorite programming language?"),
                config,
        )

        fmt.Println(result.Text())
}
