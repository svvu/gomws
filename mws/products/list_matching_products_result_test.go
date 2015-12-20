package products

import (
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/svvu/gomws/gmws"
	"github.com/svvu/gomws/mwsHttps"
)

var (
	ListMatchingProductsResponse, _ = ioutil.ReadFile("./examples/ListMatchingProducts.xml")
)

func prepareListMatchingProductsResult() *ListMatchingProductsResult {
	resp := &mwsHttps.Response{Result: string(ListMatchingProductsResponse)}
	xmlParser := gmws.NewXMLParser(resp)
	lmpResult := ListMatchingProductsResult{}
	xmlParser.Parse(&lmpResult)
	return &lmpResult
}

func Test_ListMatchingProductsResult(t *testing.T) {
	Convey("Request response", t, func() {
		lmpResult := prepareListMatchingProductsResult()

		Convey("Has 1 Product", func() {
			So(lmpResult.Products, ShouldHaveLength, 1)
		})
	})
}

func Test_ListMatchingProductsResult_Product1(t *testing.T) {
	lmpResult := prepareListMatchingProductsResult()

	Convey("Product 1's", t, func() {
		p1 := lmpResult.Products[0]

		Convey("Identifiers is not nil", func() {
			So(p1.Identifiers, ShouldNotBeNil)
		})

		Convey("AttributeSets has 1 ItemAttributes", func() {
			So(p1.AttributeSets, ShouldHaveLength, 1)
		})

		Convey("SalesRankings has 4 SalesRank", func() {
			So(p1.SalesRankings, ShouldHaveLength, 4)
		})
	})
}

func Test_ListMatchingProductsResult_Product1_Identifiers(t *testing.T) {
	lmpResult := prepareListMatchingProductsResult()

	Convey("Product 1 Identifiers'", t, func() {
		iden := lmpResult.Products[0].Identifiers

		Convey("MarketplaceASIN is not nil", func() {
			So(iden.MarketplaceASIN, ShouldNotBeNil)
		})

		Convey("MarketplaceASIN", func() {
			masin := iden.MarketplaceASIN

			Convey("MarketplaceId is ATVPDKIKX0DER", func() {
				So(masin.MarketplaceId, ShouldEqual, "ATVPDKIKX0DER")
			})

			Convey("ASIN is 059035342X", func() {
				So(masin.ASIN, ShouldEqual, "059035342X")
			})
		})
	})
}

