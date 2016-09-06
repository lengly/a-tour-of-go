package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var dict=make(map[string]int)
	var sl = strings.Fields(s)
	for i:=0;i<len(sl);i++ {
		dict[sl[i]] = dict[sl[i]] + 1	
	}

	return dict
}

func main() {
	wc.Test(WordCount)
}
