package products

import "encoding/xml"

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

// MultiProductsResult the result from the operation, contains meta info for the result.
// MultiProductsResult contains one of more products.
type MultiProductsResult struct {
	Products []Product `xml:">Product"`
	ID       string    `xml:"Id,attr"`
	IDType   string    `xml:"IdType,attr"`
	Status   string    `xml:"status,attr"`
}

// ProductResult the result from the operation, contains meta info for the result.
// ProductResult contains only one product.
type ProductResult struct {
	Product Product `xml:"Product"`
	Status  string  `xml:"status,attr"`
}

// LowestOfferListingProductResult the result from the operation.
// Simliar to ProductResult, but with extra fields
type LowestOfferListingProductResult struct {
	AllOfferListingsConsidered bool `xml:"AllOfferListingsConsidered"`
	ProductResult
	Status string `xml:"status,attr"`
}

type LowestPricedOffersProductResult struct {
	Identifier OfferIdentifier
	Summary    OffersSummary
	Offers     []Offer `xml:">Offer"`
	Status     string  `xml:"status,attr"`
}

// Product contains basic, relationship to other products, pricing, and offers info.
type Product struct {
	Identifiers         Identifiers
	AttributeSets       []ItemAttributes `xml:">ItemAttributes"`
	Relationships       Relationships
	CompetitivePricing  CompetitivePricing
	SalesRankings       []SalesRank          `xml:">SalesRank"`
	LowestOfferListings []LowestOfferListing `xml:">LowestOfferListing"`
	Offers              []SellerOffer        `xml:">Offer"`
}

// OffersSummary contains price information about the product.
// Including the LowestPrices and BuyBoxPrices, the ListPrice,
// 	the SuggestedLowerPricePlusShipping, and NumberOfOffers and
// 	NumberOfBuyBoxEligibleOffers.
type OffersSummary struct {
	// The number of unique offers contained in NumberOfOffers.
	TotalOfferCount int
	// A list that contains the total number of offers for the item for the given
	// conditions and fulfillment channels.
	NumberOfOffers []OfferCount `xml:">OfferCount"`
	// A list of the lowest prices for the item.
	LowestPrices []OfferPrice `xml:">LowestPrice"`
	// A list of item prices.
	BuyBoxPrices []OfferPrice `xml:">BuyBoxPrice"`
	// The list price of the item as suggested by the manufacturer.
	ListPrice Money
	// The suggested lower price of the item, including shipping.
	// The suggested lower price is based on a range of factors,
	// including historical selling prices, recent Buy Box-eligible prices,
	// and input from customers for your products.
	SuggestedLowerPricePlusShipping Money
	// A list that contains the total number of offers that are eligible for
	// the Buy Box for the given conditions and fulfillment channels.
	BuyBoxEligibleOffers []OfferCount `xml:">OfferCount"`
}

// Relationships contains product variation information, if applicable.
// If your search criteria match a product that is identified by a variation
// parent ASIN, the related VariationChild elements are contained in the Relationships
// element.
// If your search criteria match a specific variation child ASIN, the related
// VariationParent element is contained in the Relationships element.
type Relationships struct {
	Parents  []VariationParent `xml:"VariationParent"`
	Children []VariationChild  `xml:"VariationChild"`
}

// CompetitivePricing Contains pricing information for the product
type CompetitivePricing struct {
	CompetitivePrices     []CompetitivePrice `xml:">CompetitivePrice"`
	NumberOfOfferListings []OfferCount       `xml:">OfferListingCount"`
	TradeInValue          Money
}

// SalesRank information for the product by product category.
type SalesRank struct {
	ProductCategoryId string
	Rank              int
}

