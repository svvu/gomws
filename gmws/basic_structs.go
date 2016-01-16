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

// LoadExample load example xml and parse it by passed in struct v
func LoadExample(filePath string, v interface{}) (interface{}, error) {
	response, ferr := ioutil.ReadFile(filePath)
	if ferr != nil {
		return v, ferr
	}
	resp := &mwsHttps.Response{Result: string(response)}
	xmlParser := NewXMLParser(resp)
	err := xmlParser.Parse(v)
	if err != nil {
		return v, err
	}
	return v, nil
}
