package main

import (
	"testing"
)

func TestUrlEncode(t *testing.T) {
	in := "This is a string"
	result := urlEncode(in)
	expected := "This+is+a+string"

	if result != expected {
		t.Errorf("Expected [%v] got [%v]", expected, result)
	}
}
