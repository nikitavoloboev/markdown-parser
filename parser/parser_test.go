package parser

import "testing"

func TeststartsWithSlash(t *testing.T) {
	v := "# Intro"
	if startsWithSlash(v) != true {
		t.Error("Expected True, got ", startsWithSlash(v))
	}
}
