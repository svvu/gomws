package marketplace

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("Bad region", t, func() {
		client, err := New("Bad")

		Convey("Unknown region error returned", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Invalid region: Bad")
		})

		Convey("Nil client returned", func() {
			So(client, ShouldBeNil)
		})
	})

	Convey("Good region", t, func() {
		client, err := New("US")

		Convey("No error returned", func() {
			So(err, ShouldBeNil)
		})

		Convey("Client returned has match marketplace id", func() {
			So(client.Id, ShouldEqual, "ATVPDKIKX0DER")
		})

		Convey("Client returned has match endpoint", func() {
			So(client.EndPoint, ShouldEqual, "mws.amazonservices.com")
		})
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
