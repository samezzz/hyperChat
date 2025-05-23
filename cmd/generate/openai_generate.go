package main

import (
	"context"
	"fmt"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func openaiGenerate() {
	// Load the OpenAI API key from the environment variable
	apiKey := os.Getenv("TEST_HYPERCHAT_API_KEY")
	if apiKey == "" {
		log.Fatal("TEST_HYPERCHAT_API_KEY environment variable is not set")
	}

	// Initialize the OpenAI client with the API key
	client := openai.NewClient(apiKey)

	// Create a completion request with a prompt
	req := openai.CompletionRequest{
		Model:       openai.GPT3Davinci002, // Use a GPT model like "text-davinci-003"
		Prompt:      "What is hypertension?",
		MaxTokens:   100, // Adjust max tokens for response length
		Temperature: 0.5, // Adjust for creativity (0.0 to 1.0)
	}

	// Make the request to OpenAI
	response, err := client.CreateCompletion(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to create completion: %v", err)
	}

	// Print the response
	fmt.Println(response.Choices[0].Text)
}
