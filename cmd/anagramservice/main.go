package main

import (
	"fmt"
)

// [^0-9|\s|[:punct:]]+

func main() {
	word := "好你"
	wordSum := 0
	for i := 0; i < len(word); i++ {
		wordSum += int(word[i])
	}

	anagrams := []string{
		"elloh",
		"lloeh",
		"lehlo",
		"qwert",
		"你好",
	}

	for _, anagram := range anagrams {
		currWordSum := 0
		for i := 0; i < len(anagram); i++ {
			currWordSum += int(anagram[i])
		}

		if currWordSum == wordSum {
			fmt.Printf("%s is anagram of %s\n", anagram, word)
		}
	}
}
