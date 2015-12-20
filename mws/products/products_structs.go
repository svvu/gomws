package products

import "encoding/xml"

// References:
// http://g-ecx.images-amazon.com/images/G/01/mwsportal/doc/en_US/products/ProductsAPI_Response.xsd
// http://g-ecx.images-amazon.com/images/G/01/mwsportal/doc/en_US/products/default.xsd

// ListMatchingProductsResult the result for the ListMatchingProducts operation.
type ListMatchingProductsResult struct {
	XMLName  xml.Name  `xml:"ListMatchingProductsResponse"`
	Products []Product `xml:"ListMatchingProductsResult>Products>Product"`
}

// Product contains basic, relationship to other products, pricing, and offers info
type Product struct {
	Identifiers         Identifiers
	AttributeSets       []ItemAttributes `xml:">ItemAttributes"`
	Relationships       []Relationship
	CompetitivePricing  CompetitivePricing
	SalesRankings       []SalesRank `xml:">SalesRank"`
	LowestOfferListings []LowestOfferListing
	Offers              []Offer
}

// Identifiers contains the unique identifies for a product.
type Identifiers struct {
	MarketplaceASIN MarketplaceASIN
	SKUIdentifier   SKUIdentifier
}

// Relationship contains product variation information, if applicable.
// If your search criteria match a product that is identified by a variation
// parent ASIN, the related VariationChild elements are contained in the Relationships
// element.
// If your search criteria match a specific variation child ASIN, the related
// VariationParent element is contained in the Relationships element.
type Relationship struct {
	VariationParent VariationParent
	VariationChild  VariationChild
}

type CompetitivePricing struct {
	CompetitivePrices     []CompetitivePrice
	NumberOfOfferListings []OfferListingCount
	TradeInValue          Money
}

// SalesRank contains sales rank information for the product by product category.
type SalesRank struct {
	ProductCategoryId string
	Rank              int
}

type LowestOfferListing struct {
	Qualifiers                      Qualifiers
	NumberOfOfferListingsConsidered int
	SellerFeedbackCount             int
	Price                           Price
	MultipleOffersAtLowestPrice     string
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

type MarketplaceASIN struct {
	MarketplaceId string
	ASIN          string
}

type SKUIdentifier struct {
	MarketplaceId string
	SellerId      string
	SellerSKU     string
}

type VariationParent struct {
	Identifiers Identifiers
}

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

type CompetitivePrice struct {
	CompetitivePriceId string
	Price              Price
	Condition          string // attributes
	Subcondition       string // attributes
	BelongsToRequester bool   // attributes
}

type Price struct {
	LandedPrice  Money
	ListingPrice Money
	Shipping     Money
}

type OfferListingCount struct {
	Condition string // attributes
	Value     int
}

type DimensionType struct {
	Height DecimalWithUnits
	Length DecimalWithUnits
	Width  DecimalWithUnits
	Weight DecimalWithUnits
}

type DecimalWithUnits struct {
	Units string  `xml:"Units,attr"`
	Value float64 `xml:",chardata"`
}

type CreatorType struct {
	Role  string `xml:"Role,attr"`
	Value string `xml:",chardata"`
}

type Language struct {
	Name        string
	Type        string
	AudioFormat string
}

type Image struct {
	URL    string
	Height DecimalWithUnits
	Width  DecimalWithUnits
}

type Money struct {
	Amount       float64
	CurrencyCode string
}

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
