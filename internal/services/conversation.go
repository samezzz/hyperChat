package services

import (
	"github.com/samezzz/hyperchat/internal/models"
	"github.com/samezzz/hyperchat/internal/repository"
)

// HandleUserResponse processes user messages and handles the conversation flow
func HandleUserResponse(user string, messageBody string) {
	userState, exists := repository.GetUserState(user)
	if !exists {
		userState = models.NewUserState()
		repository.SaveUserState(user, userState)

		// First time interaction, prompt to begin onboarding
		SendMessage(user, "Hey there! I'm your personal health assistant for managing hypertension. Type 'Begin' to start or 'Skip' to skip onboarding.")
	}
}
