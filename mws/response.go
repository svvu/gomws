package mws

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// Response a custom http response with additional helper methods.
type Response struct {
	*http.Response

	Error error
}

// NewResponse return a new mws response.
func NewResponse(resp *http.Response) *Response {
	response := &Response{Response: resp}
	response.Error = parseResponseError(response)

	return response
}

// ResultParser create a new node parser for response body.
func (resp *Response) ResultParser() (*ResultParser, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return NewResultParser(body)
}

// WriteBodyTo write the response body to the destination output.
func (resp *Response) WriteBodyTo(out io.Writer) error {
	_, err := io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// Close make sure the body drained, and then close the connection to make the
// connection can be resue.
func (resp *Response) Close() {
	ioutil.ReadAll(resp.Body)
	resp.Body.Close()
}

// parseResponseError parse the xml body to extract the detail error msg for status code below.
// If fail to extract the info, it will default to response status.
//
// https://docs.developer.amazonservices.com/en_US/dev_guide/DG_Errors.html#ErrorMessages_Service_errors
// Table 1. Common HTTP error status codes
// Error code				HTTP status code	Description
// InputStreamDisconnected		400				There was an error reading the input stream.
// InvalidParameterValue		400				An invalid parameter value was used, or the request size exceeded the maximum accepted size, or the request expired.
// AccessDenied					401				Access was denied.
// InvalidAccessKeyId			403				An invalid AWSAccessKeyId value was used.
// SignatureDoesNotMatch		403				The signature used does not match the server's calculated signature value.
// InvalidAddress				404				An invalid API section or operation value was used, or an invalid path was used.
// InternalError				500				There was an internal service failure.
// QuotaExceeded				503				The total number of requests in an hour was exceeded.
// RequestThrottled				503				The frequency of requests was greater than allowed.
func parseResponseError(resp *Response) error {
	switch resp.StatusCode {
	case 200:
		return nil
	case 400, 401, 403, 404, 500, 503:
		baseErr := fmt.Errorf("Request not success. Reason: %v", resp.Status)

		// Reset the body, so parsing error won't darin the body.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return baseErr
		}
		resp.Body.Close()
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		node, err := NewResultParser(body)
		if err != nil || !node.HasErrorNodes() {
			return baseErr
		}

		mwsErrors, err := node.GetMWSErrors()
		if err != nil {
			return baseErr
		}

		msgs := []string{}
		for _, mwsErr := range mwsErrors {
			msgs = append(msgs, mwsErr.Message)
		}

		return errors.Wrap(errors.New(strings.Join(msgs, "\n")), baseErr.Error())
	default:
		return fmt.Errorf("Request not success. Reason: %v", resp.Status)
	}
}
