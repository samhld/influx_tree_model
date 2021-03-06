package main

import (
	"sort"
)

func sortByCardinality(cardMap map[string]int64) []string {
	// var cardsToBeSorted []int
	pairs := createTagCardinalityPairs(cardMap)
	sort.Sort(pairs)
	tagsSorted := createTagListFromPairs(pairs)
	return tagsSorted
}

func createTagCardinalityPairs(cardMap map[string]int64) TagCardPairList {
	var pairs TagCardPairList
	for k, v := range cardMap {
		pair := tagCardPair{k, v}
		pairs = append(pairs, pair)
	}
	return pairs
}

func createTagListFromPairs(pairs TagCardPairList) []string {
	var tagsSorted []string
	for _, pair := range pairs {
		tagsSorted = append(tagsSorted, pair.Tag)
	}
	return tagsSorted
}

type tagCardPair struct {
	Tag         string
	Cardinality int64
}

type TagCardPairList []tagCardPair

// implement Sort interface on TagPairList
// --------------------------------------------
func (p TagCardPairList) Len() int           { return len(p) }
func (p TagCardPairList) Less(i, j int) bool { return p[i].Cardinality < p[j].Cardinality }
func (p TagCardPairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// ----------------------------------------------------
