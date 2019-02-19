package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLink(t *testing.T) {
	// Define assertion
	assert := assert.New(t)
	// Link with description
	v := "- [Effective Go](https://golang.org/doc/effective_go.html) - Amazing doc."
	result := ParseLink(v)
	assert.Equal(result["Title"], "Effective Go", "Check to see if the title is correct")
	assert.Equal(result["Link"], "https://golang.org/doc/effective_go.html", "Check to see if the link is correct")
	assert.Equal(result["Description"], "Amazing doc.", "Check to see if the description is correct")
	// Link without description
	v2 := "- [Effective Go](https://golang.org/doc/effective_go.html)"
	result2 := ParseLink(v2)
	assert.Equal(result2["Title"], "Effective Go", "Check to see if the title is correct")
	assert.Equal(result2["Link"], "https://golang.org/doc/effective_go.html", "Check to see if the link is correct")
	assert.Equal(result2["Description"], "", "Check to see if the description is correct")
}

func TestParseMarkdownFile(t *testing.T) {
	// result, err := ParseMarkdownFile("websites.md")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("test")
}
