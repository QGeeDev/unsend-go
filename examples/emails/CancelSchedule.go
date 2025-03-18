package examples

import (
	"context"
	"fmt"
	"os"

	"github.com/QGeeDev/unsend-go"
)

func CancelSchedule() {
	client, err := unsend.NewClient()

	if err != nil {
		fmt.Printf(fmt.Sprintf("[ERROR] - %s\n", err.Error()))
		os.Exit(1)
	}

	request := &unsend.CancelScheduleRequest{
		EmailId: "example123",
	}

	response, _ := client.Emails.CancelSchedule(context.Background(), *request)

	fmt.Println(response)
}
