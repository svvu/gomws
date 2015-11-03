package gmws

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/svvu/gomws/mwsHttps"
	"strings"
	"testing"
)

func TestOptionalParams(t *testing.T) {
	Convey("Given a valid parameters", t, func() {
		params := []mwsHttps.Parameters{{"key1": "value1", "key2": "value2"}}

		Convey("When accept keys are empty", func() {
			acceptKeys := []string{}
			result := OptionalParams(acceptKeys, params)

			Convey("Empty parameter is returned", func() {
				So(result, ShouldBeEmpty)
			})
		})

		Convey("When accept keys not in passed in parameters", func() {
			acceptKeys := []string{"a", "b"}
			result := OptionalParams(acceptKeys, params)

			Convey("Empty parameter is returned", func() {
				So(result, ShouldBeEmpty)
			})
		})

		Convey("When accept keys exist in passed in parameters", func() {
			acceptKeys := []string{"key1", "key2"}
			result := OptionalParams(acceptKeys, params)

			Convey("Accept keys with title case are returned", func() {
				for _, ak := range acceptKeys {
					So(result, ShouldContainKey, strings.Title(ak))
				}
			})

			Convey("Values for accepy keys are returned", func() {
				for _, ak := range acceptKeys {
					So(result[strings.Title(ak)], ShouldEqual, params[0][ak])
				}
			})
		})
	})

	Convey("Given an emptry array of parameters", t, func() {
		params := []mwsHttps.Parameters{{}}
		acceptKeys := []string{"key1", "key2"}
		result := OptionalParams(acceptKeys, params)

		Convey("Empty parameter is returned", func() {
			So(result, ShouldBeEmpty)
		})
	})

	Convey("Given array of many parameters", t, func() {
		params := []mwsHttps.Parameters{
			{"key1": "value1", "key2": "value2"},
			{"key2": "value22", "key3": "value3"},
			{"key4": "value4", "key5": "value5"},
		}

		Convey("When accept keys in diff elements of passed in params", func() {
			acceptKeys := []string{"key1", "key3", "key4"}
			result := OptionalParams(acceptKeys, params)

			Convey("Values for accepy keys are returned", func() {
				for _, ak := range acceptKeys {
					So(result, ShouldContainKey, strings.Title(ak))
				}
			})
		})

		Convey("When same keys are in diff elements", func() {
			acceptKeys := []string{"key1", "key2", "key4"}
			result := OptionalParams(acceptKeys, params)

			Convey("Values from later parameter is returned", func() {
				So(result["Key2"], ShouldEqual, "value22")
			})
		})
	})
}
