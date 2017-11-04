package parser

import "testing"

func TestStartsWithHash(t *testing.T) {
	v := "# Intro"
	result := startsWithHash(v)
	expected := true
	if expected != result {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
