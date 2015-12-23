package products

import (
	"io/ioutil"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/svvu/gomws/gmws"
	"github.com/svvu/gomws/mwsHttps"
)

var GetMatchingProductResultResponse, _ = ioutil.ReadFile(
	"./examples/GetMatchingProduct.xml",
)

func prepareGetMatchingProductResult() *GetMatchingProductResult {
	resp := &mwsHttps.Response{Result: string(GetMatchingProductResultResponse)}
	xmlParser := gmws.NewXMLParser(resp)
	gmpResult := GetMatchingProductResult{}
	xmlParser.Parse(&gmpResult)
	return &gmpResult
}

func Test_GetMatchingProductResult(t *testing.T) {
	Convey("Request response", t, func() {
		gmpResult := prepareGetMatchingProductResult()

		Convey("Has 1 Results", func() {
			So(gmpResult.Results, ShouldHaveLength, 1)
		})

		Convey("ProductResult has 1 product", func() {
			So(gmpResult.Results[0].Products, ShouldHaveLength, 1)
		})
	})
}

func Test_GetMatchingProductResult_Product1(t *testing.T) {
	gmpResult := prepareGetMatchingProductResult()

	Convey("Product 1", t, func() {
		p1 := gmpResult.Results[0].Products[0]

		Convey("Identifiers is not nil", func() {
			So(p1.Identifiers, ShouldNotBeNil)
		})

		Convey("AttributeSets has 1 ItemAttributes", func() {
			So(p1.AttributeSets, ShouldHaveLength, 1)
		})

		Convey("Relationships has 0 Parents", func() {
			So(p1.Relationships.Parents, ShouldHaveLength, 0)
		})

		Convey("Relationships has 5 Children", func() {
			So(p1.Relationships.Children, ShouldHaveLength, 5)
		})

		Convey("SalesRankings has 3 SalesRank", func() {
			So(p1.SalesRankings, ShouldHaveLength, 3)
		})
	})
}

// SKIP Product's Identifiers, AttributeSets, SalesRankings.
// Tested in ListOrderItemsResult_test.go

func Test_GetMatchingProductResult_Product1_Relationships(t *testing.T) {
	gmpResult := prepareGetMatchingProductResult()

	Convey("Product 1 Relationships", t, func() {
		relation := gmpResult.Results[0].Products[0].Relationships

		children := []map[string]string{
			{
				"ASIN": "B002KT3XQC",
				"Size": "Small",
			},
			{
				"ASIN": "B002KT3XQW",
				"Size": "Medium",
			},
			{
				"ASIN": "B002KT3XQM",
				"Size": "Large",
			},
			{
				"ASIN": "B002KT3XR6",
				"Size": "X-Large",
			},
			{
				"ASIN": "B002KT3XRG",
				"Size": "XX-Large",
			},
		}

		for i, child := range children {
			Convey("Children "+strconv.Itoa(i+1), func() {
				child1 := relation.Children[i]

				Convey("Identifiers", func() {
					masin := child1.Identifiers.MarketplaceASIN

					Convey("MarketplaceId is ATVPDKIKX0DER", func() {
						So(masin.MarketplaceId, ShouldEqual, "ATVPDKIKX0DER")
					})

					Convey("ASIN is "+child["ASIN"], func() {
						So(masin.ASIN, ShouldEqual, child["ASIN"])
					})
				})

				Convey("Color is Black", func() {
					So(child1.Color, ShouldEqual, "Black")
				})

				Convey("Size is "+child["Size"], func() {
					So(child1.Size, ShouldEqual, child["Size"])
				})
			})
		}
	})
}
