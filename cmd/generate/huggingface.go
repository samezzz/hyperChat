package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const API_URL = "https://api-inference.huggingface.co/models/Hyperchat/llama3-8b-hyperModel-gguf" // Replace with your model

// Struct to define the request body (input to the model)
type RequestBody struct {
	Inputs string `json:"inputs"`
}

// Struct to handle the response
type ResponseBody struct {
	GeneratedText string `json:"generated_text"`
}

func queryHuggingFaceAPI(prompt string) (string, error) {
	// Replace with your Hugging Face API token
	apiToken := os.Getenv("HF_ACCESS_TOKEN")
	if apiToken == "" {
		return "", fmt.Errorf("Hugging Face API token not set in environment variables")
	}

	// Prepare the request body
	body := RequestBody{Inputs: prompt}
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", API_URL, bytes.NewBuffer(bodyJson))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers, including authorization
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	// Decode the response
	var responseBody []ResponseBody
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	// Return the generated text from the response
	if len(responseBody) > 0 {
		return responseBody[0].GeneratedText, nil
	}

	return "", fmt.Errorf("empty response from Hugging Face API")
}

func main() {
	// Example prompt
	prompt := "What is hypertension?"

	// Query Hugging Face API
	response, err := queryHuggingFaceAPI(prompt)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Print the generated response
	fmt.Println("Generated Response:")
	fmt.Println(response)
}
