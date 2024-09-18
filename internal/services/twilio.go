package services

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendMessage(to string, body string) error {
	client := twilio.NewRestClient()

	params := &api.CreateMessageParams{}
	params.SetFrom("whatsapp:+14155238886") // Twilio WhatsApp Sandbox number
	params.SetTo("whatsapp:+233553865162")
	params.SetBody(body)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return err
	}

	if resp.Body != nil {
		fmt.Println("Message sent: ", *resp.Body)
	} else {
		fmt.Println("Message sent but no response body")
	}

	return nil
}

func SendContentTemplate(sid string) error {
	client := twilio.NewRestClient()

	params := &api.CreateMessageParams{}
	params.SetContentSid(sid)
	params.SetTo("whatsapp:+233553865162")
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
	client := twilio.NewRestClient()

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
