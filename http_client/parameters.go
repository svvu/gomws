package mwsHttpClient

import (
	"net/url"
	"strconv"
	"strings"
)

type NormalizedParameters url.Values

func (params NormalizedParameters) UrlEncode() string {
	return strings.Replace(params.Encode(), "+", "%20", -1)
}

func formatParameterKey(keys ...string) string {
	return strings.Join(keys, ".")
}

type Parameters map[string]interface{}

func (params *Parameters) Merge(parameters Parameters) {
	for key, val := range parameters {
		params[key] = val
	}
}

func (params *Parameters) StructureKeys(baseKey string, keys ...string) (params Parameters) {
	data, ok := params[baseKey]
	if !ok {
		//TODO log
		return
	}

	delete(params, baseKey)

	switch t := data.(type) {
	default:
		key = formatParameterKey(baseKey, keys...)
		params[key] = val
	case []string:
		for i, val := range data {
			nkeys := append(keys, strconv.Itoa(i))
			key = formatParameterKey(baseKey, nkeys...)
			params[key] = val
		}
	case Parameters:
		for k, val := range data {
			nkeys := append(keys, k)
			key = formatParameterKey(baseKey, nkeys...)
			params[key] = val
		}
	}
	return
}

func (params *Parameters) ToNormalizedParameters() (NormalizedParameters, error) {
	sparams := NormalizedParameters{}
	for key, val := range params {
		switch t := val.(type) {
		default:
			err := fmt.Errorf("Unexpected type %T", t)
			return nil, err
		case bool:
			val = strconv.FormatBool(val)
		case int:
			val = strconv.Itoa(val)
		case float32:
			val = strconv.FormatFloat(val, 'f', 2, 32)
		case float64:
			val = strconv.FormatFloat(val, 'f', 2, 64)
		case string:
			val = val
		}
		sparams.Add(key, val)
	}
	return sparams, nil
}
