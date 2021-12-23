package main

import (
	"sort"
)

func sortByCardinality(cardMap map[string]int) []string {
	// var cardsToBeSorted []int
	pairs := createTagCardinalityPairs(cardMap)
	sort.Sort(pairs)

	var tagsSorted []string
	for _, pair := range pairs {
		tagsSorted = append(tagsSorted, pair.Tag)
	}

	return tagsSorted
}

func createTagCardinalityPairs(cardMap map[string]int) TagCardPairList {
	var pairs TagCardPairList
	for k, v := range cardMap {
		pair := TagCardPair{k, v}
		pairs = append(pairs, pair)
	}
	return pairs
}

type TagCardPair struct {
	Tag         string
	Cardinality int
}

type TagCardPairList []TagCardPair

// implement Sort interface on TagPairList
// --------------------------------------------
func (p TagCardPairList) Len() int           { return len(p) }
func (p TagCardPairList) Less(i, j int) bool { return p[i].Cardinality < p[j].Cardinality }
func (p TagCardPairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// ----------------------------------------------------
