package parser

import (
	"fmt"
	"log"
	"testing"
)

func assert(tb testing.TB, condition bool, msg string)
func ok(tb testing.TB, err error)
func equals(tb testing.TB, exp, act interface{})

func TestParseLinkWithDescription(t *testing.T) {
	v := "- [Effective Go](https://golang.org/doc/effective_go.html) - Amazing doc."
	ParseLinkWithDescription(v)
	// result := TestParseLinkWithDescription(v)
}

func TestParseMarkdownFile(t *testing.T) {
	result, err := ParseMarkdownFile("websites.md")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result["a: alfred"])
}
