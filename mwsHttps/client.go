package mwsHttps

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// http client wrapper to handle request to mws.
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

// Request send the http request to mws server.
// If the query is indicated un signed, an error will return.
func (client *Client) Request() (Result, error) {
	signatureErr := client.checkSignStatus()
	if signatureErr != nil {
		return "", signatureErr
	}

	encodedParams := client.parameters.Encode()
	req, err := http.NewRequest(
		"POST",
		client.EndPoint(),
		bytes.NewBufferString(encodedParams),
	)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(encodedParams)))

	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return Result(body), nil
}

const (
	Iso8061Format = time.RFC3339 // "2006-01-02T15:04:05Z07:00"
)

// Current timestamp in iso8061 format.
func now() string {
	return time.Now().UTC().Format(Iso8061Format)
}
