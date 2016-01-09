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

// MultiProductsResult the result from the operation, contains meta info for the result.
// MultiProductsResult contains one of more products.
type MultiProductsResult struct {
	Products []Product `xml:">Product"`
	Status   string    `xml:"status,attr"`
	ID       string    `xml:"Id,attr"`
	IDType   string    `xml:"IdType,attr"`
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
}

// Product contains basic, relationship to other products, pricing, and offers info.
type Product struct {
	Identifiers         Identifiers
	AttributeSets       []ItemAttributes `xml:">ItemAttributes"`
	Relationships       Relationships
	CompetitivePricing  CompetitivePricing
	SalesRankings       []SalesRank          `xml:">SalesRank"`
	LowestOfferListings []LowestOfferListing `xml:">LowestOfferListing"`
	Offers              []Offer              `xml:">Offer"`
}

// Identifiers contains the unique identifies for a product.
type Identifiers struct {
	MarketplaceASIN MarketplaceASIN
	SKUIdentifier   SKUIdentifier
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
	CompetitivePrices     []CompetitivePrice  `xml:">CompetitivePrice"`
	NumberOfOfferListings []OfferListingCount `xml:">OfferListingCount"`
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

type Offer struct {
	BuyingPrice        Price
	RegularPrice       Money
	FulfillmentChannel string
	ItemCondition      string
	ItemSubCondition   string
	SellerId           string
	SellerSKU          string
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

// Price info for the product.
type Price struct {
	LandedPrice  Money
	ListingPrice Money
	Shipping     Money
}

// OfferListingCount The number of active offer listings.
// The listing count is returned by condition, one for each listing condition value that is returned.
// Possible listing condition values are: Any, New, Used, Collectible, Refurbished, or Club.
type OfferListingCount struct {
	Condition string `xml:"condition,attr"`
	Value     int    `xml:",chardata"`
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

// Money info.
type Money struct {
	Amount       float64
	CurrencyCode string
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

type ShippingTime struct {
	Max string
}
