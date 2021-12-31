package main

import (
	"regexp"
	"strings"
)

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

func createBranches(tiers Tiers, fieldKeys []string) []Branch {
	var branches []Branch
	for f, _ := range fieldKeys {
		var branch []string
		t := 0
		for t < len(tiers) {
			node := tiers[t]
			switch v := node.(type) {
			case *Tag:
				branch[t] = v.key.text
				branch[t+1] = v.value.text
			case *Field:
				branch[t] = fieldKeys[f]
			}
		}
	}
	return branches
}

func MapTokensToData(measAPI *MeasurementAPI, tokenizedRule *TokenizedRule) Tiers {
	tiers := make(Tiers)
	for i, word := range tokenizedRule.words {
		switch word.text {
		case "MEASUREMENT":
			tiers[i] = &Measurement{measAPI.measurement, i}
		case "FIELD":
			fieldKeys := measAPI.getFieldKeys()
			// tiers[i] = []&Field{"FIELD", i}
			tiers[i] = fieldKeys
		default:
			vals := measAPI.getTagKeyValues(word.text)
			tiers[i] = &Key{word.text, i, vals, nil, nil}
		}
	}

	return tiers
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
