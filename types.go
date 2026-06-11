package zipnova

type Item struct {
	Sku                 string `json:"sku,omitempty"`
	Weight              int    `json:"weight,omitempty"`
	Height              int    `json:"height,omitempty"`
	Width               int    `json:"width,omitempty"`
	Length              int    `json:"length,omitempty"`
	ClassificationID    string `json:"classification_id,omitempty"`
	Description         string `json:"description,omitempty"`
	MustKeepVertical    bool   `json:"must_keep_vertical,omitempty"`
}

type Package struct {
	Sku              string  `json:"sku,omitempty"`
	SkuID            int     `json:"sku_id,omitempty"`
	Weight           int     `json:"weight,omitempty"`
	Height           int     `json:"height,omitempty"`
	Width            int     `json:"width,omitempty"`
	Length           int     `json:"length,omitempty"`
	ClassificationID string  `json:"classification_id,omitempty"`
	Description1     string  `json:"description_1,omitempty"`
	Description2     string  `json:"description_2,omitempty"`
	Description3     string  `json:"description_3,omitempty"`
	ContainerID      int     `json:"container_id,omitempty"`
	Items            []Item  `json:"items,omitempty"`
}

type Destination struct {
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Zipcode      string `json:"zipcode,omitempty"`
	Street       string `json:"street,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
	ID           int    `json:"id,omitempty"`
}

type ShipmentDestination struct {
	Name         string `json:"name"`
	Document     string `json:"document"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Street       string `json:"street,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
	StreetExtras string `json:"street_extras,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Zipcode      string `json:"zipcode,omitempty"`
	ID           int    `json:"id,omitempty"`
	PointID      int    `json:"point_id,omitempty"`
}

type QuoteResponse struct {
	SortedBy     string                 `json:"sorted_by"`
	Origin       QuoteOrigin            `json:"origin"`
	Destination  QuoteDest              `json:"destination"`
	DeclaredValue float64               `json:"declared_value"`
	Packages     []QuoteRespPackage     `json:"packages"`
	Results      map[string]QuoteOption `json:"results"`
	AllResults   []QuoteOption          `json:"all_results"`
}

type QuoteOrigin struct {
	LocationID  int              `json:"location_id"`
	Name        string           `json:"name"`
	Geolocation *QuoteGeolocation `json:"geolocation,omitempty"`
}

type QuoteDest struct {
	ID       int              `json:"id"`
	City     string           `json:"city"`
	State    string           `json:"state"`
	Country  string           `json:"country"`
	Zipcode  string           `json:"zipcode"`
	Geolocation *QuoteGeolocation `json:"geolocation,omitempty"`
}

type QuoteGeolocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type QuoteRespPackage struct {
	Descriptions     []string           `json:"descriptions"`
	Weight           int                `json:"weight"`
	Height           int                `json:"height"`
	Width            int                `json:"width"`
	Length           int                `json:"length"`
	Volume           int                `json:"volume"`
	ClassificationID int                `json:"classification_id"`
}

type QuoteOption struct {
	Selectable   bool               `json:"selectable"`
	LogisticType string             `json:"logistic_type"`
	Carrier      QuoteCarrier       `json:"carrier"`
	ServiceType  QuoteServiceType   `json:"service_type"`
	DeliveryTime QuoteDeliveryTime  `json:"delivery_time"`
	Amounts      QuoteAmounts       `json:"amounts"`
	Tags         []string           `json:"tags"`
	PickupPoints []QuotePickupPoint `json:"pickup_points,omitempty"`
}

type QuoteCarrier struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Rating int    `json:"rating"`
	Logo   string `json:"logo,omitempty"`
}

