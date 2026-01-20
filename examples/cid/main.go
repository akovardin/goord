package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"gohome.4gophers.ru/kovardin/goord/ord"
)

var create = false

func main() {
	client, _ := ord.NewClient(
		ord.WithBase("https://api-sandbox.ord.vk.com"),
		ord.WithToken(os.Getenv("TOKEN")),
	)

	fmt.Println("Getting list of CIDs...")

	cidList, err := client.GetCIDList(context.Background(), 0, 10)
	if err != nil {
		log.Printf("Error getting CID list: %v\n", err)
	} else {
		fmt.Printf("Retrieved %d CIDs (total: %d)\n", len(cidList.CIDs), cidList.TotalItemsCount)
		for i, id := range cidList.CIDs {
			fmt.Printf("  %d. %s\n", i+1, id)
		}
	}

	cidValue := "123e4567-e89b-12d3-a456-426655440000"

	if create {
		fmt.Println("Creating/updating a CID...")

		cid := ord.CID{
			CID:  cidValue,
			Name: "Test CID",
		}

		err = client.CreateCID(context.Background(), cidValue, cid)
		if err != nil {
			log.Printf("Error creating CID: %v\n", err)
		} else {
			fmt.Printf("CID %s created/updated successfully\n", cidValue)
		}
	}

	fmt.Println("Getting a specific CID...")

	retrievedCID, err := client.GetCID(context.Background(), cidValue)
	if err != nil {
		log.Printf("Error getting CID: %v\n", err)
	} else {
		fmt.Printf("Retrieved CID: %s\n", retrievedCID.CID)
		fmt.Printf("CID name: %s\n", retrievedCID.Name)
	}
}
