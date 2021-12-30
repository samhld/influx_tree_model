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

func MapTokensToData(measAPI *MeasurementAPI, tokenizedRule *TokenizedRule) Tree {
	tree := make(Tree)
	for i, word := range tokenizedRule.words {
		switch word.text {
		case "MEASUREMENT":
			tree[i] = &Measurement{measAPI.measurement, i}
		case "FIELD":
			tree[i] = &Field{"FIELD", i}
		default:
			vals := measAPI.getTagKeyValues(word.text)
			tree[i] = &Key{word.text, i, vals, nil, nil}
		}
	}

	return tree
}

func (t *RuleTokenizer) Tokenize(rule string) *TokenizedRule {
	matches := t.re.FindAllStringSubmatch(rule, -1)
	tokenized := &TokenizedRule{}
	for i, match := range matches {
		if match[1] != "" { // 2nd position of match tuple represents a 'word' if not zero-value
			word := Word{match[1], i, nil}
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
