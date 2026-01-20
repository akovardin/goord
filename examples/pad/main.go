package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"gohome.4gophers.ru/kovardin/goord/ord"
)

var create = true

func main() {
	client, _ := ord.NewClient(
		ord.WithBase("https://api-sandbox.ord.vk.com"),
		ord.WithToken(os.Getenv("TOKEN")),
	)

	fmt.Println("Getting list of pads...")

	pads, err := client.GetPads(context.Background(), 0, 10, "")
	if err != nil {
		log.Printf("Error getting pads: %v\n", err)
	} else {
		fmt.Printf("Retrieved %d pads (total: %d)\n", len(pads.ExternalIDs), pads.TotalItemsCount)
		for i, id := range pads.ExternalIDs {
			fmt.Printf("  %d. %s\n", i+1, id)
		}
	}

	fmt.Println("Getting restricted pads...")

	restrictedPads, err := client.GetRestrictedPads(context.Background())
	if err != nil {
		log.Printf("Error getting restricted pads: %v\n", err)
	} else {
		fmt.Printf("Retrieved %d restricted pads\n", len(restrictedPads))
		for i, url := range restrictedPads {
			fmt.Printf("  %d. %s\n", i+1, url)
		}
	}

	padExternalID := "test-pad-001"

	if create {
		fmt.Println("Creating/updating a pad...")
		pad := ord.Pad{
			PersonExternalID: "test-person-001",
			IsOwner:          true,
			Type:             ord.PadTypeWeb,
			Name:             "Test Pad",
			URL:              ord.StringPtr("https://example.com"),
		}

		err = client.CreatePad(context.Background(), padExternalID, pad)
		if err != nil {
			log.Printf("Error creating pad: %v\n", err)
		} else {
			fmt.Printf("Pad %s created/updated successfully\n", padExternalID)
		}
	}

	fmt.Println("Getting a specific pad...")

	retrievedPad, err := client.GetPad(context.Background(), padExternalID)
	if err != nil {
		log.Printf("Error getting pad: %v\n", err)
	} else {
		fmt.Printf("Retrieved pad name: %s\n", retrievedPad.Name)
		fmt.Printf("Pad type: %s\n", retrievedPad.Type)
		fmt.Printf("Pad person ID: %s\n", retrievedPad.PersonExternalID)
	}
}
