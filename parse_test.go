package main

import (
	"testing"
)

// type validParse struct {
// 	measurement string
// 	tags        map[string]string
// 	fields      map[string]interface{}
// }

func TestParseTree(t *testing.T) {

}

func TestTokenizationRule(t *testing.T) {
	// rule := "MEASUREMENT,region,app FIELD"
	point := "stats,region=us-west,app=cart count=12"
	wantedMeas := "stats"
	// wantedTagKeys := []string{"region", "app"}
	// wantedTagVals := []string{"us-west", "cart"}
	// wantedFieldKeys := []string{"count"}
	// wantedFieldVals := []inferface{}{12}
	// wantedField := map[string]interface{}{
	// 	"count": 12,
	// }

	t.Run("measurement", func(t *testing.T) {
		gotMeas := ParseMeas(point)
		assertEqualStrings(t, gotMeas, wantedMeas)
	})
	// t.Run("get indices", func(t *testing.T) {
	// 	measIndex := MeasIndex(point)
	// tagsIndex := Rule.TagsIndex()
	// fieldsIndex := Rule.FieldsIndex()
	// })
	// t.Run("tags", func(t *testing.T) {
	// 	tags := ParseTags(point)
	// 	gotTagKeys := ParseKeys(tags)
	// 	gotTagVals := ParseVals(tags)
	// })

}

func assertEqualStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
