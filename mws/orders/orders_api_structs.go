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
