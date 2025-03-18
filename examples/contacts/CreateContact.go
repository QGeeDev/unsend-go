package examples

import (
	"context"
	"fmt"
	"os"

	"github.com/QGeeDev/unsend-go"
)

func CreateContact() {
	client, err := unsend.NewClient()

	if err != nil {
		fmt.Printf(fmt.Sprintf("[ERROR] - %s\n", err.Error()))
		os.Exit(1)
	}

	request := &unsend.CreateContactRequest{
		ContactBookId: "book123",
		FirstName:     "Bill",
		LastName:      "Gates",
		Email:         "billgates@microsoft.com",
		Subscribed:    false,
	}

	response, _ := client.Contacts.CreateContact(context.Background(), *request)

	fmt.Println(response)
}
