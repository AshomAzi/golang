package main

import (
	"fmt"
)

func UniqueCharacterChecker(input string) bool {
	seen := make(map[rune]bool)

	for _, val := range input {
		if seen[val] {
			fmt.Println(seen)
			return false

		}else {
			seen[val] = true
		}
	}

	fmt.Println(seen)
	return true
}

func main() {
	fmt.Println(UniqueCharacterChecker("apple"))
}
