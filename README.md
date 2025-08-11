# HyperChat – AI Health Assistant for Hypertension Management 📱💬

HyperChat is an AI-powered WhatsApp chatbot designed to help individuals manage hypertension.  
It provides personalized health tips, onboarding questionnaires, multilingual support, and can integrate with existing health apps.

Built with **Go** and powered by **Twilio WhatsApp API**, HyperChat runs seamlessly in the cloud via services like **Railway**.

---

## 🚀 Features

- **Personalized Onboarding** – Collects essential user info to tailor health advice.
- **Hypertension Management Tips** – Evidence-based lifestyle recommendations.
- **Multi-language Support** – Communicate in multiple languages.
- **App Redirection** – Redirects to existing health tracking apps (or download link if not installed).
- **AI-Generated Responses** – Uses AI to provide accurate, concise, and context-aware answers.
- **WhatsApp Integration** – Fully functional WhatsApp bot using Twilio API.

---

## 🛠 Tech Stack

- **Language**: Go (Golang)
- **Messaging API**: Twilio WhatsApp API
- **Deployment**: Railway / Docker
- **AI Processing**: OpenAI API (or other LLM backend)
- **Data Storage**: In-memory store (can be swapped for DB)

---

## 📂 Project Structure

hyperchat/
├── cmd/
│ └── server/
│ └── main.go # Entry point of the server
├── internal/
│ ├── handlers/ # HTTP handlers (Twilio webhook, etc.)
│ ├── models/ # Data models & state
│ ├── repository/ # User state storage
│ └── services/ # Twilio, AI, and utility services
├── .env # Environment variables (not committed)
├── Dockerfile # Container build file
├── go.mod # Go modules file
└── README.md # This file

---

## ⚙️ Setup & Installation

### 1️⃣ Prerequisites

- [Go 1.22+](https://go.dev/dl/)
- [Docker](https://www.docker.com/) (for containerization)
- [Twilio Account](https://www.twilio.com/try-twilio) with WhatsApp Sandbox enabled
- OpenAI API key (if using AI)

---

### 2️⃣ Clone Repository

```bash
git clone https://github.com/yourusername/hyperchat.git
cd hyperchat

```

Set Environment Variables
Create a .env file:
TWILIO_ACCOUNT_SID=ACxxxxxxxxxxxxxxxxxxxxxx
TWILIO_AUTH_TOKEN=your_auth_token
TWILIO_WHATSAPP_NUMBER=whatsapp:+14155238886
OPENAI_API_KEY=sk-xxxxxxxxxxxxxxxx

Run Locally (Without Docker)
docker build -t hyperchat .

Run the container
docker run -p 8080:8080 --env-file .env hyperchat
