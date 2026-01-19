package main

import (
	"context"
	"fmt"
	"log"

	"gohome.4gophers.ru/kovardin/goord/ord"
)

func main() {
	// Initialize the client with your token
	client := ord.NewClient(ord.Config{
		BaseURL: "https://api-sandbox.ord.vk.com", // Use sandbox for testing
		Token:   "your-jwt-token-here",            // Replace with your actual JWT token
	})

	// Example 1: Get a list of persons
	fmt.Println("Getting list of persons...")
	persons, err := client.GetPersons(context.Background(), 0, 10)
	if err != nil {
		log.Printf("Error getting persons: %v", err)
	} else {
		fmt.Printf("Retrieved %d persons (total: %d)\n", len(persons.ExternalIDs), persons.TotalItemsCount)
		for i, id := range persons.ExternalIDs {
			fmt.Printf("  %d. %s\n", i+1, id)
		}
	}

	// Example 2: Create a new person
	fmt.Println("\nCreating/updating a person...")
	juridicalDetails := ord.JuridicalDetails{
		Type: "juridical",
	}

	// Create a person with minimal required fields
	person := ord.Person{
		Name:             "Test Company",
		Roles:            []string{"advertiser"},
		JuridicalDetails: juridicalDetails,
	}

	// Use a unique external ID for testing
	externalID := "test-person-001"

	err = client.CreatePerson(context.Background(), externalID, person)
	if err != nil {
		log.Printf("Error creating person: %v", err)
	} else {
		fmt.Printf("Person %s created/updated successfully\n", externalID)
	}

	// Example 3: Get a specific person
	fmt.Println("\nGetting a specific person...")
	retrievedPerson, err := client.GetPerson(context.Background(), externalID)
	if err != nil {
		log.Printf("Error getting person: %v", err)
	} else {
		fmt.Printf("Retrieved person: %s\n", retrievedPerson.Name)
		fmt.Printf("Person roles: %v\n", retrievedPerson.Roles)
		fmt.Printf("Person type: %s\n", retrievedPerson.JuridicalDetails.Type)
	}

	// Example 4: Get a list of contracts
	fmt.Println("\nGetting list of contracts...")
	contracts, err := client.GetContracts(context.Background(), 0, 10)
	if err != nil {
		log.Printf("Error getting contracts: %v", err)
	} else {
		fmt.Printf("Retrieved %d contracts (total: %d)\n", len(contracts.ExternalIDs), contracts.TotalItemsCount)
		for i, id := range contracts.ExternalIDs {
			fmt.Printf("  %d. %s\n", i+1, id)
		}
	}

	// Example 5: Create a new contract
	fmt.Println("\nCreating/updating a contract...")
	contract := ord.CreateContractRequest{
		Type:                 "advertising",
		ClientExternalID:     "test-person-001",
		ContractorExternalID: "test-contractor-001",
		Date:                 "2026-01-19",
		SubjectType:          "advertising",
	}

	// Use a unique external ID for testing
	contractExternalID := "test-contract-001"

	err = client.CreateContract(context.Background(), contractExternalID, contract)
	if err != nil {
		log.Printf("Error creating contract: %v", err)
	} else {
		fmt.Printf("Contract %s created/updated successfully\n", contractExternalID)
	}

	// Example 6: Get a specific contract
	fmt.Println("\nGetting a specific contract...")
	retrievedContract, err := client.GetContract(context.Background(), contractExternalID)
	if err != nil {
		log.Printf("Error getting contract: %v", err)
	} else {
		fmt.Printf("Retrieved contract type: %s\n", retrievedContract.Type)
		fmt.Printf("Contract client ID: %s\n", retrievedContract.ClientExternalID)
		fmt.Printf("Contract contractor ID: %s\n", retrievedContract.ContractorExternalID)
	}

	// Example 7: Create a CID for a contract
	fmt.Println("\nCreating a CID for contract...")
	err = client.CreateCID(context.Background(), contractExternalID)
	if err != nil {
		log.Printf("Error creating CID: %v", err)
	} else {
		fmt.Printf("CID created successfully\n")
	}
}
