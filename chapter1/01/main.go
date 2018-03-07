
package main

import "fmt"

func Step(s string) string {
	step := 2
	runes := make([]rune, 0)
	i := 0
	for _, rune := range s {
		if (i % step == 0) {
			runes = append(runes, rune)
		}
		i++
	}
	return string(runes)
}

func main() {
	fmt.Println(Step("パタトクカシーー"))
	fmt.Println(Step("パタトクカシー"))
}

