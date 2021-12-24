package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("map keys to values", func(t *testing.T) {
		keys := []string{"k1", "k2"}
		k1v := []string{"v1", "v2", "v3"}
		k2v := []string{"v1", "v2"}
		allVals := [][]string{k1v, k2v}

		gotMap := mapKeysToValues()
		wantMap := map[string][]string{
			"k1": []string{"v1", "v2", "v3"},
			"k2": []string{"v1", "v2"},
		}

		assertEqual(t, gotMap, wantMap)
	})
}
