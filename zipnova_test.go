package zipnova

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testClient(handler http.HandlerFunc) (*Client, *httptest.Server) {
	srv := httptest.NewServer(handler)
	c := &Client{
		httpClient: srv.Client(),
		baseURL:    srv.URL,
		apiToken:   "test-token",
		apiSecret:  "test-secret",
	}
	return c, srv
}

func TestQuote(t *testing.T) {
	c, srv := testClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/shipments/quote" {
			t.Fatalf("expected /shipments/quote, got %s", r.URL.Path)
		}
		user, pass, ok := r.BasicAuth()
		if !ok || user != "test-token" || pass != "test-secret" {
			t.Fatal("bad auth")
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Fatal("missing Accept header")
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Fatal("missing Content-Type header")
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"sorted_by": "price",
			"origin": map[string]any{
				"location_id": 1,
				"name":        "Warehouse",
			},
			"destination": map[string]any{
				"id": 30, "city": "CABA", "state": "CABA", "country": "Argentina", "zipcode": "1424",
			},
			"declared_value": 5000,
			"packages": []map[string]any{{
				"descriptions": []string{"Product"}, "weight": 500, "height": 10, "width": 10, "length": 10, "volume": 1000, "classification_id": 1,
			}},
			"results": map[string]any{
				"standard_delivery": map[string]any{
					"selectable":    true,
					"logistic_type": "crossdock",
					"carrier":       map[string]any{"id": 1, "name": "Test Carrier", "rating": 1},
					"service_type":  map[string]any{"id": 1, "code": "standard_delivery", "name": "Delivery", "is_urgent": 0},
					"delivery_time": map[string]any{"estimated_delivery": "2026-07-01T00:00:00Z"},
					"amounts":       map[string]any{"price_shipment": 1000, "price_insurance": 50, "price": 1050, "price_incl_tax": 1270.50, "seller_price": 1050, "seller_price_incl_tax": 1270.50},
					"tags":          []string{},
				},
			},
			"all_results": []map[string]any{{
				"selectable":    true,
				"logistic_type": "crossdock",
				"carrier":       map[string]any{"id": 1, "name": "Test Carrier", "rating": 1},
				"service_type":  map[string]any{"id": 1, "code": "standard_delivery", "name": "Delivery", "is_urgent": 0},
				"delivery_time": map[string]any{"estimated_delivery": "2026-07-01T00:00:00Z"},
				"amounts":       map[string]any{"price_shipment": 1000, "price_insurance": 50, "price": 1050, "price_incl_tax": 1270.50, "seller_price": 1050, "seller_price_incl_tax": 1270.50},
				"tags":          []string{},
			}},
		})
	})
	defer srv.Close()

	quotes, err := c.Quote(&QuoteRequest{
		AccountID:     1,
		Source:        "test",
		DeclaredValue: 5000,
		Destination:   Destination{City: "CABA", State: "CABA", Zipcode: "1424"},
		Items:         []Item{{Weight: 500, Height: 10, Width: 10, Length: 10}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(quotes.AllResults) != 1 {
		t.Fatalf("expected 1 result, got %d", len(quotes.AllResults))
	}
	if quotes.AllResults[0].Carrier.Name != "Test Carrier" {
		t.Fatalf("expected Test Carrier, got %s", quotes.AllResults[0].Carrier.Name)
	}
	if quotes.Results["standard_delivery"].Amounts.Price != 1050 {
		t.Fatalf("expected price 1050, got %f", quotes.Results["standard_delivery"].Amounts.Price)
	}
}

func TestCreateShipment(t *testing.T) {
	c, srv := testClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/shipments" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{
			"id":              27558239,
			"external_id":     "order-001",
			"status":          "new",
			"status_name":     "Procesando",
			"account_id":      1,
			"service_type":    "standard_delivery",
			"logistic_type":   "crossdock",
			"price":           1050,
			"price_incl_tax":  1270.50,
			"total_weight":    500,
			"total_volume":    1000,
			"total_packages":  1,
			"declared_value":  5000,
			"tracking":        "https://app.zipnova.com.ar/track/abc123",
			"created_at":      "2026-06-11T16:20:08+00:00",
		})
	})
	defer srv.Close()

	shipment, err := c.CreateShipment(&CreateShipmentRequest{
		AccountID:     1,
		ExternalID:    "order-001",
		ServiceType:   "standard_delivery",
		OriginID:      "1",
		DeclaredValue: 5000,
		Source:        "test",
		Destination: ShipmentDestination{
			Name: "John Doe", Document: "12345678",
			Email: "john@test.com", Phone: "+54123456789",
			Street: "Av Test", StreetNumber: "123",
			City: "CABA", State: "CABA", Zipcode: "1424",
		},
		Items: []Item{{Weight: 500, Height: 10, Width: 10, Length: 10}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if shipment.ID != 27558239 {
		t.Fatalf("expected ID 27558239, got %d", shipment.ID)
	}
	if shipment.Status != "new" {
		t.Fatalf("expected status new, got %s", shipment.Status)
	}
	if shipment.Tracking != "https://app.zipnova.com.ar/track/abc123" {
		t.Fatalf("unexpected tracking URL: %s", shipment.Tracking)
	}
}

func TestGetShipment(t *testing.T) {
	c, srv := testClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/shipments/27558239" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"id": 27558239, "external_id": "order-001", "status": "new",
			"tracking": "https://app.zipnova.com.ar/track/abc123",
		})
	})
	defer srv.Close()

	shipment, err := c.GetShipment(27558239)
	if err != nil {
		t.Fatal(err)
	}
	if shipment.ID != 27558239 {
		t.Fatalf("expected ID 27558239, got %d", shipment.ID)
	}
}

func TestGetTracking(t *testing.T) {
	c, srv := testClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/shipments/27558239/tracking" {
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		if r.URL.Query().Get("sort") != "newest" {
			t.Fatalf("expected sort=newest, got %s", r.URL.Query().Get("sort"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]map[string]any{{
			"occurred_at": "2026-06-11T16:20:08+00:00",
			"created_at":  "2026-06-11T16:20:08+00:00",
			"status": map[string]any{
				"code": "new", "name": "Procesando", "visible_name": "Procesando",
			},
		}})
	})
	defer srv.Close()

	tracking, err := c.GetTracking(27558239, TrackingNewest)
	if err != nil {
		t.Fatal(err)
	}
	if len(*tracking) != 1 {
		t.Fatalf("expected 1 event, got %d", len(*tracking))
	}
	events := *tracking
	if events[0].Status.Code != "new" {
		t.Fatalf("expected status new, got %s", events[0].Status.Code)
	}
}

func TestAuthError(t *testing.T) {
	c, srv := testClient(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"status":"error","message":"Unauthenticated."}`))
	})
	defer srv.Close()

	_, err := c.Quote(&QuoteRequest{
		AccountID: 1, Source: "test", DeclaredValue: 100,
		Destination: Destination{City: "CABA", State: "CABA"},
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestValidationError(t *testing.T) {
	c, srv := testClient(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"error","message":"The given data was invalid.","errors":{"destination.zipcode":["The destination.zipcode field is required."]}}`))
	})
	defer srv.Close()

	_, err := c.Quote(&QuoteRequest{
		AccountID: 1, Source: "test", DeclaredValue: 100,
		Destination: Destination{City: "CABA", State: "CABA"},
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
