package main

import "testing"

func TestCal(t *testing.T) {
	result, _ := cal(10, 5, "+")
	if result != 15 {
		t.Errorf("Expected 15, got %f", result)
	}
