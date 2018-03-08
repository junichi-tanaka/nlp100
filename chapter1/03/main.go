
package main

import (
	"fmt"
	"strings"
	"regexp"
	"unicode/utf8"
)

type Item struct {
	Word	string
	Length	int
}

func Word(s string) string {
	result := ""
	reg := regexp.MustCompile(`([A-Za-z0-9]*)`)
	matched := reg.FindAllString(s, 1)
	if len(matched) > 0 {
		result = matched[0]
	}
	return result
}

func Nlp03(s string) []Item {
	words := strings.Split(s, " ")
	results := make([]Item, len(words));

	for i, word := range words {
		w := Word(word)
		results[i] = Item{w, utf8.RuneCountInString(w)}
	}
	return results
}

func main() {
	results := Nlp03("Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics.")
	for _, item := range results {
		fmt.Println(item.Word, item.Length)
	}
}

