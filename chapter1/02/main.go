
package main

import "fmt"

func Combine(s1 string, s2 string) string {
	runes1 := []rune(s1)
	runes2 := []rune(s2)
	runes := make([]rune, 0)

	var rune rune
	for len(runes1) > 0 || len(runes2) > 0 {
		if len(runes1) > 0 {
			rune, runes1 = runes1[0], runes1[1:]
			runes = append(runes, rune)
		}
		if len(runes2) > 0 {
			rune, runes2 = runes2[0], runes2[1:]
			runes = append(runes, rune)
		}
	}
	return string(runes)
}

func main() {
	fmt.Println(Combine("パトカー", "タクシー"))
	fmt.Println(Combine("123", "１２３４"))
	fmt.Println(Combine("1234", "１２３"))
}

