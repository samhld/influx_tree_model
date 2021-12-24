package main

import (
	"context"
	"fmt"
	"os"

	api "github.com/influxdata/influxdb-client-go/v2/api"
)

func getTagKeyValueCounts(queryAPI api.QueryAPI, flux, bucket, measurement string) map[string]int64 {
	result, err := queryAPI.Query(context.Background(), flux)
	checkQueryError(err)

	keyValCountMap := make(map[string]int64)
	for result.Next() {
		record := result.Record()
		tag := fmt.Sprintf("%v", record.ValueByKey("tag")) //"tag" is a column injected via the Flux query
		val := record.Value().(int64)
		keyValCountMap[tag] = val
		// resultList = append(resultList, recordString)
	}

	return keyValCountMap
}

func mapKeysToValues(tagKeys []string, allVals [][]string) map[string][]string {
	// assumes tagKeys and allValls have relating indices
	m := make(map[string][]string)
	for i, key := range tagKeys {
		m[key] = allVals[i]
	}
	return m
}

func getTagKeyValues(queryAPI api.QueryAPI, flux string) []string {
	result, err := queryAPI.Query(context.Background(), flux)
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
