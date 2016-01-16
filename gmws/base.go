package gmws

import (
	"fmt"

	"github.com/svvu/gomws/marketplace"
	"github.com/svvu/gomws/mwsHttps"
)

// MwsBase contains the basic information for the api client.
type MwsBase struct {
	// Seller's Amazon id.
	SellerId string
	// Auth token for developer to use the API.
	AuthToken string
	// Region of the marketplace in two character.
	Region        string
	MarketPlaceId string
	Host          string
	// The API version.
	Version string
	// The API name.
	Name      string
	accessKey string
	secretKey string
}

// NewMwsBase create a new mws base.
// MwsConfig is the configuration struct.
// 	Contains value for: SellerId, AuthToken, Region.
func NewMwsBase(config MwsConfig, version, name string) (*MwsBase, error) {
	if config.SellerId == "" {
		return nil, fmt.Errorf("No seller id provided")
	}

	if config.AuthToken == "" {
		return nil, fmt.Errorf("No auth token provided")
	}

	region := config.Region
	if region == "" {
		region = "US"
	}

	marketPlace, mError := marketplace.New(region)
	if mError != nil {
		return nil, mError
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
	return &base, nil
}

// Path generate the url path to the api endpoint
func (base MwsBase) Path() string {
	path := ""
	if base.Name != "" {
		path += "/" + base.Name
	}
	if base.Version != "" {
		path += "/" + base.Version
	}
	return path
}

// SignatureMethod return the HmacSHA256 signature method string.
func (base MwsBase) SignatureMethod() string {
	return "HmacSHA256"
}

// SignatureVersion return version 2.
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
	}

	return GetCredential()
}

// HTTPClient return an http client with pass in querys, and ready for send of
//  request to the server.
func (base MwsBase) HTTPClient(values mwsHttps.Values) *mwsHttps.Client {
	httpClient := mwsHttps.NewClient(base.Host, base.Path())
	httpClient.SetParameters(values)
	httpClient.SetSecretKey(base.getCredential().SecretKey)
	httpClient.AugmentParameters(base.paramsToAugment())
	return httpClient
}

// SendRequest accept a structured params and send the request to the API.
func (base MwsBase) SendRequest(structuredParams Parameters) *mwsHttps.Response {
	normalizedParams, err := structuredParams.Normalize()

	if err != nil {
		return &mwsHttps.Response{Error: err}
	}

	httpClient := base.HTTPClient(normalizedParams)
	return httpClient.Send()
}
