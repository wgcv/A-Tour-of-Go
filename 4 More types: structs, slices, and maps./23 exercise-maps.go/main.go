package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var mapWord = make(map[string]int)
	sm := strings.Fields(s)
	for _, word := range sm {
		if _, exist := mapWord[word]; exist {
			mapWord[word] += 1
		} else {
			mapWord[word] = 1
		}
	}
	return mapWord
}

func main() {
	wc.Test(WordCount)
}
