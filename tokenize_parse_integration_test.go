package main

import (
	"os"
	"testing"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func TestTokenizeParse(t *testing.T) {
	measAPI, tokenizedRule := setup()
	field := "value"
	stub := &Stub{
		[]string{"app", "region"},
		[]string{"v1", "v2", "v3"},
		[]string{"v1", "v2"},
		3,
		2,
	}
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
	measAPI.setKeyValsMap()

	return measAPI, tokenizedRule
}
