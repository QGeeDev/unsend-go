package examples

import (
	"context"
	"fmt"
	"os"

	"github.com/QGeeDev/unsend-go"
)

func UpsertContact() {
	client, err := unsend.NewClient()

	if err != nil {
		fmt.Printf(fmt.Sprintf("[ERROR] - %s\n", err.Error()))
		os.Exit(1)
	}

	request := &unsend.UpsertContactRequest{
		ContactBookId: "cm8ath8d20001s3p3if0mhoq7",
		ContactId:     "cm8athj930003s3p34wfpnkue",
		Email:         "qg@qgdev.co.uk",
		FirstName:     "QGeeDev",
	}

	response, _ := client.Contacts.UpsertContact(context.Background(), *request)

	fmt.Println(response)
}
