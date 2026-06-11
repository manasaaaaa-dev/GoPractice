package main

import (
    "fmt"
    "strings"
    "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
    fields := strings.Fields(s)      
    counts := make(map[string]int)

	// for _, word := range fields {
	// 	Here range fields means - "go through each element of fields one by one".

    for index, word := range fields {
        counts[word]++
    }

    return counts
}

func main() {
    wc.Test(WordCount)
}