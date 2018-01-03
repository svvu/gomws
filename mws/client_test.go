package mws

import (
	"io/ioutil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/svvu/gomws/mws/mock"
)

const (
	testSellerID      = "SellerID"
	testAuthToken     = "AuthToken"
	testRegion        = "US"
	testAccessKey     = "AccessKey"
	testSecretKey     = "SecretKey"
	testTimestamp     = "2015-10-20T22:46:07Z"
	testASIN          = "ASIN"
	testMarketplaceID = "Marketplace"
	testAction        = "GetMyPriceForASIN"
	testVersion       = "2011-10-01"
	testClientName    = "Products"
)

func testConfig() Config {
	return Config{
		SellerId:  testSellerID,
		AuthToken: testAuthToken,
		Region:    testRegion,
		AccessKey: testAccessKey,
		SecretKey: testSecretKey,
	}
}

var testParams = struct {
	params    Parameters
	values    Values
	signature string
}{
	params: Parameters{
		"Action":          testAction,
		"MarketplaceId":   testMarketplaceID,
		"ASINList.ASIN.1": testASIN,
	},
	values: Values{url.Values{
		"Action":           []string{testAction},
		"SellerId":         []string{testSellerID},
		"AWSAccessKeyId":   []string{testAccessKey},
		"MWSAuthToken":     []string{testAuthToken},
		"SignatureVersion": []string{"2"},
		"Timestamp":        []string{testTimestamp},
		"Version":          []string{testVersion},
		"SignatureMethod":  []string{"HmacSHA256"},
		"MarketplaceId":    []string{testMarketplaceID},
		"ASINList.ASIN.1":  []string{testASIN},
	}},
	signature: "MSwoBGqrM1h7IqQ8QIZo3sNvCKuV3zvTUKO/FAAWNt0=",
}

func TestNewClient(t *testing.T) {
	Convey("When Region is not provide", t, func() {
		tconfig := testConfig()
		tconfig.Region = ""

		client, _ := NewClient(tconfig, "V1", "Test")

		Convey("Region defualt to US", func() {
			So(client.Region, ShouldEqual, "US")
		})
	})

	Convey("When Unknow Region is provide", t, func() {
		tconfig := testConfig()
		tconfig.Region = "UnKnown"

		client, err := NewClient(tconfig, "V1", "Test")

		Convey("Nil Client returned", func() {
			So(client, ShouldBeNil)
		})

		Convey("MarketPlace error returned", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Invalid region: UnKnown")
		})
	})

	Convey("When Seller Id is not provide", t, func() {
		tconfig := testConfig()
		tconfig.SellerId = ""

		client, err := NewClient(tconfig, "V1", "Test")

		Convey("Nil Client returned", func() {
			So(client, ShouldBeNil)
		})

		Convey("Seller id error returned", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "No seller id provided")
		})
	})

	Convey("When AuthToken is not provide", t, func() {
		tconfig := testConfig()
		tconfig.AuthToken = ""

		client, err := NewClient(tconfig, "V1", "Test")

		Convey("no error return", func() {
			So(err, ShouldBeNil)
		})

		Convey("Client returned", func() {
			So(client, ShouldNotBeNil)
		})

		Convey("Auth token is empty", func() {
			So(client.AuthToken, ShouldEqual, "")
		})
	})

	Convey("When AccessKey and SecretKey are not provide", t, func() {
		tconfig := testConfig()
		tconfig.AccessKey = ""
		tconfig.SecretKey = ""

		Convey("when credential not set in env variable", func() {
			client, err := NewClient(tconfig, "V1", "Test")

			Convey("Nil Client returned", func() {
				So(client, ShouldBeNil)
			})

			Convey("Credential error returned", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Can't find mws credential information")
			})
		})

		Convey("when credential set in env variable", func() {
			os.Setenv("AWS_ACCESS_KEY", "a_key")
			os.Setenv("AWS_SECRET_KEY", "s_key")

			client, err := NewClient(tconfig, "V1", "Test")

			Convey("No err returned", func() {
				So(err, ShouldBeNil)
			})

			Convey("Client returned with value read from env variables", func() {
				So(client.accessKey, ShouldEqual, "a_key")
				So(client.secretKey, ShouldEqual, "s_key")
			})
		})
	})

	Convey("Valid configuration", t, func() {
		client, err := NewClient(testConfig(), "V1", "Test")

		Convey("No err returned", func() {
			So(err, ShouldBeNil)
		})

		Convey("Client returned has expect value set", func() {
			So(client.SellerId, ShouldEqual, testSellerID)
			So(client.AuthToken, ShouldEqual, testAuthToken)
			So(client.Region, ShouldEqual, "US")
			So(client.MarketPlaceId, ShouldEqual, "ATVPDKIKX0DER")
			So(client.Host, ShouldEqual, "mws.amazonservices.com")
			So(client.Version, ShouldEqual, "V1")
			So(client.Name, ShouldEqual, "Test")
			So(client.accessKey, ShouldEqual, testAccessKey)
			So(client.secretKey, ShouldEqual, testSecretKey)
			So(client.Client, ShouldNotBeNil)
		})
	})
}

