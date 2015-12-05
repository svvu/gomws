package gmws

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/svvu/gomws/mwsHttps"
)

type XmlParser struct {
	XmlString string
}

func NewXmlParser(response *mwsHttps.Response) *XmlParser {
	return &XmlParser{XmlString: response.Result}
}

func (xmlp *XmlParser) PrettyPrint() {
	decoder := xml.NewDecoder(strings.NewReader(xmlp.XmlString))
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
