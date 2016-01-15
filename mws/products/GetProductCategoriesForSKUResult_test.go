package products

import (
	"fmt"
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/svvu/gomws/gmws"
	"github.com/svvu/gomws/mwsHttps"
)

var GetProductCategoriesForSKUResultResponse, _ = ioutil.ReadFile(
	"./examples/GetProductCategoriesForSKU.xml",
)

func prepareGetProductCategoriesForSKUResult() *GetProductCategoriesForSKUResult {
	resp := &mwsHttps.Response{Result: string(GetProductCategoriesForSKUResultResponse)}
	xmlParser := gmws.NewXMLParser(resp)
	gcResult := GetProductCategoriesForSKUResult{}
	err := xmlParser.Parse(&gcResult)
	if err != nil {
		fmt.Println(err)
	}
	return &gcResult
}

func Test_GetProductCategoriesForSKUResult(t *testing.T) {
	Convey("Request response", t, func() {
		gcResult := prepareGetProductCategoriesForSKUResult()

		Convey("Has Result", func() {
			So(gcResult.Result, ShouldNotBeNil)
		})

		Convey("Result has 2 product categoies", func() {
			So(gcResult.Result.ProductCategories, ShouldHaveLength, 2)
		})
	})
}

func Test_GetProductCategoriesForSKUResult_ProductCategory(t *testing.T) {
	gcResult := prepareGetProductCategoriesForSKUResult()

	Convey("ProductCategory 1", t, func() {
		c1 := gcResult.Result.ProductCategories[0]
		productCategoryAsserter(c1, "271578011", "Project Management")

		Convey("Parent", func() {
			p1 := c1.Parent
			productCategoryAsserter(*p1, "2675", "Management & Leadership")

			Convey("Parent", func() {
				p2 := p1.Parent
				productCategoryAsserter(*p2, "3", "Business & Investing")

				Convey("Parent", func() {
					p3 := p2.Parent
					productCategoryAsserter(*p3, "1000", "Subjects")

					Convey("Parent", func() {
						p4 := p3.Parent
						productCategoryAsserter(*p4, "283155", "Subjects")

						Convey("Parent is nil", func() {
							So(p4.Parent, ShouldBeNil)
						})
					})
				})
			})
		})
	})

	Convey("ProductCategory 2", t, func() {
		c2 := gcResult.Result.ProductCategories[1]
		productCategoryAsserter(c2, "684248011", "Management")

		Convey("Parent", func() {
			p1 := c2.Parent
			productCategoryAsserter(*p1, "468220", "Business & Finance")

			Convey("Parent", func() {
				p2 := p1.Parent
				productCategoryAsserter(*p2, "465600", "New, Used & Rental Textbooks")

				Convey("Parent", func() {
					p3 := p2.Parent
					productCategoryAsserter(*p3, "2349030011", "Specialty Boutique")

					Convey("Parent", func() {
						p4 := p3.Parent
						productCategoryAsserter(*p4, "283155", "Specialty Boutique")

						Convey("Parent is nil", func() {
							So(p4.Parent, ShouldBeNil)
						})
					})
				})
			})
		})
	})
}
