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
	keyValCountMap map[string]int64
	// keyValsMap       make(map[string][]string)
	fluxGetValCounts string
	fluxGetVals      string
}

func NewMeasurementAPI(queryAPI api.QueryAPI, bucket, measurement string) *MeasurementAPI {
	mAPI := &MeasurementAPI{}
	mAPI.api = queryAPI
	mAPI.bucket = bucket
	mAPI.measurement = measurement
	mAPI.keyValCountMap = make(map[string]int64)
	flux := readFlux("tag_key_value_counts_by_measurement.flux")
	mAPI.fluxGetValCounts = fmt.Sprintf(flux, bucket, measurement)
	mAPI.fluxGetVals = readFlux("all_tag_values_by_tag.flux") // store injectable flux

	return mAPI

	// mAPI.fluxGetVals := readFlux("all_tag_value_by_tag.flux")
}

func (m *MeasurementAPI) getTagKeyValueCounts() map[string]int64 {
	result, err := m.api.Query(context.Background(), m.fluxGetValCounts)
	checkQueryError(err)

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
	m.fluxGetVals = fmt.Sprintf(m.fluxGetVals, m.bucket, tag)
	result, err := m.api.Query(context.Background(), m.fluxGetVals)
	checkQueryError(err)

	var vals []string
	for result.Next() {
		checkQueryError(result.Err())
		strVal := fmt.Sprintf("%v", result.Record().Value())
		vals = append(vals, strVal)
	}
	return vals
}

func checkQueryError(err error) {
	if err != nil {
		fmt.Printf("Error querying for tag values: %v", err)
	}
}

func readFlux(fileName string) string {
	byteFlux, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("error reading file: %s, err: %v", fileName, err)
	}
	return string(byteFlux)
}
