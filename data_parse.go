package main

import (
	"context"
	"fmt"

	api "github.com/influxdata/influxdb-client-go/v2/api"
)

func getTagKeyValueCounts(queryAPI api.QueryAPI, flux, bucket, measurement string) map[string]int64 {
	result, err := queryAPI.Query(context.Background(), flux)
	if err != nil {
		fmt.Printf("Error querying for tag values: %v", err)
	}
	keyValMap := make(map[string]int64)
	for result.Next() {
		record := result.Record()
		tag := fmt.Sprintf("%v", record.ValueByKey("tag")) //"tag" is a column injected via the Flux query
		val := record.Value().(int64)
		keyValMap[tag] = val
		// resultList = append(resultList, recordString)
	}

	fmt.Println(keyValMap)
}
