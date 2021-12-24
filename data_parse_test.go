package main

import (
	"testing"
)

type Stub struct {
	keys []string
	k1Vals1 []string
	k2Vals2 []string
	k1Count int
	k2Count int
}

func (s *Stub) getTagKeyValues(s string) []string {
	s.
}

func TestParse(t *testing.T) {
	stub := &Stub{
		[]string{"k1", "k2"},
		[]string{"v1", "v2", "v3"},
		[]string{"v1", "v2"},
		3,
		2,
	}
	t.Run("map keys to values", func(t *testing.T) {
		api := &API{}
		keys := []string{"k1", "k2"}
		k1v := []string{"v1", "v2", "v3"}
		k2v := []string{"v1", "v2"}
		allVals := [][]string{k1v, k2v}

		gotMap := mapKeysToValues(keys, allVals)
		wantMap := map[string][]string{
			"k1": []string{"v1", "v2", "v3"},
			"k2": []string{"v1", "v2"},
		}

		assertEqual(t, gotMap, wantMap)
	})
}
