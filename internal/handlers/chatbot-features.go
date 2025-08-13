package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/bregydoc/gtranslate"
	"github.com/samezzz/hyperchat/internal/models"
	"github.com/samezzz/hyperchat/internal/repository"
	"github.com/samezzz/hyperchat/internal/services"
)

// Define a struct to parse the JSON response
var parsedResponse struct {
	Result string `json:"result"`
}

func HandleChatbotFeatures(from, messageBody string) {
	userState, exists := repository.GetUserState(from)
	if !exists {
		log.Printf("User state not found for %s", from)
		services.SendMessage(from, "Sorry, there was an error. Please try again.")
		return
	}

	if !userState.FeatureSelected {
		feature := selectFeature(from, messageBody)
		if feature != "" {
			userState.FeatureSelected = true
			userState.CurrentFeature = feature
			repository.SaveUserState(from, userState)
		} else {
			showFeatureMenu(from)
			return
		}
	}

	switch userState.CurrentFeature {
	case "health_tips":
		handleHealthTips(from, userState)
	case "log_bp":
		handleBloodPressureLogging(from, messageBody, userState)
	case "check_bp":
		handleCheckBloodPressure(from, messageBody, userState)
	case "chatbot":
		handleChatbot(from, messageBody)
	case "multilanguage":
		handleMultilanguage(from, messageBody, userState)
	case "onboard":
		HandleOnboarding(from, messageBody, userState)
	default:
		services.SendMessage(from, "I'm not sure how to help with that. Let's go back to the main menu.")
		resetUserState(from, userState)
	}
}

func selectFeature(from, messageBody string) string {
	switch strings.ToLower(messageBody) {
	case "1", "health tips", "tips", "health_tips", "👩‍⚕️ Health Tips\nReceive personalized advice to manage your blood pressure effectively":
		return "health_tips"
	case "2", "log blood pressure", "log_bloold_pressure", "log bp", "📊 Log Blood Pressure\nKeep track of your blood pressure readings and monitor changes over time":
		return "log_bp"
	case "3", "chat", "chatbot", "🤖 Conversation with Bot\nInteractive conversation with our chatbot to answer all your questions":
		return "chatbot"
	case "4", "check blood pressure", "check_blood_pressure", "check bp", "🩺 Check Blood Pressure\nOpen our website and check your blood pressure immediately":
		return "check_bp"
	case "5", "select other language", "multilanguage", "🌍 Select other language\nHave conversations in our catalogue of languages":
		return "multilanguage"
	case "6", "onboard", "👋 Onboard\nAnswer few health related questions to tailor recommendations for you.":
		return "onboard"
	case "begin":
		return "onboard"
	default:
		if messageBody == "Thank you for completing the onboarding process. You can now start using the features of the app!" || messageBody == "Skip" {
			return ""
		} else {
			services.SendMessage(from, "Sorry I didn't quite get that")
			return ""
		}
	}
}

// BLOOD PRESSURE LOGGING
func handleBloodPressureLogging(from, messageBody string, userState *models.UserState) {
	if userState.BPLogStage == 0 {
		services.SendMessage(from, "Please enter your blood pressure reading (e.g., 120/80).")
		userState.BPLogStage = 1
		repository.SaveUserState(from, userState)
	} else {
		// Process the blood pressure reading
		services.SendMessage(from, "Thank you for logging your blood pressure.")
		response, err := services.GenerateResponse("Hey gpt, this is my blood pressure reading, " + messageBody + ". I want you to give me a brief advice on what to do with a blood pressure like this. Don't talk much. I really don't want you to generate more than 5 lines. Just give me tips on how to manage blood pressure like this")
		if err != nil {
			log.Printf("Error generating response: %v", err)
			services.SendMessage(from, "Sorry, I encountered an error. Please try again.")
			return
		}

		// Parse the JSON response to extract only the 'result' field
		err = json.Unmarshal([]byte(response), &parsedResponse)
		if err != nil {
			log.Printf("Error parsing response: %v", err)
			services.SendMessage(from, "Sorry, I encountered an error while processing the response. Please try again.")
			return
		}

		// Send only the 'result' value to the user
		if parsedResponse.Result != "" {
			services.SendMessage(from, parsedResponse.Result)
		} else {
			// In case the result field is empty or there's an issue with the response
			services.SendMessage(from, "I'm sorry, I couldn't generate a response. Please try again.")
		}

		resetUserState(from, userState)
	}
}

