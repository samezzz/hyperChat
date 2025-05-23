package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// func GenerateResponse(prompt string) (string, error) {
// 	modelPath := "/mnt/c/Users/SamEss/Desktop/programming/go/projects/hyperModel/model/pretrained_model/llama3-8b-hyperModel-q4-gguf-original/unsloth.Q4_K_M.gguf"
// 	numTokens := 5
// 	temperature := 0.9
//
// 	// Omit the timeout context to allow long-running commands to finish
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
//
// 	// Prepare llama-cli command with appropriate arguments
// 	cmd := exec.CommandContext(ctx, "llama-cli",
// 		"-m", modelPath,
// 		"-p", prompt,
// 		"-n", fmt.Sprint(numTokens),
// 		"--temp", fmt.Sprintf("%.1f", temperature),
// 		"--ctx_size", "2048",
// 		"--batch_size", "16",
// 		"--top_k", "50",
// 		"--top_p", "0.9",
// 	)
//
// 	var stdout, stderr bytes.Buffer
// 	cmd.Stdout = &stdout
// 	cmd.Stderr = &stderr
//
// 	// Execute the command and capture stdout and stderr
// 	if err := cmd.Run(); err != nil {
// 		return "", fmt.Errorf("failed to execute llama-cli: %v, stderr: %v", err, stderr.String())
// 	}
//
// 	return stdout.String(), nil
// }
//
// func ExtractResponse(text string) string {
// 	// Split the text into lines
// 	lines := strings.Split(text, "\n")
//
// 	// Define unwanted keywords
// 	keywords := []string{
// 		"llama", "kv", "loader", "model", "init",
// 		"token", "sampler", "context", "graph", "CPU",
// 	}
//
// 	// Filter out lines that contain unwanted keywords
// 	var filteredLines []string
// 	for _, line := range lines {
// 		shouldExclude := false
// 		for _, keyword := range keywords {
// 			if strings.Contains(line, keyword) {
// 				shouldExclude = true
// 				break
// 			}
// 		}
// 		if !shouldExclude {
// 			filteredLines = append(filteredLines, line)
// 		}
// 	}
//
// 	// Join the filtered lines back together and trim whitespace
// 	response := strings.Join(filteredLines, "\n")
// 	return strings.TrimSpace(response)
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	// Extract prompt from URL query
// 	prompt := r.URL.Query().Get("prompt")
// 	if prompt == "" {
// 		http.Error(w, "Prompt is required", http.StatusBadRequest)
// 		return
// 	}
//
// 	log.Printf("Received prompt: %s", prompt)
//
// 	// Generate response using llama-cli
// 	response, err := generateResponse(prompt)
// 	if err != nil {
// 		log.Printf("Error generating response: %v", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
//
// 	// Filter and return generated response to the client
// 	filteredResponse := extractResponse(response)
// 	fmt.Fprintln(w, filteredResponse)
// }

func GenerateResponse(prompt string) (string, error) {
	// Define the request body to match the curl request
	body := map[string]interface{}{
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
		"web_access": false, // Add this field to match the curl request
	}

	// Marshal the body into JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("Error marshalling request body: %v", err)
	}

	// Create a new HTTP request with the correct URL
	req, err := http.NewRequest("POST", "https://chatgpt-42.p.rapidapi.com/chatgpt", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set the headers to match the curl request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-rapidapi-host", "chatgpt-42.p.rapidapi.com")
	// req.Header.Set("x-rapidapi-key", "ea992e16c5msh521cca4f62eacacp11f229jsn0a18ae384699") // Replace this with your actual API key
	req.Header.Set("x-rapidapi-key", "e3fe1637e0mshb24301e54891c05p1ab219jsn25766d5046ff") // Replace this with your actual API key

	// Create a new HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Print and return the response
	fmt.Println(string(bodyBytes))
	return string(bodyBytes), nil
}

// func GenerateResponse(prompt string) (string, error) {
// 	// ✅ Define the request body, including the "model" field (required!)
// 	body := map[string]interface{}{
// 		"model":       "grok-2-latest", // This was missing!
// 		"stream":      false,
// 		"temperature": 0,
// 		"messages": []map[string]string{
// 			{
// 				"role":    "user",
// 				"content": prompt,
// 			},
// 		},
// 	}
//
// 	// ✅ Marshal the body into JSON
// 	jsonBody, err := json.Marshal(body)
// 	if err != nil {
// 		return "", fmt.Errorf("error marshalling request body: %v", err)
// 	}
//
// 	// ✅ Create a new HTTP request
// 	req, err := http.NewRequest("POST", "https://api.x.ai/v1/chat/completions", bytes.NewBuffer(jsonBody))
// 	if err != nil {
// 		return "", fmt.Errorf("error creating request: %v", err)
// 	}
//
// 	// ✅ Set headers
// 	apiKey := "xai-mCIA8AoHc6Xe8nkklqoaJDdF6C8E7R5j4fRFyQdKPTuIHRZLCIgH5wkZsLWEpVXcfwe9uVUUTVQPzAVW"
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", "Bearer "+apiKey)
//
// 	// ✅ Create client and send the request
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return "", fmt.Errorf("error sending request: %v", err)
// 	}
// 	defer resp.Body.Close()
//
// 	// ✅ Read the response body
// 	bodyBytes, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", fmt.Errorf("error reading response body: %v", err)
// 	}
//
// 	// Optional: Check status code for debugging
// 	if resp.StatusCode != http.StatusOK {
// 		return "", fmt.Errorf("received non-200 response: %d\n%s", resp.StatusCode, bodyBytes)
// 	}
//
// 	// ✅ Print and return the response body as a string
// 	fmt.Println("Response Body:")
// 	fmt.Println(string(bodyBytes))
// 	return string(bodyBytes), nil
// }
