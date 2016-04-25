// Deprecated

package gmws

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func errorReponseResult() *ErrorResult {
	result, err := LoadExample("./examples/ErrorResponse.xml", &ErrorResult{})

	if err != nil {
		fmt.Println(err)
	}
	return result.(*ErrorResult)
}

func serviceStatusResult() *GetServiceStatusResult {
	result, err := LoadExample("./examples/GetServiceStatus.xml", &GetServiceStatusResult{})

	if err != nil {
		fmt.Println(err)
	}
	return result.(*GetServiceStatusResult)
}

func Test_ErrorResult(t *testing.T) {
	Convey("Request response", t, func() {
		responseErr := errorReponseResult().Error

		Convey("Type is Sender", func() {
			So(responseErr.Type, ShouldEqual, "Sender")
		})

		Convey("Code is InvalidParameterValue", func() {
			So(responseErr.Code, ShouldEqual, "InvalidParameterValue")
		})

		Convey("Message is 'Value for parameter MarketplaceId is not valid: ATVPDKIKX0DE'", func() {
			So(responseErr.Message, ShouldEqual, "Value for parameter MarketplaceId is not valid: ATVPDKIKX0DE")
		})
	})
}

func Test_GetServiceStatusResult(t *testing.T) {
	Convey("Request response", t, func() {
		status := serviceStatusResult()

		Convey("Status is GREEN_I", func() {
			So(status.Status, ShouldEqual, "GREEN_I")
		})

		Convey("Timestamp is 2013-09-05T18%3A12%3A21", func() {
			So(status.Timestamp, ShouldEqual, "2013-09-05T18%3A12%3A21")
		})

		Convey("MessageId is 173964729I", func() {
			So(status.MessageId, ShouldEqual, "173964729I")
		})

		Convey("Message", func() {
			msg := status.Messages[0]

			Convey("Locale is en_US", func() {
				So(msg.Locale, ShouldEqual, "en_US")
			})

			Convey("Text is We are experiencing high latency in UK because of heavy traffic.", func() {
				So(msg.Text, ShouldEqual, "We are experiencing high latency in UK because of heavy traffic.")
			})
		})
	})
}