func TestClient_Path(t *testing.T) {
	testCases := []struct {
		desc         string
		inputVersion string
		inputName    string
		output       string
	}{
		{
			desc:         "Client has name and version",
			inputVersion: "V1",
			inputName:    "Test",
			output:       "/Test/V1",
		},
		{
			desc:         "Client only has name",
			inputVersion: "",
			inputName:    "Test",
			output:       "/Test",
		},
		{
			desc:         "Client only has version",
			inputVersion: "V1",
			inputName:    "",
			output:       "/V1",
		},
	}

	for _, testCase := range testCases {
		Convey(testCase.desc, t, func() {
			client, _ := NewClient(testConfig(), testCase.inputVersion, testCase.inputName)
			path := client.Path()

			So(path, ShouldEqual, testCase.output)
		})
	}
}

func TestClient_EndPoint(t *testing.T) {
	Convey("Expect host + path return", t, func() {
		client, _ := NewClient(testConfig(), "V1", "Test")
		endPoint := client.EndPoint()

		So(endPoint, ShouldEqual, "https://mws.amazonservices.com/Test/V1")
	})
}

func TestClient_SignatureMethod(t *testing.T) {
	Convey("Signature method HmacSHA256 returned", t, func() {
		client, _ := NewClient(testConfig(), testVersion, testClientName)
		sm := client.SignatureMethod()
		So(sm, ShouldEqual, "HmacSHA256")
	})
}

func TestClient_SignatureVersion(t *testing.T) {
	Convey("Signature version 2 returned", t, func() {
		client, _ := NewClient(testConfig(), testVersion, testClientName)
		sv := client.SignatureVersion()
		So(sv, ShouldEqual, "2")
	})
}

func TestClient_SendRequest(t *testing.T) {
	server := mock.NewServer()
	defer server.Close()

	client, _ := NewClient(testConfig(), testVersion, testClientName)
	client.Host = server.Host()
	client.Transport = mock.NoVerifyTransport()

	Convey("When create request error", t, func() {
		badClient, _ := NewClient(testConfig(), testVersion, testClientName)
		badClient.Host = "bad host"

		resp, err := badClient.SendRequest(testParams.params)

		Convey("Error is not nil", func() {
			So(err, ShouldNotBeNil)
		})

		Convey("Response is nil", func() {
			So(resp, ShouldBeNil)
		})
	})

	Convey("When Response with error", t, func() {
		server.SetResponse(mock.NewResponse(404, "Not Found"))

		resp, err := client.SendRequest(testParams.params)
		defer resp.Body.Close()

		Convey("SendRequest not return error", func() {
			So(err, ShouldBeNil)
		})

		Convey("Response with status code 404", func() {
			So(resp.StatusCode, ShouldEqual, 404)
		})

		Convey("Response with error page not found", func() {
			So(resp.Error.Error(), ShouldEqual, "Request not success. Reason: 404 Not Found")
		})
	})

	Convey("When request success", t, func() {
		server.SetResponse(mock.NewResponse(200, "Hello, client"))

		resp, err := client.SendRequest(testParams.params)
		defer resp.Body.Close()

		Convey("SendRequest not return error", func() {
			So(err, ShouldBeNil)
		})

		Convey("Response with status code 200", func() {
			So(resp.StatusCode, ShouldEqual, 200)
		})

		Convey("Response with nil error", func() {
			So(resp.Error, ShouldBeNil)
		})

		Convey("Response with result msg", func() {
			// Response msg end with new line character
			body, _ := ioutil.ReadAll(resp.Body)
			So(strings.TrimSpace(string(body)), ShouldEqual, "Hello, client")
		})
	})
}

