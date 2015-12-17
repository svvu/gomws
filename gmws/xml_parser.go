package gmws

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/svvu/gomws/mwsHttps"
)

// XMLParser use parse the xml string
type XMLParser struct {
	XMLString string
}

// NewXMLParser Create a new parse for the response, seting the response result to XMLString
func NewXMLParser(response *mwsHttps.Response) *XMLParser {
	return &XMLParser{XMLString: response.Result()}
}

// PrettyPrint Print the xml in indent format.
// Note: The method will ignore namespace, attributes, comments for the tag
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

// Parse unmarshal the xml string to target struct
func (xmlp *XMLParser) Parse(v interface{}) error {
	return xml.Unmarshal([]byte(xmlp.XMLString), v)
}
