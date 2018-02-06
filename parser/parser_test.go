package parser

import (
	"fmt"
	"log"
	"testing"
)

func TestParseLinkWithDescription(t *testing.T) {
	v := "- [The art of computer programming (1968)](http://www.goodreads.com/book/show/112239.The_Art_of_Computer_Programming_Volumes_1_3_Boxed_Set)"
	ParseLinkWithDescription(v)
	// result := TestParseLinkWithDescription(v)
}

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
