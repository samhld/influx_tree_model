package main

import "testing"

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
	// wantedTags := map[string]string{
	// 	"region": "us-west",
	// 	"app":    "cart",
	// }
	// wantedField := map[string]interface{}{
	// 	"count": 12,
	// }

	t.Run("parse measurement", func(t *testing.T) {
		gotMeas := ParseMeas(point)
		assertEqualStrings(t, gotMeas, wantedMeas)
	})

}

func assertEqualStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
