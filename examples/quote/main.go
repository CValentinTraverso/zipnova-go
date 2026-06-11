package main

import (
	"fmt"
	"log"

	"github.com/CValentinTraverso/zipnova-go"
)

func main() {
	client := zipnova.NewClient("your-api-token", "your-api-secret", zipnova.CountryArgentina)

	quotes, err := client.Quote(&zipnova.QuoteRequest{
		AccountID:     123,
		Source:        "example-app",
		DeclaredValue: 5000,
		Destination: zipnova.Destination{
			City:  "CABA",
			State: "CABA",
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

	for _, opt := range quotes.Data {
		fmt.Printf("Carrier: %s (%d)\n", opt.CarrierName, opt.CarrierID)
		fmt.Printf("  Service:     %s\n", opt.ServiceType)
		fmt.Printf("  Price:       %.2f\n", opt.Price)
		fmt.Printf("  Price (tax): %.2f\n", opt.PriceInclTax)
		fmt.Printf("  Weight:      %dg\n", opt.TotalWeight)
		fmt.Printf("  Packages:    %d\n", opt.TotalPackages)
		fmt.Println()
	}
}
