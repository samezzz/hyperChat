package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/samezzz/hyperchat/internal/services"
)

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

	fmt.Println("Extracted Body String:", bodyString)
	fmt.Println("Extracted Message From:", from)
	fmt.Println("Extracted Message Body:", messageBody)

	response := services.HandleUserResponse(from, messageBody)
	fmt.Fprintf(w, response)
}
