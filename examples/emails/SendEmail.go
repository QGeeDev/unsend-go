package examples

import (
	"context"
	"fmt"
	"os"

	"github.com/QGeeDev/unsend-go"
)

func SendEmail() {
	client, err := unsend.NewClient()

	if err != nil {
		fmt.Printf(fmt.Sprintf("[ERROR] - %s\n", err.Error()))
		os.Exit(1)
	}

	request := &unsend.SendEmailRequest{
		To: []string{"qg@qgdev.co.uk"},
		From: "hello@updates.qgtest.dev",
		Subject: "Hello, World!",
		Text: "This is a test email",
	}

	response, _ := client.Emails.SendEmail(context.Background(), *request)

	fmt.Println(response)
	fmt.Println()
}	