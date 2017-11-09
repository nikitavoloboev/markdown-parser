package parser

import (
	"fmt"
	"log"
	"testing"
)

func TestStartsWithHash(t *testing.T) {
	v := "# Intro"
	result := startsWithHash(v)
	expected := true
	if expected != result {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestParseMarkdownFile(t *testing.T) {
	result, err := ParseMarkdownFile("websites.md")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result["a: alfred"])
}
