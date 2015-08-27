package mwsHttpClient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type MwsHttpClient struct {
	Host       string
	Path       string
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
	stringToSign.WriteString(client.Host)
	stringToSign.WriteString("\n")
	stringToSign.WriteString(client.Path)
	stringToSign.WriteString("\n")
	stringToSign.WriteString(client.Parameters.Encode())

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

func (client *MwsHttpClient) EndPoint() string {
	return "https://" + client.Host + client.Path
}

func (client *MwsHttpClient) Request() string {
	if !client.signed {
		return ""
	}

	encodedParams := client.Parameters.Encode()
	req, err := http.NewRequest(
		"POST",
		client.EndPoint(),
		bytes.NewBufferString(encodedParams),
	)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(encodedParams)))

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		// TODO Clean up
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	// TODO Clean up
	fmt.Println(string(body))

	return "sth"
}

const (
	Iso8061Format = time.RFC3339 // "2006-01-02T15:04:05Z07:00"
)

func now() string {
	return time.Now().UTC().Format(Iso8061Format)
}
