package mwsHttps

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckStatusCode(t *testing.T) {
	Convey("When code is 1xx", t, func() {
		pass := CheckStatusCode(100)
		Convey("Should return true", func() {
			So(pass, ShouldBeTrue)
		})
	})

	Convey("When code is 2xx", t, func() {
		pass := CheckStatusCode(200)
		Convey("Should return true", func() {
			So(pass, ShouldBeTrue)
		})
	})

	Convey("When code is 3xx", t, func() {
		pass := CheckStatusCode(333)
		Convey("Should return false", func() {
			So(pass, ShouldBeFalse)
		})
	})

	Convey("When code is 4xx", t, func() {
		pass := CheckStatusCode(404)
		Convey("Should return false", func() {
			So(pass, ShouldBeFalse)
		})
	})

	Convey("When code is 5xx", t, func() {
		pass := CheckStatusCode(503)
		Convey("Should return false", func() {
			So(pass, ShouldBeFalse)
		})
	})

	Convey("When code is not 3 digit", t, func() {
		pass := CheckStatusCode(2003)
		Convey("Should return false", func() {
			So(pass, ShouldBeFalse)
		})
	})
}
