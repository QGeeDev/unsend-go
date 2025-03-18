package examples

import (
	"context"
	"fmt"
	"os"

	"github.com/QGeeDev/unsend-go"
)

func GetContact() {
	client, err := unsend.NewClient()

	if err != nil {
		fmt.Printf(fmt.Sprintf("[ERROR] - %s\n", err.Error()))
		os.Exit(1)
	}

	request := &unsend.GetContactRequest{
		ContactBookId: "cm8ath8d20001s3p3if0mhoq7",
		ContactId:     "cm8athj930003s3p34wfpnkue",
	}

	response, _ := client.Contacts.GetContact(context.Background(), *request)

	fmt.Println(response)
}
