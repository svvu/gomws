package orders

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/svvu/gomws/gmws"
	"github.com/svvu/gomws/mwsHttps"
)

func TestListOrderItemsByNextTokenResult(t *testing.T) {
	Convey("Request result", t, func() {
		resp := &mwsHttps.Response{Body: []byte(loadExample("ListOrderItemsByNextToken"))}
		xmlParser := gmws.NewXMLParser(resp)
		loiResult := ListOrderItemsByNextTokenResult{}
		xmlParser.Parse(&loiResult)

		Convey("has NextToken", func() {
			So(loiResult.NextToken, ShouldEqual, "MRgZW55IGNhcm5hbCBwbGVhc3VyZS6=")
		})

		Convey("has AmazonOrderId", func() {
			So(loiResult.AmazonOrderId, ShouldEqual, "058-1233752-8214740")
		})

		Convey("has 2 OrderItems", func() {
			So(loiResult.OrderItems, ShouldHaveLength, 2)
		})

		Convey("Order 1", func() {
			o1 := loiResult.OrderItems[0]
			Convey("has ASIN", func() {
				So(o1.ASIN, ShouldEqual, "BT0093TELA")
			})

			Convey("has OrderItemId", func() {
				So(o1.OrderItemId, ShouldEqual, "68828574383266")
			})

			Convey("has SellerSKU", func() {
				So(o1.SellerSKU, ShouldEqual, "CBA_OTF_1")
			})

			Convey("has Title", func() {
				So(o1.Title, ShouldEqual, "Example item name")
			})

			Convey("has QuantityOrdered", func() {
				So(o1.QuantityOrdered, ShouldEqual, 1)
			})

			Convey("has QuantityShipped", func() {
				So(o1.QuantityShipped, ShouldEqual, 1)
			})

			Convey("has ItemPrice CurrencyCode", func() {
				So(o1.ItemPrice.CurrencyCode, ShouldEqual, "USD")
			})

			Convey("has ItemPrice Amount", func() {
				So(o1.ItemPrice.Amount, ShouldEqual, "25.99")
			})

			Convey("has ShippingPrice CurrencyCode", func() {
				So(o1.ShippingPrice.CurrencyCode, ShouldEqual, "USD")
			})

			Convey("has ShippingPrice Amount", func() {
				So(o1.ShippingPrice.Amount, ShouldEqual, "1.26")
			})

			Convey("has ScheduledDeliveryEndDate", func() {
				So(o1.ScheduledDeliveryEndDate, ShouldEqual, "2013-09-091T01:30:00.000-06:00")
			})

			Convey("has ScheduledDeliveryStartDate", func() {
				So(o1.ScheduledDeliveryStartDate, ShouldEqual, "2013-09-06T02:00:00.000-06:00")
			})

			Convey("has CODFee CurrencyCode", func() {
				So(o1.CODFee.CurrencyCode, ShouldEqual, "USD")
			})

			Convey("has CODFee Amount", func() {
				So(o1.CODFee.Amount, ShouldEqual, "10.00")
			})

			Convey("has CODFeeDiscount CurrencyCode", func() {
				So(o1.CODFeeDiscount.CurrencyCode, ShouldEqual, "USD")
			})

			Convey("has CODFeeDiscount Amount", func() {
				So(o1.CODFeeDiscount.Amount, ShouldEqual, "1.00")
			})

			Convey("has GiftMessageText", func() {
				So(o1.GiftMessageText, ShouldEqual, "For you!")
			})

			Convey("has GiftWrapPrice CurrencyCode", func() {
				So(o1.GiftWrapPrice.CurrencyCode, ShouldEqual, "USD")
			})

			Convey("has GiftWrapPrice Amount", func() {
				So(o1.GiftWrapPrice.Amount, ShouldEqual, "1.99")
			})

			Convey("has GiftWrapLevel", func() {
				So(o1.GiftWrapLevel, ShouldEqual, "Classic")
			})

			Convey("has PriceDesignation", func() {
				So(o1.PriceDesignation, ShouldEqual, "BusinessPrice")
			})
		})

		Convey("Order 2", func() {
			o2 := loiResult.OrderItems[1]
			Convey("has ASIN", func() {
				So(o2.ASIN, ShouldEqual, "BCTU1104UEFB")
			})

			Convey("has OrderItemId", func() {
				So(o2.OrderItemId, ShouldEqual, "79039765272157")
			})

			Convey("has SellerSKU", func() {
				So(o2.SellerSKU, ShouldEqual, "CBA_OTF_5")
			})

			Convey("has Title", func() {
				So(o2.Title, ShouldEqual, "Example item name")
			})

			Convey("has QuantityOrdered", func() {
				So(o2.QuantityOrdered, ShouldEqual, 2)
			})

			Convey("has ItemPrice CurrencyCode", func() {
				So(o2.ItemPrice.CurrencyCode, ShouldEqual, "USD")
			})

			Convey("has ItemPrice Amount", func() {
				So(o2.ItemPrice.Amount, ShouldEqual, "17.95")
			})

			Convey("has PromotionId", func() {
				poid := o2.PromotionIds[0]
				So(poid, ShouldEqual, "FREESHIP")
			})

			Convey("has ConditionId", func() {
				So(o2.ConditionId, ShouldEqual, "Used")
			})

			Convey("has ConditionSubtypeId", func() {
				So(o2.ConditionSubtypeId, ShouldEqual, "Mint")
			})

			Convey("has ConditionNote", func() {
				So(o2.ConditionNote, ShouldEqual, "Example ConditionNote")
			})

			Convey("has PriceDesignation", func() {
				So(o2.PriceDesignation, ShouldEqual, "BusinessPrice")
			})
		})
	})
}
