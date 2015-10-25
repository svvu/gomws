package mwsHttps

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFormatParameterKey(t *testing.T) {
	baseKey := "baseKey"

	Convey("No optional keys passed in", t, func() {
		key := formatParameterKey(baseKey)

		Convey("Result is same as base key", func() {
			So(key, ShouldEqual, baseKey)
		})
	})

	Convey("With optional keys passed in", t, func() {
		optionalKeys := []string{"k1", "k2", "k3"}
		key := formatParameterKey(baseKey, optionalKeys...)

		Convey("Result is base key combined with optonal keys", func() {
			So(key, ShouldEqual, "baseKey.k1.k2.k3")
		})
	})
}

func TestValuesEncode(t *testing.T) {
	values := NewValues()
	values.Add("key1", "a b c")

	Convey("Space should be replaced by %20", t, func() {
		encodeValue := values.Encode()
		So(encodeValue, ShouldEqual, "key1=a%20b%20c")
	})
}

func TestParametersMerge(t *testing.T) {
	params1 := Parameters{"key1": "value1", "key2": "value2"}
	params2 := Parameters{"key3": "value3", "key4": "value4"}

	params1.Merge(params2)

	Convey("Merged param should have 4 keys", t, func() {
		So(len(params1), ShouldEqual, 4)
	})

	Convey("Merged param should have key1", t, func() {
		So(params1["key1"], ShouldEqual, "value1")
	})

	Convey("Merged param should have key2", t, func() {
		So(params1["key2"], ShouldEqual, "value2")
	})

	Convey("Merged param should have key3", t, func() {
		So(params1["key3"], ShouldEqual, "value3")
	})

	Convey("Merged param should have key4", t, func() {
		So(params1["key4"], ShouldEqual, "value4")
	})
}

func TestParametersStructureKeys(t *testing.T) {
	Convey("When values is another parameters", t, func() {
		params := Parameters{
			"outterKey1": Parameters{
				"paramKey1": "value1",
				"paramKey2": "value2",
			},
		}

		Convey("Inner parameter's key added to result key", func() {
			resultParam := params.StructureKeys("outterKey1", "param")

			Convey("Key outterKey1.param.paramKey1 has value value1", func() {
				So(resultParam["outterKey1.param.paramKey1"], ShouldEqual, "value1")
			})

			Convey("Key outterKey2.param.paramKey2 has value value2", func() {
				So(resultParam["outterKey1.param.paramKey2"], ShouldEqual, "value2")
			})
		})
	})

	Convey("When values is another map", t, func() {
		params := Parameters{
			"outterKey2": map[string]int{
				"mapKey1": 1,
				"mapKey2": 2,
			},
		}

		Convey("Map's key added to result key", func() {
			resultParam := params.StructureKeys("outterKey2", "map")

			Convey("Key outterKey2.map.mapKey1 has value 1", func() {
				So(resultParam["outterKey2.map.mapKey1"], ShouldEqual, 1)
			})

			Convey("Key outterKey2.map.mapKey2 has value 2", func() {
				So(resultParam["outterKey2.map.mapKey2"], ShouldEqual, 2)
			})
		})
	})

	Convey("When values is slice", t, func() {
		params := Parameters{"slice": []string{"value1", "value2", "value3"}}

		Convey("Slice is flatten and index add to key", func() {
			resultParam := params.StructureKeys("slice", "key")

			Convey("Key slice.key.1 has value value1", func() {
				So(resultParam["slice.key.1"], ShouldEqual, "value1")
			})

			Convey("Key slice.key.2 has value value1", func() {
				So(resultParam["slice.key.2"], ShouldEqual, "value2")
			})

			Convey("Key slice.key.3 has value value1", func() {
				So(resultParam["slice.key.3"], ShouldEqual, "value3")
			})

		})
	})

	Convey("When values is neither slice or map", t, func() {
		type testStruct struct{ a string }
		params := Parameters{
			"int":    1,
			"string": "string value",
			"struct": testStruct{"struct filed"},
		}

		Convey("Base key join with optional keys", func() {
			Convey("Key int.a.b.c has value 1", func() {
				resultParam := params.StructureKeys("int", "a", "b", "c")
				So(resultParam["int.a.b.c"], ShouldEqual, 1)
			})

			Convey("Key string has value 'string value'", func() {
				resultParam := params.StructureKeys("string")
				So(resultParam["string"], ShouldEqual, "string value")
			})

			Convey("Key struct.k has field a equal to 'struct filed'", func() {
				resultParam := params.StructureKeys("struct", "k")
				So(resultParam["struct.k"].(testStruct).a, ShouldEqual, "struct filed")
			})
		})
	})
}

func TestParametersNormalize(t *testing.T) {
	Convey("When value is bool", t, func() {
		params := Parameters{"boolKey": true}
		Convey("Value converted to string", func() {
			resultParam, err := params.Normalize()
			So(err, ShouldBeNil)
			So(resultParam.Get("boolKey"), ShouldEqual, "true")
		})
	})

	Convey("When value is int", t, func() {
		params := Parameters{"intKey": 12}
		Convey("Value converted to string", func() {
			resultParam, err := params.Normalize()
			So(err, ShouldBeNil)
			So(resultParam.Get("intKey"), ShouldEqual, "12")
		})
	})

	Convey("When value is float32", t, func() {
		params := Parameters{"float32Key": float32(1.11)}
		Convey("Value converted to string", func() {
			resultParam, err := params.Normalize()
			So(err, ShouldBeNil)
			So(resultParam.Get("float32Key"), ShouldEqual, "1.11")
		})
	})

	Convey("When value is float32", t, func() {
		params := Parameters{"float32Key": float32(1.1)}
		Convey("Value converted to string with 2 precision", func() {
			resultParam, err := params.Normalize()
			So(err, ShouldBeNil)
			So(resultParam.Get("float32Key"), ShouldEqual, "1.10")
		})
	})

	Convey("When value is float64", t, func() {
		params := Parameters{"float64Key": float64(1.1)}
		Convey("Value converted to string with 2 precision", func() {
			resultParam, err := params.Normalize()
			So(err, ShouldBeNil)
			So(resultParam.Get("float64Key"), ShouldEqual, "1.10")
		})
	})

	Convey("When value is string", t, func() {
		params := Parameters{"stringKey": "A string"}
		Convey("Value is string", func() {
			resultParam, err := params.Normalize()
			So(err, ShouldBeNil)
			So(resultParam.Get("stringKey"), ShouldEqual, "A string")
		})
	})

	Convey("When value is other type", t, func() {
		params := Parameters{"sliceKey": []int{1, 2}}
		Convey("An error raise", func() {
			_, err := params.Normalize()
			So(err.Error(), ShouldEqual, "Unexpected type []int")
		})
	})
}
