package main

import "testing"

func TestMysum(t *testing.T) {
	x := mySum(2, 3, 5)
	if x != 10 {
		t.Error("Expected", 10, "Got", x)
	}
}
