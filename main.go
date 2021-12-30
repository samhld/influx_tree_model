package main

import (
	"fmt"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	measAPI, tokenizedRule := setup()
	// field := "value"
	// stub := &Stub{
	// 	[]string{"app", "region"},
	// 	[]string{"v1", "v2", "v3"},
	// 	[]string{"v1", "v2"},
	// 	3,
	// 	2,
	// }
	measAPI.setKeyValsMap()
	// keyValsMap := measAPI.keyValsMap
	tree := make(Tree)
	for i, word := range tokenizedRule.words {
		switch word.text {
		case "MEASUREMENT":
			tree[i] = &Measurement{measAPI.measurement, i}
		case "FIELD":
			tree[i] = &Field{"FIELD", i}
		default:
			vals := measAPI.getTagKeyValues(word.text)
			fmt.Printf("vals: %q\n", vals)
			tree[i] = &Key{word.text, i, vals, nil, nil}
		}
	}

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
	measAPI.setKeyValsMap()

	return measAPI, tokenizedRule
}
