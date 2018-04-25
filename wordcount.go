// Word Count https://tour.golang.org/moretypes/23
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) (wordCount map[string]int) {

	// string to list of words
	wordList := strings.Fields(s)

	// init output
	// can we derive the type to initiate from return declaration?
	wordCount = make(map[string]int)

	// count each word in list
	for _, word := range wordList {
		// increment the count for the current word
		wordCount[word] += 1
	}

	return
}

func main() {
	wc.Test(WordCount)
}
