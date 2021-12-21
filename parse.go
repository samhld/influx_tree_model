package main

import "strings"

func ParseMeas(point string) string {
	substrings := strings.Split(point, ",")
	return substrings[0]
}