func Test_ListMatchingProductsResult_Product1_ItemAttributes1(t *testing.T) {
	lmpResult := prepareListMatchingProductsResult()

	Convey("Product 1 ItemAttributes 1", t, func() {
		attr := lmpResult.Products[0].AttributeSets[0]

		Convey("Lang is en-US", func() {
			So(attr.Lang, ShouldEqual, "en-US")
		})

		Convey("Author", func() {
			author := attr.Author

			Convey("Has length 1", func() {
				So(author, ShouldHaveLength, 1)
			})

			Convey("Author 1 is Rowling, J.K.", func() {
				So(author[0], ShouldEqual, "Rowling, J.K.")
			})
		})

		Convey("Binding is Paperback", func() {
			So(attr.Binding, ShouldEqual, "Paperback")
		})

		Convey("Brand is Scholastic Press", func() {
			So(attr.Brand, ShouldEqual, "Scholastic Press")
		})

		Convey("Creator", func() {
			creator := attr.Creator[0]

			Convey("Role is Illustrator", func() {
				So(creator.Role, ShouldEqual, "Illustrator")
			})

			Convey("Value is GrandPré, Mary", func() {
				So(creator.Value, ShouldEqual, "GrandPré, Mary")
			})
		})

		Convey("Edition is 1st", func() {
			So(attr.Edition, ShouldEqual, "1st")
		})

		Convey("Feature", func() {
			fea := attr.Feature

			Convey("Has length 1", func() {
				So(fea, ShouldHaveLength, 1)
			})

			Convey("Feature 1 is Recommended Age: 9 years and up", func() {
				So(fea[0], ShouldEqual, "Recommended Age: 9 years and up")
			})
		})

		Convey("ItemDimensions", func() {
			idim := attr.ItemDimensions

			Convey("Height", func() {
				height := idim.Height

				Convey("Units is inches", func() {
					So(height.Units, ShouldEqual, "inches")
				})

				Convey("Value is 0.8", func() {
					So(height.Value, ShouldEqual, 0.8)
				})
			})

			Convey("Length", func() {
				height := idim.Length

				Convey("Units is inches", func() {
					So(height.Units, ShouldEqual, "inches")
				})

				Convey("Value is 7.5", func() {
					So(height.Value, ShouldEqual, 7.5)
				})
			})

			Convey("Width", func() {
				height := idim.Width

				Convey("Units is inches", func() {
					So(height.Units, ShouldEqual, "inches")
				})

				Convey("Value is 5.2", func() {
					So(height.Value, ShouldEqual, 5.20)
				})
			})

			Convey("Weight", func() {
				height := idim.Weight

				Convey("Units is pounds", func() {
					So(height.Units, ShouldEqual, "pounds")
				})

				Convey("Value is 0.5", func() {
					So(height.Value, ShouldEqual, 0.5)
				})
			})
		})

		Convey("IsAutographed is false", func() {
			So(attr.IsAutographed, ShouldBeFalse)
		})

		Convey("IsMemorabilia is false", func() {
			So(attr.IsMemorabilia, ShouldBeFalse)
		})

		Convey("Label is Scholastic Paperbacks", func() {
			So(attr.Label, ShouldEqual, "Scholastic Paperbacks")
		})

		Convey("Languages", func() {
			languages := attr.Languages

			Convey("Has length 3", func() {
				So(languages, ShouldHaveLength, 3)
			})

			Convey("Language 1", func() {
				language1 := languages[0]

				Convey("Name is english", func() {
					So(language1.Name, ShouldEqual, "english")
				})

				Convey("Type is Unknown", func() {
					So(language1.Type, ShouldEqual, "Unknown")
				})
			})

			Convey("Language 2", func() {
				language1 := languages[1]

				Convey("Name is english", func() {
					So(language1.Name, ShouldEqual, "english")
				})

				Convey("Type is Original Language", func() {
					So(language1.Type, ShouldEqual, "Original Language")
				})
			})

			Convey("Language 3", func() {
				language1 := languages[2]

				Convey("Name is english", func() {
					So(language1.Name, ShouldEqual, "english")
				})

				Convey("Type is Published", func() {
					So(language1.Type, ShouldEqual, "Published")
				})
			})
		})

		Convey("ListPrice", func() {
			lp := attr.ListPrice

			Convey("Amount is 10.99", func() {
				So(lp.Amount, ShouldEqual, 10.99)
			})

			Convey("CurrencyCode is USD", func() {
				So(lp.CurrencyCode, ShouldEqual, "USD")
			})
		})

		Convey("Manufacturer is Scholastic Paperbacks", func() {
			So(attr.Manufacturer, ShouldEqual, "Scholastic Paperbacks")
		})

		Convey("NumberOfItems is 1", func() {
			So(attr.NumberOfItems, ShouldEqual, 1)
		})

		Convey("NumberOfPages is 320", func() {
			So(attr.NumberOfPages, ShouldEqual, 320)
		})

		Convey("PackageDimensions", func() {
			idim := attr.PackageDimensions

			Convey("Height", func() {
				height := idim.Height

				Convey("Units is inches", func() {
					So(height.Units, ShouldEqual, "inches")
				})

				Convey("Value is 1", func() {
					So(height.Value, ShouldEqual, 1)
				})
			})

			Convey("Length", func() {
				height := idim.Length

				Convey("Units is inches", func() {
					So(height.Units, ShouldEqual, "inches")
				})

				Convey("Value is 7.5", func() {
					So(height.Value, ShouldEqual, 7.5)
				})
			})

			Convey("Width", func() {
				height := idim.Width

				Convey("Units is inches", func() {
					So(height.Units, ShouldEqual, "inches")
				})

				Convey("Value is 5.2", func() {
					So(height.Value, ShouldEqual, 5.20)
				})
			})

			Convey("Weight", func() {
				height := idim.Weight

				Convey("Units is pounds", func() {
					So(height.Units, ShouldEqual, "pounds")
				})

				Convey("Value is 0.5", func() {
					So(height.Value, ShouldEqual, 0.5)
				})
			})
		})

		Convey("PackageQuantity is 1", func() {
			So(attr.PackageQuantity, ShouldEqual, 1)
		})

		Convey("PartNumber is 9780590353427", func() {
			So(attr.PartNumber, ShouldEqual, "9780590353427")
		})

		Convey("ProductGroup is Book", func() {
			So(attr.ProductGroup, ShouldEqual, "Book")
		})

		Convey("ProductTypeName is ABIS_BOOK", func() {
			So(attr.ProductTypeName, ShouldEqual, "ABIS_BOOK")
		})

		Convey("PublicationDate is 1999-10-01", func() {
			So(attr.PublicationDate, ShouldEqual, "1999-10-01")
		})

		Convey("Publisher is Scholastic Paperbacks", func() {
			So(attr.Publisher, ShouldEqual, "Scholastic Paperbacks")
		})

		Convey("ReleaseDate is 1999-09-08", func() {
			So(attr.ReleaseDate, ShouldEqual, "1999-09-08")
		})

		Convey("SmallImage", func() {
			img := attr.SmallImage

			Convey("Url has expect value", func() {
				So(img.URL, ShouldEqual, "http://ecx.images-amazon.com/images/I/51MU5VilKpL._SL75_.jpg")
			})

			Convey("Height", func() {
				height := img.Height

				Convey("Units is pixels", func() {
					So(height.Units, ShouldEqual, "pixels")
				})

				Convey("Value is 75", func() {
					So(height.Value, ShouldEqual, 75)
				})
			})

			Convey("Width", func() {
				height := img.Width

				Convey("Units is pixels", func() {
					So(height.Units, ShouldEqual, "pixels")
				})

				Convey("Value is 51", func() {
					So(height.Value, ShouldEqual, 51)
				})
			})
		})

		Convey("Studio is Scholastic Paperbacks", func() {
			So(attr.Studio, ShouldEqual, "Scholastic Paperbacks")
		})

		Convey("Title is Harry Potter and the Sorcerer's Stone (Book 1)", func() {
			So(attr.Title, ShouldEqual, "Harry Potter and the Sorcerer's Stone (Book 1)")
		})
	})
}

