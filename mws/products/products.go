package products

// Reference http://docs.developer.amazonservices.com/en_US/products/Products_Overview.html

import (
	"github.com/svvu/gomws/mws"
)

// Products is the client for the api
type Products struct {
	*mws.Client
}

// NewClient generate a new product client
func NewClient(config mws.Config) (*Products, error) {
	prodcuts := new(Products)
	base, err := mws.NewClient(config, prodcuts.Version(), prodcuts.Name())
	if err != nil {
		return nil, err
	}
	prodcuts.Client = base
	return prodcuts, nil
}

// Version return the current version of api
func (p Products) Version() string {
	return "2011-10-01"
}

// Name return the name of the api
func (p Products) Name() string {
	return "Products"
}

// GetServiceStatus Returns the operational status of the Products API section.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetServiceStatus.html
func (p Products) GetServiceStatus() (*mws.Response, error) {
	params := mws.Parameters{
		"Action": "GetServiceStatus",
	}
	return p.SendRequest(params)
}

// ListMatchingProducts Returns a list of products and their attributes, based on a search query.
// Optional Parameters:
// 	QueryContextId - string
// http://docs.developer.amazonservices.com/en_US/products/Products_ListMatchingProducts.html
func (p Products) ListMatchingProducts(query string, optional ...mws.Parameters) (*mws.Response, error) {
	op := mws.OptionalParams([]string{"QueryContextId"}, optional)
	params := mws.Parameters{
		"Action":        "ListMatchingProducts",
		"Query":         query,
		"MarketplaceId": p.MarketPlaceId,
	}.Merge(op)

	return p.SendRequest(params)
}

// GetMatchingProduct Returns a list of products and their attributes, based on a list of ASIN values.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetMatchingProduct.html
func (p Products) GetMatchingProduct(asinList []string) (*mws.Response, error) {
	params := mws.Parameters{
		"Action":        "GetMatchingProduct",
		"ASINList":      asinList,
		"MarketplaceId": p.MarketPlaceId,
	}
	structuredParams := params.StructureKeys("ASINList", "ASIN")

	return p.SendRequest(structuredParams)
}

// GetMatchingProductForId Returns a list of products and their attributes, based on a list of ASIN, GCID, SellerSKU, UPC, EAN, ISBN, and JAN values.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetMatchingProductForId.html
func (p Products) GetMatchingProductForId(idType string, idList []string) (*mws.Response, error) {
	params := mws.Parameters{
		"Action":        "GetMatchingProductForId",
		"IdType":        idType,
		"IdList":        idList,
		"MarketplaceId": p.MarketPlaceId,
	}
	structuredParams := params.StructureKeys("IdList", "Id")

	return p.SendRequest(structuredParams)
}

// GetCompetitivePricingForSKU Returns the current competitive price of a product, based on SellerSKU.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetCompetitivePricingForSKU.html
func (p Products) GetCompetitivePricingForSKU(sellerSKUList []string) (*mws.Response, error) {
	params := mws.Parameters{
		"Action":        "GetCompetitivePricingForSKU",
		"SellerSKUList": sellerSKUList,
		"MarketplaceId": p.MarketPlaceId,
	}
	structuredParams := params.StructureKeys("SellerSKUList", "SellerSKU")

	return p.SendRequest(structuredParams)
}

// GetCompetitivePricingForASIN Returns the current competitive price of a product, based on ASIN.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetCompetitivePricingForASIN.html
func (p Products) GetCompetitivePricingForASIN(asinList []string) (*mws.Response, error) {
	params := mws.Parameters{
		"Action":        "GetCompetitivePricingForASIN",
		"ASINList":      asinList,
		"MarketplaceId": p.MarketPlaceId,
	}
	structuredParams := params.StructureKeys("ASINList", "ASIN")

	return p.SendRequest(structuredParams)
}

// GetLowestOfferListingsForSKU Returns pricing information for the lowest-price active offer listings for up to 20 products, based on SellerSKU.
// Optional Parameters:
// 	ItemCondition - string
// 	ExcludeMe - bool
// http://docs.developer.amazonservices.com/en_US/products/Products_GetLowestOfferListingsForSKU.html
func (p Products) GetLowestOfferListingsForSKU(sellerSKUList []string, optional ...mws.Parameters) (*mws.Response, error) {
	op := mws.OptionalParams([]string{"ItemCondition", "ExcludeMe"}, optional)
	params := mws.Parameters{
		"Action":        "GetLowestOfferListingsForSKU",
		"SellerSKUList": sellerSKUList,
		"MarketplaceId": p.MarketPlaceId,
	}.Merge(op)
	structuredParams := params.StructureKeys("SellerSKUList", "SellerSKU")

	return p.SendRequest(structuredParams)
}

