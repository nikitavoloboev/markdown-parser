package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLinkWithDescription(t *testing.T) {
	assert := assert.New(t)
	v := "- [Effective Go](https://golang.org/doc/effective_go.html) - Amazing doc."
	result := ParseLinkWithDescription(v)
	assert.Equal(result["Title"], "Effective Go", "Check to see if the title is correct")
	assert.Equal(result["Link"], "https://golang.org/doc/effective_go.html", "Check to see if the link is correct")
	assert.Equal(result["Description"], "Amazing doc.", "Check to see if the description is correct")
}

func TestParseLink(t *testing.T) {
	assert := assert.New(t)
	v := "- [Effective Go](https://golang.org/doc/effective_go.html)"
	result := ParseLinkWithDescription(v)
	assert.Equal(result["Title"], "Effective Go", "Check to see if the title is correct")
	assert.Equal(result["Link"], "https://golang.org/doc/effective_go.html", "Check to see if the link is correct")
	assert.Equal(result["Description"], "", "Check to see if the description is correct")
}

func TestParseMarkdownFile(t *testing.T) {
	// result, err := ParseMarkdownFile("websites.md")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("test")
}
