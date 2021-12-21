package main

import (
	"strings"
	"unicode"
)

type Rule struct {
	line             string
	measurementIndex int
	tagsIndices      []int
	fieldsIndices    []int
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

tokenizeRule(userRule string) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	strings.FieldsFunc(line, f)
}

func MeasIndex(line string) {

}
func ParseMeas(point string) string {
	substrings := strings.Split(point, ",")
	return substrings[0]
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
