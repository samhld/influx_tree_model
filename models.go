package main

type Key struct {
	key    string
	values []int
	isTag bool

}

type Value struct {
	value string
	key   string
	child &Tag|
}
