package utils

import (
	"sort"
)

type ParameterIterator struct {
	keys         []string
	data         Parameters
	currentIndex int
}

func NewParameterIterator(params Parameters) *ParameterIterator {
	keys := make([]string, len(params))

	i := 0
	for k := range params {
		keys[i] = k
		i += 1
	}

	return &ParameterIterator{data: params, keys: keys, currentIndex: 0}
}

func (itera *ParameterIterator) Sort() {
	sort.Strings(itera.keys)
}

func (itera *ParameterIterator) HasNext() bool {
	return currentIndex < len(keys)
}

func (itera *ParameterIterator) Next() (string, string) {
	if !itera.HasNext() {
		return nil, nil
	}

	key := keys[currentIndex]

	if val, ok := itera.data[key]; ok {
		return key, val
	} else {
		return key, nil
	}
}
