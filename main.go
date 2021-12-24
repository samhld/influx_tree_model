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

	rawFlux := readFlux("tag_key_value_counts_by_measurement.flux")
	flux := fmt.Sprintf(rawFlux, bucket, measurement)

	keyValMap := getTagKeyValueCounts(queryAPI, flux, bucket, measurement)

	sorted := sortByCardinality(keyValMap)

	fmt.Println(sorted)

	tag := "tag1"
	rawFlux = readFlux("all_tag_values_by_tag.flux")
	flux = fmt.Sprintf(rawFlux, bucket, tag)
	keyVals := getTagKeyValues(queryAPI, flux)
	fmt.Println(keyVals)
}
