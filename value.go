package main

type Value struct {
	text   string
	tier   int
	parent *Node
	child  *Node
}

func NewValue(text string, parent *Node, tier int) *Value {
	return &Value{text, tier, parent, nil}
}

func (v *Value) Tier() int {
	return v.tier
}

func (v *Value) Text() string {
	return v.text
}
