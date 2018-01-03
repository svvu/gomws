package mock

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

// Server a mock server to capture request to API.
type Server struct {
	*httptest.Server
	responseHandler func(r *http.Request) *http.Response
}

// SetResponse set the response for the incoming requests.
// The method cached the response body value and reset the body to support
// 	multiple times of requests.
func (ms *Server) SetResponse(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	ms.responseHandler = func(r *http.Request) *http.Response {
		if err == nil {
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
		return resp
	}
}

// SetResponseHandler set a handler to update reponse base on request info.
func (ms *Server) SetResponseHandler(handler func(r *http.Request) *http.Response) {
	ms.responseHandler = handler
}

// Host return the url without the method (URN).
func (ms *Server) Host() string {
	return strings.Replace(ms.Server.URL, "https://", "", -1)
}

// NewServer create and start a new mock Server.
func NewServer() *Server {
	server := &Server{}

	handler := func(w http.ResponseWriter, r *http.Request) {
		var resp *http.Response

		if server.responseHandler == nil {
			resp = &http.Response{}
		} else {
			resp = server.responseHandler(r)
		}

		for k, v := range resp.Header {
			w.Header()[k] = v
		}

		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
		} else {
			w.WriteHeader(resp.StatusCode)
			fmt.Fprint(w, string(body))
		}
	}

	server.Server = httptest.NewTLSServer(http.HandlerFunc(handler))

	return server
}

// NoVerifyTransport return transport that skip the certificate verification.
func NoVerifyTransport() *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
}

// NewResponse create a new response.
func NewResponse(statusCode int, body string) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
		Header:     http.Header{},
	}
}
