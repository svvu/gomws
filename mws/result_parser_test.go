package mws

import (
	"fmt"
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewResultParser(t *testing.T) {
	Convey("when create XML parser fail", t, func() {
		rp, err := NewResultParser([]byte(""))

		Convey("return nil ResultParser", func() {
			So(rp, ShouldBeNil)
		})

		Convey("return error", func() {
			So(err, ShouldNotBeNil)
		})
	})

	Convey("when create XML parser success", t, func() {
		rp, err := NewResultParser([]byte("<foo>bar</foo>"))

		Convey("return created ResultParser", func() {
			So(rp, ShouldNotBeNil)
		})

		Convey("return no error", func() {
			So(err, ShouldBeNil)
		})
	})
}

func errorTestExample() []byte {
	response, ferr := ioutil.ReadFile("./exampleResponses/ErrorResponse.xml")

	if ferr != nil {
		fmt.Println(ferr)
		return []byte{}
	}

	return response
}

func noErrorTestExample() []byte {
	response, ferr := ioutil.ReadFile("./exampleResponses/NoErrorResponse.xml")

	if ferr != nil {
		fmt.Println(ferr)
		return []byte{}
	}

	return response
}

func TestResultParser_HasErrorNodes(t *testing.T) {
	Convey("When response has error", t, func() {
		rp, _ := NewResultParser(errorTestExample())

		Convey("Has error should be true", func() {
			So(rp.HasErrorNodes(), ShouldBeTrue)
		})
	})

	Convey("When response has no error", t, func() {
		rp, _ := NewResultParser(noErrorTestExample())

		Convey("Has error should be false", func() {
			So(rp.HasErrorNodes(), ShouldBeFalse)
		})
	})
}

func TestResultParser_GetMWSErrors(t *testing.T) {
	Convey("When response has error", t, func() {
		rp, _ := NewResultParser(errorTestExample())
		errors, _ := rp.GetMWSErrors()

		Convey("Errors should not be blank", func() {
			So(len(errors), ShouldBeGreaterThan, 0)
		})

		Convey("Error should has Type Sender", func() {
			So(errors[0].Type, ShouldEqual, "Sender")
		})

		Convey("Error should has Code InvalidParameterValue", func() {
			So(errors[0].Code, ShouldEqual, "InvalidParameterValue")
		})

		Convey("Error should has Message Value for parameter MarketplaceId is not valid: ATVPDKIKXXX", func() {
			So(errors[0].Message, ShouldEqual, "Value for parameter MarketplaceId is not valid: ATVPDKIKXXX")
		})

		Convey("Error should has Detail No comment", func() {
			So(errors[0].Detail, ShouldEqual, "No comment")
		})
	})

	Convey("When response has no error", t, func() {
		rp, _ := NewResultParser(noErrorTestExample())
		errors, _ := rp.GetMWSErrors()

		Convey("No errors should returned", func() {
			So(len(errors), ShouldEqual, 0)
		})
	})
}
