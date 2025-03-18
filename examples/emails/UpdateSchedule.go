package examples

import (
	"context"
	"fmt"
	"os"

	"github.com/QGeeDev/unsend-go"
)

func UpdateSchedule() {
	client, err := unsend.NewClient()

	if err != nil {
		fmt.Printf(fmt.Sprintf("[ERROR] - %s\n", err.Error()))
		os.Exit(1)
	}

	request := &unsend.UpdateScheduleRequest{
		EmailId:     "example123",
		ScheduledAt: "2024-12-31T23:59:59Z",
	}

	response, _ := client.Emails.UpdateSchedule(context.Background(), *request)

	fmt.Println(response)
}