// HEALTH TIPS
func handleHealthTips(from string, userState *models.UserState) {
	// Generate the chatbot response from the service
	response, err := services.GenerateResponse("Hey gpt, I want you to provide some health tips to manage hypertension. So if you're to be a medical health practitioner with tons of experience, how would you summarize some health tips that would benefit me if I want to manage hypertension. This is a WhatsApp message so the reader doesn't want any long talk. Don't talk much. Just tell me what's necessary. I really don't want you to generate more than 5 lines. Provide concise, clear, and evidence-based answers about hypertension management. Focus on key points such as diagnosis, lifestyle changes, medications, monitoring, and complications. Keep responses brief and easy to understand.")
	if err != nil {
		log.Printf("Error generating response: %v", err)
		services.SendMessage(from, "Sorry, I encountered an error. Please try again.")
		return
	}

	// Parse the JSON response to extract only the 'result' field
	err = json.Unmarshal([]byte(response), &parsedResponse)
	if err != nil {
		log.Printf("Error parsing response: %v", err)
		services.SendMessage(from, "Sorry, I encountered an error while processing the response. Please try again.")
		return
	}

	// Send only the 'result' value to the user
	if parsedResponse.Result != "" {
		services.SendMessage(from, parsedResponse.Result)
	} else {
		// In case the result field is empty or there's an issue with the response
		services.SendMessage(from, "I'm sorry, I couldn't generate a response. Please try again.")
	}

	resetUserState(from, userState)
}

// CHATBOT
func handleChatbot(from, messageBody string) {
	// Retrieve user state
	userState, exists := repository.GetUserState(from)
	if !exists {
		log.Printf("User state not found for %s", from)
		services.SendMessage(from, "Sorry, there was an error. Please try again.")
		return
	}

	// Default to English if no language is set
	language := userState.LanguageSelected
	if language == "" {
		language = "en"
	}

	// If user requests the 'menu', reset their state and offer the main menu
	if strings.ToLower(messageBody) == "menu" {
		resetUserState(from, nil)
		return
	}

	// Generate the chatbot response
	response, err := services.GenerateResponse(
		"Hey gpt, you're a specialist in hypertension management. I have a question for you about hypertension. Don't talk much. Just tell me what's necessary. Provide concise, clear, and evidence-based answers about the question I have for you. Focus on key points such as diagnosis, lifestyle changes, medications, monitorthe question I have for you Keep responses brief and easy to understand. I really don't want you to generate more than 5 lines. Now wait for me to ask my question. Always think deeply for insights you can share. Your responses should be unique. " + messageBody,
	)
	if err != nil {
		log.Printf("Error generating response: %v", err)
		services.SendMessage(from, "Sorry, I encountered an error. Please try again.")
		return
	}

	// Try parsing as JSON (new API format)
	var parsed struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}
	parsedResult := strings.TrimSpace(response)

	// If it looks like JSON, try unmarshalling
	if strings.HasPrefix(parsedResult, "{") {
		if err := json.Unmarshal([]byte(parsedResult), &parsed); err == nil {
			if len(parsed.Candidates) > 0 && len(parsed.Candidates[0].Content.Parts) > 0 {
				parsedResult = parsed.Candidates[0].Content.Parts[0].Text
			}
		}
	}

	// Translate the response based on user preference
	if parsedResult != "" {
		translatedText, err := gtranslate.TranslateWithParams(parsedResult, gtranslate.TranslationParams{
			From: "auto",
			To:   language, // Use the stored language
		})
		if err != nil {
			log.Printf("Translation error: %v", err)
			services.SendMessage(from, "Sorry, I encountered an error while translating the response. Please try again.")
			return
		}

		// Send translated response
		services.SendMessage(from, translatedText)
		services.SendMessage(from, "Type 'menu' to return to the main menu or continue chatting.")
	} else {
		services.SendMessage(from, "I'm sorry, I couldn't generate a response. Please try again.")
	}
}

