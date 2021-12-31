package main

import (
	"context"
	"fmt"

	api "github.com/influxdata/influxdb-client-go/v2/api"
)

const (
	fluxGetKeys      = "flux/tag_keys_by_measurement.flux"
	fluxGetValCounts = "flux/tag_key_value_counts_by_measurement.flux"
	fluxGetVals      = "flux/all_tag_values_by_tag.flux"
)

type dataGetter interface {
	getTagKeys() []string
	getTagKeyValues(flux string) []string
	getTagKeyValueCounts(flux string) map[string]int64
}

type MeasurementAPI struct {
	api            api.QueryAPI
	bucket         string
	measurement    string
	tagKeys        []string
	tags           map[string][]Tag
	keyValCountMap map[string]int64 // map so not sorted -- must sort ad hoc to use
	keyValsMap     map[string][]string
}

func NewMeasurementAPI(queryAPI api.QueryAPI, bucket, measurement string) *MeasurementAPI {
	mAPI := &MeasurementAPI{}
	mAPI.api = queryAPI
	mAPI.bucket = bucket
	mAPI.measurement = measurement
	mAPI.keyValCountMap = make(map[string]int64)
	mAPI.keyValsMap = make(map[string][]string)
	mAPI.setTagKeys()
	mAPI.tags = make(map[string][]Tag)
	for _, key := range mAPI.tagKeys {
		mAPI.setTagKeyValues(key)
	}
	return mAPI
}

func (m *MeasurementAPI) getTagKeys() []string {
	if m.tagKeys != nil {
		return m.tagKeys
	} else {
		m.setTagKeys()
		return m.tagKeys
	}
}

func (m *MeasurementAPI) setTagKeys() {
	flux := fmt.Sprintf(readFlux(fluxGetKeys), m.bucket, m.measurement)
	result, err := m.api.Query(context.Background(), flux)
	if err != nil {
		fmt.Printf("error querying for tag keys: %v", err)
	}
	var keys []string
	for result.Next() {
		key := fmt.Sprintf("%v", result.Record().Value())
		if key == "_measurement" || key == "_field" {
			continue
		} else {
			keys = append(keys, key)
		}
	}
	m.tagKeys = keys
}

func (m *MeasurementAPI) getTagKeyValues(key string) []string {
	if tags, ok := m.tags[key]; ok {
		var tagVals []string
		for _, tag := range tags {
			tagVals = append(tagVals, tag.value.text)
		}
		return tagVals
	} else {
		fmt.Errorf("err: key %s doesn't exist in tag set", key)
	}

	return nil
}

func (m *MeasurementAPI) getTagsByKey(key string) []Tag {
	if tags, ok := m.tags[key]; ok {
		return tags
	} else {
		fmt.Errorf("err: key %s doesn't exist in tag set", key)
	}
	return nil
}

func (m *MeasurementAPI) setTagKeyValues(key string) {
	flux := fmt.Sprintf(readFlux(fluxGetVals), m.bucket, key)
	result, err := m.api.Query(context.Background(), flux)
	if err != nil {
		fmt.Printf("error querying for tag key values: %v", err)
	}
	var tags []Tag
	for result.Next() {
		// checkQueryError(result.Err())
		tag := Tag{}
		strVal := fmt.Sprintf("%v", result.Record().Value())
		tag.key = Key{key, 0, nil, nil, nil}
		tag.value = Value{strVal, 0, nil, nil}
		tags = append(tags, tag)
	}
	m.tags[key] = tags
}

func (m *MeasurementAPI) getTagKeyValueCounts() map[string]int64 {
	flux := fmt.Sprintf(readFlux(fluxGetValCounts), m.bucket, m.measurement)
	result, err := m.api.Query(context.Background(), flux)
	if err != nil {
		fmt.Printf("error querying for tag key value counts: %v", err)
	}

	for result.Next() {
		record := result.Record()
		tag := fmt.Sprintf("%v", record.ValueByKey("tag")) //"tag" is a column injected via the Flux query
		if tag == "_measurement" || tag == "_field" {
			continue
		}
		val := record.Value().(int64)
		m.keyValCountMap[tag] = val
	}
	return m.keyValCountMap
}
