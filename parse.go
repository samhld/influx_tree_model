package main

import (
	"regexp"
	"strings"
)

type RuleTokenizer struct {
<<<<<<< HEAD
	// line string
	re *regexp.Regexp
=======
	line string
	re   *regexp.Regexp
>>>>>>> ddf0be69c869da6c5baea6197110d12fbd2744f0
}

type RuleMap struct {
	words []Word
	ops   []Op
}

type Word struct {
	text  string
	index int
}

type Op struct {
	text  string
	index int
}

func NewRuleTokenizer() *RuleTokenizer {
	return &RuleTokenizer{
		regexp.MustCompile(`(?P<words>[a-z A-Z _]+)|(?P<ops>[|>\s])`),
	}
}
func (t *RuleTokenizer) Tokenize(rule string) *RuleMap {
	matches := t.re.FindAllStringSubmatch(rule, -1)
	ruleMap := &RuleMap{}
	for i, match := range matches {
		if match[1] != "" { // 2nd position is a 'word'
			word := Word{match[1], i}
			ruleMap.words = append(ruleMap.words, word)
		} else {
			op := Op{match[2], i} // 3rd positioin is an 'op'
			ruleMap.ops = append(ruleMap.ops, op)
		}
	}
	return ruleMap
}

func MeasIndex(line string) {

}
func ParseMeas(point string) string {
	substrings := strings.Split(point, ",")
	return substrings[0]
=======

}

type Line struct {
	Measurement Measurement
	Tags        []Tag
	Fields      []Field
}

type Measurement string

type Tag struct {
	Key   string
	Value string
	Index int
}

type Field struct {
	Key   string
	Value interface{}
	Index int
}
type Tags map[string]string
type Fields map[string]interface{}

<<<<<<< HEAD
=======
func NewRuleTokenizer(line string) *RuleTokenizer {
	return &RuleTokenizer{
		line,
		regexp.MustCompile(`(?P<words>[a-z A-Z _]+)|(?P<ops>[|>\s])`),
	}
}
func (t *RuleTokenizer) Tokenize(rule string) *RuleMap {
	matches := t.re.FindAllStringSubmatch(rule, -1)
	ruleMap := &RuleMap{}
	for i, match := range matches {
		if match[1] != "" { // 2nd position is a 'word'
			word := Word{match[1], i}
			ruleMap.words = append(ruleMap.words, word)
		} else {
			op := Op{match[2], i}
			ruleMap.ops = append(ruleMap.ops, op)
		}
	}
	return ruleMap
}

func MeasIndex(line string) {

}
func ParseMeas(point string) string {
	substrings := strings.Split(point, ",")
	return substrings[0]
}

>>>>>>> ddf0be69c869da6c5baea6197110d12fbd2744f0
// func (t Tags) Keys() []string {
// 	var keys []string
// 	for k, _ := range t {
// 		keys = append(keys, k)
// 	}
// 	return keys
// }

// func (t Tags) Values() []string {
// 	var vals []string
// 	for _, v := range t {
// 		vals = append(vals, v)
// 	}
// 	return vals
// }

// func ParseTags(tagsMap Tags) ([]string, []string) {
// 	tagsMap := parseTagsToMap(tags)

// 	var keys []string
// 	var vals []string

// 	for k, v := range tagsMap {
// 		keys = append(keys, k)
// 		vals = append(vals, v)
// 	}
// 	return keys, vals
// }

// func parseTagsToMap(point string) Tags {
// 	tagsString := strings.Split(point, " ")[0]
// 	newTagsString := strings.TrimPrefix(tagsString, ParseMeas(tagsString)+",")
// 	var tags Tags
// 	commaSep := strings.Split(newTagsString, ",")
// 	for _, tag := range commaSep {
// 		tagSplit := strings.Split(tag, "=")
// 		tags[tagSplit[0]] = tagSplit[1]
// 	}

// 	return tags

// }
