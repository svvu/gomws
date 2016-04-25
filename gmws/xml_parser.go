// Deprecated

package gmws

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/svvu/gomws/mwsHttps"
)

// XMLParser use parse the XML string.
type XMLParser struct {
	XMLString string
}

// NewXMLParser Create a new parse for the response, seting the response result to XMLString.
func NewXMLParser(response *mwsHttps.Response) *XMLParser {
	return &XMLParser{XMLString: response.Result()}
}

// PrettyPrint Print the XML in indent format.
// Note: The method will ignore namespace, attributes, comments for the tag.
func (xmlp *XMLParser) PrettyPrint() {
	decoder := xml.NewDecoder(strings.NewReader(xmlp.XMLString))
	decoder.Strict = false

	rBuffer := &bytes.Buffer{}
	encoder := xml.NewEncoder(rBuffer)
	encoder.Indent("", "  ")

	for {
		token, err := decoder.RawToken()
		if err == io.EOF {
			encoder.Flush()
			break
		}
		if err != nil {
			break
		}
		err = encoder.EncodeToken(token)
		if err != nil {
			break
		}
	}

	fmt.Println(string(rBuffer.Bytes()))
}

// Parse unmarshal the XML string to target struct
func (xmlp *XMLParser) Parse(v interface{}) error {
	err := xml.Unmarshal([]byte(xmlp.XMLString), v)

	xrHandler, ok := v.(XMLResultHandler)
	if ok {
		xrHandler.ParseCallback()
	}

	return err
}

// HasError check whether or nor API send back error, not http error
func (xmlp *XMLParser) HasError() bool {
	err := xmlp.GetError()
	return err.Error != nil
}

// GetError return the ErrorResult contains error type, code, and detail message
func (xmlp *XMLParser) GetError() *ErrorResult {
	err := ErrorResult{}
	xmlp.Parse(&err)
	return &err
}

// XMLResultHandler is interface to allow parsed XML result to have callback
type XMLResultHandler interface {
	ParseCallback()
}
