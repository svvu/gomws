package mws

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Values is a url.Values for custom encoding.
type Values struct {
	url.Values
}

// NewValues initilize the Values struct with default value.
func NewValues() Values {
	return Values{url.Values{}}
}

// Encode encode the parameters and replace '+' by '%20'.
func (params Values) Encode() string {
	return strings.Replace(params.Values.Encode(), "+", "%20", -1)
}

// Client the basic client handle request send to API endpoint.
type Client struct {
	// The api host for the region.
	Host string
	// Region of the marketplace in two character.
	Region string
	// Marketplace identitier for the region.
	MarketPlaceId string
	// The API version.
	Version string
	// The API name.
	Name string
	// Seller's Amazon id.
	SellerId string
	// Auth token for developer to use the API.
	AuthToken string
	// Credential for requests.
	accessKey string
	secretKey string

	*http.Client
}

// NewClient create a new mws base client.
// Config is the configuration struct.
// 	Contains value for: SellerId, AuthToken, Region, and optional credential.
// If region is not set in config, will be default to US.
func NewClient(config Config, version, name string) (*Client, error) {
	if config.SellerId == "" {
		return nil, fmt.Errorf("No seller id provided")
	}

	region := config.Region
	if region == "" {
		region = "US"
	}

	marketPlace, mError := NewMarketPlace(region)
	if mError != nil {
		return nil, mError
	}

	credential := config.Credential()
	if credential.AccessKey == "" || credential.SecretKey == "" {
		return nil, fmt.Errorf("Can't find mws credential information")
	}

	base := Client{
		SellerId:      config.SellerId,
		AuthToken:     config.AuthToken,
		Region:        region,
		MarketPlaceId: marketPlace.Id,
		Host:          marketPlace.EndPoint,
		Version:       version,
		Name:          name,
		accessKey:     credential.AccessKey,
		secretKey:     credential.SecretKey,
		Client:        new(http.Client),
	}

	return &base, nil
}

// Path generate the url path for the api endpoint.
func (base Client) Path() string {
	path := ""
	if base.Name != "" {
		path += "/" + base.Name
	}
	if base.Version != "" {
		path += "/" + base.Version
	}
	return path
}

// EndPoint generate the endpoint for the request by combinding the host and path.
func (base Client) EndPoint() string {
	return "https://" + base.Host + base.Path()
}

// SignatureMethod return the HmacSHA256 signature method string.
func (base Client) SignatureMethod() string {
	return "HmacSHA256"
}

// SignatureVersion return version 2.
func (base Client) SignatureVersion() string {
	return "2"
}

// SendRequest accept a structured params and send the request to the API.
func (base Client) SendRequest(structuredParams Parameters) (*Response, error) {
	request, err := base.buildRequest(structuredParams)
	if err != nil {
		return nil, err
	}

	resp, err := base.Client.Do(request)
	if err != nil {
		return nil, err
	}

	return NewResponse(resp), nil
}

// buildRequest prepare the requet to send to the api.
// The method will create a post request with encoded body of signed parameters.
func (base Client) buildRequest(structuredParams Parameters) (*http.Request, error) {
	params, err := structuredParams.Normalize()
	if err != nil {
		return nil, err
	}

	encodedParams := base.signQuery(params).Encode()
	req, err := http.NewRequest(
		"POST",
		base.EndPoint(),
		bytes.NewBufferString(encodedParams),
	)

	if err != nil {
		return nil, err
	}

	// Add content headers.
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(encodedParams)))

	return req, nil
}

// signQuery generate the signature and add the signature to the http parameters.
func (base Client) signQuery(params Values) Values {
	// Add client info to the query params.
	params.Set("SellerId", base.SellerId)
	if base.AuthToken != "" {
		params.Set("MWSAuthToken", base.AuthToken)
	}
	params.Set("SignatureMethod", base.SignatureMethod())
	params.Set("SignatureVersion", base.SignatureVersion())
	params.Set("AWSAccessKeyId", base.accessKey)
	params.Set("Version", base.Version)

	params.Set("Timestamp", now())

	signature := base.generateSignature(params)
	params.Set("Signature", signature)
	return params
}

// signature generate the signature by the parameters and the secretKey using HmacSHA256.
func (base Client) generateSignature(params Values) string {
	stringToSign := base.generateStringToSignV2(params)
	signature2 := SignV2(stringToSign, base.secretKey)
	return signature2
}

// generateStringToSignV2 Generate the string to sign for the query.
func (base Client) generateStringToSignV2(params Values) string {
	var stringToSign bytes.Buffer

	stringToSign.WriteString("POST\n")
	stringToSign.WriteString(base.Host)
	stringToSign.WriteString("\n")
	stringToSign.WriteString(base.Path())
	stringToSign.WriteString("\n")
	stringToSign.WriteString(params.Encode())

	return stringToSign.String()
}
