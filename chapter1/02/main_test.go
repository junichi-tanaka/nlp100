
package main

import "testing"

func TestCombine(t *testing.T) {
	actual := Combine("パトカー", "タクシー")
	expected := "パタトクカシーー"
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	} 

}
