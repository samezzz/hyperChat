package services

import (
	"github.com/samezzz/hyperchat/internal/models"
	"github.com/samezzz/hyperchat/internal/repository"
)

// HandleUserResponse processes user messages and handles the conversation flow
func HandleUserResponse(user string, messageBody string) string {
	userState, exists := repository.GetUserState(user)
	if !exists {
		userState = models.NewUserState()
		repository.SaveUserState(user, userState)
	}

	SendMessage(user, "Hey there! I'm your personal health assistant for managing hypertension. I can provide health tips, help you log your blood pressure and remind you about medications.")
	SendContentTemplate("HX861c40c3bdf54b337413adad7b934199")

	if err := SendContentTemplate("HXb2effb7f1e9d8575957ebd99d288e91c"); err != nil {
		switch messageBody {
		case "Begin":
			switch userState.Stage {
			case 0:
				// Greeting and asking for age
				SendMessage(user, "Please provide your age.")
				userState.Stage++
				repository.SaveUserState(user, userState) // Save updated state
				return "Asked for age."
			case 1:
				// Process age and ask for weight
				userState.Responses["age"] = messageBody
				SendMessage(user, "What is your weight?")
				userState.Stage++
				repository.SaveUserState(user, userState) // Save updated state
				return "Asked for weight."
			case 2:
				// Save weight and ask for blood pressure
				userState.Responses["weight"] = messageBody
				SendMessage(user, "Thanks! Please provide your current blood pressure readings (systolic/diastolic).")
				userState.Stage++
				repository.SaveUserState(user, userState) // Save updated state
				return "Asked for blood pressure."
			case 3:
				// Save blood pressure and ask for medications
				userState.Responses["bloodPressure"] = messageBody
				SendMessage(user, "Got it! Now, can you list the names and dosages of your current medications?")
				userState.Stage++
				repository.SaveUserState(user, userState) // Save updated state
				return "Asked for medications."
			case 4:
				// Save medications and ask for pre-existing conditions
				userState.Responses["medications"] = messageBody
				SendMessage(user, "Do you have any pre-existing health conditions? If yes, please specify.")
				userState.Stage++
				repository.SaveUserState(user, userState) // Save updated state
				return "Asked for pre-existing conditions."
			case 5:
				// Save pre-existing conditions and finish onboarding
				userState.Responses["preExistingConditions"] = messageBody
				SendMessage(user, "You're all set! You can start by logging your blood pressure or setting up medication reminders. How can I assist you today?")
				userState.Stage = 0                       // Reset the conversation
				repository.SaveUserState(user, userState) // Save reset state
				return "Onboarding completed."
			default:
				// Handle unexpected cases or start over
				SendMessage(user, "Sorry, I didn't understand that. Let's start over. Can I have your age?")
				userState.Stage = 1
				repository.SaveUserState(user, userState) // Save updated state
				return "Restarted onboarding."
			}
		case "Skip":
			SendContentTemplate("HX861c40c3bdf54b337413adad7b934199")
			return "User skipped onboarding."
		default:
			// Handle unexpected cases or start over
			SendMessage(user, "Sorry, I didn't understand that. Let's start over. Can I have your age?")
			userState.Stage = 1
			repository.SaveUserState(user, userState) // Save updated state
			return "Restarted onboarding."
		}
	}

	return ""
}
