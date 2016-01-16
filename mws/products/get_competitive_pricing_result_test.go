package products

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func prepareGetCompetitivePricingForSKUResult() *GetCompetitivePricingForSKUResult {
	return loadExample("GetCompetitivePricingForSKU").(*GetCompetitivePricingForSKUResult)
}

func Test_GetCompetitivePricingForSKUResult(t *testing.T) {
	Convey("Request response", t, func() {
		gcpResult := prepareGetCompetitivePricingForSKUResult()

		Convey("Has 1 Results", func() {
			So(gcpResult.Results, ShouldHaveLength, 1)
		})

		Convey("ProductResult has product", func() {
			So(gcpResult.Results[0].Product, ShouldNotBeNil)
		})
	})
}

func Test_GetCompetitivePricingForSKUResult_Product(t *testing.T) {
	gmpResult := prepareGetCompetitivePricingForSKUResult()

	Convey("Product", t, func() {
		p := gmpResult.Results[0].Product

		Convey("Identifiers is not nil", func() {
			So(p.Identifiers, ShouldNotBeNil)
		})

		Convey("CompetitivePricing is not nil", func() {
			So(p.CompetitivePricing, ShouldNotBeNil)
		})

		Convey("CompetitivePricing has 2 CompetitivePrice", func() {
			So(p.CompetitivePricing.CompetitivePrices, ShouldHaveLength, 2)
		})

		Convey("CompetitivePricing has 3 NumberOfOfferListings", func() {
			So(p.CompetitivePricing.NumberOfOfferListings, ShouldHaveLength, 3)
		})

		Convey("CompetitivePricing has TradeInValue", func() {
			So(p.CompetitivePricing.TradeInValue, ShouldNotBeNil)
		})

		Convey("SalesRankings has 4 SalesRank", func() {
			So(p.SalesRankings, ShouldHaveLength, 4)
		})
	})
}

func Test_GetCompetitivePricingForSKUResult_Product_Identifiers(t *testing.T) {
	gmpResult := prepareGetCompetitivePricingForSKUResult()

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

func Test_GetCompetitivePricingForSKUResult_Product_CompetitivePricing(t *testing.T) {
	gmpResult := prepareGetCompetitivePricingForSKUResult()
	comPricing := gmpResult.Results[0].Product.CompetitivePricing

	Convey("CompetitivePrice 1", t, func() {
		cp1 := comPricing.CompetitivePrices[0]

		Convey("BelongsToRequester is true", func() {
			So(cp1.BelongsToRequester, ShouldBeTrue)
		})

		Convey("Condition is New", func() {
			So(cp1.Condition, ShouldEqual, "New")
		})

		Convey("Subcondition is New", func() {
			So(cp1.Subcondition, ShouldEqual, "New")
		})

		Convey("CompetitivePriceId is 1", func() {
			So(cp1.CompetitivePriceId, ShouldEqual, "1")
		})

		Convey("Price", func() {
			price := cp1.Price
			Convey("LandedPrice", func() {
				moneyAsserter(price.LandedPrice, "USD", 40.03)
			})

			Convey("ListingPrice", func() {
				moneyAsserter(price.ListingPrice, "USD", 40.03)
			})

			Convey("Shipping", func() {
				moneyAsserter(price.Shipping, "USD", 0)
			})
		})
	})

	Convey("CompetitivePrice 2", t, func() {
		cp2 := comPricing.CompetitivePrices[1]

		Convey("BelongsToRequester is false", func() {
			So(cp2.BelongsToRequester, ShouldBeFalse)
		})

		Convey("Condition is Used", func() {
			So(cp2.Condition, ShouldEqual, "Used")
		})

		Convey("Subcondition is Good", func() {
			So(cp2.Subcondition, ShouldEqual, "Good")
		})

		Convey("CompetitivePriceId is 2", func() {
			So(cp2.CompetitivePriceId, ShouldEqual, "2")
		})

		Convey("Price", func() {
			price := cp2.Price
			Convey("LandedPrice", func() {
				moneyAsserter(price.LandedPrice, "USD", 30.50)
			})

			Convey("ListingPrice", func() {
				moneyAsserter(price.ListingPrice, "USD", 30.50)
			})

			Convey("Shipping", func() {
				moneyAsserter(price.Shipping, "USD", 0)
			})
		})
	})

	Convey("NumberOfOfferListings", t, func() {
		Convey("OfferListingCount 1", func() {
			ol1 := comPricing.NumberOfOfferListings[0]
			Convey("Condition is Any", func() {
				So(ol1.Condition, ShouldEqual, "Any")
			})

			Convey("Value is 296", func() {
				So(ol1.Value, ShouldEqual, 296)
			})
		})

		Convey("OfferListingCount 2", func() {
			ol2 := comPricing.NumberOfOfferListings[1]
			Convey("Condition is Used", func() {
				So(ol2.Condition, ShouldEqual, "Used")
			})

			Convey("Value is 145", func() {
				So(ol2.Value, ShouldEqual, 145)
			})
		})

		Convey("OfferListingCount 3", func() {
			ol3 := comPricing.NumberOfOfferListings[2]
			Convey("Condition is New", func() {
				So(ol3.Condition, ShouldEqual, "New")
			})

			Convey("Value is 151", func() {
				So(ol3.Value, ShouldEqual, 151)
			})
		})
	})

	Convey("TradeInValue", t, func() {
		moneyAsserter(comPricing.TradeInValue, "USD", 17.05)
	})
}
