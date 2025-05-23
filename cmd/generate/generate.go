package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func generate() {
	// Define the model path and prompt
	modelPath := "/mnt/c/Users/SamEss/Desktop/programming/go/projects/hyperModel/model/pretrained_model/llama3-8b-hyperModel-q4-gguf-original/unsloth.Q4_K_M.gguf"
	prompt := "What is hypertension?"
	numTokens := 50
	temperature := 0.2

	// Create the command
	cmd := exec.Command("llama-cli",
		"-m", modelPath,
		"-p", prompt,
		"-n", fmt.Sprint(numTokens),
		"--temp", fmt.Sprintf("%.1f", temperature),
	)

	// Capture the output
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	// Run the command
	if err := cmd.Run(); err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	// Print the output
	fmt.Println("Response from llama-cli:")
	fmt.Println(out.String())
}
