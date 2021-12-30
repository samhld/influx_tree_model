package main

import (
	"fmt"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	measAPI, tokenizedRule := setup()
	tree := MapTokensToData(measAPI, tokenizedRule)
	fmt.Printf("%q", tree)
}

func setup() (*MeasurementAPI, *TokenizedRule) {
	measurement := "test"
	rule := "MEASUREMENT>region>app>FIELD"
	tokenizer := NewRuleTokenizer()
	tokenizedRule := tokenizer.Tokenize(rule)
	bucket := "test"
	org := os.Getenv("INFLUX_REMOTE_ORG")
	token := os.Getenv("INFLUX_REMOTE_TOKEN")
	url := os.Getenv("INFLUX_REMOTE_HOST")
	client := influxdb2.NewClient(url, token)
	queryAPI := client.QueryAPI(org)

	measAPI := NewMeasurementAPI(queryAPI, bucket, measurement)

	return measAPI, tokenizedRule
}
