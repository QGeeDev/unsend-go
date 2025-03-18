package examples

import (
	"context"
	"fmt"
	"os"

	"github.com/QGeeDev/unsend-go"
)

func GetDomains() {
	client, err := unsend.NewClient()

	if err != nil {
		fmt.Printf(fmt.Sprintf("[ERROR] - %s\n", err.Error()))
		os.Exit(1)
	}

	response, _ := client.Domains.GetDomains(context.Background())

	fmt.Println(response)
}
