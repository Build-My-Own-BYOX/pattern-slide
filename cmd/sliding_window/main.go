package main

import (
	"fmt"
	"github.com/Build-My-Own-BYOX/pattern-slide/algo"
)

func main() {
	targetStr := "cat"
	randomStr := "catLorem Ipsum iscatcat catdogcat simply dummcat caty text cat and dogcat"
	occurs := algo.BruteForceSearch(randomStr, targetStr)
	for pos := range occurs {
		fmt.Println("Found match at", occurs[pos])
	}
	fmt.Println("Total number of matches: ", len(occurs))
}
