package services

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func cleanNumber(num string) string {
	// Remove all spaces
	num = strings.ReplaceAll(num, " ", "")

	// Remove any non-printable characters
	num = strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, num)

	// If it already starts with whatsapp:, assume it's already in E.164
	if strings.HasPrefix(num, "whatsapp:") {
		return num
	}

	// Otherwise, ensure it starts with +
	if !strings.HasPrefix(num, "+") {
		num = "+" + num
	}

	// Then prepend whatsapp:
	return "whatsapp:" + num
}

func newTwilioClient() *twilio.RestClient {
	fmt.Println("TWILIO_ACCOUNT_SID: ", os.Getenv("TWILIO_ACCOUNT_SID"))
	fmt.Println("TWILIO_AUTH_TOKEN: ", os.Getenv("TWILIO_AUTH_TOKEN"))

	return twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
}

func SendMessage(to string, body string) error {
	client := newTwilioClient()
	params := &api.CreateMessageParams{}

	params.SetFrom("whatsapp:+14155238886") // Twilio Sandbox number
	params.SetTo(cleanNumber(to))
	params.SetBody(body)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return err
	}

	if resp.Body != nil {
		fmt.Println("Message sent:", *resp.Body)
	} else {
		fmt.Println("Message sent but no response body")
	}

	return nil
}

func SendContentTemplate(to string, sid string) error {
	client := newTwilioClient()

	params := &api.CreateMessageParams{}
	params.SetContentSid(sid)
	params.SetTo(cleanNumber(to))
	params.SetFrom("whatsapp:+14155238886")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		if resp.Body != nil {
			fmt.Println(*resp.Body)
		} else {
			fmt.Println(resp.Body)
		}
	}
	return nil
}

func ReviewTemplate(sid string) {
	client := newTwilioClient()

	resp, err := client.ContentV1.FetchContent(sid)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		if resp.DateCreated != nil {
			fmt.Println(*resp.DateCreated)
		} else {
			fmt.Println(resp.DateCreated)
		}
	}
}
