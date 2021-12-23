package main

import (
	"context"
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

	fluxGetTagValuesPerKey := fmt.Sprintf(`import "influxdata/influxdb/schema"
											schema.tagValues(bucket: "%s", tag: "%s")
											|> count()`,
		bucket,
		tag)

	tagKeysResult, err := getBucketTagKeys(queryAPI, bucket)

	keyResultList, _ := readResults(tagKeysResult)
	fmt.Println(keyResultList)

	// if err == nil {
	// 	// Iterate over query response
	// 	for tagKeysResult.Next() {
	// 		fmt.Printf("key result: %q", tagKeysResult)
	// 	}
	// } else {
	// 	panic(err)
	// }

	// fmt.Printf("tag keys: %q\n", tagKeys)

	result, err := queryAPI.Query(context.Background(), fluxGetTagValuesPerKey)
	fmt.Printf("result: %v\n", result)
	if err == nil {
		// Iterate over query response
		for result.Next() {
			fmt.Printf("value: %v\n", result.Record().Value())
		}
		// Check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
	// Ensures background processes finishes
	client.Close()
}
