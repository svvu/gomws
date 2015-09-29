package gmws

import (
	"../marketplace"
	. "../mwsHttpClient"
)

type MwsBase struct {
	SellerId      string // Seller's Amazon id
	AuthToken     string // Auth token for developer to use the api
	Region        string // Region of the marketplace in two character
	MarketPlaceId string
	Host          string
	Version       string // The api's version
	Name          string // The api's name
	accessKey     string
	secretKey     string
}

// sellerId, authToken, region string
func NewMwsBase(config MwsConfig, version, name string) *MwsBase {
	if config.SellerId == "" {
		// Log
		return nil
	}

	if config.AuthToken == "" {
		// Log
		return nil
	}

	region := config.Region
	if region == "" {
		region = "US"
	}

	marketPlace, mError := marketplace.New(region)
	if mError != nil {
		// TODO
	}

	base := MwsBase{
		SellerId:      config.SellerId,
		AuthToken:     config.AuthToken,
		Region:        region,
		MarketPlaceId: marketPlace.Id,
		Host:          marketPlace.EndPoint,
		Version:       version,
		Name:          name,
		accessKey:     config.AccessKey,
		secretKey:     config.SecretKey,
	}
	return &base
}

func (base MwsBase) Path() string {
	return "/" + base.Name + "/" + base.Version
}

func (base MwsBase) SignatureMethod() string {
	return "HmacSHA256"
}

func (base MwsBase) SignatureVersion() string {
	return "2"
}

// paramsToAugment generate a list of client information add to the query.
func (base MwsBase) paramsToAugment() map[string]string {
	clientInfo := map[string]string{
		"SellerId":         base.SellerId,
		"MWSAuthToken":     base.AuthToken,
		"SignatureMethod":  base.SignatureMethod(),
		"SignatureVersion": base.SignatureVersion(),
		"AWSAccessKeyId":   base.getCredential().AccessKey,
		"Version":          base.Version,
	}
	return clientInfo
}

// getCredential return the mws credential, if not set, it will try to retrieve
//  the information from env variables.
// Using env variables is recommanded and more secure.
func (base MwsBase) getCredential() Credential {
	if base.accessKey != "" && base.secretKey != "" {
		return Credential{base.accessKey, base.secretKey}
	} else {
		return GetCredential()
	}
}

// HttpClient return an http client with pass in querys, and ready for send of
//  request to the server
func (base MwsBase) HttpClient(values Values) *MwsHttpClient {
	httpClient := MwsHttpClient{
		Host:       base.Host,
		Path:       base.Path(),
		Parameters: values,
	}

	httpClient.AugmentParameters(base.paramsToAugment())
	httpClient.SignQuery(base.getCredential().SecretKey)
	return &httpClient
}
