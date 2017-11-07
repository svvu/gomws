package mws

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// formatParameterKey combine the base key and the augument keys by '.'.
func formatParameterKey(baseKey string, keys ...string) string {
	paramKeys := append([]string{baseKey}, keys...)
	return strings.Join(paramKeys, ".")
}

// Parameters is the parameters pass to the gomws API.
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
// 	p := Parameters{
// 		"slice": []string{"a", "b"},
// 	}
// 	p.StructureKeys("slice", "fields")
// result:
// 	Parameters{
// 		"slice.fields.1": "a",
// 		"slice.fields.2": "a",
// 	}
//
// If the value is another Parameters, the Parameters' keys will used to augument.
// the keys
// Ex:
// 	p := Parameters{
// 		"params": Parameters{"a": 1, "b": 2},
// 	}
// 	p.StructureKeys("params", "fields")
// result:
// 	Parameters{
// 		"params.fields.a": 1,
// 		"params.fields.b": 2,
//	}
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
// convert to string, an error will be returned.
// Float will round to 2 decimal precision.
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
		case time.Time:
			isoT := val.(time.Time)
			stringVal = isoT.UTC().Format(time.RFC3339) // "2006-01-02T15:04:05Z07:00"
		}
		nParams.Set(key, stringVal)
	}
	return nParams, nil
}

// OptionalParams get the values from the pass in parameters.
// Only values for keys that are accepted will be returned.
//
// Note: The keys returned will be in title case.
//
// If the key appear in mulit parameters, later one will override the previous.
// Ex:
// 		ps := []Parameters{
// 			{"key1": "value1", "key2": "value2"},
// 			{"key1": "newValue1", "key3": "value3"},
// 		}
// 		acceptKeys := []string{"key1", "key2"}
// 		resultParams := OptionalParams(acceptKeys, ps)
// result:
// 		resultParams -> {"Key1": "newValue1", "Key2": "value2"}
func OptionalParams(acceptKeys []string, ops []Parameters) Parameters {
	param := Parameters{}
	op := Parameters{}

	if len(ops) == 0 {
		return param
	}

	for _, p := range ops {
		op.Merge(p)
	}

	for _, key := range acceptKeys {
		value, ok := op[key]
		if ok {
			param[strings.Title(key)] = value
			delete(op, key)
		}
	}

	return param
}
