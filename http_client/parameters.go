package mwsHttpClient

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type NormalizedParameters struct {
	url.Values
}

// Constructor for NormalizedParameters
func NewNormalizedParameters() NormalizedParameters {
	return NormalizedParameters{url.Values{}}
}

func (params NormalizedParameters) Encode() string {
	return strings.Replace(params.Values.Encode(), "+", "%20", -1)
}

func (params NormalizedParameters) Set(key, value string) {
	params.Values.Set(key, value)
}

// formatParameterKey combine the base key and the augument keys by '.'
func formatParameterKey(baseKey string, keys ...string) string {
	return baseKey + "." + strings.Join(keys, ".")
}

// The parameters pass to the gmws api
type Parameters map[string]interface{}

// Merge merge the target Parameters to current Parameters
func (params Parameters) Merge(parameters Parameters) {
	for key, val := range parameters {
		params[key] = val
	}
}

// StructureKeys structure the keys for the parameters
// The basekey is the key for the current parameters
// keys are the string to augument the base key
//
// If the value is a slice, then additonal index will be used to augument the keys
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
// If the value is another Parameters, the Parameters' keys will used to augument
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
// If the value is other type, other keys will be used to structure the keys
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

	// switch data.(type) {
	// default:
	// 	key := formatParameterKey(baseKey, keys...)
	// 	params[key] = data
	// case Parameters:
	// 	for k, val := range data.(Parameters) {
	// 		nkeys := append(keys, k)
	// 		key := formatParameterKey(baseKey, nkeys...)
	// 		params[key] = val
	// 	}
	// case []string:
	// 	for i, val := range data.([]string) {
	// 		nkeys := append(keys, strconv.Itoa(i+1))
	// 		key := formatParameterKey(baseKey, nkeys...)
	// 		params[key] = val
	// 	}
	// case map[string]string:
	// 	for k, val := range data.(map[string]string) {
	// 		nkeys := append(keys, k)
	// 		key := formatParameterKey(baseKey, nkeys...)
	// 		params[key] = val
	// 	}
	// }
	return params
}

// NormalizeParameters convert all the values to string, if a value can't not
// convert to string, an error will be returned
func (params Parameters) NormalizeParameters() (NormalizedParameters, error) {
	nParams := NewNormalizedParameters()
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
