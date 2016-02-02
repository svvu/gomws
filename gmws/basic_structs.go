package gmws

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/svvu/gomws/mwsHttps"
)

// GetServiceStatusResult the result for the GetServiceStatus operation.
type GetServiceStatusResult struct {
	XMLName   xml.Name  `xml:"GetServiceStatusResponse"`
	Status    string    `xml:"GetServiceStatusResult>Status"`
	Timestamp string    `xml:"GetServiceStatusResult>Timestamp"`
	MessageId string    `xml:"GetServiceStatusResult>MessageId"`
	Messages  []Message `xml:"GetServiceStatusResult>Messages>Message"`
}

// Message the operational status message.
type Message struct {
	Locale string
	Text   string
}

// ErrorResult error message from the API, most of time its bad request error.
type ErrorResult struct {
	XMLName xml.Name `xml:"ErrorResponse"`
	Error   *Error   `xml:"Error"`
}

// Error represents the error message from the API.
type Error struct {
	// Error type. Values: Sender, Server.
	Type string
	// Amazon error code.
	Code string
	// Text explain the error.
	Message string
}

// LoadExample load example xml and parse it by passed in struct v
func LoadExample(filePath string, v interface{}) (interface{}, error) {
	response, ferr := ioutil.ReadFile(filePath)
	if ferr != nil {
		return v, ferr
	}
	resp := &mwsHttps.Response{Body: response}
	xmlParser := NewXMLParser(resp)
	err := xmlParser.Parse(v)
	if err != nil {
		return v, err
	}
	return v, nil
}
