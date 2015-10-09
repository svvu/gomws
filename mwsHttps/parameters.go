package mwsHttps

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// formatParameterKey combine the base key and the augument keys by '.'.
func formatParameterKey(baseKey string, keys ...string) string {
	return baseKey + "." + strings.Join(keys, ".")
}

// Values is url.Values for custom encoding.
type Values struct {
	url.Values
}

func NewValues() Values {
	return Values{url.Values{}}
}

func (params Values) Encode() string {
	return strings.Replace(params.Values.Encode(), "+", "%20", -1)
}

// Set sets the key to value. It replaces any existing values.
func (params Values) Set(key, value string) {
	params.Values.Set(key, value)
}

// The parameters pass to the gomws api.
type Parameters map[string]interface{}

// Merge merge the target Parameters to current Parameters.
func (params Parameters) Merge(parameters Parameters) Parameters {
	for key, val := range parameters {
		params[key] = val
	}
	return params
}

// StructureKeys structure the keys for the parameters.
// The basekey is the key for the current parameters.
// keys are the string to augument the base key.
//
// If the value is a slice, then additonal index will be used to augument the keys.
// Ex:
// p := Parameters{
// 		"slice": []string{"a", "b"},
// }
// p.StructureKeys("arrayFiled", "fields")
// -> Parameters{
// 		"slice.fields.1": "a",
// 		"slice.fields.2": "a",
// }
//
// If the value is another Parameters, the Parameters' keys will used to augument.
// the keys
// Ex:
// p := Parameters{
// 		"params": Parameters{"a": 1, "b": 2},
// }
// p.StructureKeys("params", "fields")
// -> Parameters{
// 		"params.fields.a": 1,
// 		"params.fields.b": 2,
// }
//
// If the value is other type, other keys will be used to structure the keys.
func (params Parameters) StructureKeys(baseKey string, keys ...string) Parameters {
	data, ok := params[baseKey]
	if !ok {
		//TODO log
		return params
	}

	delete(params, baseKey)

	switch reflect.TypeOf(data).Kind() {
	default:
		key := formatParameterKey(baseKey, keys...)
		params[key] = data
	case reflect.Map:
		valueMap := reflect.ValueOf(data)
		for _, k := range valueMap.MapKeys() {
			nkeys := append(keys, k.String())
			key := formatParameterKey(baseKey, nkeys...)
			params[key] = valueMap.MapIndex(k).Interface()
		}
	case reflect.Slice:
		valueSlice := reflect.ValueOf(data)
		for i := 0; i < valueSlice.Len(); i++ {
			nkeys := append(keys, strconv.Itoa(i+1))
			key := formatParameterKey(baseKey, nkeys...)
			params[key] = valueSlice.Index(i).Interface()
		}
	}

	return params
}

// Normalize convert all the values to string, if a value can't not
// convert to string, an error will be returned.
func (params Parameters) Normalize() (Values, error) {
	nParams := NewValues()
	var stringVal string
	for key, val := range params {
		switch t := val.(type) {
		default:
			err := fmt.Errorf("Unexpected type %T", t)
			return nParams, err
		case bool:
			stringVal = strconv.FormatBool(val.(bool))
		case int:
			stringVal = strconv.Itoa(val.(int))
		case float32:
			stringVal = strconv.FormatFloat(float64(val.(float32)), 'f', 2, 32)
		case float64:
			stringVal = strconv.FormatFloat(val.(float64), 'f', 2, 64)
		case string:
			stringVal = val.(string)
		}
		nParams.Set(key, stringVal)
	}
	return nParams, nil
}
