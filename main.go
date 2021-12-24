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
	keyValMap := measAPI.getTagKeyValueCounts()
	sorted := sortByCardinality(keyValMap)
	fmt.Println(sorted)

	tag := "tag1"
	keyVals := measAPI.getTagKeyValues(tag)
	fmt.Println(keyVals)
}
