package products

import (
	"fmt"
	"strconv"

	"github.com/smartystreets/goconvey/convey"
	"github.com/svvu/gomws/gmws"
)

func loadExample(name string) interface{} {
	var (
		result interface{}
		err    error
	)

	switch name {
	case "ListMatchingProducts":
		result, err = gmws.LoadExample("./examples/"+name+".xml", &ListMatchingProductsResult{})
	case "GetMatchingProduct":
		result, err = gmws.LoadExample("./examples/"+name+".xml", &GetMatchingProductResult{})
	case "GetMatchingProductForId":
		result, err = gmws.LoadExample("./examples/"+name+".xml", &GetMatchingProductForIdResult{})
	case "GetCompetitivePricingForSKU":
		result, err = gmws.LoadExample("./examples/"+name+".xml", &GetCompetitivePricingForSKUResult{})
	case "GetCompetitivePricingForASIN_ClientError":
		result, err = gmws.LoadExample("./examples/"+name+".xml", &GetCompetitivePricingForASINResult{})
	case "GetLowestOfferListingsForSKU":
		result, err = gmws.LoadExample("./examples/"+name+".xml", &GetLowestOfferListingsForSKUResult{})
	case "GetLowestPricedOffersForSKU", "GetLowestPricedOffersForSKU_NoOffers", "GetLowestPricedOffersForSKU_ServerError":
		result, err = gmws.LoadExample("./examples/"+name+".xml", &GetLowestPricedOffersForSKUResult{})
	case "GetMyPriceForSKU":
		result, err = gmws.LoadExample("./examples/"+name+".xml", &GetMyPriceForSKUResult{})
	case "GetProductCategoriesForSKU":
		result, err = gmws.LoadExample("./examples/"+name+".xml", &GetProductCategoriesForSKUResult{})
	}

	if err != nil {
		fmt.Println(err)
	}

	return result
}

func decimalWithUnitsAsserter(dwn DecimalWithUnits, unit string, value float64) {
	convey.Convey("Units is "+unit, func() {
		convey.So(dwn.Units, convey.ShouldEqual, unit)
	})

	convey.Convey("Value is "+strconv.FormatFloat(value, 'f', 2, 64), func() {
		convey.So(dwn.Value, convey.ShouldEqual, value)
	})
}

func dimensionsAsserter(dim DimensionType, lengthUnit, weightUnit string, expectValue map[string]float64) {
	convey.Convey("Height", func() {
		decimalWithUnitsAsserter(dim.Height, lengthUnit, expectValue["Height"])
	})

	convey.Convey("Length", func() {
		decimalWithUnitsAsserter(dim.Length, lengthUnit, expectValue["Length"])
	})

	convey.Convey("Width", func() {
		decimalWithUnitsAsserter(dim.Width, lengthUnit, expectValue["Width"])
	})

	convey.Convey("Weight", func() {
		decimalWithUnitsAsserter(dim.Weight, weightUnit, expectValue["Weight"])
	})
}

func moneyAsserter(money Money, currencyCode string, amount float64) {
	convey.Convey("CurrencyCode is "+currencyCode, func() {
		convey.So(money.CurrencyCode, convey.ShouldEqual, currencyCode)
	})

	convey.Convey("Amount is "+strconv.FormatFloat(amount, 'f', 2, 64), func() {
		convey.So(money.Amount, convey.ShouldEqual, amount)
	})
}

func productCategoryAsserter(pc ProductCategory, id, name string) {
	convey.Convey("ProductCategoryId is "+id, func() {
		convey.So(pc.Id, convey.ShouldEqual, id)
	})

	convey.Convey("ProductCategoryName is "+name, func() {
		convey.So(pc.Name, convey.ShouldEqual, name)
	})
}
