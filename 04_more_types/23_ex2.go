package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var m = make(map[string]int)
	for _, i := range strings.Fields(s) {
		m[i]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
