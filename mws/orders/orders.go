// Reference http://docs.developer.amazonservices.com/en_US/orders/2013-09-01/Orders_Overview.html

package orders

import (
	"github.com/svvu/gomws/gmws"
	"github.com/svvu/gomws/mwsHttps"
)

// Orders is the client for the api
type Orders struct {
	*gmws.MwsBase
}

// NewClient generate a new orders client
func NewClient(config gmws.MwsConfig) (*Orders, error) {
	orders := new(Orders)
	base, err := gmws.NewMwsBase(config, orders.Version(), orders.Name())
	if err != nil {
		return nil, err
	}
	orders.MwsBase = base
	return orders, nil
}

// Version return the current version of api
func (o Orders) Version() string {
	return "2013-09-01"
}

// Name return the name of the api
func (o Orders) Name() string {
	return "Orders"
}

// GetServiceStatus Returns the operational status of the Orders API section.
// http://docs.developer.amazonservices.com/en_US/orders/2013-09-01/MWS_GetServiceStatus.html
func (o Orders) GetServiceStatus() *mwsHttps.Response {
	params := gmws.Parameters{
		"Action": "GetServiceStatus",
	}

	return o.SendRequest(params)
}

// ListOrders Returns orders created or updated during a time frame that you specify.
//
// Note: When calling this operation, either CreatedAfter or LastUpdatedAfter must be specify.
// Specify both will return an error.
//
// other params:
// 	CreatedAfter - string, ISO-8601 date format.
// 		Required, if LastUpdatedAfter is not specified.
// 	CreatedBefore - string, ISO-8601 date format.
// 		Can only be specified if CreatedAfter is specified.
//	LastUpdatedAfter - string, ISO-8601 date format.
// 		Required, if CreatedAfter is not specified.
// 		If LastUpdatedAfter is specified, then BuyerEmail and SellerOrderId cannot be specified.
// 	LastUpdatedBefore - string, ISO-8601 date format.
// 		Can only be specified if LastUpdatedAfter is specified.
// 	OrderStatus - []string.
// 		Values: PendingAvailability, Pending, Unshipped, PartiallyShipped, Shipped,
// 		InvoiceUnconfirmed, Canceled, Unfulfillable. Default: All.
//	FulfillmentChannel - []string.
// 		Values: AFN, MFN. Default: All.
// 	PaymentMethod - []string.
// 		Values: COD, CVS, Other. Default: All.
// 	BuyerEmail - string.
// 		If BuyerEmail is specified, then FulfillmentChannel, OrderStatus, PaymentMethod,
// 		LastUpdatedAfter, LastUpdatedBefore, and SellerOrderId cannot be specified.
// 	SellerOrderId - string.
// 		If SellerOrderId is specified, then FulfillmentChannel, OrderStatus,
// 		PaymentMethod, LastUpdatedAfter, LastUpdatedBefore, and BuyerEmail cannot be specified.
// 	MaxResultsPerPage - int.
// 		Value 1 - 100. Default 100.
// 	TFMShipmentStatus - []string.
// 		Values: PendingPickUp, LabelCanceled, PickedUp, AtDestinationFC,
// 		Delivered, RejectedByBuyer, Undeliverable, ReturnedToSeller, Lost.
func (o Orders) ListOrders(others ...gmws.Parameters) *mwsHttps.Response {
	op := gmws.OptionalParams([]string{
		"CreatedAfter", "CreatedBefore",
		"LastUpdatedAfter", "LastUpdatedBefore",
		"OrderStatus", "FulfillmentChannel", "PaymentMethod",
		"SellerOrderId", "BuyerEmail",
		"TFMShipmentStatus", "MaxResultsPerPage",
	}, others)
	params := gmws.Parameters{
		"Action":        "ListOrders",
		"MarketplaceId": []string{o.MarketPlaceId},
	}.Merge(op)

	structuredParams := params.StructureKeys("MarketplaceId", "Id")
	if _, ok := structuredParams["OrderStatus"]; ok {
		structuredParams = params.StructureKeys("OrderStatus", "Status")
	}
	if _, ok := structuredParams["FulfillmentChannel"]; ok {
		structuredParams = params.StructureKeys("FulfillmentChannel", "Channel")
	}
	if _, ok := structuredParams["OrderStatus"]; ok {
		structuredParams = params.StructureKeys("PaymentMethod", "Method")
	}
	if _, ok := structuredParams["OrderStatus"]; ok {
		structuredParams = params.StructureKeys("TFMShipmentStatus", "Status")
	}

	return o.SendRequest(structuredParams)
}

// ListOrdersByNextToken Returns the next page of orders using the NextToken parameter.
// http://docs.developer.amazonservices.com/en_US/orders/2013-09-01/Orders_ListOrdersByNextToken.html
func (o Orders) ListOrdersByNextToken(nextToken string) *mwsHttps.Response {
	params := gmws.Parameters{
		"Action":    "ListOrdersByNextToken",
		"NextToken": nextToken,
	}

	return o.SendRequest(params)
}

// GetOrder Returns orders based on the AmazonOrderId values that you specify.
// Maximum 50 ids.
// http://docs.developer.amazonservices.com/en_US/orders/2013-09-01/Orders_GetOrder.html
func (o Orders) GetOrder(amazonOrderIds []string) *mwsHttps.Response {
	params := gmws.Parameters{
		"Action":        "GetOrder",
		"AmazonOrderId": amazonOrderIds,
	}
	structuredParams := params.StructureKeys("AmazonOrderId", "Id")

	return o.SendRequest(structuredParams)
}

// ListOrderItems Returns order items based on the AmazonOrderId that you specify.
// http://docs.developer.amazonservices.com/en_US/orders/2013-09-01/Orders_ListOrderItems.html
func (o Orders) ListOrderItems(amazonOrderID string) *mwsHttps.Response {
	params := gmws.Parameters{
		"Action":        "ListOrderItems",
		"AmazonOrderId": amazonOrderID,
	}

	return o.SendRequest(params)
}

// ListOrderItemsByNextToken Returns the next page of order items using the NextToken parameter.
// http://docs.developer.amazonservices.com/en_US/orders/2013-09-01/Orders_ListOrderItemsByNextToken.html
func (o Orders) ListOrderItemsByNextToken(nextToken string) *mwsHttps.Response {
	params := gmws.Parameters{
		"Action":    "ListOrderItemsByNextToken",
		"NextToken": nextToken,
	}

	return o.SendRequest(params)
}
