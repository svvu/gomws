package gmws

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/svvu/gomws/mwsHttps"
)

// formatParameterKey combine the base key and the augument keys by '.'.
func formatParameterKey(baseKey string, keys ...string) string {
	paramKeys := append([]string{baseKey}, keys...)
	return strings.Join(paramKeys, ".")
}

// Parameters is the parameters pass to the gomws api.
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
// 	convert to string, an error will be returned.
// Float will round to 2 decimal precision.
func (params Parameters) Normalize() (mwsHttps.Values, error) {
	nParams := mwsHttps.NewValues()
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
