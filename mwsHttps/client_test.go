package mwsHttps

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func getTestParams() Values {
	return Values{url.Values{
		"Action":           []string{"GetMyPriceForASIN"},
		"SellerId":         []string{"SellerID"},
		"AWSAccessKeyId":   []string{"AccessKey"},
		"MWSAuthToken":     []string{"AuthToken"},
		"SignatureVersion": []string{"2"},
		"Timestamp":        []string{"2015-10-20T22:46:07Z"},
		"Version":          []string{"2011-10-01"},
		"SignatureMethod":  []string{"HmacSHA256"},
		"MarketplaceId":    []string{"Marketplace"},
		"ASINList.ASIN.1":  []string{"ASIN"},
	}}
}

var (
	expectedParams = "ASINList.ASIN.1=ASIN&AWSAccessKeyId=AccessKey&" +
		"Action=GetMyPriceForASIN&MWSAuthToken=AuthToken&MarketplaceId=Marketplace&" +
		"SellerId=SellerID&SignatureMethod=HmacSHA256&SignatureVersion=2&" +
		"Timestamp=2015-10-20T22%3A46%3A07Z&Version=2011-10-01"
)

func TestCheckSignStatus(t *testing.T) {
	httpClient := NewClient("http://test", "/path")
	httpClient.SetParameters(NewValues())

	Convey("When query not signed and dont have secret key", t, func() {
		err := httpClient.checkSignStatus()
		Convey("An error raise", func() {
			So(err, ShouldNotBeNil)
		})
	})

	Convey("When query not signed and have secret key", t, func() {
		httpClient.SetSecretKey("SecretKey")
		err := httpClient.checkSignStatus()
		Convey("No error raise", func() {
			So(err, ShouldBeNil)
		})

		Convey("Parameters have signature value", func() {
			So(httpClient.parameters.Get("Signature"), ShouldNotBeNil)
		})
	})
}

func TestCalculateStringToSignV2(t *testing.T) {
	httpClient := NewClient("mws.amazonservices.com", "/Products/2011-10-01")
	expectedString := "POST\nmws.amazonservices.com\n" +
		"/Products/2011-10-01\n" + expectedParams
	httpClient.SetParameters(getTestParams())

	stringToSign := httpClient.calculateStringToSignV2()
	Convey("Expect string returned", t, func() {
		So(stringToSign, ShouldEqual, expectedString)
	})
}

func TestSignature(t *testing.T) {
	httpClient := NewClient("mws.amazonservices.com", "/Products/2011-10-01")
	expectedString := "MSwoBGqrM1h7IqQ8QIZo3sNvCKuV3zvTUKO%2FFAAWNt0%3D"
	httpClient.SetParameters(getTestParams())

	signature := url.QueryEscape(httpClient.signature("SecretKey"))
	Convey("Expect string returned", t, func() {
		So(signature, ShouldEqual, expectedString)
	})
}

func TestSignQuery(t *testing.T) {
	httpClient := NewClient("mws.amazonservices.com", "/Products/2011-10-01")
	httpClient.SetParameters(getTestParams())

	httpClient.SignQuery("SecretKey")
	Convey("Signature added to parameters", t, func() {
		So(httpClient.parameters.Get("Signature"), ShouldNotBeNil)
	})
}

func TestSetSecretKey(t *testing.T) {
	httpClient := NewClient("mws.amazonservices.com", "/Products/2011-10-01")
	httpClient.SetSecretKey("secretKey")
	Convey("signatureKey setted", t, func() {
		So(httpClient.signatureKey, ShouldEqual, "secretKey")
	})
}

func TestSetParameters(t *testing.T) {
	httpClient := NewClient("mws.amazonservices.com", "/Products/2011-10-01")
	params := NewValues()
	params.Set("A", "1")
	httpClient.SetParameters(params)
	Convey("Parameters setted", t, func() {
		So(httpClient.parameters, ShouldResemble, params)
	})
}

func TestAugmentParameters(t *testing.T) {
	httpClient := NewClient("mws.amazonservices.com", "/Products/2011-10-01")
	httpClient.signed = true
	httpClient.SetParameters(NewValues())
	paramsToAugment := map[string]string{"key": "value"}
	httpClient.AugmentParameters(paramsToAugment)

	Convey("Signed changed to false", t, func() {
		So(httpClient.signed, ShouldBeFalse)
	})

	Convey("Parameters has key from augment parameters", t, func() {
		So(httpClient.parameters.Get("key"), ShouldNotBeNil)
	})
}

