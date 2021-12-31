package main

import (
	"fmt"
)

func (m *MeasurementAPI) setKeyValsMap() {
	tagKeys := m.getTagKeys()
	for _, key := range tagKeys {
		keyVals := m.getTagKeyValues(key)
		m.keyValsMap[key] = keyVals
	}
}

func mapKeysToValues(tagKeys []string, allVals [][]string) map[string][]string {
	// assumes tagKeys and allValls have relating indices
	m := make(map[string][]string)
	for i, key := range tagKeys {
		m[key] = allVals[i]
	}
	return m
}

func checkQueryError(err error) {
	if err != nil {
		fmt.Printf("Error querying: %v", err)
	}
}
