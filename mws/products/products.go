//Reference http://docs.developer.amazonservices.com/en_US/products/Products_Overview.html

package products

import (
	"../../gmws"
	"../../mwsHttps"
)

type Products struct {
	*gmws.MwsBase
}

func NewClient(config gmws.MwsConfig) (*Products, error) {
	prodcuts := new(Products)
	base, err := gmws.NewMwsBase(config, prodcuts.Version(), prodcuts.Name())
	if err != nil {
		return nil, err
	}
	prodcuts.MwsBase = base
	return prodcuts, nil
}

func (p Products) Version() string {
	return "2011-10-01"
}

func (p Products) Name() string {
	return "Products"
}

// Returns the operational status of the Products API section.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetServiceStatus.html
func (p Products) GetServiceStatus() (mwsHttps.Result, error) {
	params := mwsHttps.Parameters{
		"Action": "GetServiceStatus",
	}
	structedParams, err := params.Normalize()

	if err != nil {
		return "", err
	}

	httpClient := p.HttpClient(structedParams)
	return httpClient.Request()
}

// Returns a list of products and their attributes, based on a search query.
// http://docs.developer.amazonservices.com/en_US/products/Products_ListMatchingProducts.html
// Optional Parameters:
// 	queryContextId string
func (p Products) ListMatchingProducts(query string, optional ...mwsHttps.Parameters) (mwsHttps.Result, error) {
	op := gmws.OptionalParams([]string{"queryContextId"}, optional)
	params := mwsHttps.Parameters{
		"Action":        "ListMatchingProducts",
		"Query":         query,
		"MarketplaceId": p.MarketPlaceId,
	}.Merge(op)
	structedParams, err := params.Normalize()

	if err != nil {
		return "", err
	}

	httpClient := p.HttpClient(structedParams)
	return httpClient.Request()
}

// Returns a list of products and their attributes, based on a list of ASIN values.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetMatchingProduct.html
func (p Products) GetMatchingProduct(asinList []string) (mwsHttps.Result, error) {
	params := mwsHttps.Parameters{
		"Action":        "GetMatchingProduct",
		"ASINList":      asinList,
		"MarketplaceId": p.MarketPlaceId,
	}
	structedParams, err := params.StructureKeys("ASINList", "ASIN").Normalize()

	if err != nil {
		return "", err
	}

	httpClient := p.HttpClient(structedParams)
	return httpClient.Request()
}

// Returns a list of products and their attributes, based on a list of ASIN, GCID, SellerSKU, UPC, EAN, ISBN, and JAN values.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetMatchingProductForId.html
func (p Products) GetMatchingProductForId(idType string, idList []string) (mwsHttps.Result, error) {
	params := mwsHttps.Parameters{
		"Action":        "GetMatchingProductForId",
		"IdType":        idType,
		"IdList":        idList,
		"MarketplaceId": p.MarketPlaceId,
	}
	structedParams, err := params.StructureKeys("IdList", "Id").Normalize()

	if err != nil {
		return "", err
	}

	httpClient := p.HttpClient(structedParams)
	return httpClient.Request()
}

// Returns the current competitive price of a product, based on SellerSKU.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetCompetitivePricingForSKU.html
func (p Products) GetCompetitivePricingForSKU(sellerSKUList []string) (mwsHttps.Result, error) {
	params := mwsHttps.Parameters{
		"Action":        "GetCompetitivePricingForSKU",
		"SellerSKUList": sellerSKUList,
		"MarketplaceId": p.MarketPlaceId,
	}
	structedParams, err := params.StructureKeys("SellerSKUList", "SellerSKU").Normalize()

	if err != nil {
		return "", err
	}

	httpClient := p.HttpClient(structedParams)
	return httpClient.Request()
}

// Returns the current competitive price of a product, based on ASIN.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetCompetitivePricingForASIN.html
func (p Products) GetCompetitivePricingForASIN(asinList []string) (mwsHttps.Result, error) {
	params := mwsHttps.Parameters{
		"Action":        "GetCompetitivePricingForASIN",
		"ASINList":      asinList,
		"MarketplaceId": p.MarketPlaceId,
	}
	structedParams, err := params.StructureKeys("ASINList", "ASIN").Normalize()

	if err != nil {
		return "", err
	}

	httpClient := p.HttpClient(structedParams)
	return httpClient.Request()
}

// Returns pricing information for the lowest-price active offer listings for up to 20 products, based on SellerSKU.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetLowestOfferListingsForSKU.html
// Optional Parameters:
// 	itemCondition string
// 	excludeMe bool
func (p Products) GetLowestOfferListingsForSKU(sellerSKUList []string, optional ...mwsHttps.Parameters) (mwsHttps.Result, error) {
	op := gmws.OptionalParams([]string{"itemCondition", "excludeMe"}, optional)
	params := mwsHttps.Parameters{
		"Action":        "GetLowestOfferListingsForSKU",
		"SellerSKUList": sellerSKUList,
		"MarketplaceId": p.MarketPlaceId,
	}.Merge(op)
	structedParams, err := params.StructureKeys("SellerSKUList", "SellerSKU").Normalize()

	if err != nil {
		return "", err
	}

	httpClient := p.HttpClient(structedParams)
	return httpClient.Request()
}

