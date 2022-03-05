package main

import "testing"

func TestSqrt(t *testing.T) {
	got, _ := sqrt(4)
	if got != 2 {
		t.Errorf("sqrt(4) = %f, want 2", got)
	}
}
