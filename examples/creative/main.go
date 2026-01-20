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

	fmt.Println("Getting list of creatives...")

	creatives, err := client.GetCreatives(context.Background(), 0, 10)
	if err != nil {
		log.Printf("Error getting creatives: %v\n", err)
	} else {
		fmt.Printf("Retrieved %d creatives (total: %d)\n", len(creatives.ExternalIDs), creatives.TotalItemsCount)
		for i, id := range creatives.ExternalIDs {
			fmt.Printf("  %d. %s\n", i+1, id)
		}
	}

	externalID := "test-creative-001"

	if create {
		fmt.Println("Creating/updating a creative (v3)...")

		creative := ord.CreateCreativeV3Request{
			PersonExternalID: ord.StringPtr("test-person-001"),
			KKTUs:            []string{"5.2.3"},
			Name:             ord.StringPtr("Тестовый креатив"),
			Form:             ord.CreativeFormBanner,
			TargetURLs:       &[]string{"https://example.com"},
			MediaExternalIDs: &[]string{"test-media-001"},
			PayType:          ord.StringPtr(ord.CreativePayTypeCPM),
		}

		err = client.CreateCreativeV3(context.Background(), externalID, creative)
		if err != nil {
			log.Printf("Error creating creative: %v\n", err)
		} else {
			fmt.Printf("Creative %s created/updated successfully\n", externalID)
		}
	}

	fmt.Println("Getting a specific creative...")

	retrievedCreative, err := client.GetCreativeV3(context.Background(), externalID)
	if err != nil {
		log.Printf("Error getting creative: %v\n", err)
	} else {
		fmt.Printf("Retrieved creative: %s\n", *retrievedCreative.Name)
		fmt.Printf("Creative form: %s\n", retrievedCreative.Form)
		fmt.Printf("Creative KKTVs: %v\n", retrievedCreative.KKTUs)
	}

	if create {
		fmt.Println("Adding texts to creative...")

		texts := []string{"Текст рекламы 1", "Текст рекламы 2"}

		err = client.AddTextsToCreative(context.Background(), externalID, texts)
		if err != nil {
			log.Printf("Error adding texts to creative: %v\n", err)
		} else {
			fmt.Printf("Texts added to creative %s successfully\n", externalID)
		}
	}

	if create {
		fmt.Println("Adding media to creative...")

		mediaExternalIDs := []string{"test-media-001"}

		err = client.AddMediaToCreative(context.Background(), externalID, mediaExternalIDs)
		if err != nil {
			log.Printf("Error adding media to creative: %v\n", err)
		} else {
			fmt.Printf("Media added to creative %s successfully\n", externalID)
		}
	}

	fmt.Println("Getting creative ERIDs...")

	erids, err := client.GetCreativeERIDs(context.Background(), 0, 10)
	if err != nil {
		log.Printf("Error getting creative ERIDs: %v\n", err)
	} else {
		fmt.Printf("Retrieved %d ERIDs (total: %d)\n", len(erids.ERIDs), erids.TotalItemsCount)
		for i, erid := range erids.ERIDs {
			fmt.Printf("  %d. %s\n", i+1, erid)
		}
	}

	fmt.Println("Getting creative ERID/External ID pairs...")

	pairs, err := client.GetCreativeERIDExternalIDPairs(context.Background(), 0, 10)
	if err != nil {
		log.Printf("Error getting creative ERID/External ID pairs: %v\n", err)
	} else {
		fmt.Printf("Retrieved %d pairs (total: %d)\n", len(pairs.Items), pairs.TotalItemsCount)
		for i, pair := range pairs.Items {
			fmt.Printf("  %d. ERID: %s, External ID: %s\n", i+1, pair.ERID, pair.ExternalID)
		}
	}
}
