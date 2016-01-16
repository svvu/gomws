package products

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func prepareGetMyPriceForSKUResult() *GetMyPriceForSKUResult {
	return loadExample("GetMyPriceForSKU").(*GetMyPriceForSKUResult)
}

func Test_GetMyPriceForSKUResult(t *testing.T) {
	Convey("Request response", t, func() {
		gmpResult := prepareGetMyPriceForSKUResult()

		Convey("Has 1 Results", func() {
			So(gmpResult.Results, ShouldHaveLength, 1)
		})

		Convey("ProductResult is not nil", func() {
			So(gmpResult.Results[0].Product, ShouldNotBeNil)
		})

		Convey("ProductResult 1 status is Success", func() {
			So(gmpResult.Results[0].Status, ShouldEqual, "Success")
		})
	})
}

func Test_GetMyPriceForSKUResult_Product(t *testing.T) {
	gmpResult := prepareGetMyPriceForSKUResult()

	Convey("Product 1", t, func() {
		p1 := gmpResult.Results[0].Product

		Convey("Identifiers is not nil", func() {
			So(p1.Identifiers, ShouldNotBeNil)
		})

		Convey("Offers has 1 offer", func() {
			So(p1.Offers, ShouldHaveLength, 1)
		})
	})
}

func Test_GetMyPriceForSKUResult_Product_Identifiers(t *testing.T) {
	gmpResult := prepareGetMyPriceForSKUResult()

	Convey("Product Identifiers", t, func() {
		iden := gmpResult.Results[0].Product.Identifiers

		Convey("MarketplaceASIN is not nil", func() {
			So(iden.MarketplaceASIN, ShouldNotBeNil)
		})

		Convey("MarketplaceASIN", func() {
			masin := iden.MarketplaceASIN

			Convey("MarketplaceId is ATVPDKIKX0DER", func() {
				So(masin.MarketplaceId, ShouldEqual, "ATVPDKIKX0DER")
			})

			Convey("ASIN is 1933890517", func() {
				So(masin.ASIN, ShouldEqual, "1933890517")
			})
		})

		Convey("SKUIdentifier", func() {
			skuIden := iden.SKUIdentifier

			Convey("MarketplaceId is ATVPDKIKX0DER", func() {
				So(skuIden.MarketplaceId, ShouldEqual, "ATVPDKIKX0DER")
			})

			Convey("SellerId is A1IMEXAMPLEWRC", func() {
				So(skuIden.SellerId, ShouldEqual, "A1IMEXAMPLEWRC")
			})

			Convey("SellerSKU is SKU2468", func() {
				So(skuIden.SellerSKU, ShouldEqual, "SKU2468")
			})
		})
	})
}

func Test_GetMyPriceForSKUResult_Product_Offers(t *testing.T) {
	gmpResult := prepareGetMyPriceForSKUResult()

	Convey("Product Offer", t, func() {
		offer := gmpResult.Results[0].Product.Offers[0]

		Convey("BuyingPrice", func() {
			bp := offer.BuyingPrice

			Convey("LandedPrice", func() {
				moneyAsserter(bp.LandedPrice, "USD", 303.99)
			})

			Convey("ListingPrice", func() {
				moneyAsserter(bp.ListingPrice, "USD", 300.00)
			})

			Convey("Shipping", func() {
				moneyAsserter(bp.Shipping, "USD", 3.99)
			})
		})

		Convey("RegularPrice", func() {
			moneyAsserter(offer.RegularPrice, "USD", 300.00)
		})

		Convey("FulfillmentChannel is MERCHANT", func() {
			So(offer.FulfillmentChannel, ShouldEqual, "MERCHANT")
		})

		Convey("ItemCondition is Used", func() {
			So(offer.ItemCondition, ShouldEqual, "Used")
		})

		Convey("ItemSubCondition is Acceptable", func() {
			So(offer.ItemSubCondition, ShouldEqual, "Acceptable")
		})

		Convey("SellerId is A1IMEXAMPLEWRC", func() {
			So(offer.SellerId, ShouldEqual, "A1IMEXAMPLEWRC")
		})

		Convey("SellerSKU is SKU2468", func() {
			So(offer.SellerSKU, ShouldEqual, "SKU2468")
		})
	})
}
