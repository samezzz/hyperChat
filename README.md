# HyperChat – AI Health Assistant for Hypertension Management 📱💬

**HyperChat** is an AI-powered WhatsApp chatbot designed to help individuals manage hypertension.  
It provides **personalized health tips**, onboarding questionnaires, multilingual support, and can even redirect users to integrated health tracking apps.

Built with **Go**, powered by **Twilio WhatsApp API**, and deployable to the cloud with **Railway**, HyperChat is lightweight, fast, and reliable.

---

## 🚀 Features

- **Personalized Onboarding** – Tailors health advice to each user’s profile.
- **Hypertension Management Tips** – Evidence-based lifestyle recommendations.
- **Multi-language Support** – Interact in your preferred language.
- **App Redirection** – Opens an installed health app or sends a download link.
- **AI-Generated Responses** – Context-aware, concise, and helpful answers.
- **WhatsApp Integration** – Seamless messaging via Twilio API.

---

## 🛠 Tech Stack

- **Language**: Go (Golang)
- **Messaging API**: Twilio WhatsApp API
- **AI Processing**: OpenAI API (or other LLMs)
- **Deployment**: Railway / Docker
- **Storage**: In-memory store (can be replaced with a DB)

---

## 📂 Project Structure

hyperchat/
├── cmd/
│ └── server/
│ └── main.go # Server entry point
├── internal/
│ ├── handlers/ # HTTP handlers (Twilio webhook, etc.)
│ ├── models/ # Data models & state definitions
│ ├── repository/ # User state storage
│ └── services/ # Twilio, AI, and utility services
├── .env # Environment variables (not committed)
├── Dockerfile # Container build file
├── go.mod # Go modules file
└── README.md # This file

yaml
Copy
Edit

---

## ⚙️ Setup & Installation

### 1️⃣ Prerequisites

- [Go 1.22+](https://go.dev/dl/)
- [Docker](https://www.docker.com/) (optional, for containerization)
- [Twilio Account](https://www.twilio.com/try-twilio) with WhatsApp Sandbox enabled
- **(Optional)** OpenAI API key for AI features

---

### 2️⃣ Clone the Repository

```bash
git clone https://github.com/yourusername/hyperchat.git
cd hyperchat
3️⃣ Configure Environment Variables
```

## Create a .env file in the project root:

TWILIO_ACCOUNT_SID=ACxxxxxxxxxxxxxxxxxxxxxx
TWILIO_AUTH_TOKEN=your_auth_token
TWILIO_WHATSAPP_NUMBER=whatsapp:+14155238886
OPENAI_API_KEY=sk-xxxxxxxxxxxxxxxx

## Run Locally (Without Docker)

```bash
go mod tidy
go run cmd/server/main.go
Server will start on localhost:8080.
```

## Run with Docker

Build the image:

```bash
docker build -t hyperchat .
```

## Run the container:

docker run -p 8080:8080 --env-file .env hyperchat

## 📡 Twilio Webhook Setup

Go to Messaging → WhatsApp Sandbox in the Twilio Console.

## Set the Webhook URL for incoming messages:

📌 Example Interaction
User: Hi
Bot: Hey there! I'm your personal health assistant for managing hypertension...
User: health_tips
Bot: 1. Follow the DASH diet...
📜 License
MIT License © 2025 [Samess]

---

```

```