func TestClient_buildRequest(t *testing.T) {
	now = func() string { return testTimestamp }

	Convey("When create request error", t, func() {
		client, _ := NewClient(testConfig(), testVersion, testClientName)
		p := Parameters{"bad_params": []interface{}{}}

		request, err := client.buildRequest(p)

		Convey("Error retuened", func() {
			So(err, ShouldNotBeNil)
		})

		Convey("No request created", func() {
			So(request, ShouldBeNil)
		})
	})

	Convey("When create request success", t, func() {
		client, _ := NewClient(testConfig(), testVersion, testClientName)
		values, _ := testParams.params.Normalize()
		expectedSignedQuery := client.signQuery(values).Encode()

		request, err := client.buildRequest(testParams.params)

		Convey("Error is nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("Request body have pass in params with signature", func() {
			body, _ := ioutil.ReadAll(request.Body)
			So(string(body), ShouldEqual, expectedSignedQuery)
		})

		Convey("Request has content headers", func() {
			So(
				request.Header.Get("Content-Type"),
				ShouldEqual,
				"application/x-www-form-urlencoded",
			)
			So(
				request.Header.Get("Content-Length"),
				ShouldEqual,
				strconv.Itoa(len(expectedSignedQuery)),
			)
		})

		Convey("Request has client url", func() {
			So(request.URL.String(), ShouldEqual, client.EndPoint())
		})
	})
}

func TestClient_signQuery(t *testing.T) {
	now = func() string { return testTimestamp }
	config := testConfig()
	client, _ := NewClient(config, testVersion, testClientName)
	values, _ := testParams.params.Normalize()

	values = client.signQuery(values)

	Convey("Client info added to parameters", t, func() {
		So(values.Get("SellerId"), ShouldEqual, config.SellerId)
		So(values.Get("MWSAuthToken"), ShouldEqual, config.AuthToken)
		So(values.Get("SignatureMethod"), ShouldEqual, client.SignatureMethod())
		So(values.Get("SignatureVersion"), ShouldEqual, client.SignatureVersion())
		So(values.Get("AWSAccessKeyId"), ShouldEqual, config.AccessKey)
		So(values.Get("Version"), ShouldEqual, testVersion)
		So(values.Get("Timestamp"), ShouldEqual, now())
	})

	Convey("Signature added to parameters", t, func() {
		So(values.Get("Signature"), ShouldEqual, testParams.signature)
	})
}

func TestClient_generateSignature(t *testing.T) {
	client, _ := NewClient(testConfig(), testVersion, testClientName)

	signature := client.generateSignature(testParams.values)

	Convey("Expect string returned", t, func() {
		So(signature, ShouldEqual, testParams.signature)
	})
}

func TestClient_generateStringToSignV2(t *testing.T) {
	queryString := "ASINList.ASIN.1=ASIN&AWSAccessKeyId=AccessKey&" +
		"Action=GetMyPriceForASIN&MWSAuthToken=AuthToken&MarketplaceId=Marketplace&" +
		"SellerId=SellerID&SignatureMethod=HmacSHA256&SignatureVersion=2&" +
		"Timestamp=2015-10-20T22%3A46%3A07Z&Version=2011-10-01"
	expectedString := "POST\n" +
		"mws.amazonservices.com\n" +
		"/Products/2011-10-01\n" +
		queryString
	client, _ := NewClient(testConfig(), testVersion, testClientName)

	stringToSign := client.generateStringToSignV2(testParams.values)

	Convey("Expect string returned", t, func() {
		So(stringToSign, ShouldEqual, expectedString)
	})
}
