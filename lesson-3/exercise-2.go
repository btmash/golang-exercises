package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var splitString []string
	wordMap := make(map[string]int)
	splitString = strings.Fields(s)
	for _, x := range splitString {
		counter, ok := wordMap[x]
		if ok {
			counter++
		} else {
			counter = 1
		}
		wordMap[x] = counter

	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
