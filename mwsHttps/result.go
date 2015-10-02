package mwsHttps

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

type Result string

func (r Result) ToJSON() {

}

func (r Result) ToStruct() {

}

// PrettyPrint print the xml with indentation added to the start tag of xml.
// namespace will be ignored.
func (r Result) PrettyPrint() {
	decoder := xml.NewDecoder(strings.NewReader(string(r)))
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
