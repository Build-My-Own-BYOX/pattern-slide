package main

import (
	"fmt"
	"github.com/Build-My-Own-BYOX/pattern-slide/algo"
)

func main() {
	targetStr := "cat"
	randomStr := "Lorem Ipsum iscat simply dummy text cat and dog"
	occurs := algo.SingleSlidingWindow(randomStr, targetStr)
	for pos := range occurs {
		fmt.Println("Found match at", occurs[pos])
	}
	fmt.Println("Total number of matches: ", len(occurs))
}
