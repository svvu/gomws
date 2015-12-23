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

// ListOrdersByNextTokenResult the result for the ListOrdersByNextToken operation.
type ListOrdersByNextTokenResult struct {
	XMLName           xml.Name `xml:"ListOrdersByNextTokenResponse"`
	NextToken         string   `xml:"ListOrdersByNextTokenResult>NextToken"`
	CreatedBefore     string   `xml:"ListOrdersByNextTokenResult>CreatedBefore"`
	LastUpdatedBefore string   `xml:"ListOrdersByNextTokenResult>LastUpdatedBefore"`
	Orders            []Order  `xml:"ListOrdersByNextTokenResult>Orders>Order"`
}

// GetOrderResult the result for the GetOrder operation.
type GetOrderResult struct {
	XMLName           xml.Name `xml:"GetOrderResponse"`
	NextToken         string   `xml:"GetOrderResult>NextToken"`
	CreatedBefore     string   `xml:"GetOrderResult>CreatedBefore"`
	LastUpdatedBefore string   `xml:"GetOrderResult>LastUpdatedBefore"`
	Orders            []Order  `xml:"GetOrderResult>Orders>Order"`
}

// ListOrderItemsResult the result for the ListOrderItems operation.
type ListOrderItemsResult struct {
	XMLName       xml.Name    `xml:"ListOrderItemsResponse"`
	NextToken     string      `xml:"ListOrderItemsResult>NextToken"`
	AmazonOrderId string      `xml:"ListOrderItemsResult>AmazonOrderId"`
	OrderItems    []OrderItem `xml:"ListOrderItemsResult>OrderItems>OrderItem"`
}

// ListOrderItemsByNextTokenResult the result for the ListOrderItemsByNextToken operation.
type ListOrderItemsByNextTokenResult struct {
	XMLName       xml.Name    `xml:"ListOrderItemsByNextTokenResponse"`
	NextToken     string      `xml:"ListOrderItemsByNextTokenResult>NextToken"`
	AmazonOrderId string      `xml:"ListOrderItemsByNextTokenResult>AmazonOrderId"`
	OrderItems    []OrderItem `xml:"ListOrderItemsByNextTokenResult>OrderItems>OrderItem"`
}

