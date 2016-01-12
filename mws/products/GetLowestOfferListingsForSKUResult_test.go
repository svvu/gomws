package products

import (
	"io/ioutil"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/svvu/gomws/gmws"
	"github.com/svvu/gomws/mwsHttps"
)

var GetLowestOfferListingsForSKUResultResponse, _ = ioutil.ReadFile(
	"./examples/GetLowestOfferListingsForSKU.xml",
)

func prepareGetLowestOfferListingsForSKUResult() *GetLowestOfferListingsForSKUResult {
	resp := &mwsHttps.Response{Result: string(GetLowestOfferListingsForSKUResultResponse)}
	xmlParser := gmws.NewXMLParser(resp)
	glolResult := GetLowestOfferListingsForSKUResult{}
	xmlParser.Parse(&glolResult)
	return &glolResult
}

func Test_GetLowestOfferListingsForSKUResult(t *testing.T) {
	Convey("Request response", t, func() {
		glolResult := prepareGetLowestOfferListingsForSKUResult()

		Convey("Has 1 Results", func() {
			So(glolResult.Results, ShouldHaveLength, 1)
		})

		Convey("AllOfferListingsConsidered is false", func() {
			So(glolResult.Results[0].AllOfferListingsConsidered, ShouldBeFalse)
		})

		Convey("ProductResult has product", func() {
			So(glolResult.Results[0].Product, ShouldNotBeNil)
		})
	})
}

func Test_GetLowestOfferListingsForSKUResult_Product(t *testing.T) {
	glolResult := prepareGetLowestOfferListingsForSKUResult()

	Convey("Product", t, func() {
		p := glolResult.Results[0].Product

		Convey("Identifiers is not nil", func() {
			So(p.Identifiers, ShouldNotBeNil)
		})

		Convey("LowestOfferListings is not nil", func() {
			So(p.LowestOfferListings, ShouldNotBeNil)
		})

		Convey("LowestOfferListings has 3 LowestOfferListing", func() {
			So(p.LowestOfferListings, ShouldHaveLength, 3)
		})

	})
}

func Test_GetLowestOfferListingsForSKUResult_Product_Identifiers(t *testing.T) {
	glolResult := prepareGetLowestOfferListingsForSKUResult()

	Convey("Product Identifiers", t, func() {
		iden := glolResult.Results[0].Product.Identifiers

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

func Test_GetLowestOfferListingsForSKUResult_Product_LowestOfferListing1(t *testing.T) {
	glolResult := prepareGetLowestOfferListingsForSKUResult()

	expectedListing := []map[string]string{
		{
			"ItemCondition":                   "Used",
			"ItemSubcondition":                "Acceptable",
			"FulfillmentChannel":              "Merchant",
			"ShipsDomestically":               "True",
			"ShippingTime.Max":                "0-2 days",
			"SellerPositiveFeedbackRating":    "95-97%",
			"NumberOfOfferListingsConsidered": "3",
			"SellerFeedbackCount":             "8900",
			"LandedPrice":                     "28.68",
			"ListingPrice":                    "24.69",
			"Shipping":                        "3.99",
			"MultipleOffersAtLowestPrice":     "True",
		},
		{
			"ItemCondition":                   "Used",
			"ItemSubcondition":                "Good",
			"FulfillmentChannel":              "Amazon",
			"ShipsDomestically":               "True",
			"ShippingTime.Max":                "0-2 days",
			"SellerPositiveFeedbackRating":    "90-94%",
			"NumberOfOfferListingsConsidered": "1",
			"SellerFeedbackCount":             "1569694",
			"LandedPrice":                     "30.50",
			"ListingPrice":                    "30.50",
			"Shipping":                        "0",
			"MultipleOffersAtLowestPrice":     "False",
		},
		{
			"ItemCondition":                   "Used",
			"ItemSubcondition":                "Good",
			"FulfillmentChannel":              "Merchant",
			"ShipsDomestically":               "True",
			"ShippingTime.Max":                "0-2 days",
			"SellerPositiveFeedbackRating":    "95-97%",
			"NumberOfOfferListingsConsidered": "3",
			"SellerFeedbackCount":             "7732",
			"LandedPrice":                     "30.99",
			"ListingPrice":                    "27.00",
			"Shipping":                        "3.99",
			"MultipleOffersAtLowestPrice":     "False",
		},
	}

	for i, listingResult := range expectedListing {
		Convey("Product LowestOfferListing "+strconv.Itoa(i+1), t, func() {
			listing := glolResult.Results[0].Product.LowestOfferListings[i]

			Convey("Qualifiers", func() {
				q := listing.Qualifiers

				Convey("ItemCondition is Used", func() {
					So(q.ItemCondition, ShouldEqual, listingResult["ItemCondition"])
				})

				Convey("ItemSubcondition is Acceptable", func() {
					So(q.ItemSubcondition, ShouldEqual, listingResult["ItemSubcondition"])
				})

				Convey("FulfillmentChannel is Merchant", func() {
					So(q.FulfillmentChannel, ShouldEqual, listingResult["FulfillmentChannel"])
				})

				Convey("ShipsDomestically is True", func() {
					So(q.ShipsDomestically, ShouldEqual, listingResult["ShipsDomestically"])
				})

				Convey("ShippingTime Max is 0-2 days", func() {
					So(q.ShippingTime.MaximumDayRange, ShouldEqual, listingResult["ShippingTime.Max"])
				})

				Convey("SellerPositiveFeedbackRating is 95-97%", func() {
					So(q.SellerPositiveFeedbackRating, ShouldEqual, listingResult["SellerPositiveFeedbackRating"])
				})
			})

			Convey("NumberOfOfferListingsConsidered is 3", func() {
				numOffer, _ := strconv.Atoi(listingResult["NumberOfOfferListingsConsidered"])
				So(listing.NumberOfOfferListingsConsidered, ShouldEqual, numOffer)
			})

			Convey("SellerFeedbackCount is 8900", func() {
				feedback, _ := strconv.Atoi(listingResult["SellerFeedbackCount"])
				So(listing.SellerFeedbackCount, ShouldEqual, feedback)
			})

			Convey("Price", func() {
				price := listing.Price

				Convey("LandedPrice", func() {
					landingPrice, _ := strconv.ParseFloat(listingResult["LandedPrice"], 64)
					moneyAsserter(price.LandedPrice, "USD", landingPrice)
				})

				Convey("ListingPrice", func() {
					listingPrice, _ := strconv.ParseFloat(listingResult["ListingPrice"], 64)
					moneyAsserter(price.ListingPrice, "USD", listingPrice)
				})

				Convey("Shipping", func() {
					shipping, _ := strconv.ParseFloat(listingResult["Shipping"], 64)
					moneyAsserter(price.Shipping, "USD", shipping)
				})
			})

			Convey("MultipleOffersAtLowestPrice is True", func() {
				So(listing.MultipleOffersAtLowestPrice, ShouldEqual, listingResult["MultipleOffersAtLowestPrice"])
			})
		})
	}
}
