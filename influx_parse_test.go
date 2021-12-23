package main

import (
	"testing"
)

func TestOrdering(t *testing.T) {
	cards := []int{5, 6, 1, 2, 7, 3}
	sorted(cards)
	wantSorted := []int{1, 2, 3, 5, 6, 7}
	assertEqual(t, cards, wantSorted)
}