// Returns pricing information for the lowest-price active offer listings for up to 20 products, based on ASIN.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetLowestOfferListingsForASIN.html
// Optional Parameters:
// 	itemCondition string
// 	excludeMe bool
func (p Products) GetLowestOfferListingsForASIN(asinList []string, optional ...mwsHttps.Parameters) (mwsHttps.Result, error) {
	op := gmws.OptionalParams([]string{"itemCondition", "excludeMe"}, optional)
	params := mwsHttps.Parameters{
		"Action":        "GetLowestOfferListingsForASIN",
		"ASINList":      asinList,
		"MarketplaceId": p.MarketPlaceId,
	}.Merge(op)
	structedParams, err := params.StructureKeys("ASINList", "ASIN").Normalize()

	if err != nil {
		return "", err
	}

	httpClient := p.HttpClient(structedParams)
	return httpClient.Request()
}

// Returns lowest priced offers for a single product, based on SellerSKU.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetLowestPricedOffersForSKU.html
func (p Products) GetLowestPricedOffersForSKU(sellerSKU, itemCondition string) (mwsHttps.Result, error) {
	params := mwsHttps.Parameters{
		"Action":        "GetLowestPricedOffersForSKU",
		"SellerSKU":     sellerSKU,
		"ItemCondition": itemCondition,
		"MarketplaceId": p.MarketPlaceId,
	}
	structedParams, err := params.Normalize()

	if err != nil {
		return "", err
	}

	httpClient := p.HttpClient(structedParams)
	return httpClient.Request()
}

// Returns lowest priced offers for a single product, based on ASIN.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetLowestPricedOffersForASIN.html
func (p Products) GetLowestPricedOffersForASIN(asin, itemCondition string) (mwsHttps.Result, error) {
	params := mwsHttps.Parameters{
		"Action":        "GetLowestPricedOffersForASIN",
		"ASIN":          asin,
		"ItemCondition": itemCondition,
		"MarketplaceId": p.MarketPlaceId,
	}
	structedParams, err := params.Normalize()

	if err != nil {
		return "", err
	}

	httpClient := p.HttpClient(structedParams)
	return httpClient.Request()
}

// Returns pricing information for your own offer listings, based on SellerSKU.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetMyPriceForSKU.html
func (p Products) GetMyPriceForSKU(sellerSKUList []string, optional ...mwsHttps.Parameters) (mwsHttps.Result, error) {
	op := gmws.OptionalParams([]string{"itemCondition"}, optional)
	params := mwsHttps.Parameters{
		"Action":        "GetMyPriceForSKU",
		"SellerSKUList": sellerSKUList,
		"MarketplaceId": p.MarketPlaceId,
	}.Merge(op)
	structedParams, err := params.StructureKeys("SellerSKUList", "SellerSKU").Normalize()

	if err != nil {
		return "", err
	}

	httpClient := p.HttpClient(structedParams)
	return httpClient.Request()
}

// Returns pricing information for your own offer listings, based on ASIN.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetMyPriceForASIN.html
func (p Products) GetMyPriceForASIN(asinList []string, optional ...mwsHttps.Parameters) (mwsHttps.Result, error) {
	op := gmws.OptionalParams([]string{"itemCondition"}, optional)
	params := mwsHttps.Parameters{
		"Action":        "GetMyPriceForASIN",
		"ASINList":      asinList,
		"MarketplaceId": p.MarketPlaceId,
	}.Merge(op)
	structedParams, err := params.StructureKeys("ASINList", "ASIN").Normalize()

	if err != nil {
		return "", err
	}

	httpClient := p.HttpClient(structedParams)
	return httpClient.Request()
}

// Returns the parent product categories that a product belongs to, based on SellerSKU.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetProductCategoriesForSKU.html
func (p Products) GetProductCategoriesForSKU(sellerSKU string) (mwsHttps.Result, error) {
	params := mwsHttps.Parameters{
		"Action":        "GetProductCategoriesForSKU",
		"SellerSKU":     sellerSKU,
		"MarketplaceId": p.MarketPlaceId,
	}
	structedParams, err := params.Normalize()

	if err != nil {
		return "", err
	}

	httpClient := p.HttpClient(structedParams)
	return httpClient.Request()
}

// Returns the parent product categories that a product belongs to, based on ASIN.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetProductCategoriesForASIN.html
func (p Products) GetProductCategoriesForASIN(asin string) (mwsHttps.Result, error) {
	params := mwsHttps.Parameters{
		"Action":        "GetProductCategoriesForASIN",
		"ASIN":          asin,
		"MarketplaceId": p.MarketPlaceId,
	}
	structedParams, err := params.Normalize()

	if err != nil {
		return "", err
	}

	httpClient := p.HttpClient(structedParams)
	return httpClient.Request()
}
