package main

import "testing"

func TestParse(t *testing.T) {
	tagKeys := []string{"key1", "key2", "key3"}
	key1Values := []string{"val1", "val2", "val3"}
	key2Values := []string{"val1", "val2"}
	key3Values := []string{"val1", "val2", "val3", "val4"}

	t.Run("test creating map of keys to values", func(t *testing.T) {
		wantMap := map[string][]string{
			"key1": []string{"val1", "val2", "val3"},
			"key2": []string{"val1", "val2"},
			"key3": []string{"val1", "val2", "val3", "val4"},
		}

		gotMap := createKeyValuesMap(tagKeys)

		assertEqual(t, gotMap, wantMap)
	})
	// t.Run("test ordering tags by cardinality", func(t *testing.T) {

	// }
}
