package examples

import (
	"context"
	"fmt"
	"os"

	"github.com/QGeeDev/unsend-go"
)

func GetEmail() {
	client, err := unsend.NewClient()

	if err != nil {
		fmt.Printf(fmt.Sprintf("[ERROR] - %s\n", err.Error()))
		os.Exit(1)
	}

	request := &unsend.GetEmailRequest{
		EmailId: "example123",
	}

	response, _ := client.Emails.GetEmail(context.Background(), *request)

	fmt.Println(response)
}
