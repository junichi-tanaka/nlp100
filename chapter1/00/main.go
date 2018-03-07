
package main

import	"fmt"

func Reverse(s string) string {
	// https://github.com/golang/go/wiki/SliceTricks#reversing
	a := []rune(s)
	for i := len(a)/2-1; i >= 0; i-- {
		opp := len(a)-1-i
		a[i], a[opp] = a[opp], a[i]
	}
	return string(a)
}

func main() {
	fmt.Println(Reverse("stressed"))
	fmt.Println(Reverse("あいうえお"))
	fmt.Println(Reverse("Hello, 世界"))
}
