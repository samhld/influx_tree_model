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
	// timeRange := "5m"
	tag := "tag1"

	client := influxdb2.NewClient(url, token)

	queryAPI := client.QueryAPI(org)

	tagKeysResult, err := getBucketTagKeys(queryAPI, bucket)

	keyResultList, _ := readResults(tagKeysResult)
	fmt.Println(keyResultList)

	tagValues, err := getTagKeyValues(queryAPI, bucket, tag)
	if err != nil {
		panic(err)
	}
	valuesList, err := readResults(tagValues)
	if err != nil {
		panic(err)
	}

	fmt.Println(valuesList)
	// fmt.Printf("result: %v\n", result)
	// if err == nil {
	// 	// Iterate over query response
	// 	for result.Next() {
	// 		fmt.Printf("value: %v\n", result.Record().Value())
	// 	}
	// 	// Check for an error
	// 	if result.Err() != nil {
	// 		fmt.Printf("query parsing error: %s\n", result.Err().Error())
	// 	}
	// } else {
	// 	panic(err)
	// }
	// Ensures background processes finishes
	client.Close()
}
