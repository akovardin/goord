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

	fmt.Println("Getting list of persons...")

	persons, err := client.GetPersons(context.Background(), 0, 10)
	if err != nil {
		log.Printf("Error getting persons: %v\n", err)
	} else {
		fmt.Printf("Retrieved %d persons (total: %d)\n", len(persons.ExternalIDs), persons.TotalItemsCount)
		for i, id := range persons.ExternalIDs {
			fmt.Printf("  %d. %s\n", i+1, id)
		}
	}

	externalID := "test-person-001"

	if create {
		fmt.Println("Creating/updating a person...")

		juridicalDetails := ord.JuridicalDetails{
			Type: ord.PersonTypePhysical,
			INN:  "910810615691",
		}

		person := ord.Person{
			Name:             "Ковардин Артем Сергеевич",
			Roles:            []string{"advertiser"},
			JuridicalDetails: juridicalDetails,
		}

		err = client.CreatePerson(context.Background(), externalID, person)
		if err != nil {
			log.Printf("Error creating person: %v\n", err)
		} else {
			fmt.Printf("Person %s created/updated successfully\n", externalID)
		}
	}

	fmt.Println("Getting a specific person...")

	retrievedPerson, err := client.GetPerson(context.Background(), externalID)
	if err != nil {
		log.Printf("Error getting person: %v\n", err)
	} else {
		fmt.Printf("Retrieved person: %s\n", retrievedPerson.Name)
		fmt.Printf("Person roles: %v\n", retrievedPerson.Roles)
		fmt.Printf("Person type: %s\n", retrievedPerson.JuridicalDetails.Type)
	}
}
