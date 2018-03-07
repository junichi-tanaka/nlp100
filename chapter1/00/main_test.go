
package main

import "testing"

func TestReverseStressed(t *testing.T) {
	subject := "stressed"
	actual := Reverse(subject)
	expected := "desserts"
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	} 

}

func TestReverseAiueo(t *testing.T) {
	subject := "あいうえお"
	actual := Reverse(subject)
	expected := "おえういあ"
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	} 
}

func TestReverseHelloWorld(t *testing.T) {
	subject := "Hello, 世界"
	actual := Reverse(subject)
	expected := "界世 ,olleH"
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	} 
}