// LowestOfferListing Contains pricing information for the lowest offer listing for each offer listing group.
type LowestOfferListing struct {
	// Qualifiers Contains the six qualifiers:
	// 	ItemCondition, ItemSubcondition, FulfillmentChannel, ShipsDomestically,
	// 	ShippingTime, and SellerPositiveFeedbackRating.
	// These qualifiers identify the offer listing group from which the lowest offer listing was taken.
	Qualifiers Qualifiers
	// Of the offer listings considered, this number indicates how many belonged to this offer listing group.
	NumberOfOfferListingsConsidered int
	// The number of seller feedback ratings that have been submitted for the seller with the lowest-priced offer listing in the group.
	SellerFeedbackCount int
	// Pricing information for the lowest offer listing in the group.
	Price Price
	// Indicates if there is more than one listing with the lowest listing price in the group.
	MultipleOffersAtLowestPrice string
}

type SellerOffer struct {
	BuyingPrice        Price
	RegularPrice       Money
	FulfillmentChannel string
	ItemCondition      string
	ItemSubCondition   string
	SellerId           string
	SellerSKU          string
}

type Offer struct {
	// true if this is your offer.
	MyOffer bool
	// The subcondition of the item.
	// Values: New, Mint, Very Good, Good, Acceptable, Poor, Club, OEM, Warranty,
	// 	Refurbished Warranty, Refurbished, Open Box, or Other.
	SubCondition string
	// Information about the seller's feedback, including the percentage of
	// positive feedback, and the total count of feedback received.
	SellerFeedbackRating SellerFeedbackRating
	// The maximum time within which the item will likely be shipped once an order has been placed.
	ShippingTime ShippingTime
	// The price of the item.
	ListingPrice Money
	// The number of Amazon Points offered with the purchase of an item.
	Points int
	// The shipping cost.
	Shipping Money
	// The state and country from where the item is shipped.
	ShipsFrom Address
	// true if the offer is fulfilled by Amazon.
	IsFulfilledByAmazon bool
	// true if the offer is currently in the Buy Box.
	// There can be up to two Buy Box winners at any time per ASIN,
	// one that is eligible for Prime and one that is not eligible for Prime.
	IsBuyBoxWinner bool
	// true if the seller of the item is eligible to win the Buy Box.
	IsFeaturedMerchant bool
}

// OfferIdentifier contains the unique identifier for a product.
type OfferIdentifier struct {
	// An encrypted, Amazon-defined marketplace identifier.
	MarketplaceId string
	// The Seller SKU of the item.
	SellerSKU string
	// The item condition.
	ItemCondition string
	// The update time for the offer.
	TimeOfOfferChange string
}

// Identifiers contains the unique identifies for a product.
type Identifiers struct {
	MarketplaceASIN MarketplaceASIN
	SKUIdentifier   SKUIdentifier
}

// MarketplaceASIN contains ASIN for the product in the Marketplace.
type MarketplaceASIN struct {
	MarketplaceId string
	ASIN          string
}

// SKUIdentifier contains the SKU info for the products in the Marketplace.
type SKUIdentifier struct {
	MarketplaceId string
	SellerId      string
	SellerSKU     string
}

// VariationParent parents for the product.
// Contains basic info for the parent product.
type VariationParent struct {
	Identifiers Identifiers
}

// VariationChild children for the product.
// Contains basic info and children variation info.
type VariationChild struct {
	Identifiers            Identifiers
	Color                  string
	Edition                string
	Flavor                 string
	GemType                []string
	GolfClubFlex           string
	GolfClubLoft           DecimalWithUnits
	HandOrientation        string
	HardwarePlatform       string
	ItemDimensions         DimensionType
	MaterialType           []string
	MetalType              string
	Model                  string
	OperatingSystem        []string
	PackageQuantity        int
	ProductTypeSubcategory string
	RingSize               string
	ShaftMaterial          string
	Scent                  string
	Size                   string
	SizePerPearl           string
	TotalDiamondWeight     DecimalWithUnits
	TotalGemWeight         DecimalWithUnits
}