// GetLowestOfferListingsForASIN Returns pricing information for the lowest-price active offer listings for up to 20 products, based on ASIN.
// Optional Parameters:
// 	ItemCondition - string
// 	ExcludeMe - bool
// http://docs.developer.amazonservices.com/en_US/products/Products_GetLowestOfferListingsForASIN.html
func (p Products) GetLowestOfferListingsForASIN(asinList []string, optional ...mws.Parameters) (*mws.Response, error) {
	op := mws.OptionalParams([]string{"ItemCondition", "ExcludeMe"}, optional)
	params := mws.Parameters{
		"Action":        "GetLowestOfferListingsForASIN",
		"ASINList":      asinList,
		"MarketplaceId": p.MarketPlaceId,
	}.Merge(op)
	structuredParams := params.StructureKeys("ASINList", "ASIN")

	return p.SendRequest(structuredParams)
}

// GetLowestPricedOffersForSKU Returns lowest priced offers for a single product, based on SellerSKU.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetLowestPricedOffersForSKU.html
func (p Products) GetLowestPricedOffersForSKU(sellerSKU, itemCondition string) (*mws.Response, error) {
	params := mws.Parameters{
		"Action":        "GetLowestPricedOffersForSKU",
		"SellerSKU":     sellerSKU,
		"ItemCondition": itemCondition,
		"MarketplaceId": p.MarketPlaceId,
	}

	return p.SendRequest(params)
}

// GetLowestPricedOffersForASIN Returns lowest priced offers for a single product, based on ASIN.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetLowestPricedOffersForASIN.html
func (p Products) GetLowestPricedOffersForASIN(asin, itemCondition string) (*mws.Response, error) {
	params := mws.Parameters{
		"Action":        "GetLowestPricedOffersForASIN",
		"ASIN":          asin,
		"ItemCondition": itemCondition,
		"MarketplaceId": p.MarketPlaceId,
	}

	return p.SendRequest(params)
}

// GetMyPriceForSKU Returns pricing information for your own offer listings, based on SellerSKU.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetMyPriceForSKU.html
func (p Products) GetMyPriceForSKU(sellerSKUList []string, optional ...mws.Parameters) (*mws.Response, error) {
	op := mws.OptionalParams([]string{"ItemCondition"}, optional)
	params := mws.Parameters{
		"Action":        "GetMyPriceForSKU",
		"SellerSKUList": sellerSKUList,
		"MarketplaceId": p.MarketPlaceId,
	}.Merge(op)
	structuredParams := params.StructureKeys("SellerSKUList", "SellerSKU")

	return p.SendRequest(structuredParams)
}

// GetMyPriceForASIN Returns pricing information for your own offer listings, based on ASIN.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetMyPriceForASIN.html
func (p Products) GetMyPriceForASIN(asinList []string, optional ...mws.Parameters) (*mws.Response, error) {
	op := mws.OptionalParams([]string{"ItemCondition"}, optional)
	params := mws.Parameters{
		"Action":        "GetMyPriceForASIN",
		"ASINList":      asinList,
		"MarketplaceId": p.MarketPlaceId,
	}.Merge(op)
	structuredParams := params.StructureKeys("ASINList", "ASIN")

	return p.SendRequest(structuredParams)
}

// GetProductCategoriesForSKU Returns the parent product categories that a product belongs to, based on SellerSKU.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetProductCategoriesForSKU.html
func (p Products) GetProductCategoriesForSKU(sellerSKU string) (*mws.Response, error) {
	params := mws.Parameters{
		"Action":        "GetProductCategoriesForSKU",
		"SellerSKU":     sellerSKU,
		"MarketplaceId": p.MarketPlaceId,
	}

	return p.SendRequest(params)
}

// GetProductCategoriesForASIN Returns the parent product categories that a product belongs to, based on ASIN.
// http://docs.developer.amazonservices.com/en_US/products/Products_GetProductCategoriesForASIN.html
func (p Products) GetProductCategoriesForASIN(asin string) (*mws.Response, error) {
	params := mws.Parameters{
		"Action":        "GetProductCategoriesForASIN",
		"ASIN":          asin,
		"MarketplaceId": p.MarketPlaceId,
	}

	return p.SendRequest(params)
}
