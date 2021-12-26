package main

import (
	"testing"
)

func TestOrdering(t *testing.T) {
	t.Run("test sorting out-of-order", func(t *testing.T) {
		cardMap := map[string]int64{
			"tag1": 4,
			"tag2": 3,
			"tag3": 5,
			"tag4": 1,
		}
		gotSortedTags := sortByCardinality(cardMap)
		wantSortedTags := []string{"tag4", "tag2", "tag1", "tag3"}

		assertEqual(t, gotSortedTags, wantSortedTags)
	})
	t.Run("test return already-sorted", func(t *testing.T) {
		cardMap := map[string]int64{
			"tag1": 1,
			"tag2": 2,
			"tag3": 3,
			"tag4": 4,
		}
		gotSortedTags := sortByCardinality(cardMap)
		wantSortedTags := []string{"tag1", "tag2", "tag3", "tag4"}

		assertEqual(t, gotSortedTags, wantSortedTags)
	})
}