// CompetitivePrice Contains pricing information.
type CompetitivePrice struct {
	// CompetitivePriceId the pricing model for each price that is returned.
	// Valid values: 1, 2.
	// Value definitions: 1 = New Buy Box Price, 2 = Used Buy Box Price.
	CompetitivePriceId string
	// Pricing information for a given CompetitivePriceId value.
	Price Price
	// Indicates the condition of the item whose pricing information is returned.
	// Possible values are: New, Used, Collectible, Refurbished, or Club.
	Condition string `xml:"condition,attr"`
	// Indicates the subcondition of the item whose pricing information is returned.
	// Possible values are:
	// 	New, Mint, Very Good, Good, Acceptable, Poor, Club, OEM, Warranty,
	// 	Refurbished Warranty, Refurbished, Open Box, or Other.
	Subcondition string `xml:"subcondition,attr"`
	// Indicates whether or not the pricing information is for an offer listing that belongs to the requester.
	BelongsToRequester bool `xml:"belongsToRequester,attr"`
}

// OfferPrice price info for the offer, with condition and fulfillment channel.
type OfferPrice struct {
	Condition          string `xml:"condition,attr"`
	FulfillmentChannel string `xml:"fulfillmentChannel,attr"`
	Points             int
	Price
}

// Price info for the product.
type Price struct {
	// The price of the item plus the shipping cost.
	LandedPrice Money
	// The price of the item.
	ListingPrice Money
	// The shipping cost.
	Shipping Money
}

// Money an amount of money in a specified currency.
type Money struct {
	// The total value.
	Amount float64
	// The currency code.
	CurrencyCode string
}

// OfferCount The number of offer listings.
// The listing count is returned by condition, one for each listing condition value that is returned.
// Possible listing condition values are: Any, New, Used, Collectible, Refurbished, or Club.
type OfferCount struct {
	// Indicates the condition of the item.
	// Values: New, Used, Collectible, Refurbished, or Club.
	Condition string `xml:"condition,attr"`
	// Indicates whether the item is fulfilled by Amazon or by the seller.
	// Values: Amazon, Merchant
	FulfillmentChannel string `xml:"fulfillmentChannel,attr"`
	Value              int    `xml:",chardata"`
}

// DimensionType dimension info, contains Height, Weight, Length, Width.
type DimensionType struct {
	Height DecimalWithUnits
	Length DecimalWithUnits
	Width  DecimalWithUnits
	Weight DecimalWithUnits
}

// DecimalWithUnits contains the value and unit for the value.
type DecimalWithUnits struct {
	Units string  `xml:"Units,attr"`
	Value float64 `xml:",chardata"`
}

// CreatorType Creator info.
type CreatorType struct {
	Role  string `xml:"Role,attr"`
	Value string `xml:",chardata"`
}

// Language type.
type Language struct {
	Name        string
	Type        string
	AudioFormat string
}

// Image atrributes.
type Image struct {
	URL    string
	Height DecimalWithUnits
	Width  DecimalWithUnits
}

// Qualifiers identify the offer listing group from which the lowest offer listing was taken.
type Qualifiers struct {
	ItemCondition                string
	ItemSubcondition             string
	FulfillmentChannel           string
	ShipsDomestically            string
	ShippingTime                 ShippingTime
	SellerPositiveFeedbackRating string
}

// SellerFeedbackRating Information about the seller's feedback,
// including the percentage of positive feedback, and the total count of feedback received.
type SellerFeedbackRating struct {
	// The percentage of positive feedback for the seller in the past 365 days.
	SellerPositiveFeedbackRating float64
	// The count of feedback received about the seller.
	FeedbackCount int
}

