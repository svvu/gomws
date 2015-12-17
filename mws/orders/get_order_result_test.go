package orders

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/svvu/gomws/gmws"
	"github.com/svvu/gomws/mwsHttps"
)

func TestGetOrderResult(t *testing.T) {
	Convey("Request result", t, func() {
		resp := &mwsHttps.Response{Body: []byte(loadExample("GetOrder"))}
		xmlParser := gmws.NewXMLParser(resp)
		goResult := GetOrderResult{}
		xmlParser.Parse(&goResult)

		Convey("has NextToken", func() {
			So(goResult.NextToken, ShouldEqual, "2YgYW55IGNhcm5hbCBwbGVhc3VyZS4=")
		})

		Convey("has LastUpdatedBefore", func() {
			So(goResult.LastUpdatedBefore, ShouldEqual, "2013-09-25T18%3A10%3A21.687Z")
		})

		Convey("Order 1", func() {
			o1 := goResult.Orders[0]
			Convey("has AmazonOrderId", func() {
				So(o1.AmazonOrderId, ShouldEqual, "058-1233752-8214740")
			})

			Convey("has PurchaseDate", func() {
				So(o1.PurchaseDate, ShouldEqual, "2013-09-05T00%3A06%3A07.000Z")
			})

			Convey("has LastUpdateDate", func() {
				So(o1.LastUpdateDate, ShouldEqual, "2013-09-07T12%3A43%3A16.000Z")
			})

			Convey("has OrderStatus", func() {
				So(o1.OrderStatus, ShouldEqual, "Unshipped")
			})

			Convey("has OrderType", func() {
				So(o1.OrderType, ShouldEqual, "StandardOrder")
			})

			Convey("has ShipServiceLevel", func() {
				So(o1.ShipServiceLevel, ShouldEqual, "Std JP Kanto8")
			})

			Convey("has FulfillmentChannel", func() {
				So(o1.FulfillmentChannel, ShouldEqual, "MFN")
			})

			Convey("has OrderTotal CurrencyCode", func() {
				So(o1.OrderTotal.CurrencyCode, ShouldEqual, "JPY")
			})

			Convey("has OrderTotal Amount", func() {
				So(o1.OrderTotal.Amount, ShouldEqual, "1507.00")
			})

			Convey("has ShippingAddress Name", func() {
				So(o1.ShippingAddress.Name, ShouldEqual, "Jane Smith")
			})

			Convey("has ShippingAddress AddressLine1", func() {
				So(o1.ShippingAddress.AddressLine1, ShouldEqual, "1-2-10 Akasaka")
			})

			Convey("has ShippingAddress City", func() {
				So(o1.ShippingAddress.City, ShouldEqual, "Tokyo")
			})

			Convey("has ShippingAddress PostalCode", func() {
				So(o1.ShippingAddress.PostalCode, ShouldEqual, "107-0053")
			})

			Convey("has ShippingAddress CountryCode", func() {
				So(o1.ShippingAddress.CountryCode, ShouldEqual, "JP")
			})

			Convey("has NumberOfItemsShipped", func() {
				So(o1.NumberOfItemsShipped, ShouldEqual, 0)
			})

			Convey("has NumberOfItemsUnshipped", func() {
				So(o1.NumberOfItemsUnshipped, ShouldEqual, 1)
			})

			Convey("has 3 PaymentExecutionDetail", func() {
				So(o1.PaymentExecutionDetail, ShouldHaveLength, 3)
			})

			Convey("PaymentExecutionDetailItem 1", func() {
				pd1 := o1.PaymentExecutionDetail[0]
				Convey("has Payment Amount", func() {
					So(pd1.Payment.Amount, ShouldEqual, "10.00")
				})

				Convey("has Payment CurrencyCode", func() {
					So(pd1.Payment.CurrencyCode, ShouldEqual, "JPY")
				})

				Convey("has PaymentMethod", func() {
					So(pd1.PaymentMethod, ShouldEqual, "PointsAccount")
				})
			})

			Convey("PaymentExecutionDetailItem 2", func() {
				pd2 := o1.PaymentExecutionDetail[1]
				Convey("has Payment Amount", func() {
					So(pd2.Payment.Amount, ShouldEqual, "317.00")
				})

				Convey("has Payment CurrencyCode", func() {
					So(pd2.Payment.CurrencyCode, ShouldEqual, "JPY")
				})

				Convey("has PaymentMethod", func() {
					So(pd2.PaymentMethod, ShouldEqual, "GC")
				})
			})

			Convey("PaymentExecutionDetailItem 3", func() {
				pd3 := o1.PaymentExecutionDetail[2]
				Convey("has Payment Amount", func() {
					So(pd3.Payment.Amount, ShouldEqual, "1180.00")
				})

				Convey("has Payment CurrencyCode", func() {
					So(pd3.Payment.CurrencyCode, ShouldEqual, "JPY")
				})

				Convey("has PaymentMethod", func() {
					So(pd3.PaymentMethod, ShouldEqual, "COD")
				})
			})

			Convey("has PaymentMethod", func() {
				So(o1.PaymentMethod, ShouldEqual, "COD")
			})

			Convey("has MarketplaceId", func() {
				So(o1.MarketplaceId, ShouldEqual, "ATVPDKIKX0DER")
			})

			Convey("has BuyerName", func() {
				So(o1.BuyerName, ShouldEqual, "Jane Smith")
			})

			Convey("has BuyerEmail", func() {
				So(o1.BuyerEmail, ShouldEqual, "5vlhEXAMPLEh9h5@marketplace.amazon.com")
			})

			Convey("has ShipmentServiceLevelCategory", func() {
				So(o1.ShipmentServiceLevelCategory, ShouldEqual, "Standard")
			})

			Convey("has IsBusinessOrder", func() {
				So(o1.IsBusinessOrder, ShouldBeFalse)
			})

			Convey("has IsPrime", func() {
				So(o1.IsPrime, ShouldBeFalse)
			})

			Convey("has IsPremiumOrder", func() {
				So(o1.IsPremiumOrder, ShouldBeFalse)
			})
		})
	})
}
