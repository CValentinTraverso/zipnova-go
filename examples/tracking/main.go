package main

import (
	"fmt"
	"log"

	"github.com/CValentinTraverso/zipnova-go"
)

func main() {
	client := zipnova.NewClient("your-api-token", "your-api-secret", zipnova.CountryArgentina)

	tracking, err := client.GetTracking(42, zipnova.TrackingNewest)
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range tracking.Data {
		fmt.Printf("[%s] %s\n", event.OccurredAt, event.Status.Name)
		if event.Status.Substatus != "" {
			fmt.Printf("  Substatus: %s\n", event.Status.Substatus)
		}
		fmt.Printf("  (%s)\n", event.Status.VisibleName)
		fmt.Println()
	}
}
