package main

import (
	"context"
	"fmt"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	bucket := "test"
	org := os.Getenv("INFLUX_LOCAL_ORG")
	token := os.Getenv("INFLUX_LOCAL_TOKEN")
	url := os.Getenv("INFLUX_LOCAL_HOST")
	timeRange := "5m"

	client := influxdb2.NewClient(url, token)

	queryAPI := client.QueryAPI(org)

	flux := fmt.Sprintf(`from(bucket: "%s") |> range(start: -%s)`, bucket, timeRange)
	fmt.Println(flux)

	result, err := queryAPI.Query(context.Background(), flux)
	if err == nil {
		// Iterate over query response
		for result.Next() {
			// Notice when group key has changed
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// Access data
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

	// fmt.Printf("Result: %q", result)

}
