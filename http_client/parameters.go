package mwsHttpClient

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type NormalizedParameters struct {
	url.Values
}

func NewNormalizedParameters() NormalizedParameters {
	return NormalizedParameters{url.Values{}}
}

func (params NormalizedParameters) Encode() string {
	return strings.Replace(params.Values.Encode(), "+", "%20", -1)
}

func (params NormalizedParameters) Set(key, value string) {
	params.Values.Set(key, value)
}

func formatParameterKey(baseKey string, keys ...string) string {
	return baseKey + "." + strings.Join(keys, ".")
}

type Parameters map[string]interface{}

func (params Parameters) Merge(parameters Parameters) {
	for key, val := range parameters {
		params[key] = val
	}
}

func (params Parameters) StructureKeys(baseKey string, keys ...string) Parameters {
	data, ok := params[baseKey]
	if !ok {
		//TODO log
		return params
	}

	delete(params, baseKey)

	switch data.(type) {
	default:
		key := formatParameterKey(baseKey, keys...)
		params[key] = data
	case []string:
		for i, val := range data.([]string) {
			nkeys := append(keys, strconv.Itoa(i+1))
			key := formatParameterKey(baseKey, nkeys...)
			params[key] = val
		}
	case Parameters:
		for k, val := range data.(Parameters) {
			nkeys := append(keys, k)
			key := formatParameterKey(baseKey, nkeys...)
			params[key] = val
		}
	}
	return params
}

func (params Parameters) ToNormalizedParameters() (NormalizedParameters, error) {
	sparams := NewNormalizedParameters()
	var stringVal string
	for key, val := range params {
		switch t := val.(type) {
		default:
			err := fmt.Errorf("Unexpected type %T", t)
			return sparams, err
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
		sparams.Set(key, stringVal)
	}
	return sparams, nil
}
