package main

import "fmt"

func UniqueCharacterChecker(input string) map[rune]bool {
	seen := map[rune]bool{}
	for _, v := range input {
		seen[v] = true
		for u := range seen {
			// fmt.Println(u)
			if v == u {
				seen[v] = false
				fmt.Println("Seen")
				break
			
			}
		}
	}
	return seen
}

func main() {
	fmt.Println(UniqueCharacterChecker("apple"))
}
