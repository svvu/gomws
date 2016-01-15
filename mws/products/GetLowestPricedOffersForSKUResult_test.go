package products

import (
	"fmt"
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/svvu/gomws/gmws"
	"github.com/svvu/gomws/mwsHttps"
)

var GetLowestPricedOffersForSKUResultResponse, _ = ioutil.ReadFile(
	"./examples/GetLowestPricedOffersForSKU.xml",
)

func prepareGetLowestPricedOffersForSKUResult() *GetLowestPricedOffersForSKUResult {
	resp := &mwsHttps.Response{Result: string(GetLowestPricedOffersForSKUResultResponse)}
	xmlParser := gmws.NewXMLParser(resp)
	glplResult := GetLowestPricedOffersForSKUResult{}
	err := xmlParser.Parse(&glplResult)
	if err != nil {
		fmt.Println(err)
	}
	return &glplResult
}

func Test_GetLowestPricedOffersForSKUResult(t *testing.T) {
	Convey("Request response", t, func() {
		glplResult := prepareGetLowestPricedOffersForSKUResult()

		Convey("Result is not nil", func() {
			So(glplResult.Result, ShouldNotBeNil)
		})

		Convey("Result has Identifier", func() {
			So(glplResult.Result.Identifier, ShouldNotBeNil)
		})

		Convey("Result has Summary", func() {
			So(glplResult.Result.Summary, ShouldNotBeNil)
		})

		Convey("Result has 1 Offer", func() {
			So(glplResult.Result.Offers, ShouldHaveLength, 1)
		})
	})
}

func Test_GetLowestPricedOffersForSKUResult_Identifier(t *testing.T) {
	glplResult := prepareGetLowestPricedOffersForSKUResult()

	Convey("Identifier", t, func() {
		iden := glplResult.Result.Identifier

		Convey("MarketplaceId is ATVPDKIKX0DER", func() {
			So(iden.MarketplaceId, ShouldEqual, "ATVPDKIKX0DER")
		})

		Convey("SellerSKU is GE Product", func() {
			So(iden.SellerSKU, ShouldEqual, "GE Product")
		})

		Convey("ItemCondition is New", func() {
			So(iden.ItemCondition, ShouldEqual, "New")
		})

		Convey("TimeOfOfferChange is 2015-07-19T23:15:11.859Z", func() {
			So(iden.TimeOfOfferChange, ShouldEqual, "2015-07-19T23:15:11.859Z")
		})

	})
}

func Test_GetLowestPricedOffersForSKUResult_Summary(t *testing.T) {
	glplResult := prepareGetLowestPricedOffersForSKUResult()

	Convey("Summary", t, func() {
		summary := glplResult.Result.Summary

		Convey("TotalOfferCount is 1", func() {
			So(summary.TotalOfferCount, ShouldEqual, 1)
		})

		Convey("NumberOfOffers Has 1 OfferCount", func() {
			So(summary.NumberOfOffers, ShouldHaveLength, 1)
		})

		Convey("OfferCount", func() {
			oc := summary.NumberOfOffers[0]

			Convey("Condition is new", func() {
				So(oc.Condition, ShouldEqual, "new")
			})

			Convey("FulfillmentChannel is Amazon", func() {
				So(oc.FulfillmentChannel, ShouldEqual, "Amazon")
			})

			Convey("Value is 1", func() {
				So(oc.Value, ShouldEqual, 1)
			})
		})

		Convey("LowestPrices Has 1 LowestPrice", func() {
			So(summary.LowestPrices, ShouldHaveLength, 1)
		})

		Convey("LowestPrice", func() {
			lp := summary.LowestPrices[0]

			Convey("Condition is new", func() {
				So(lp.Condition, ShouldEqual, "new")
			})

			Convey("FulfillmentChannel is Amazon", func() {
				So(lp.FulfillmentChannel, ShouldEqual, "Amazon")
			})

			Convey("LandedPrice", func() {
				moneyAsserter(lp.LandedPrice, "USD", 32.99)
			})

			Convey("ListingPrice", func() {
				moneyAsserter(lp.ListingPrice, "USD", 32.99)
			})

			Convey("Shipping", func() {
				moneyAsserter(lp.Shipping, "USD", 0.00)
			})
		})

		Convey("BuyBoxPrices", func() {
			bp := summary.BuyBoxPrices[0]

			Convey("Condition is New", func() {
				So(bp.Condition, ShouldEqual, "New")
			})

			Convey("LandedPrice", func() {
				moneyAsserter(bp.LandedPrice, "USD", 32.99)
			})

			Convey("ListingPrice", func() {
				moneyAsserter(bp.ListingPrice, "USD", 32.99)
			})

			Convey("Shipping", func() {
				moneyAsserter(bp.Shipping, "USD", 0.00)
			})
		})

		Convey("ListPrice", func() {
			moneyAsserter(summary.ListPrice, "USD", 58.34)
		})

		Convey("SuggestedLowerPricePlusShipping", func() {
			moneyAsserter(summary.SuggestedLowerPricePlusShipping, "USD", 32.99)
		})

		Convey("BuyBoxEligibleOffers", func() {
			oc := summary.BuyBoxEligibleOffers[0]

			Convey("Condition is new", func() {
				So(oc.Condition, ShouldEqual, "new")
			})

			Convey("FulfillmentChannel is Amazon", func() {
				So(oc.FulfillmentChannel, ShouldEqual, "Amazon")
			})

			Convey("Value is 1", func() {
				So(oc.Value, ShouldEqual, 1)
			})
		})
	})
}

func Test_GetLowestPricedOffersForSKUResult_Offers(t *testing.T) {
	glplResult := prepareGetLowestPricedOffersForSKUResult()

	Convey("Offer", t, func() {
		offer := glplResult.Result.Offers[0]

		Convey("MyOffer is false", func() {
			So(offer.MyOffer, ShouldBeFalse)
		})

		Convey("SubCondition is new", func() {
			So(offer.SubCondition, ShouldEqual, "new")
		})

		Convey("SellerFeedbackRating", func() {
			sfr := offer.SellerFeedbackRating

			Convey("SellerPositiveFeedbackRating is 100.0", func() {
				So(sfr.SellerPositiveFeedbackRating, ShouldEqual, 100.0)
			})

			Convey("FeedbackCount is 1", func() {
				So(sfr.FeedbackCount, ShouldEqual, 1)
			})
		})

		Convey("ShippingTime", func() {
			st := offer.ShippingTime

			Convey("MinimumHours is 0", func() {
				So(st.MinimumHours, ShouldEqual, "0")
			})

			Convey("MaximumHours is 0", func() {
				So(st.MaximumHours, ShouldEqual, "0")
			})

			Convey("AvailabilityType is NOW", func() {
				So(st.AvailabilityType, ShouldEqual, "NOW")
			})
		})

		Convey("ListingPrice", func() {
			moneyAsserter(offer.ListingPrice, "USD", 32.99)
		})

		Convey("Shipping", func() {
			moneyAsserter(offer.Shipping, "USD", 0)
		})

		Convey("IsFulfilledByAmazon is true", func() {
			So(offer.IsFulfilledByAmazon, ShouldBeTrue)
		})

		Convey("IsBuyBoxWinner is true", func() {
			So(offer.IsBuyBoxWinner, ShouldBeTrue)
		})

		Convey("IsFeaturedMerchant is true", func() {
			So(offer.IsFeaturedMerchant, ShouldBeTrue)
		})
	})
}
