package gmws

import "encoding/xml"

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
