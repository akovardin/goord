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

	fmt.Println("Getting list of statistics...")

	statistics, err := client.GetStatisticsList(context.Background(), 0, 10)
	if err != nil {
		log.Printf("Error getting statistics: %v\n", err)
	} else {
		fmt.Printf("Retrieved %d statistics (total: %d)\n", len(statistics.Items), statistics.TotalItemsCount)
		for i, item := range statistics.Items {
			fmt.Printf("  %d. Creative: %s, Pad: %s, Shows: %d\n", i+1, item.CreativeExternalID, item.PadExternalID, item.ShowsCount)
		}
	}

	if create {
		fmt.Println("Creating statistics...")

		var showsCount uint64 = 1000
		amount := ord.StatisticsAmount{
			ExcludingVAT: "800",
			VATRate:      "20",
			VAT:          "160",
			IncludingVAT: "960",
		}

		statisticsItems := ord.StatisticsV2ItemsArray{
			Items: []ord.StatisticsV2Item{
				{
					CreativeExternalID: "01",
					PadExternalID:      "test-pad-001",
					ShowsCount:         showsCount,
					InvoiceShowsCount:  &showsCount,
					Amount:             &amount,
					AmountPerEvent:     ord.StringPtr("0.96"),
					PayType:            ord.StringPtr(ord.StatisticsPayTypeCPM),
					DateStartPlanned:   ord.StringPtr("2023-01-01"),
					DateEndPlanned:     ord.StringPtr("2023-01-31"),
					DateStartActual:    "2023-01-01",
					DateEndActual:      "2023-01-31",
				},
			},
		}

		externalIDs, err := client.CreateStatisticsV2(context.Background(), statisticsItems)
		if err != nil {
			log.Printf("Error creating statistics: %v\n", err)
		} else {
			fmt.Printf("Statistics created successfully, external IDs: %v\n", externalIDs)
		}
	}

	deleteReq := ord.DeleteStatisticsRequest{
		Items: []struct {
			CreativeExternalID string `json:"creative_external_id"`
			PadExternalID      string `json:"pad_external_id"`
			DateStartActual    string `json:"date_start_actual"`
		}{
			{
				CreativeExternalID: "test-creative-001",
				PadExternalID:      "test-pad-001",
				DateStartActual:    "2023-01-01",
			},
		},
	}

	err = client.DeleteStatisticsV3(context.Background(), deleteReq)
	if err != nil {
		log.Printf("Error deleting statistics: %v\n", err)
	} else {
		fmt.Println("Statistics deleted successfully")
	}
}
