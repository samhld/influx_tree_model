package main

import (
	"fmt"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	bucket := "test"
	org := os.Getenv("INFLUX_REMOTE_ORG")
	token := os.Getenv("INFLUX_REMOTE_TOKEN")
	url := os.Getenv("INFLUX_REMOTE_HOST")
	measurement := "test"
	client := influxdb2.NewClient(url, token)
	queryAPI := client.QueryAPI(org)

	measAPI := NewMeasurementAPI(queryAPI, bucket, measurement)
	measAPI.keyValCountMap = measAPI.getTagKeyValueCounts()
	sorted := sortByCardinality(measAPI.keyValCountMap)
	fmt.Println(sorted)

	tagKeys := measAPI.getTagKeys()
	fmt.Println(tagKeys)

	for _, key := range tagKeys {
		keyVals := measAPI.getTagKeyValues(key)
		measAPI.keyValsMap[key] = keyVals
	}

	fmt.Printf("keyValCountMap: %v\n", measAPI.keyValCountMap)
	fmt.Printf("keyValMap: %v\n", measAPI.keyValsMap)
}
