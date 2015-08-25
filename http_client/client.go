package mwsHttpClient

import (
	"bytes"
	"net/http"
	"time"
)

type MwsHttpClient struct {
	EndPoint   string
	Parameters NormalizedParameters
	signed     bool
}

func (client *MwsHttpClient) AugmentParameters(params map[string]string) {
	for k, v := range params {
		client.Parameters.Set(k, v)
	}

	client.Parameters.Set("Timestamp", now())
	client.signed = false
}

func (client *MwsHttpClient) calculateStringToSignV2() string {
	var stringToSign bytes.Buffer

	stringToSign.WriteString("POST\n")
	stringToSign.WriteString(client.EndPoint)
	stringToSign.WriteString("\n/")
	stringToSign.WriteString("\n")
	stringToSign.WriteString(client.Parameters.UrlEncode())

	return stringToSign.String()
}

func (client *MwsHttpClient) signature(params NormalizedParameters, secretKey string) string {
	stringToSign := client.calculateStringToSignV2()
	signature2 := SignV2(stringToSign, secretKey)
	return signature2
}

func (client *MwsHttpClient) SignQuery(secretKey string) {
	signature := client.signature(client.Parameters, secretKey)
	client.Parameters.Set("Signature", signature)
	client.signed = true
}

func (client *MwsHttpClient) Request() string {
	if !client.signed {
		return ""
	}

	encodedParams := client.Parameters.UrlEncode()
	req, err := http.NewRequest(
		"POST",
		client.EndPoint,
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

func now() string {
	return time.Now().UTC().Format(Iso8061Format)
}
