package mwsHttpClient

import (
	"bytes"
	"net/http"
	"time"
)

type mwsHttpClient struct {
	EndPoint   string
	Parameters NormalizedParameters
}

func newHttpClient(endPoint string, params NormalizedParameters) *mwsHttpClient {
	return &mwsHttpClient{
		EndPoint:   endPoint,
		Parameters: params,
	}
}

func (client *mwsHttpClient) AugmentParameters(params map[string]string) {
	for k, v := range params {
		client.Parameters.Add(k, v)
	}

	client.Parameters.Add("Timestamp", now())
}

func (client *mwsHttpClient) calculateStringToSignV2() string {
	var stringToSign bytes.Buffer

	stringToSign.WriteString("POST\n")
	stringToSign.WriteString(client.Endpoint())
	stringToSign.WriteString("\n/")
	stringToSign.WriteString("\n")
	stringToSign.WriteString(client.Parameters.UrlEncode())

	return stringToSign.String()
}

func (client *mwsHttpClient) signature(params NormalizedParameters, credential Credential) string {
	stringToSign := calculateStringToSignV2(params)
	signature2 := SignV2(stringToSign, credential.SecretKey)
	return signature2
}

func (client *mwsHttpClient) SignQuery(credential Credential) {
	signature := client.Signature(client.Parameters, credential)
	client.Parameters.Add("Signature", signature)
}

func (client *mwsHttpClient) Request() string {
	encodedParams := client.Parameters.UrlEncode()
	req, err := http.NewRequest(
		"POST",
		client.Endpoint(),
		bytes.NewBufferString(encodedParams),
	)

	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Add("Content-Length", strconv.Itoa(len(encodedParams)))

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {

	}
	defer resp.Body.Close()

	return "sth"
}

const (
	Iso8061Format = time.RFC3339 // "2006-01-02T15:04:05Z07:00"
)

func now() time.Time {
	return time.Now().UTC().Format(Iso8061Format)
}