func TestBuildRequest(t *testing.T) {
	httpClient := NewClient("mws.amazonservices.com", "/Products/2011-10-01")
	httpClient.SetParameters(getTestParams())

	request, err := httpClient.buildRequest()

	Convey("Error is nil", t, func() {
		So(err, ShouldBeNil)
	})

	Convey("Request body have pass in params with signature", t, func() {
		body, _ := ioutil.ReadAll(request.Body)
		So(string(body), ShouldEqual, expectedParams)
	})

	Convey("Request has content type urlencoded", t, func() {
		So(request.Header.Get("Content-Type"), ShouldEqual, "application/x-www-form-urlencoded")
	})

	Convey("Request has content type urlencoded", t, func() {
		So(request.Header.Get("Content-Length"), ShouldEqual, strconv.Itoa(len(expectedParams)))
	})

	Convey("Request has client url", t, func() {
		So(request.URL.String(), ShouldEqual, "https://mws.amazonservices.com/Products/2011-10-01")
	})
}

func TestSend(t *testing.T) {
	Convey("Query not signed and secret key is empty", t, func() {
		httpClient := NewClient("mws.amazonservices.com", "/Products/2011-10-01")
		resp := httpClient.Send()

		Convey("Result is empty", func() {
			So(resp.Result(), ShouldBeBlank)
		})

		Convey("Error is not nil", func() {
			So(resp.Error, ShouldNotBeNil)
		})
	})

	Convey("Response with error", t, func() {
		Convey("404 not found", func() {
			ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.NotFound(w, r)
			}))
			defer ts.Close()

			noVerifyTransport := &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			}

			httpClient := NewClient(strings.Replace(ts.URL, "https://", "", -1), "")
			httpClient.Transport = noVerifyTransport
			httpClient.signed = true

			response := httpClient.Send()

			Convey("Response with status code 404", func() {
				So(response.StatusCode, ShouldEqual, 404)
			})

			Convey("Response with error page not found", func() {
				So(response.Error.Error(), ShouldEqual, "Request not success. Reason: 404 Not Found")
			})
		})
	})

	Convey("Good response", t, func() {
		ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello, client")
		}))
		defer ts.Close()

		noVerifyTransport := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}

		httpClient := NewClient(strings.Replace(ts.URL, "https://", "", -1), "")
		httpClient.Transport = noVerifyTransport
		httpClient.signed = true

		response := httpClient.Send()

		Convey("Response with status code 200", func() {
			So(response.StatusCode, ShouldEqual, 200)
		})

		Convey("Response with nil error", func() {
			So(response.Error, ShouldBeNil)
		})

		Convey("Response with result msg", func() {
			// Response msg end with new line character
			So(strings.TrimSpace(response.Result()), ShouldEqual, "Hello, client")
		})
	})
}

func TestParseResponse(t *testing.T) {
	httpClient := NewClient("mws.amazonservices.com", "/Products/2011-10-01")
	Convey("When response with 503", t, func() {
		httpResponse := http.Response{
			Status:     "Service unavailable",
			StatusCode: 503,
			Body:       ioutil.NopCloser(bytes.NewBufferString("503 Service unavailable")),
		}
		response := httpClient.parseResponse(&Response{}, &httpResponse)
		Convey("Response with status code 503", func() {
			So(response.StatusCode, ShouldEqual, 503)
		})

		Convey("Response with error page Service unavailable", func() {
			So(response.Error.Error(), ShouldEqual, "Request not success. Reason: Service unavailable")
		})
	})

	Convey("When response with valid message", t, func() {
		httpResponse := http.Response{
			Status:     "OK",
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString("Request OK")),
		}
		response := httpClient.parseResponse(&Response{}, &httpResponse)
		Convey("Response with status code 200", func() {
			So(response.StatusCode, ShouldEqual, 200)
		})

		Convey("Response with result message", func() {
			So(response.Result(), ShouldEqual, "Request OK")
		})

		Convey("Response with nil error", func() {
			So(response.Error, ShouldBeNil)
		})
	})
}
