package mwsHttps

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Response struct {
	Result     string
	Error      error
	StatusCode int
	Status     string
}

// PrettyPrint print the xml with indentation added to the start tag of xml.
// namespace will be ignored.
func (resp *Response) PrettyPrint() {
	decoder := xml.NewDecoder(strings.NewReader(string(resp.Result)))
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

// CheckStatusCode check whether or not the status indicate a request error
// When the code not start with 1 or 2, false returned
func CheckStatusCode(code int) bool {
	scode := strconv.Itoa(code)
	greenStatus := regexp.MustCompile(`^[1-2][0-9]{2}$`)

	return greenStatus.MatchString(scode)
}
