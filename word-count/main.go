package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum`

	splittedText := strings.Fields(text)

	wordCount := map[string]int{}

	for _, word := range splittedText {
		lowerWord := strings.ToLower(word)

		count := wordCount[lowerWord]
		wordCount[lowerWord] = count+1
	}

	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
