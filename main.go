package main

func main() {

}

// bucket := "test"
// org := os.Getenv("INFLUX_LOCAL_ORG")
// token := os.Getenv("INFLUX_LOCAL_TOKEN")
// url := os.Getenv("INFLUX_LOCAL_HOST")
// // timeRange := "5m"
// tag := "tag1"

// client := influxdb2.NewClient(url, token)

// queryAPI := client.QueryAPI(org)

// flux := fmt.Sprintf(`import "influxdata/influxdb/schema"
// 					schema.tagValues(bucket: "%s", tag: "%s")`,
// 	bucket,
// 	tag)

// fmt.Println(flux)

// result, err := queryAPI.Query(context.Background(), flux)
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
// // Ensures background processes finishes
// client.Close()

// }
