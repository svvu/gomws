package orders

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/svvu/gomws/gmws"
	"github.com/svvu/gomws/mwsHttps"
)

func TestListOrdersByNextTokenResult(t *testing.T) {
	Convey("Request result", t, func() {
		resp := &mwsHttps.Response{Body: []byte(loadExample("ListOrdersByNextToken"))}
		xmlParser := gmws.NewXMLParser(resp)
		loinResult := ListOrdersByNextTokenResult{}
		xmlParser.Parse(&loinResult)

		Convey("has NextToken", func() {
			So(loinResult.NextToken, ShouldEqual, "2YgYW55IGNhcm5hbCBwbGVhc3VyZS4=")
		})

		Convey("has LastUpdatedBefore", func() {
			So(loinResult.LastUpdatedBefore, ShouldEqual, "2013-09-25T18%3A10%3A21.687Z")
		})

		Convey("has 2 Orders", func() {
			So(loinResult.Orders, ShouldHaveLength, 2)
		})

		Convey("Order 1", func() {
			o1 := loinResult.Orders[0]
			Convey("has ShipmentServiceLevelCategory", func() {
				So(o1.ShipmentServiceLevelCategory, ShouldEqual, "Standard")
			})

			Convey("has ShipServiceLevel", func() {
				So(o1.ShipServiceLevel, ShouldEqual, "Std JP Kanto8")
			})

			Convey("has EarliestShipDate", func() {
				So(o1.EarliestShipDate, ShouldEqual, "2013-08-20T19:51:16Z")
			})

			Convey("has LatestShipDate", func() {
				So(o1.LatestShipDate, ShouldEqual, "2013-08-25T19:49:35Z")
			})

			Convey("has MarketplaceId", func() {
				So(o1.MarketplaceId, ShouldEqual, "A1VC38T7YXB528")
			})

			Convey("has SalesChannel", func() {
				So(o1.SalesChannel, ShouldEqual, "Amazon.com")
			})

			Convey("has OrderType", func() {
				So(o1.OrderType, ShouldEqual, "Preorder")
			})

			Convey("has BuyerEmail", func() {
				So(o1.BuyerEmail, ShouldEqual, "5vlhEXAMPLEh9h5@marketplace.amazon.com")
			})

			Convey("has FulfillmentChannel", func() {
				So(o1.FulfillmentChannel, ShouldEqual, "MFN")
			})

			Convey("has OrderStatus", func() {
				So(o1.OrderStatus, ShouldEqual, "Pending")
			})

			Convey("has BuyerName", func() {
				So(o1.BuyerName, ShouldEqual, "John Jones")
			})

			Convey("has LastUpdateDate", func() {
				So(o1.LastUpdateDate, ShouldEqual, "2013-08-20T19:49:35Z")
			})

			Convey("has PurchaseDate", func() {
				So(o1.PurchaseDate, ShouldEqual, "2013-08-20T19:49:35Z")
			})

			Convey("has NumberOfItemsShipped", func() {
				So(o1.NumberOfItemsShipped, ShouldEqual, 0)
			})

			Convey("has NumberOfItemsUnshipped", func() {
				So(o1.NumberOfItemsUnshipped, ShouldEqual, 0)
			})

			Convey("has AmazonOrderId", func() {
				So(o1.AmazonOrderId, ShouldEqual, "902-3159896-1390916")
			})

			Convey("has PaymentMethod", func() {
				So(o1.PaymentMethod, ShouldEqual, "Other")
			})

			Convey("has IsBusinessOrder", func() {
				So(o1.IsBusinessOrder, ShouldBeTrue)
			})

			Convey("has PurchaseOrderNumber", func() {
				So(o1.PurchaseOrderNumber, ShouldEqual, "PO12345678")
			})

			Convey("has IsPrime", func() {
				So(o1.IsPrime, ShouldBeFalse)
			})

			Convey("has IsPremiumOrder", func() {
				So(o1.IsPremiumOrder, ShouldBeFalse)
			})
		})

		Convey("Order 2", func() {
			o2 := loinResult.Orders[1]
			Convey("has AmazonOrderId", func() {
				So(o2.AmazonOrderId, ShouldEqual, "058-1233752-8214740")
			})

			Convey("has PurchaseDate", func() {
				So(o2.PurchaseDate, ShouldEqual, "2013-09-05T00%3A06%3A07.000Z")
			})

			Convey("has LastUpdateDate", func() {
				So(o2.LastUpdateDate, ShouldEqual, "2013-09-07T12%3A43%3A16.000Z")
			})

			Convey("has OrderStatus", func() {
				So(o2.OrderStatus, ShouldEqual, "Unshipped")
			})

			Convey("has OrderType", func() {
				So(o2.OrderType, ShouldEqual, "StandardOrder")
			})

			Convey("has ShipServiceLevel", func() {
				So(o2.ShipServiceLevel, ShouldEqual, "Std JP Kanto8")
			})

			Convey("has FulfillmentChannel", func() {
				So(o2.FulfillmentChannel, ShouldEqual, "MFN")
			})

			Convey("has OrderTotal CurrencyCode", func() {
				So(o2.OrderTotal.CurrencyCode, ShouldEqual, "JPY")
			})

			Convey("has OrderTotal Amount", func() {
				So(o2.OrderTotal.Amount, ShouldEqual, "1507.00")
			})

			Convey("has ShippingAddress Name", func() {
				So(o2.ShippingAddress.Name, ShouldEqual, "Jane Smith")
			})

			Convey("has ShippingAddress AddressLine1", func() {
				So(o2.ShippingAddress.AddressLine1, ShouldEqual, "1-2-10 Akasaka")
			})

			Convey("has ShippingAddress City", func() {
				So(o2.ShippingAddress.City, ShouldEqual, "Tokyo")
			})

			Convey("has ShippingAddress PostalCode", func() {
				So(o2.ShippingAddress.PostalCode, ShouldEqual, "107-0053")
			})

			Convey("has ShippingAddress CountryCode", func() {
				So(o2.ShippingAddress.CountryCode, ShouldEqual, "JP")
			})

			Convey("has NumberOfItemsShipped", func() {
				So(o2.NumberOfItemsShipped, ShouldEqual, 0)
			})

			Convey("has NumberOfItemsUnshipped", func() {
				So(o2.NumberOfItemsUnshipped, ShouldEqual, 1)
			})

			Convey("has 3 PaymentExecutionDetail", func() {
				So(o2.PaymentExecutionDetail, ShouldHaveLength, 3)
			})

			Convey("PaymentExecutionDetailItem 1", func() {
				ped1 := o2.PaymentExecutionDetail[0]
				Convey("has Payment Amount", func() {
					So(ped1.Payment.Amount, ShouldEqual, "10.00")
				})

				Convey("has Payment CurrencyCode", func() {
					So(ped1.Payment.CurrencyCode, ShouldEqual, "JPY")
				})

				Convey("has PaymentMethod", func() {
					So(ped1.PaymentMethod, ShouldEqual, "PointsAccount")
				})
			})

			Convey("PaymentExecutionDetailItem 2", func() {
				ped2 := o2.PaymentExecutionDetail[1]
				Convey("has Payment Amount", func() {
					So(ped2.Payment.Amount, ShouldEqual, "317.00")
				})

				Convey("has Payment CurrencyCode", func() {
					So(ped2.Payment.CurrencyCode, ShouldEqual, "JPY")
				})

				Convey("has PaymentMethod", func() {
					So(ped2.PaymentMethod, ShouldEqual, "GC")
				})
			})

			Convey("PaymentExecutionDetailItem 3", func() {
				ped3 := o2.PaymentExecutionDetail[2]
				Convey("has Payment Amount", func() {
					So(ped3.Payment.Amount, ShouldEqual, "1180.00")
				})

				Convey("has Payment CurrencyCode", func() {
					So(ped3.Payment.CurrencyCode, ShouldEqual, "JPY")
				})

				Convey("has PaymentMethod", func() {
					So(ped3.PaymentMethod, ShouldEqual, "COD")
				})
			})

			Convey("has PaymentMethod", func() {
				So(o2.PaymentMethod, ShouldEqual, "COD")
			})

			Convey("has MarketplaceId", func() {
				So(o2.MarketplaceId, ShouldEqual, "ATVPDKIKX0DER")
			})

			Convey("has BuyerName", func() {
				So(o2.BuyerName, ShouldEqual, "Jane Smith")
			})

			Convey("has BuyerEmail", func() {
				So(o2.BuyerEmail, ShouldEqual, "5vlhEXAMPLEh9h5@marketplace.amazon.com")
			})

			Convey("has ShipmentServiceLevelCategory", func() {
				So(o2.ShipmentServiceLevelCategory, ShouldEqual, "Standard")
			})

			Convey("has IsBusinessOrder", func() {
				So(o2.IsBusinessOrder, ShouldBeFalse)
			})

			Convey("has IsPrime", func() {
				So(o2.IsPrime, ShouldBeFalse)
			})

			Convey("has IsPremiumOrder", func() {
				So(o2.IsPremiumOrder, ShouldBeFalse)
			})

		})
	})
}
