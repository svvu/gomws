package gmws

import (
	"fmt"
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/svvu/gomws/mwsHttps"
)

func getErrorResponse() *mwsHttps.Response {
	response, ferr := ioutil.ReadFile("./examples/ErrorResponse.xml")
	if ferr != nil {
		fmt.Println(ferr)
	}
	resp := &mwsHttps.Response{Body: response}
	return resp
}

func getNormalResponse() *mwsHttps.Response {
	response, ferr := ioutil.ReadFile("./examples/GetServiceStatus.xml")
	if ferr != nil {
		fmt.Println(ferr)
	}
	resp := &mwsHttps.Response{Body: response}
	return resp
}

func Test_HasError(t *testing.T) {
	Convey("When response has error tag", t, func() {
		response := getErrorResponse()
		xmlParser := NewXMLParser(response)

		Convey("Has error is true", func() {
			So(xmlParser.HasError(), ShouldBeTrue)
		})
	})

	Convey("When response doesnt have error tag", t, func() {
		response := getNormalResponse()
		xmlParser := NewXMLParser(response)

		Convey("Has error is false", func() {
			So(xmlParser.HasError(), ShouldBeFalse)
		})
	})
}

func Test_GetError(t *testing.T) {
	Convey("When response has error tag", t, func() {
		response := getErrorResponse()
		xmlParser := NewXMLParser(response)

		Convey("Error is not nil", func() {
			So(xmlParser.GetError().Error, ShouldNotBeNil)
		})
	})

	Convey("When response doesnt have error tag", t, func() {
		response := getNormalResponse()
		xmlParser := NewXMLParser(response)

		Convey("Error should be nil", func() {
			So(xmlParser.GetError().Error, ShouldBeNil)
		})
	})
}
