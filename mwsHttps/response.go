package mwsHttps

import (
	"net/http"
	"regexp"
	"strconv"
)

// Response is the reponse from the API.
type Response struct {
	Body          []byte
	Error         error
	StatusCode    int
	Status        string
	Header        http.Header
	RequestHeader string
}

// Result return the response data as string
func (resp Response) Result() string {
	return string(resp.Body)
}

// CheckStatusCode check whether or not the status indicate a request error.
// When the code not start with 1 or 2, false returned.
func CheckStatusCode(code int) bool {
	scode := strconv.Itoa(code)
	greenStatus := regexp.MustCompile(`^[1-2][0-9]{2}$`)

	return greenStatus.MatchString(scode)
}
