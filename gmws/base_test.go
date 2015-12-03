package gmws

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/svvu/gomws/mwsHttps"
)

var testConfig = MwsConfig{
	SellerId:  "SellerA",
	AuthToken: "Authtoken",
	Region:    "US",
	AccessKey: "accesskey",
	SecretKey: "secretKey",
}

func TestNewMwsBase(t *testing.T) {
	Convey("When Seller Id is not provide", t, func() {
		testConfig.SellerId = ""
		client, err := NewMwsBase(testConfig, "V1", "Test")

		Convey("Nil Client returned", func() {
			So(client, ShouldBeNil)
		})

		Convey("Seller id error returned", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "No seller id provided")
		})

		testConfig.SellerId = "SellerA"
	})

	Convey("When AuthToken is not provide", t, func() {
		testConfig.AuthToken = ""
		client, err := NewMwsBase(testConfig, "V1", "Test")

		Convey("Nil Client returned", func() {
			So(client, ShouldBeNil)
		})

		Convey("Auth token error returned", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "No auth token provided")
		})

		testConfig.AuthToken = "authtoken"
	})

	Convey("Valid configuration", t, func() {
		client, err := NewMwsBase(testConfig, "V1", "Test")

		Convey("No err returned", func() {
			So(err, ShouldBeNil)
		})

		Convey("New client returned", func() {
			So(client, ShouldNotBeNil)
		})
	})
}

func TestPath(t *testing.T) {
	Convey("Client has name and version", t, func() {
		client, _ := NewMwsBase(testConfig, "V1", "Test")
		path := client.Path()

		Convey("Path with name and version returned", func() {
			So(path, ShouldEqual, "/Test/V1")
		})
	})

	Convey("Client only has name", t, func() {
		client, _ := NewMwsBase(testConfig, "", "Test")
		path := client.Path()

		Convey("Path with name retuend", func() {
			So(path, ShouldEqual, "/Test")
		})
	})

	Convey("Client only has version", t, func() {
		client, _ := NewMwsBase(testConfig, "V1", "")
		path := client.Path()

		Convey("Path with version retuend", func() {
			So(path, ShouldEqual, "/V1")
		})
	})
}

func TestSignatureMethod(t *testing.T) {
	Convey("Signature method HmacSHA256 returned", t, func() {
		client, _ := NewMwsBase(testConfig, "V1", "Test")
		sm := client.SignatureMethod()
		So(sm, ShouldEqual, "HmacSHA256")
	})
}

func TestSignatureVersion(t *testing.T) {
	Convey("Signature version 2 returned", t, func() {
		client, _ := NewMwsBase(testConfig, "V1", "Test")
		sv := client.SignatureVersion()
		So(sv, ShouldEqual, "2")
	})
}

func TestParamsToAugment(t *testing.T) {
	Convey("Expected info returned", t, func() {
		client, _ := NewMwsBase(testConfig, "V1", "Test")
		pta := client.paramsToAugment()
		So(pta["SellerId"], ShouldEqual, testConfig.SellerId)
		So(pta["MWSAuthToken"], ShouldEqual, testConfig.AuthToken)
		So(pta["SignatureMethod"], ShouldEqual, "HmacSHA256")
		So(pta["SignatureVersion"], ShouldEqual, "2")
		So(pta["AWSAccessKeyId"], ShouldEqual, testConfig.AccessKey)
		So(pta["Version"], ShouldEqual, "V1")
	})
}

func TestGetCredential(t *testing.T) {
	Convey("Client have access key and secret key", t, func() {
		client, _ := NewMwsBase(testConfig, "V1", "Test")
		credential := client.getCredential()

		Convey("Keys from client retured", func() {
			So(credential.AccessKey, ShouldEqual, testConfig.AccessKey)
			So(credential.SecretKey, ShouldEqual, testConfig.SecretKey)
		})
	})

	// TODO
	// Add test for get credential from env
}

func TestHttpClient(t *testing.T) {
	client, _ := NewMwsBase(testConfig, "V1", "Test")
	params := mwsHttps.NewValues()
	httpClient := client.HttpClient(params)

	Convey("Http client has expected host", t, func() {
		So(httpClient.Host, ShouldEqual, "mws.amazonservices.com")
	})

	Convey("Http client has expected path", t, func() {
		So(httpClient.Path, ShouldEqual, "/Test/V1")
	})
}
