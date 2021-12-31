package main

import (
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// This is an API server that proxies InfluxDB and holds (in-mem) a hierarchical structure of Line Protocl elements

// func Serve(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// }

type RuleServer struct {
	rule          string
	measurement   string
	measAPI       *MeasurementAPI
	tokenizedRule *TokenizedRule
}

func makeTreeFromRule(rule, measurement string) Tree {
	ruleServer := NewRuleServer(rule, measurement)
	return MapTokensToData(ruleServer.measAPI, ruleServer.tokenizedRule)
}

func makeTreeNoRule(measurement string) Tree {
	ruleServer := NewRuleServer("", measurement)
	keyValCountMap := ruleServer.measAPI.getTagKeyValueCounts()
	sorted := sortByCardinality(keyValCountMap)
	tree := make(Tree)
	tree[0] = &Measurement{measurement, 0}
	for i, key := range sorted {
		vals := ruleServer.measAPI.getTagKeyValues(key)
		tier := i + 1
		tree[tier] = &Key{key, tier, vals, nil, nil}
	}
	tier := len(sorted) + 1
	tree[tier] = &Field{"FIELD", tier}

	return tree
}

func NewRuleServer(rule, measurement string) *RuleServer {
	ruleServer := &RuleServer{}
	tokenizer := NewRuleTokenizer()
	tokenizedRule := tokenizer.Tokenize(rule)
	ruleServer.tokenizedRule = tokenizedRule
	// make Influx client and wrap in MeasurementAPI
	bucket := "test"
	org := os.Getenv("INFLUX_REMOTE_ORG")
	token := os.Getenv("INFLUX_REMOTE_TOKEN")
	url := os.Getenv("INFLUX_REMOTE_HOST")
	client := influxdb2.NewClient(url, token)
	queryAPI := client.QueryAPI(org)
	ruleServer.measAPI = NewMeasurementAPI(queryAPI, bucket, measurement)

	return ruleServer
}
