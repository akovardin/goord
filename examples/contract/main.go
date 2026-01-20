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

	fmt.Println("Getting list of contracts...")

	contracts, err := client.GetContracts(context.Background(), 0, 10)
	if err != nil {
		log.Printf("Error getting contracts: %v\n", err)
	} else {
		fmt.Printf("Retrieved %d contracts (total: %d)\n", len(contracts.ExternalIDs), contracts.TotalItemsCount)
		for i, id := range contracts.ExternalIDs {
			fmt.Printf("  %d. %s\n", i+1, id)
		}
	}

	contractExternalID := "test-contract-001"

	if create {
		fmt.Println("Creating/updating a contract...")
		contract := ord.CreateContractRequest{
			Type:                 ord.ContractTypeService,
			ClientExternalID:     "test-person-001",
			ContractorExternalID: "i5c6aq2r0rs-1jaocdo32",
			Date:                 "2026-01-19",
			SubjectType:          ord.ContractSubjectTypeOther,
		}

		err = client.CreateContract(context.Background(), contractExternalID, contract)
		if err != nil {
			log.Printf("Error creating contract: %v\n", err)
		} else {
			fmt.Printf("Contract %s created/updated successfully\n", contractExternalID)
		}
	}

	fmt.Println("Getting a specific contract...")

	retrievedContract, err := client.GetContract(context.Background(), contractExternalID)
	if err != nil {
		log.Printf("Error getting contract: %v\n", err)
	} else {
		fmt.Printf("Retrieved contract type: %s\n", retrievedContract.Type)
		fmt.Printf("Contract client ID: %s\n", retrievedContract.ClientExternalID)
		fmt.Printf("Contract contractor ID: %s\n", retrievedContract.ContractorExternalID)
	}

	fmt.Println("Requesting a CID for contract...")

	err = client.RequestCID(context.Background(), contractExternalID)
	if err != nil {
		log.Printf("Error requesting CID: %v\n", err)
	} else {
		fmt.Printf("CID requested successfully\n")
	}
}
