package main

import (
	"reflect"
	"testing"
)

func TestTokenizationRule(t *testing.T) {
	// line := "stats,region=us-west,app=cart count=12"
	ruleTokenizer := NewRuleTokenizer()
	t.Run("test parsing rule to map with indexed tokens", func(t *testing.T) {
		rule := "MEASUREMENT>region>app>FIELD"
		wantedRuleMap := &RuleMap{
			[]Word{
				{"MEASUREMENT", 0},
				{"region", 2},
				{"app", 4},
				{"FIELD", 6},
			},
			[]Op{
				{">", 1},
				{">", 3},
				{">", 5},
			},
		}
		gotRuleMap := ruleTokenizer.Tokenize(rule)
		if !reflect.DeepEqual(gotRuleMap, wantedRuleMap) {
			t.Errorf("got %q want %q", gotRuleMap, wantedRuleMap)
		}

	})

	// t.Run("measurement", func(t *testing.T) {
	// 	gotMeas := ParseMeas(point)
	// 	assertEqualStrings(t, gotMeas, wantedMeas)
	// })
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

// func assertEqualStrings(t testing.TB, got, want string) {
// 	t.Helper()
// 	if got != want {
// 		t.Errorf("got %s, want %s", got, want)
// 	}
// }