// ShippingTime The time range in which an item will likely be shipped once an order has been placed.
type ShippingTime struct {
	// Max day range the item will be shipped
	MaximumDayRange string `xml:"Max"`
	// The minimum time, in hours, that the item will likely be shipped after the order has been placed.
	MinimumHours string `xml:"minimumHours,attr"`
	// The maximum time, in hours, that the item will likely be shipped after the order has been placed.
	MaximumHours string `xml:"maximumHours,attr"`
	// The date when the item will be available for shipping.
	// Only displayed for items that are not currently available for shipping.
	AvailableDate string `xml:"availableDate,attr"`
	// Indicates whether the item is available for shipping now,
	// or on a known or an unknown date in the future.
	// If known, the availableDate attribute indicates the date that the item will be available for shipping.
	// Values:
	// 	NOW - The item is available for shipping now.
	// 	FUTURE_WITHOUT_DATE - The item will be available for shipping on an unknown date in the future.
	// 	FUTURE_WITH_DATE - The item will be available for shipping on a known date in the future.
	AvailabilityType string `xml:"availabilityType,attr"`
}

type Address struct {
	State   string
	Country string
}

// ItemAttributes is product's attributes infomation.
type ItemAttributes struct {
	// Lang is not attributes of the items, its the language the attributes in
	Lang                                 string `xml:"lang,attr"`
	Actor                                []string
	Artist                               []string
	AspectRatio                          string
	AudienceRating                       string
	Author                               []string
	BackFinding                          string
	BandMaterialType                     string
	Binding                              string
	BlurayRegion                         string
	Brand                                string
	CEROAgeRating                        string
	ChainType                            string
	ClaspType                            string
	Color                                string
	CPUManufacturer                      string
	CPUSpeed                             DecimalWithUnits
	CPUType                              string
	Creator                              []CreatorType
	Department                           string
	Director                             []string
	DisplaySize                          DecimalWithUnits
	Edition                              string
	EpisodeSequence                      string
	ESRBAgeRating                        string
	Feature                              []string
	Flavor                               string
	Format                               []string
	GemType                              []string
	Genre                                string
	GolfClubFlex                         string
	GolfClubLoft                         DecimalWithUnits
	HandOrientation                      string
	HardDiskInterface                    string
	HardDiskSize                         DecimalWithUnits
	HardwarePlatform                     string
	HazardousMaterialType                string
	ItemDimensions                       DimensionType
	IsAdultProduct                       bool
	IsAutographed                        bool
	IsEligibleForTradeIn                 bool
	IsMemorabilia                        bool
	IssuesPerYear                        string
	ItemPartNumber                       string
	Label                                string
	Languages                            []Language `xml:">Language"`
	LegalDisclaimer                      string
	ListPrice                            Money
	Manufacturer                         string
	ManufacturerMaximumAge               DecimalWithUnits
	ManufacturerMinimumAge               DecimalWithUnits
	ManufacturerPartsWarrantyDescription string
	MaterialType                         []string
	MaximumResolution                    DecimalWithUnits
	MediaType                            []string
	MetalStamp                           string
	MetalType                            string
	Model                                string
	NumberOfDiscs                        int
	NumberOfIssues                       int
	NumberOfItems                        int
	NumberOfPages                        int
	NumberOfTracks                       int
	OperatingSystem                      []string
	OpticalZoom                          DecimalWithUnits
	PackageDimensions                    DimensionType
	PackageQuantity                      int
	PartNumber                           string
	PegiRating                           string
	Platform                             []string
	ProcessorCount                       int
	ProductGroup                         string
	ProductTypeName                      string
	ProductTypeSubcategory               string
	PublicationDate                      string
	Publisher                            string
	RegionCode                           string
	ReleaseDate                          string
	RingSize                             string
	RunningTime                          DecimalWithUnits
	ShaftMaterial                        string
	Scent                                string
	SeasonSequence                       string
	SeikodoProductCode                   string
	Size                                 string
	SizePerPearl                         string
	SmallImage                           Image
	Studio                               string
	SubscriptionLength                   DecimalWithUnits
	SystemMemorySize                     DecimalWithUnits
	SystemMemoryType                     string
	TheatricalReleaseDate                string
	Title                                string
	TotalDiamondWeight                   DecimalWithUnits
	TotalGemWeight                       DecimalWithUnits
	Warranty                             string
	WEEETaxValue                         Money
}
