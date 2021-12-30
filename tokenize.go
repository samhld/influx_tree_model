package main

import (
	"regexp"
	"strings"
)

type TokenTag struct {
	key    string
	values []string
	parent string
	child  string
}

type RuleTokenizer struct {
	// line string
	re *regexp.Regexp
}

type TokenizedRule struct {
	// words []Word
	words []Word
	ops   []Op
}

type Word struct {
	text    string
	index   int
	sibling *Word
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

func (t *RuleTokenizer) Tokenize(rule string) *TokenizedRule {
	matches := t.re.FindAllStringSubmatch(rule, -1)
	tokenized := &TokenizedRule{}
	for i, match := range matches {
		if match[1] != "" { // 2nd position of match tuple represents a 'word' if not zero-value
			word := Word{match[1], i, nil}
			// switch word {
			// case "MEASUREMENT":
			// 	meas := Measurement{word, i}
			// case "FIELD":
			// 	field := Field{word, i}
			// default:
			// 	tag := Key{word, i, nil, nil, nil}
			// 	tokenized.tags = append(tokenized.tags, tag)
			tokenized.words = append(tokenized.words, word)
		} else {
			op := Op{match[2], i}
			tokenized.ops = append(tokenized.ops, op)
		}
	}
	return tokenized
}

func (t *RuleTokenizer) FindSiblings(rule string) [][]string {
	detectedSibs := detectSiblingTokens(rule)
	var siblings [][]string
	for _, pipeSet := range detectedSibs {
		set := strings.Split(pipeSet[0], "|")
		siblings = append(siblings, set)
	}
	return siblings
}

func detectSiblingTokens(rule string) [][]string {
	var sibs [][]string
	re := regexp.MustCompile(`[^>|]+\|[^>|]+`)
	sibs = re.FindAllStringSubmatch(rule, -1)
	return sibs
}

func MeasIndex(line string) {
}

func ParseMeas(point string) string {
	substrings := strings.Split(point, ",")
	return substrings[0]
}

// func MeasIndex(line string) {

// }
// func ParseMeas(point string) string {
// 	substrings := strings.Split(point, ",")
// 	return substrings[0]
// }

type Line struct {
	Measurement Measurement
	Tags        []Tag
	Fields      []Field
}

type Tag struct {
	Key   string
	Value string
	Index int
}

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
