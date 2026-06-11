package main

import (
	"fmt"
	"log"

	"github.com/CValentinTraverso/zipnova-go"
)

func main() {
	client := zipnova.NewClient("your-api-token", "your-api-secret", zipnova.CountryArgentina)

	shipment, err := client.CreateShipment(&zipnova.CreateShipmentRequest{
		AccountID:     123,
		ExternalID:    "order-001",
		ServiceType:   "standard",
		OriginID:      "auto",
		DeclaredValue: 5000,
		Source:        "example-app",
		Destination: zipnova.ShipmentDestination{
			Name:         "John Doe",
			Document:     "12345678",
			Email:        "john@example.com",
			Phone:        "+54123456789",
			Street:       "Av. Siempre Viva",
			StreetNumber: "123",
			City:         "CABA",
			State:        "CABA",
			Zipcode:      "1424",
		},
		Items: []zipnova.Item{
			{
				Weight: 500,
				Height: 10,
				Width:  10,
				Length: 10,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Shipment created!\n")
	fmt.Printf("  ID:             %d\n", shipment.ID)
	fmt.Printf("  External ID:    %s\n", shipment.ExternalID)
	fmt.Printf("  Status:         %s\n", shipment.Status)
	fmt.Printf("  Tracking URL:   %s\n", shipment.Tracking)
	fmt.Printf("  Price:          %.2f\n", shipment.Price)
	fmt.Printf("  Price (incl.tax): %.2f\n", shipment.PriceInclTax)
}
