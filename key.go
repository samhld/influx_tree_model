package main

type Key struct {
	text   string
	tier   int
	values []string
	parent *Node
	child  *Node
}

func NewKey(text string, parent *Node, tier int) *Key {
	return &Key{text, tier, nil, parent, nil}
}

func (k *Key) Tier() int {
	return k.tier
}

func (k *Key) Text() string {
	return k.text
}
