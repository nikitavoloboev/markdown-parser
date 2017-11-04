package parser

import "testing"

func TestStartsWithSlash(t *testing.T) {
	v := "# Intro"
	result := startsWithSlash(v)
	expected := true
	if expected != result {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
