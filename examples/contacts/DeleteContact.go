package examples

import (
	"context"
	"fmt"
	"os"

	"github.com/QGeeDev/unsend-go"
)

func DeleteContact() {
	client, err := unsend.NewClient()

	if err != nil {
		fmt.Printf(fmt.Sprintf("[ERROR] - %s\n", err.Error()))
		os.Exit(1)
	}

	request := &unsend.DeleteContactRequest{
		ContactBookId: "Book12345",
		ContactId:     "12345",
	}

	response, _ := client.Contacts.DeleteContact(context.Background(), *request)

	fmt.Println(response)
}
