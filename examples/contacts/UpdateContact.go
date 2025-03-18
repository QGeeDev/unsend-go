package examples

import (
	"context"
	"fmt"
	"os"

	"github.com/QGeeDev/unsend-go"
)

func UpdateContact() {
	client, err := unsend.NewClient()

	if err != nil {
		fmt.Printf(fmt.Sprintf("[ERROR] - %s\n", err.Error()))
		os.Exit(1)
	}

	request := &unsend.UpdateContactRequest{
		ContactBookId: "book12345",
		ContactId:     "12345",
		FirstName:     "QG",
	}

	response, _ := client.Contacts.UpdateContact(context.Background(), *request)

	fmt.Println(response)
}
