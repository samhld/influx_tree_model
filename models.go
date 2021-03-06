package main

import "fmt"

type Tree []Branch

type Branch struct {
	measurement Measurement
	tags        []Tag // appended to in order as the source will be ordered
	field       Field
}

type Tiers map[int]interface{}

// func (tr *Tree) createTiers(numTiers int)

type Node interface {
	Tier() int
	Text() string
}

// Measurement struct impls Node interface
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

type Tag struct {
	key   Key
	value Value
}

func (t *Tag) String() string {
	return fmt.Sprintf("%s: %s\n", t.key.Text(), t.value.Text())
}

func (t *Tag) Tier() int {
	return t.key.tier
}

func (t *Tag) Text() string {
	return t.key.text
}

// Key struct impls Node interface
type Key struct {
	text   string
	tier   int
	parent *Node
	child  *Node
}

func NewKey(text string, parent *Node, tier int) *Key {
	return &Key{text, tier, parent, nil}
}

func (k *Key) Tier() int {
	return k.tier
}

func (k *Key) Text() string {
	return k.text
}

// Field struct impls Node interface
type Field struct {
	text string
	tier int
}

func NewField(text string, parent *Node, tier int) *Key {
	return &Key{text, tier, parent, nil}
}

func (f *Field) Tier() int {
	return f.tier
}

func (f *Field) Text() string {
	return f.text
}

// Value struct impls Node interface
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