type QuoteServiceType struct {
	ID       int    `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	IsUrgent int    `json:"is_urgent"`
}

type QuoteDeliveryTime struct {
	EstimatedDelivery string              `json:"estimated_delivery"`
	Times             *QuoteTimeBreakdown `json:"times,omitempty"`
}

type QuoteTimeBreakdown struct {
	Preparation  string `json:"preparation"`
	Crossdocking string `json:"crossdocking"`
	Carrier      any    `json:"carrier"`
	Total        any    `json:"total"`
}

type QuoteAmounts struct {
	PriceShipment   float64 `json:"price_shipment"`
	PriceInsurance  float64 `json:"price_insurance"`
	Price           float64 `json:"price"`
	PriceInclTax    float64 `json:"price_incl_tax"`
	SellerPrice     float64 `json:"seller_price"`
	SellerPriceIncl float64 `json:"seller_price_incl_tax"`
}

type QuotePickupPoint struct {
	PointID     int                `json:"point_id"`
	Description string             `json:"description"`
	Location    *QuotePickupLoc    `json:"location,omitempty"`
}

type QuotePickupLoc struct {
	Street       string            `json:"street"`
	StreetNumber string            `json:"street_number"`
	City         string            `json:"city"`
	State        string            `json:"state"`
	Zipcode      string            `json:"zipcode"`
	Geolocation  *QuoteGeolocation `json:"geolocation,omitempty"`
}

type ShipmentResource struct {
	ID               int                  `json:"id"`
	ExternalID       string               `json:"external_id"`
	DeliveryID       string               `json:"delivery_id,omitempty"`
	CarrierTrackingID string              `json:"carrier_tracking_id,omitempty"`
	CreatedAt        string               `json:"created_at"`
	AccountID        int                  `json:"account_id"`
	LogisticType     string               `json:"logistic_type"`
	ServiceType      string               `json:"service_type"`
	Status           string               `json:"status"`
	StatusName       string               `json:"status_name"`
	Tracking         string               `json:"tracking,omitempty"`
	TrackingExternal string               `json:"tracking_external,omitempty"`
	Price            float64              `json:"price"`
	PriceInclTax     float64              `json:"price_incl_tax"`
	TotalWeight      int                  `json:"total_weight"`
	TotalVolume      int                  `json:"total_volume"`
	TotalPackages    int                  `json:"total_packages"`
	Carrier          *CarrierInfo         `json:"carrier,omitempty"`
	Destination      *DestinationInfo     `json:"destination,omitempty"`
	Origin           *OriginInfo          `json:"origin,omitempty"`
	Packages         []ShipmentPackage    `json:"packages,omitempty"`
	DeclaredValue    float64               `json:"declared_value"`
	ShipmentDelivery *ShipmentDeliveryTime `json:"delivery_time,omitempty"`
}

type CarrierInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo,omitempty"`
}

type DestinationInfo struct {
	Name         string `json:"name"`
	Document     string `json:"document,omitempty"`
	Street       string `json:"street,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zipcode      string `json:"zipcode,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Email        string `json:"email,omitempty"`
}

type OriginInfo struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Street       string `json:"street,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zipcode      string `json:"zipcode,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Email        string `json:"email,omitempty"`
}

type ShipmentPackage struct {
	ID           int    `json:"id"`
	LabelCode    string `json:"label_code,omitempty"`
	Weight       int    `json:"weight"`
	Height       int    `json:"height"`
	Width        int    `json:"width"`
	Length       int    `json:"length"`
	Volume       int    `json:"volume,omitempty"`
	Description1 string `json:"description_1,omitempty"`
}

type TrackingResponse []TrackingEvent

type TrackingEvent struct {
	OccurredAt string       `json:"occurred_at"`
	CreatedAt  string       `json:"created_at"`
	Status     TrackingStatus `json:"status"`
}

type TrackingStatus struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	VisibleName string `json:"visible_name"`
	Substatus   string `json:"substatus,omitempty"`
}

type ShipmentDeliveryTime struct {
	EstimatedDelivery string `json:"estimated_delivery,omitempty"`
	DropoffDeadlineAt string `json:"dropoff_deadline_at,omitempty"`
}


