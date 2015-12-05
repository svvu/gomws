package mwsHttps

import (
	"regexp"
	"strconv"
)

type Response struct {
	Result     string
	Error      error
	StatusCode int
	Status     string
}

// CheckStatusCode check whether or not the status indicate a request error
// When the code not start with 1 or 2, false returned
func CheckStatusCode(code int) bool {
	scode := strconv.Itoa(code)
	greenStatus := regexp.MustCompile(`^[1-2][0-9]{2}$`)

	return greenStatus.MatchString(scode)
}
