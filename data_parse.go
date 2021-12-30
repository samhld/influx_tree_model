package main

import (
	"context"
	"fmt"
	"os"

	api "github.com/influxdata/influxdb-client-go/v2/api"
)

type dataGetter interface {
	getTagKeyValues(flux string) []string
	getTagKeyValueCounts(flux string) map[string]int64
}

type MeasurementAPI struct {
	api         api.QueryAPI
	bucket      string
	measurement string
	// tagKeys          []string
	keyValCountMap   map[string]int64 // map so not sorted -- must sort ad hoc to use
	keyValsMap       map[string][]string
	fluxGetValCounts string
	fluxGetVals      string
	fluxGetKeys      string
}

func NewMeasurementAPI(queryAPI api.QueryAPI, bucket, measurement string) *MeasurementAPI {
	mAPI := &MeasurementAPI{}
	mAPI.api = queryAPI
	mAPI.bucket = bucket
	mAPI.measurement = measurement
	mAPI.keyValCountMap = make(map[string]int64)
	mAPI.keyValsMap = make(map[string][]string)
	flux := readFlux("flux/tag_key_value_counts_by_measurement.flux")
	mAPI.fluxGetValCounts = fmt.Sprintf(flux, bucket, measurement)
	mAPI.fluxGetVals = readFlux("flux/all_tag_values_by_tag.flux") // store injectable flux
	flux = readFlux("flux/tag_keys_by_measurement.flux")
	mAPI.fluxGetKeys = fmt.Sprintf(flux, bucket, measurement)

	return mAPI
}

func (m *MeasurementAPI) setKeyValsMap() {
	tagKeys := m.getTagKeys()
	for _, key := range tagKeys {
		keyVals := m.getTagKeyValues(key)
		m.keyValsMap[key] = keyVals
	}
}

func (m *MeasurementAPI) getTagKeyValueCounts() map[string]int64 {
	result, err := m.api.Query(context.Background(), m.fluxGetValCounts)
	if err != nil {
		fmt.Printf("error querying for tag key value counts: %v", err)
	}

	for result.Next() {
		record := result.Record()
		tag := fmt.Sprintf("%v", record.ValueByKey("tag")) //"tag" is a column injected via the Flux query
		val := record.Value().(int64)
		m.keyValCountMap[tag] = val
	}

	return m.keyValCountMap
}

func mapKeysToValues(tagKeys []string, allVals [][]string) map[string][]string {
	// assumes tagKeys and allValls have relating indices
	m := make(map[string][]string)
	for i, key := range tagKeys {
		m[key] = allVals[i]
	}
	return m
}

func (m *MeasurementAPI) getTagKeyValues(tag string) []string {
	fluxGetVals := fmt.Sprintf(m.fluxGetVals, m.bucket, tag)
	result, err := m.api.Query(context.Background(), fluxGetVals)
	if err != nil {
		fmt.Printf("error querying for tag key values: %v", err)
	}
	// checkQueryError(err)
	var vals []string
	for result.Next() {
		checkQueryError(result.Err())
		strVal := fmt.Sprintf("%v", result.Record().Value())
		vals = append(vals, strVal)
	}
	return vals
}

func (m *MeasurementAPI) getTagKeys() []string {
	result, err := m.api.Query(context.Background(), m.fluxGetKeys)
	if err != nil {
		fmt.Printf("error querying for tag keys: %v", err)
	}
	// checkQueryError(err)
	var keys []string
	for result.Next() {
		key := fmt.Sprintf("%v", result.Record().Value())
		keys = append(keys, key)
	}

	return keys
}

func checkQueryError(err error) {
	if err != nil {
		fmt.Printf("Error querying: %v", err)
	}
}

func readFlux(fileName string) string {
	byteFlux, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("error reading file: %s, err: %v", fileName, err)
	}
	return string(byteFlux)
}
