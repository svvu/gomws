package mwsHttps

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Values is url.Values for custom encoding.
type Values struct {
	url.Values
}

// NewValues initilize the Values struct with default value
func NewValues() Values {
	return Values{url.Values{}}
}

// Encode encode the parameters and replace + by %20
func (params Values) Encode() string {
	return strings.Replace(params.Values.Encode(), "+", "%20", -1)
}

// Client is http client wrapper to handle request to mws.
type Client struct {
	// The host of the end point
	Host string
	// The path to the operation
	Path string
	// The query parameters send to the server
	parameters Values
	// Whether or not the parameters are signed
	signed bool
	// Key use to sign the query
	signatureKey string

	*http.Client
}

func NewClient(host, path string) *Client {
	return &Client{
		Host:   host,
		Path:   path,
		Client: &http.Client{},
	}
}

func (client *Client) checkSignStatus() error {
	if !client.signed {
		if client.signatureKey == "" {
			return fmt.Errorf("Query not signed, unknow secret key")
		} else {
			client.SignQuery(client.signatureKey)
		}
	}
	return nil
}

// calculateStringToSignV2 Calculate the signature to sign the query for signature version 2.
func (client *Client) calculateStringToSignV2() string {
	var stringToSign bytes.Buffer

	stringToSign.WriteString("POST\n")
	stringToSign.WriteString(client.Host)
	stringToSign.WriteString("\n")
	stringToSign.WriteString(client.Path)
	stringToSign.WriteString("\n")
	stringToSign.WriteString(client.parameters.Encode())

	return stringToSign.String()
}

// signature generate the signature by the parameters and the secretKey using HmacSHA256.
func (client *Client) signature(secretKey string) string {
	stringToSign := client.calculateStringToSignV2()
	signature2 := SignV2(stringToSign, secretKey)
	return signature2
}

// SignQuery generate the signature and add the signature to the http parameters.
func (client *Client) SignQuery(secretKey string) {
	client.parameters.Set("Timestamp", now())

	signature := client.signature(secretKey)
	client.parameters.Set("Signature", signature)
	client.signed = true
}

// SetSecretKey update the key for siging the query
func (client *Client) SetSecretKey(secretKey string) {
	client.signatureKey = secretKey
}

// SetParameters assign the passin parameters to the client
func (client *Client) SetParameters(v Values) {
	client.parameters = v
}

// AugmentParameters add new parameters to http's query and indicate the query is not signed.
func (client *Client) AugmentParameters(params map[string]string) {
	for k, v := range params {
		client.parameters.Set(k, v)
	}

	client.signed = false
}

func (client *Client) EndPoint() string {
	return "https://" + client.Host + client.Path
}

// buildRequest prepare the requet to send to the api.
func (client *Client) buildRequest() (*http.Request, error) {
	encodedParams := client.parameters.Encode()
	req, err := http.NewRequest(
		"POST",
		client.EndPoint(),
		bytes.NewBufferString(encodedParams),
	)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(encodedParams)))

	return req, nil
}

// Send send the http request to mws server.
// If the query is indicated un signed, an error will return.
func (client *Client) Send() *Response {
	response := new(Response)
	signatureErr := client.checkSignStatus()
	if signatureErr != nil {
		response.Error = signatureErr
		return response
	}

	if client.Client == nil {
		client.Client = &http.Client{}
	}

	req, err := client.buildRequest()
	if err != nil {
		response.Error = err
		return response
	}

	resp, err := client.Client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		response.Error = err
		return response
	}

	return client.parseResponse(response, resp)
}

func (client *Client) parseResponse(response *Response, resp *http.Response) *Response {
	response.Status = resp.Status
	response.StatusCode = resp.StatusCode
	if !CheckStatusCode(resp.StatusCode) {
		response.Error = fmt.Errorf("Request not success. Reason: %v", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.Error = err
		return response
	}

	response.Result = string(body)
	return response
}

const (
	Iso8061Format = time.RFC3339 // "2006-01-02T15:04:05Z07:00"
)

// Current timestamp in iso8061 format.
func now() string {
	return time.Now().UTC().Format(Iso8061Format)
}
