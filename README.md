# zipnova-go

Go client library for [Zipnova](https://docs.zipnova.com), a shipping management service for Argentina, Chile, and Mexico.

## Installation

```shell
go get github.com/CValentinTraverso/zipnova-go
```

## Usage

```go
import "github.com/CValentinTraverso/zipnova-go"

client := zipnova.NewClient("apiToken", "apiSecret", zipnova.CountryArgentina)
```

### Quote shipping prices

```go
quotes, err := client.Quote(&zipnova.QuoteRequest{
    AccountID:     123,
    Source:        "my-app",
    OriginID:      1,
    DeclaredValue: 5000,
    Destination:   zipnova.Destination{City: "CABA", State: "CABA", Zipcode: "1424"},
    Items: []zipnova.Item{{
        Weight: 500, Height: 10, Width: 10, Length: 10,
    }},
})

for _, opt := range quotes.AllResults {
    fmt.Printf("%s - %s: $%.2f\n", opt.Carrier.Name, opt.ServiceType.Name, opt.Amounts.Price)
}
```

### Create a shipment

```go
shipment, err := client.CreateShipment(&zipnova.CreateShipmentRequest{
    AccountID:     123,
    ExternalID:    "order-001",
    ServiceType:   "standard",
    OriginID:      "auto",
    DeclaredValue: 5000,
    Source:        "my-app",
    Destination: zipnova.ShipmentDestination{
        Name: "John Doe", Document: "12345678",
        Email: "john@example.com", Phone: "+54123456789",
        Street: "Av. Siempre Viva", StreetNumber: "123",
        City: "CABA", State: "CABA", Zipcode: "1424",
    },
    Items: []zipnova.Item{{
        Weight: 500, Height: 10, Width: 10, Length: 10,
    }},
})
```

### Track a shipment

```go
tracking, err := client.GetTracking(42, zipnova.TrackingNewest)
```

### Get shipment details

```go
shipment, err := client.GetShipment(42)
```

## Supported countries

| Country | Domain |
|---|---|
| Argentina | `zipnova.CountryArgentina` |
| Chile | `zipnova.CountryChile` |
| Mexico | `zipnova.CountryMexico` |

## Authentication

Uses HTTP Basic Authentication with your API Token as username and API Secret as password.
