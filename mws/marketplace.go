package mws

import (
	"fmt"
)

// EndPoints a list of API endpoints by marketpalceID
var EndPoints = map[string]string{
	"A2EUQ1WTGCTBG2": "mws.amazonservices.ca",
	"ATVPDKIKX0DER":  "mws.amazonservices.com",
	"A1PA6795UKMFR9": "mws-eu.amazonservices.com",
	"A1RKKUPIHCS9HS": "mws-eu.amazonservices.com",
	"A13V1IB3VIYZZH": "mws-eu.amazonservices.com",
	"A21TJRUUN4KGV":  "mws.amazonservices.in",
	"APJ6JRA9NG5V4":  "mws-eu.amazonservices.com",
	"A1F83G8C2ARO7P": "mws-eu.amazonservices.com",
	"A1VC38T7YXB528": "mws.amazonservices.jp",
	"AAHKV2X7AFYLW":  "mws.amazonservices.com.cn",
}

// MarketPlaceIds a list of marketplace by region
var MarketPlaceIds = map[string]string{
	"CA": "A2EUQ1WTGCTBG2",
	"US": "ATVPDKIKX0DER",
	"DE": "A1PA6795UKMFR9",
	"ES": "A1RKKUPIHCS9HS",
	"FR": "A13V1IB3VIYZZH",
	"IN": "A21TJRUUN4KGV",
	"IT": "APJ6JRA9NG5V4",
	"UK": "A1F83G8C2ARO7P",
	"JP": "A1VC38T7YXB528",
	"CN": "AAHKV2X7AFYLW",
}

// MarketPlaceError for marketplace.
// There are two type of errors: marketplace id error and region error.
type MarketPlaceError struct {
	errorType string
	value     string
}

func (e MarketPlaceError) Error() string {
	return fmt.Sprintf("Invalid %v: %v", e.errorType, e.value)
}

// MarketPlace contains region, id, and enpoint for the marketpalce.
type MarketPlace struct {
	Region   string
	Id       string
	EndPoint string
}

// NewMarketPlace create a new marketplace base on the region.
func NewMarketPlace(region string) (*MarketPlace, error) {
	mp := MarketPlace{Region: region}

	marketPlaceId, idError := mp.MarketPlaceId()
	if idError != nil {
		return nil, idError
	}
	mp.Id = marketPlaceId

	endPoint, endPointError := mp.MarketPlaceEndPoint()
	if endPointError != nil {
		return nil, endPointError
	}
	mp.EndPoint = endPoint
	return &mp, nil
}

// MarketPlaceEndPoint get the MWS end point for the region.
func (mp *MarketPlace) MarketPlaceEndPoint() (string, error) {
	if mp.EndPoint != "" {
		return mp.EndPoint, nil
	}
	if val, ok := EndPoints[mp.Id]; ok {
		return val, nil
	}
	return "", MarketPlaceError{"marketplace id", mp.Id}
}

// MarketPlaceId get the marketpalce id for the region.
func (mp *MarketPlace) MarketPlaceId() (string, error) {
	if mp.Id != "" {
		return mp.Id, nil
	}
	if val, ok := MarketPlaceIds[mp.Region]; ok {
		return val, nil
	}
	return "", MarketPlaceError{"region", mp.Region}
}

// Encoding get the ecoding for file upload and parsing
// TODO add encoding for JP.
func Encoding(region string) string {
	switch region {
	case "CN":
		return "UTF-16"
	default:
		return "ISO-8859-1"
	}
}
