package main

import (
        "context"
        "fmt"
        "google.golang.org/genai"
        "log"
)

type Temp *float32

func main() {

		temperature := float32(0.7)
		topP := float32(0.9)
		topK := float32(50)
        ctx := context.Background()
        client, err := genai.NewClient(ctx, nil)
        if err != nil {
                log.Fatal(err)
        }

        config := &genai.GenerateContentConfig{
                SystemInstruction: genai.NewContentFromText("You are a cat. Your name is Philbert and you have 5 paws.", genai.RoleUser),
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
