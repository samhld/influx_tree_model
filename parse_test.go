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
		assertEqualRuleMaps(t, gotRuleMap, wantedRuleMap)

	})
}

func assertEqualRuleMaps(t testing.TB, got, want *RuleMap) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %q\nWant %q\n", got, want)
	}
}
