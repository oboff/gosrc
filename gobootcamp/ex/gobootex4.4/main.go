package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	//return map[string]int{"x": 1}
	wordcounts := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		wordcounts[word]++
	}
	return wordcounts
}
func main() {
	wc.Test(WordCount)
}
