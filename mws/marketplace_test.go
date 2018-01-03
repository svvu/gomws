package mws

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewMarketPlace(t *testing.T) {
	Convey("Bad region", t, func() {
		client, err := NewMarketPlace("Bad")

		Convey("Unknown region error returned", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Invalid region: Bad")
		})

		Convey("Nil client returned", func() {
			So(client, ShouldBeNil)
		})
	})

	Convey("Good region", t, func() {
		type marketPlaceTest struct {
			region   string
			id       string
			endpoint string
		}

		testCases := []marketPlaceTest{
			{
				region:   "CA",
				id:       "A2EUQ1WTGCTBG2",
				endpoint: "mws.amazonservices.ca",
			},
			{
				region:   "US",
				id:       "ATVPDKIKX0DER",
				endpoint: "mws.amazonservices.com",
			},
			{
				region:   "DE",
				id:       "A1PA6795UKMFR9",
				endpoint: "mws-eu.amazonservices.com",
			},
			{
				region:   "ES",
				id:       "A1RKKUPIHCS9HS",
				endpoint: "mws-eu.amazonservices.com",
			},
			{
				region:   "FR",
				id:       "A13V1IB3VIYZZH",
				endpoint: "mws-eu.amazonservices.com",
			},
			{
				region:   "IN",
				id:       "A21TJRUUN4KGV",
				endpoint: "mws.amazonservices.in",
			},
			{
				region:   "IT",
				id:       "APJ6JRA9NG5V4",
				endpoint: "mws-eu.amazonservices.com",
			},
			{
				region:   "UK",
				id:       "A1F83G8C2ARO7P",
				endpoint: "mws-eu.amazonservices.com",
			},
			{
				region:   "JP",
				id:       "A1VC38T7YXB528",
				endpoint: "mws.amazonservices.jp",
			},
			{
				region:   "CN",
				id:       "AAHKV2X7AFYLW",
				endpoint: "mws.amazonservices.com.cn",
			},
		}

		for _, testCase := range testCases {
			Convey("Create MarketPlace for "+testCase.region, func() {
				client, err := NewMarketPlace(testCase.region)

				Convey("No error returned", func() {
					So(err, ShouldBeNil)
				})

				Convey("Client returned has match marketplace id", func() {
					So(client.Id, ShouldEqual, testCase.id)
				})

				Convey("Client returned has match endpoint", func() {
					So(client.EndPoint, ShouldEqual, testCase.endpoint)
				})
			})
		}
	})
}

func TestEncoding(t *testing.T) {
	Convey("Region CN", t, func() {
		encoding := Encoding("CN")
		Convey("UTF-16 encoding returned", func() {
			So(encoding, ShouldEqual, "UTF-16")
		})
	})

	Convey("Other region", t, func() {
		encoding := Encoding("US")
		Convey("ISO-8859-1 encoding returned", func() {
			So(encoding, ShouldEqual, "ISO-8859-1")
		})
	})
}
