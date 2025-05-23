package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// StatusCallback represents the structure of the incoming data from Twilio

type StatusCallback struct {
	MessageSid    string `json:"MessageSid"`
	MessageStatus string `json:"MessageStatus"`
	To            string `json:"To"`
	From          string `json:"From"`
}

// statusCallbackHandler handles incoming status updates from Twilio
func StatusCallbackHandler(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON payload
	var callbackData StatusCallback
	err := json.NewDecoder(r.Body).Decode(&callbackData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Log the incoming status update
	log.Printf("Received status update for Message SID: %s, Status: %s, From: %s, To: %s",
		callbackData.MessageSid, callbackData.MessageStatus, callbackData.From, callbackData.To)

	// Handle different message statuses
	switch callbackData.MessageStatus {
	case "queued":
		// Message has been queued for sending
		log.Println("Message is queued.")
	case "sent":
		// Message has been sent
		log.Println("Message has been sent.")
	case "delivered":
		// Message has been delivered to the recipient
		log.Println("Message has been delivered.")
	case "undelivered":
		// Message was undelivered
		log.Println("Message was undelivered.")
	case "failed":
		// Message failed to send
		log.Println("Message failed to send.")
	default:
		log.Println("Unknown message status.")
	}

	// Respond to Twilio to acknowledge receipt of the callback
	w.WriteHeader(http.StatusOK)
}
