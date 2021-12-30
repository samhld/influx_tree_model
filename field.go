package main

type Field struct {
	text string
	tier int
}

func NewField(text string, parent *Node, tier int) *Key {
	return &Key{text, tier, nil, parent, nil}
}

func (f *Field) Tier() int {
	return f.tier
}

func (f *Field) Text() string {
	return f.text
}
