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

	byteFlux, err := os.ReadFile("tag_key_value_counts_by_measurement.flux")
	if err != nil {
		panic(err)
	}
	strFlux := fmt.Sprintf(string(byteFlux), bucket, measurement)
	client := influxdb2.NewClient(url, token)
	queryAPI := client.QueryAPI(org)

	keyValMap := getTagKeyValueCounts(queryAPI, strFlux, bucket, measurement)

	sorted := sortByCardinality(keyValMap)

	fmt.Println(sorted)
}