// CHECK BLOOD PRESSURE
func handleCheckBloodPressure(from, messageBody string, userState *models.UserState) {
	// Create a web link that redirects to your app
	// This will be clickable in WhatsApp and can handle the app opening logic
	webRedirectLink := "https://hyperchat.up.railway.app/open-app"

	// Fallback download link
	fallbackLink := "https://hyperchat.up.railway.app/"

	msg := fmt.Sprintf(
		"🩺 *Blood Pressure Checker*\n\n"+
			"Tap here to open the app:\n%s\n\n"+
			"📱 Don't have the app? Download it here:\n%s",
		webRedirectLink, fallbackLink,
	)

	if err := services.SendMessage(from, msg); err != nil {
		log.Printf("Error sending check BP message: %v", err)
	}
	resetUserState(from, userState)
}

// MULTILANGUAGE
func handleMultilanguage(from, messageBody string, userState *models.UserState) {
	// Check if the user has already selected a language
	if userState.LanguageSelected == "" && userState.LanguageStage == 0 {
		// Show the language selection menu
		fromNumber := strings.TrimPrefix(from, "whatsapp:")
		services.SendContentTemplate(fromNumber, "HX5be2530d319980d8f9874136ead15eda")
		userState.LanguageStage = 1
		repository.SaveUserState(from, userState)
		return
	}

	// Handle language selection
	if userState.LanguageStage == 1 {
		switch strings.ToLower(messageBody) {
		case "1", "french", "français", "🇫🇷 french (français)", "select_french", "🇫🇷 French (Français)\nContinue in French for a fully translated experience.":
			userState.LanguageSelected = "fr"
			services.SendMessage(from, "Vous avez sélectionné le français. Vous pouvez maintenant converser avec le chatbot en français.")
		case "2", "twi", "akan", "🇬🇭 twi (akan)", "select_twi", "🇬🇭 Twi (Akan)\nKɔ so wɔ Akan kasa de gye wo ho.":
			userState.LanguageSelected = "ak"
			services.SendMessage(from, "Woapaw Akan kasa. Afei, wubetumi aka nsɛm wɔ Akan mu kyerɛ chatbot no.")
		case "3", "english", "🌍 english", "select_english", "🌍 English\nStay in English and continue as usual.":
			userState.LanguageSelected = "en"
			services.SendMessage(from, "You have selected English. You can now chat with the bot in English.")
		default:
			services.SendMessage(from, "Invalid selection. Please choose a valid option:\n1. 🇫🇷 French\n2. 🇬🇭 Twi\n3. 🌍 English (default)")
			return
		}

		// Save the user's language preference and move to chatbot
		userState.LanguageStage = 0
		userState.FeatureSelected = true
		userState.CurrentFeature = "chatbot"
		repository.SaveUserState(from, userState)
		handleChatbot(from, "Hey, all responses should be summarized")
	}
}

// TODO: SET REMINDERS
func handleSetReminders(from, messageBody string, userState *models.UserState) {
	if userState.ReminderStage == 0 {
		services.SendMessage(from, "What would you like to set a reminder for?")
		userState.ReminderStage = 1
		repository.SaveUserState(from, userState)
	} else {
		// Process the reminder
		// TODO: Implement reminder setting logic
		services.SendMessage(from, "Reminder set successfully.")
		resetUserState(from, userState)
	}
}

func resetUserState(from string, userState *models.UserState) {
	if userState == nil {
		userState = &models.UserState{}
	}
	userState.FeatureSelected = false
	userState.CurrentFeature = ""
	userState.BPLogStage = 0
	userState.ReminderStage = 0
	repository.SaveUserState(from, userState)
	showFeatureMenu(from)
}

func showFeatureMenu(from string) {
	fromNumber := strings.TrimPrefix(from, "whatsapp:")
	services.SendContentTemplate(fromNumber, "HXb2182414459a8e5987e913308b3cbc1e")
}
