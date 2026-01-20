package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"gohome.4gophers.ru/kovardin/goord/ord"
)

func main() {
	client, _ := ord.NewClient(
		ord.WithBase("https://api-sandbox.ord.vk.com"),
		ord.WithToken(os.Getenv("TOKEN")),
	)

	fmt.Println("Getting KKTU codes...")

	kktuResponse, err := client.GetKKTUCodes(context.Background(), "test", "ru", 0, 10, []string{})
	if err != nil {
		log.Printf("Error getting KKTU codes: %v\n", err)
	} else {
		fmt.Printf("Retrieved %d KKTU items (total: %d)\n", len(kktuResponse.Items), kktuResponse.TotalItemsCount)
		for i, item := range kktuResponse.Items {
			fmt.Printf("  %d. %s - %s\n", i+1, item.Code, item.Name)
		}
	}

	fmt.Println("Getting ERIR message...")

	erirResponse, err := client.GetERIRMessage(context.Background(), "ru", "test")
	if err != nil {
		log.Printf("Error getting ERIR message: %v\n", err)
	} else {
		fmt.Printf("Retrieved %d ERIR messages\n", len(erirResponse.Items))
		for i, item := range erirResponse.Items {
			fmt.Printf("  %d. %s - %s\n", i+1, item.Name, item.Message)
		}
	}

	fmt.Println("Posting ERIR messages...")

	messages := []string{"test1", "test2"}
	postResponse, err := client.PostERIRMessages(context.Background(), "ru", messages)
	if err != nil {
		log.Printf("Error posting ERIR messages: %v\n", err)
	} else {
		fmt.Printf("Posted ERIR messages, received %d items\n", len(postResponse.Items))
		for i, item := range postResponse.Items {
			fmt.Printf("  %d. %s - %s\n", i+1, item.Name, item.Message)
		}
	}
}
