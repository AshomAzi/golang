package main

import (
	"fmt"
	"strings"
)

func WordFrequencyCounter(input string) map[string]int {
	frequency := make(map[string]int)
	each := strings.Split(input, " ")
	for _, v := range each {
		frequency[v] += 1
	}
	return frequency
}

func main() {
	fmt.Println(WordFrequencyCounter("apple banana apple orange banana apple"))
}