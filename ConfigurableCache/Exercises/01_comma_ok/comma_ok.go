package main

import (
	"strings"
	"fmt"
)

/*
	Maps, the “comma ok” idiom and zero values

	Build: a word counter that takes a sentence and fills a map[string]int, plus Lookup(word) (int, bool).

	Try: read a missing key with and without the second return value; delete a key that doesn’t exist.

	You’ve got it when: you know what a missing key returns for int, string and a pointer.
*/

func WordCounter(sentence string) map[string]int {
	wordCounter := make(map[string]int)
	words := strings.Fields(sentence)

	for _, w := range words {
		cleaned := strings.ToLower(w)
		wordCounter[cleaned]++
	}
	return wordCounter
}

func Lookup(wordCount map[string]int, word string) (int, bool) {
	val, ok := wordCount[strings.ToLower(word)]
	return val, ok
} 

func main() {
	counts := WordCounter("Go is great and Go is fast")

	count, ok := Lookup(counts, "go")
	fmt.Printf("Lookup 'go': count=%d, exists=%t\n", count, ok)

	count, ok = Lookup(counts, "python")
	fmt.Printf("Lookup missing key (with ok): count=%d, exists=%t\n", count, ok)

	valWithoutOk := counts["python"]
	fmt.Printf("Read missing key (without ok): %d\n", valWithoutOk)

	delete(counts, "rust")
	fmt.Println("Deleted non-existent key 'rust' without error.")
}
