package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/samezzz/hyperchat/internal/models"
	"github.com/samezzz/hyperchat/internal/repository"
	"github.com/samezzz/hyperchat/internal/services"
)

type AIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	bodyString := string(bodyBytes)

	values, err := url.ParseQuery(bodyString)
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusInternalServerError)
		return
	}

	from := values.Get("From")
	messageBody := values.Get("Body")

	// Log extracted message
	fmt.Println("Extracted Message From:", from)
	fmt.Println("Extracted Message Body:", messageBody)

	// Check if user exists in the repository
	userState, exists := repository.GetUserState(from)
	if !exists {
		// First-time user, initiate onboarding flow
		userState = models.NewUserState()
		repository.SaveUserState(from, userState)

		// Send welcome message and onboarding instructions
		services.SendContentTemplate("HX3167e0909369dbe82b39866dbb0a1d76")
		return
	}

	// Route based on user state
	if userState.Onboarding {
		HandleOnboarding(from, messageBody, userState)
	} else {
		// If onboarding is done, route to features
		HandleChatbotFeatures(from, messageBody)
	}
}
