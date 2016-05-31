package gmws

import (
	"fmt"
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ErrorTestExample() []byte {
	response, ferr := ioutil.ReadFile("./examples/ErrorResponse.xml")

	if ferr != nil {
		fmt.Println(ferr)
		return []byte{}
	}

	return response
}

func NoErrorTestExample() []byte {
	response, ferr := ioutil.ReadFile("./examples/XMLNodeTest.xml")

	if ferr != nil {
		fmt.Println(ferr)
		return []byte{}
	}

	return response
}

func TestHasErrors(t *testing.T) {
	Convey("When response has error", t, func() {
		xNode, _ := GenerateXMLNode(ErrorTestExample())

		Convey("Has error should be true", func() {
			So(HasErrors(xNode), ShouldBeTrue)
		})
	})

	Convey("When response has no error", t, func() {
		xNode, _ := GenerateXMLNode(NoErrorTestExample())

		Convey("Has error should be false", func() {
			So(HasErrors(xNode), ShouldBeFalse)
		})
	})
}

func TestGetErrors(t *testing.T) {
	Convey("When response has error", t, func() {
		xNode, _ := GenerateXMLNode(ErrorTestExample())
		errors, _ := GetErrors(xNode)

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
		xNode, _ := GenerateXMLNode(NoErrorTestExample())
		errors, _ := GetErrors(xNode)

		Convey("No errors should returned", func() {
			So(len(errors), ShouldEqual, 0)
		})
	})
}
