package main

type Tree map[int]Node

// func (tr *Tree) createTiers(numTiers int)

type Node interface {
	Tier() int
	Text() string
}

type Measurement struct {
	text string
	tier int
}

func (m *Measurement) Tier() int {
	return m.tier
}

func (m *Measurement) Text() string {
	return m.text
}

// type Tree struct {
// 	route
// }
