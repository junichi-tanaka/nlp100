
package main

import (
	"testing"
)

func TestNlp03(t *testing.T) {
	actual := Nlp03("Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics.")
	expected := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9}
	for i, item := range actual {
		if item.Length != expected[i] {
			t.Errorf("%v: expected %v, but got %v", item.Word, expected[i], item.Length)
		}
	} 

}
