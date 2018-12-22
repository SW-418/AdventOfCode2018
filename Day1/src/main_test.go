package main

import "testing"

func TestShouldAddAddition(t *testing.T) {
	if ShouldAdd("+3") != true {
		t.Error("+3 should return true")
	}
}

func TestShouldAddSubtraction(t *testing.T) {
	if ShouldAdd("-3") == true {
		t.Error("-3 should return false")
	}
}
