package main

import (
	"testing"
)

func TestOrdering(t *testing.T) {
	cardMap := map[string]int{
		"tag1": 4,
		"tag2": 3,
		"tag3": 5,
		"tag4": 1,
	}
	gotSortedTags := sortByCardinality(cardMap)
	wantSortedTags := []string{"tag4", "tag2", "tag1", "tag3"}

	assertEqual(t, gotSortedTags, wantSortedTags)

}