func Test_ListMatchingProductsResult_Product1_SalesRankings(t *testing.T) {
	lmpResult := prepareListMatchingProductsResult()

	Convey("Product 1 SalesRankings'", t, func() {
		sr := lmpResult.Products[0].SalesRankings

		Convey("SalesRank 1'", func() {
			sr1 := sr[0]

			Convey("ProductCategoryId is book_display_on_website", func() {
				So(sr1.ProductCategoryId, ShouldEqual, "book_display_on_website")
			})

			Convey("Rank is 401", func() {
				So(sr1.Rank, ShouldEqual, 401)
			})
		})

		Convey("SalesRank 2'", func() {
			sr2 := sr[1]

			Convey("ProductCategoryId is 15356791", func() {
				So(sr2.ProductCategoryId, ShouldEqual, "15356791")
			})

			Convey("Rank is 5", func() {
				So(sr2.Rank, ShouldEqual, 5)
			})
		})

		Convey("SalesRank 3'", func() {
			sr3 := sr[2]

			Convey("ProductCategoryId is 3153", func() {
				So(sr3.ProductCategoryId, ShouldEqual, "3153")
			})

			Convey("Rank is 8", func() {
				So(sr3.Rank, ShouldEqual, 8)
			})
		})

		Convey("SalesRank 4'", func() {
			sr4 := sr[3]

			Convey("ProductCategoryId is 17468", func() {
				So(sr4.ProductCategoryId, ShouldEqual, "17468")
			})

			Convey("Rank is 16", func() {
				So(sr4.Rank, ShouldEqual, 16)
			})
		})
	})
}
