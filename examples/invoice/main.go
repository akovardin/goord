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

	fmt.Println("Getting list of invoices...")

	invoices, err := client.GetInvoices(context.Background(), 0, 10)
	if err != nil {
		log.Printf("Error getting invoices: %v\n", err)
	} else {
		fmt.Printf("Retrieved %d invoices (total: %d)\n", len(invoices.ExternalIDs), invoices.TotalItemsCount)
		for i, id := range invoices.ExternalIDs {
			fmt.Printf("  %d. %s\n", i+1, id)
		}
	}

	externalID := "test-invoice-001"

	if create {
		fmt.Println("Creating/updating an invoice...")

		amount := ord.InvoiceAmount{
			Services: ord.InvoiceAmountGroup{
				ExcludingVat: "1000",
				VatRate:      "20",
				Vat:          "200",
				IncludingVat: "1200",
			},
		}

		invoice := ord.Invoice{
			ContractExternalID: "test-contract-001",
			Date:               "2023-01-01",
			DateStart:          "2023-01-01",
			DateEnd:            "2023-01-31",
			Amount:             amount,
			ClientRole:         ord.InvoiceClientRoleTypeAdvertiser,
			ContractorRole:     ord.InvoiceClientRoleTypePublisher,
			Flags:              []string{"flag1", "flag2"},
		}

		err = client.CreateInvoiceHeader(context.Background(), externalID, invoice)
		if err != nil {
			log.Printf("Error creating invoice: %v\n", err)
		} else {
			fmt.Printf("Invoice %s created/updated successfully\n", externalID)
		}
	}

	fmt.Println("Getting a specific invoice...")

	retrievedInvoice, err := client.GetInvoice(context.Background(), externalID)
	if err != nil {
		log.Printf("Error getting invoice: %v\n", err)
	} else {
		fmt.Printf("Retrieved invoice contract: %s\n", retrievedInvoice.ContractExternalID)
		fmt.Printf("Invoice date: %s\n", retrievedInvoice.Date)
		fmt.Printf("Invoice amount: %s\n", retrievedInvoice.Amount.Services.IncludingVat)
	}
}
