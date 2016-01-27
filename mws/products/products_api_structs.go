package products

import (
	"encoding/xml"

	"github.com/svvu/gomws/gmws"
)

// References:
// http://g-ecx.images-amazon.com/images/G/01/mwsportal/doc/en_US/products/ProductsAPI_Response.xsd
// http://g-ecx.images-amazon.com/images/G/01/mwsportal/doc/en_US/products/default.xsd

// ListMatchingProductsResult the result for the ListMatchingProducts operation.
type ListMatchingProductsResult struct {
	XMLName xml.Name              `xml:"ListMatchingProductsResponse"`
	Results []MultiProductsResult `xml:"ListMatchingProductsResult"`
}

// GetMatchingProductResult the result for the GetMatchingProduct operation.
type GetMatchingProductResult struct {
	XMLName        xml.Name        `xml:"GetMatchingProductResponse"`
	ProductResults []ProductResult `xml:"GetMatchingProductResult"`
	Results        []MultiProductsResult
}

// ParseCallback is callback trigger after finish parse.
// The callback will move the data from MatchProductResult to Results by converting
// MatchProductResult.Product to an array of Product in Results.
// This method convert the result to make GetMatchingProduct's result is consition
// with GetMatchingProductForIdResult. Because GetMatchingProduct return Product under result
// tag instead of Products tag.
func (mpr *GetMatchingProductResult) ParseCallback() {
	results := make([]MultiProductsResult, len(mpr.ProductResults))
	for i, r := range mpr.ProductResults {
		result := &results[i]
		result.Products = []Product{r.Product}
		result.Status = r.Status
		result.Error = r.Error
	}
	mpr.Results = results
	mpr.ProductResults = nil
}

// GetMatchingProductForIdResult the result for the GetMatchingProductForID operation.
type GetMatchingProductForIdResult struct {
	XMLName xml.Name              `xml:"GetMatchingProductForIdResponse"`
	Results []MultiProductsResult `xml:"GetMatchingProductForIdResult"`
}

// GetCompetitivePricingForSKUResult the result for the GetCompetitivePricingForSKU operation.
type GetCompetitivePricingForSKUResult struct {
	XMLName xml.Name        `xml:"GetCompetitivePricingForSKUResponse"`
	Results []ProductResult `xml:"GetCompetitivePricingForSKUResult"`
}

// GetCompetitivePricingForASINResult the result for the GetCompetitivePricingForASIN operation.
type GetCompetitivePricingForASINResult struct {
	XMLName xml.Name        `xml:"GetCompetitivePricingForASINResponse"`
	Results []ProductResult `xml:"GetCompetitivePricingForASINResult"`
}

// GetLowestOfferListingsForSKUResult the result for the GetLowestOfferListingsForSKU operation.
type GetLowestOfferListingsForSKUResult struct {
	XMLName xml.Name                          `xml:"GetLowestOfferListingsForSKUResponse"`
	Results []LowestOfferListingProductResult `xml:"GetLowestOfferListingsForSKUResult"`
}

// GetLowestOfferListingsForASINResult the result for the GetLowestOfferListingsForASIN operation.
type GetLowestOfferListingsForASINResult struct {
	XMLName xml.Name                          `xml:"GetLowestOfferListingsForASINResponse"`
	Results []LowestOfferListingProductResult `xml:"GetLowestOfferListingsForASINResult"`
}

// GetLowestPricedOffersForSKUResult the result for the GetLowestPricedOffersForSKU operation.
type GetLowestPricedOffersForSKUResult struct {
	XMLName xml.Name                        `xml:"GetLowestPricedOffersForSKUResponse"`
	Result  LowestPricedOffersProductResult `xml:"GetLowestPricedOffersForSKUResult"`
}

// GetLowestPricedOffersForASINResult the result for the GetLowestPricedOffersForASIN operation.
type GetLowestPricedOffersForASINResult struct {
	XMLName xml.Name                        `xml:"GetLowestPricedOffersForASINResponse"`
	Result  LowestPricedOffersProductResult `xml:"GetLowestPricedOffersForASINResult"`
}

// GetMyPriceForSKUResult the result for the GetMyPriceForSKU operation.
type GetMyPriceForSKUResult struct {
	XMLName xml.Name        `xml:"GetMyPriceForSKUResponse"`
	Results []ProductResult `xml:"GetMyPriceForSKUResult"`
}

// GetMyPriceForASINResult the result for the GetMyPriceForASIN operation.
type GetMyPriceForASINResult struct {
	XMLName xml.Name        `xml:"GetMyPriceForASINResponse"`
	Results []ProductResult `xml:"GetMyPriceForASINResult"`
}

// GetProductCategoriesForSKUResult the result for the GetProductCategoriesForSKU operation.
type GetProductCategoriesForSKUResult struct {
	XMLName xml.Name                `xml:"GetProductCategoriesForSKUResponse"`
	Result  ProductCategoriesResult `xml:"GetProductCategoriesForSKUResult"`
}

// GetProductCategoriesForASINResult the result for the GetProductCategoriesForASIN operation.
type GetProductCategoriesForASINResult struct {
	XMLName xml.Name                `xml:"GetProductCategoriesForASINResponse"`
	Result  ProductCategoriesResult `xml:"GetProductCategoriesForASINResult"`
}

// MultiProductsResult the result from the operation, contains meta info for the result.
// MultiProductsResult contains one of more products.
type MultiProductsResult struct {
	Products []Product   `xml:">Product"`
	ID       string      `xml:"Id,attr"`
	IDType   string      `xml:"IdType,attr"`
	Status   string      `xml:"status,attr"`
	Error    *gmws.Error `xml:"Error"`
}

// ProductResult the result from the operation, contains meta info for the result.
// ProductResult contains only one product.
type ProductResult struct {
	Product Product     `xml:"Product"`
	Status  string      `xml:"status,attr"`
	Error   *gmws.Error `xml:"Error"`
}

// LowestOfferListingProductResult the result from the operation.
// Simliar to ProductResult, but with extra fields
type LowestOfferListingProductResult struct {
	AllOfferListingsConsidered bool `xml:"AllOfferListingsConsidered"`
	ProductResult
	Status string      `xml:"status,attr"`
	Error  *gmws.Error `xml:"Error"`
}

// LowestPricedOffersProductResult contains Identifier, Offer Summary, Offers list for the product.
type LowestPricedOffersProductResult struct {
	Identifier OfferIdentifier
	Summary    OffersSummary
	Offers     []Offer     `xml:">Offer"`
	Status     string      `xml:"status,attr"`
	Error      *gmws.Error `xml:"Error"`
}

// ProductCategoriesResult a list ProductCategory.
type ProductCategoriesResult struct {
	ProductCategories []ProductCategory `xml:"Self"`
	Error             *gmws.Error       `xml:"Error"`
}
