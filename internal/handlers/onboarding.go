package handlers

import (
	"strings"

	"github.com/samezzz/hyperchat/internal/models"
	"github.com/samezzz/hyperchat/internal/repository"
	"github.com/samezzz/hyperchat/internal/services"
)

func HandleOnboarding(from, messageBody string, userState *models.UserState) {
	messageBody = strings.ToLower(messageBody)

	switch userState.Stage {
	case 0:
		if messageBody == "begin" {
			services.SendMessage(from, "Please provide your age.")
			userState.Stage = 1
		} else if messageBody == "skip" {
			skipOnboarding(from, userState)
		} else {
			services.SendMessage(from, "Sorry, I didn't quite get that. Please type 'begin' to start the onboarding process or 'skip' to skip it.")
		}
	case 1:
		userState.Age = messageBody
		services.SendMessage(from, "Great, now please tell me your weight.")
		userState.Stage = 2
	case 2:
		userState.Weight = messageBody
		services.SendMessage(from, "Got it! Please provide your current blood pressure readings (systolic/diastolic).")
		userState.Stage = 3
	case 3:
		userState.BloodPressure = messageBody
		services.SendMessage(from, "Now, can you list the names and dosages of your current medications, if any?")
		userState.Stage = 4
	case 4:
		userState.Medications = messageBody
		services.SendMessage(from, "Do you have any pre-existing health conditions? If yes, please specify.")
		userState.Stage = 5
	case 5:
		userState.PreExistingConditions = messageBody
		completeOnboarding(from, userState)
	default:
		services.SendMessage(from, "Sorry, I didn't understand. Let's continue onboarding.")
	}

	repository.SaveUserState(from, userState)
}

func completeOnboarding(from string, userState *models.UserState) {
	userState.Onboarding = false
	userState.FeatureSelected = false
	repository.SaveUserState(from, userState)
	services.SendMessage(from, "Thank you for completing the onboarding process. You can now start using the features of the app!")
	showFeatureMenu(from)
}

func skipOnboarding(from string, userState *models.UserState) {
	userState.Onboarding = false
	userState.FeatureSelected = false
	repository.SaveUserState(from, userState)
	showFeatureMenu(from)
}
