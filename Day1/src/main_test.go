package main

import "testing"

func Test_ShouldAdd_AdditionShouldReturnTrue(t *testing.T) {
	if ShouldAdd("+3") != true {
		t.Error("+3 should return true")
	}
}

func Test_ShouldAdd_SubtractionShouldReturnFalse(t *testing.T) {
	if ShouldAdd("-3") == true {
		t.Error("-3 should return false")
	}
}
