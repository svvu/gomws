package orders

import "encoding/xml"

// ListOrdersResult the result for the ListOrders operation.
type ListOrdersResult struct {
	XMLName           xml.Name `xml:"ListOrdersResponse"`
	NextToken         string   `xml:"ListOrdersResult>NextToken"`
	CreatedBefore     string   `xml:"ListOrdersResult>CreatedBefore"`
	LastUpdatedBefore string   `xml:"ListOrdersResult>LastUpdatedBefore"`
	Orders            []Order  `xml:"ListOrdersResult>Orders>Order"`
}

// Order information.
type Order struct {
	AmazonOrderId                string
	SellerOrderId                string
	PurchaseDate                 string
	LastUpdateDate               string
	OrderStatus                  string
	FulfillmentChannel           string
	SalesChannel                 string
	OrderChannel                 string
	ShipServiceLevel             string
	ShippingAddress              Address
	OrderTotal                   Money
	NumberOfItemsShipped         int64
	NumberOfItemsUnshipped       int64
	PaymentExecutionDetail       []PaymentExecutionDetailItem `xml:">PaymentExecutionDetailItem"`
	PaymentMethod                string
	MarketplaceId                string
	BuyerEmail                   string
	BuyerName                    string
	ShipmentServiceLevelCategory string
	ShippedByAmazonTFM           bool
	TFMShipmentStatus            string
	CbaDisplayableShippingLabel  string
	OrderType                    string
	EarliestShipDate             string
	LatestShipDate               string
	EarliestDeliveryDate         string
	LatestDeliveryDate           string
	IsBusinessOrder              bool
	PurchaseOrderNumber          string
	IsPrime                      bool
	IsPremiumOrder               bool
}

// Address the shipping address for the order.
type Address struct {
	Name          string
	AddressLine1  string
	AddressLine2  string
	AddressLine3  string
	City          string
	County        string
	District      string
	StateOrRegion string
	PostalCode    string
	// The two-digit country code. In ISO 3166-1-alpha 2 format.
	CountryCode string
	Phone       string
}

// Money contains Currency type and amount.
type Money struct {
	// Three-digit currency code. In ISO 4217 format.
	CurrencyCode string
	Amount       string
}

// PaymentExecutionDetailItem is information about a sub-payment method used to pay for a COD order.
type PaymentExecutionDetailItem struct {
	Payment Money
	// PaymentMethod values:
	//  COD - Cash On Delivery. Available only in China (CN) and Japan (JP).
	//  GC - Gift Card. Available only in CN and JP.
	//  PointsAccount - Amazon Points. Available only in JP.
	PaymentMethod string
}
