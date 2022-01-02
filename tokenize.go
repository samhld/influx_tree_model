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
	for f := range fieldKeys {
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

func MapTokensToData(measAPI dataGetter, tokenizedRule *TokenizedRule) Tiers {
	tiers := make(Tiers)
	tagTierTracker := 0 // increment each Tag tier; add to i to keep i aligned
	for i, word := range tokenizedRule.words {
		switch word.text {
		case "MEASUREMENT":
			tiers[i+tagTierTracker] = &Measurement{measAPI.getMeasurement(), i}
		case "FIELD":
			fields := measAPI.getFields()
			// tiers[i] = []&Field{"FIELD", i}
			tiers[i+tagTierTracker] = fields
		default:
			tiers[i+tagTierTracker] = &Key{word.text, i + tagTierTracker, nil, nil}
			vals := measAPI.getTagKeyValues(word.text)
			tagTierTracker++
			var values []*Value
			for _, val := range vals {
				values = append(values, &Value{val, i + tagTierTracker, nil, nil})
			}
			tiers[i+tagTierTracker] = values
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
