package main

import (
	"context"
	"fmt"

	api "github.com/influxdata/influxdb-client-go/v2/api"
)

func readResults(result *api.QueryTableResult) ([]interface{}, error) {
	var resultList []interface{}
	// Iterate over query response
	for result.Next() {
		resultList = append(resultList, result.Record().Value())
	}
	// Check for an error
	if result.Err() != nil {
		return nil, result.Err()
	}
	return resultList, nil
}

func getBucketTagKeys(queryAPI api.QueryAPI, bucket string) (*api.QueryTableResult, error) {
	flux := fmt.Sprintf(`import "influxdata/influxdb/schema"
						schema.tagKeys(bucket: "%s")`,
		bucket)

	return queryAPI.Query(context.Background(), flux)
}
