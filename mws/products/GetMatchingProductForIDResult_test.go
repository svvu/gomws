package products

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/svvu/gomws/gmws"
	"github.com/svvu/gomws/mwsHttps"
)

var GetMatchingProductForIdResultResponse, _ = ioutil.ReadFile(
	"./examples/GetMatchingProductForId.xml",
)

func prepareGetMatchingProductForIdResult() *GetMatchingProductForIdResult {
	resp := &mwsHttps.Response{Result: string(GetMatchingProductForIdResultResponse)}
	xmlParser := gmws.NewXMLParser(resp)
	gmpResult := GetMatchingProductForIdResult{}
	err := xmlParser.Parse(&gmpResult)
	if err != nil {
		fmt.Println(err)
	}
	return &gmpResult
}

func Test_GetMatchingProductForIDResult(t *testing.T) {
	Convey("Request response", t, func() {
		gmpResult := prepareGetMatchingProductForIdResult()

		Convey("Has 2 Results", func() {
			So(gmpResult.Results, ShouldHaveLength, 2)
		})
	})
}

func Test_GetMatchingProductForIDResult_Result1(t *testing.T) {
	gmpResult := prepareGetMatchingProductForIdResult()
	results := []map[string]string{
		{
			"products": "1",
			"ID":       "9781933988665",
			"IDType":   "ISBN",
			"Status":   "Success",
		},
		{
			"products": "2",
			"ID":       "0439708184",
			"IDType":   "ISBN",
			"Status":   "Success",
		},
	}

	for i, expectResult := range results {
		Convey("Result "+strconv.Itoa(i+1), t, func() {
			result := gmpResult.Results[i]

			Convey("Has "+expectResult["products"]+" product", func() {
				products, _ := strconv.Atoi(expectResult["products"])
				So(result.Products, ShouldHaveLength, products)
			})

			Convey("ID is "+expectResult["ID"], func() {
				So(result.ID, ShouldEqual, expectResult["ID"])
			})

			Convey("IDType is", func() {
				So(result.IDType, ShouldEqual, expectResult["IDType"])
			})

			Convey("Status is ", func() {
				So(result.Status, ShouldEqual, expectResult["Status"])
			})
		})
	}
}

// SKIP Product's Identifiers, AttributeSets, SalesRankings.
// Tested in ListOrderItemsResult_test.go
