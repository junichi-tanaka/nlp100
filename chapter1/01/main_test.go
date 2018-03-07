
package main

import "testing"

func TestStep(t *testing.T) {
	subject := "パタトクカシーー"
	actual := Step(subject)
	expected := "パトカー"
	if actual != expected {
		t.Errorf("expected %v, but got %v", expected, actual)
	} 

}