// Order information.
type Order struct {
	// An Amazon-defined order identifier, in 3-7-7 format.
	AmazonOrderId string
	// A seller-defined order identifier.
	SellerOrderId string
	// The date when the order was created.
	PurchaseDate string
	// The date when the order was last updated.
	LastUpdateDate string
	OrderStatus    string
	// How the order was fulfilled: by Amazon (AFN) or by the seller (MFN).
	FulfillmentChannel string
	// The sales channel of the first item in the order.
	SalesChannel string
	// The order channel of the first item in the order.
	OrderChannel string
	// The shipment service level of the order.
	ShipServiceLevel string
	ShippingAddress  Address
	// The total charge for the order.
	OrderTotal Money
	// The number of items shipped.
	NumberOfItemsShipped int
	// The number of items unshipped.
	NumberOfItemsUnshipped int
	// Information about sub-payment methods for a Cash On Delivery (COD) order.
	PaymentExecutionDetail []PaymentExecutionDetailItem `xml:">PaymentExecutionDetailItem"`
	// The main payment method of the order.
	// PaymentMethod values:
	// 	COD - Cash On Delivery. Available only in China (CN) and Japan (JP).
	// 	CVS - Convenience Store. Available only in JP.
	// 	Other - A payment method other than COD and CVS.
	PaymentMethod string
	// The anonymized identifier for the Marketplace where the order was placed.
	MarketplaceId string
	// The anonymized e-mail address of the buyer.
	BuyerEmail string
	// The name of the buyer.
	BuyerName string
	// The shipment service level category of the order.
	// ShipmentServiceLevelCategory values:
	// 	Expedited, FreeEconomy, NextDay, SameDay, SecondDay, Scheduled, Standard
	ShipmentServiceLevelCategory string
	// Indicates if the order was shipped by the Amazon Transportation for Merchants (Amazon TFM) service.
	ShippedByAmazonTFM bool
	// The status of the Amazon TFM order.
	TFMShipmentStatus string
	// A seller-customized shipment service level that is mapped to one of the four
	// 	standard shipping settings supported by Checkout by Amazon (CBA).
	CbaDisplayableShippingLabel string
	// The type of the order.
	// OrderType values:
	// 	StandardOrder - An order that contains items for which you currently have inventory in stock.
	// 	Preorder - An order that contains items with a release date that is in the future.
	OrderType string
	// The start of the time period that you have committed to ship the order. In ISO 8601 date format.
	EarliestShipDate string
	// The end of the time period that you have committed to ship the order. In ISO 8601 date format.
	LatestShipDate string
	// The start of the time period that you have commited to fulfill the order. In ISO 8601 date format.
	EarliestDeliveryDate string
	// The end of the time period that you have commited to fulfill the order. In ISO 8601 date format.
	LatestDeliveryDate string
	// Indicates that the order is an Amazon Business order.
	IsBusinessOrder bool
	// The purchase order (PO) number entered by the buyer at checkout.
	PurchaseOrderNumber string
	// Indicates that the order is a seller-fulfilled Amazon Prime order.
	IsPrime bool
	// Indicates that the order has a Premium Shipping Service Level Agreement.
	IsPremiumOrder bool
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
	// The currency amount.
	Amount string
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

// OrderItem information.
type OrderItem struct {
	ASIN string
	// An Amazon-defined order item identifier.
	OrderItemId string
	// The seller SKU of the item.
	SellerSKU string
	// Buyer information for custom orders from the Amazon Custom program.
	BuyerCustomizedInfo BuyerCustomizedInfo
	// The name of the item.
	Title string
	// The number of items in the order.
	QuantityOrdered int
	// The number of items shipped.
	QuantityShipped int
	// The number and value of Amazon Points granted with the purchase of an item (available only in Japan).
	PointsGranted PointsGranted
	// The selling price of the order item. Its price * qty ordered.
	ItemPrice Money
	// The shipping price of the item.
	ShippingPrice Money
	// The gift wrap price of the item.
	GiftWrapPrice Money
	// The tax on the item price.
	ItemTax Money
	// The tax on the shipping price.
	ShippingTax Money
	// The tax on the gift wrap price.
	GiftWrapTax Money
	// The discount on the shipping price.
	ShippingDiscount Money
	// The total of all promotional discounts in the offer.
	PromotionDiscount Money
	PromotionIds      []string `xml:">PromotionId"`
	// The fee charged for COD service. CODFee is a response element only in Japan (JP).
	CODFee Money
	// The discount on the COD fee. CODFeeDiscount is a response element only in Japan (JP).
	CODFeeDiscount Money
	// A gift message provided by the buyer.
	GiftMessageText string
	// The gift wrap level specified by the buyer.
	GiftWrapLevel string
	// Invoice information (available only in China).
	InvoiceData InvoiceData
	// The condition of the item as described by the seller.
	ConditionNote string
	// ConditionId values:
	// 	New, Used, Collectible, Refurbished, Preorder, Club
	ConditionId string
	// ConditionSubtypeId values:
	// 	New, Mint, Very Good Good, Acceptable, Poor, Club, OEM, Warranty,
	// 	Refurbished Warranty, Refurbished, Open Box, Any, Other,
	ConditionSubtypeId string
	// The start date of the scheduled delivery window in the time zone of the order destination. In ISO 8601 date format.
	ScheduledDeliveryStartDate string
	// The end date of the scheduled delivery window in the time zone of the order destination. In ISO 8601 date format.
	ScheduledDeliveryEndDate string
	// Indicates that the selling price is a special price that is available only for Amazon Business orders.
	PriceDesignation string
}

// BuyerCustomizedInfo the buyer information for custom orders from the Amazon Custom program.
type BuyerCustomizedInfo struct {
	// The location of a zip file containing Amazon Custom data.
	CustomizedURL string
}

// PointsGranted the number and value of Amazon Points granted with the purchase of an item.
// Available only in Japan.
type PointsGranted struct {
	// The number of Amazon Points granted with the purchase of an item.
	PointsNumber int
	// The monetary value of the Amazon Points granted.
	PointsMonetaryValue Money
}

// InvoiceData the Invoice information (available only in China).
type InvoiceData struct {
	// The invoice requirement information.
	// InvoiceRequirement values:
	// 	Individual - Buyer requested a separate invoice for each order item in the order.
	// 	Consolidated – Buyer requested one invoice to include all of the order items in the order.
	// 	MustNotSend – Buyer did not request an invoice.
	InvoiceRequirement string
	// Invoice category information selected by the buyer at the time the order was placed.
	BuyerSelectedInvoiceCategory string
	// The buyer-specified invoice title.
	InvoiceTitle string
	// InvoiceInformation values:
	// 	NotApplicable - Buyer did not request an invoice.
	// 	BuyerSelectedInvoiceCategory – Amazon recommends using the BuyerSelectedInvoiceCategory value returned by this operation for the invoice category on the invoice.
	// 	ProductTitle – Amazon recommends using the product title for invoice category on the invoice.
	InvoiceInformation string
}
