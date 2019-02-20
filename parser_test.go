package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLink(t *testing.T) {
	// Define assertion
	assert := assert.New(t)
	// 1. Link with description
	v := "- [Effective Go](https://golang.org/doc/effective_go.html) - Amazing doc."
	result := ParseLink(v)
	assert.Equal(result["Title"], "Effective Go", "1. Check to see if the title is correct")
	assert.Equal(result["Link"], "https://golang.org/doc/effective_go.html", "1. Check to see if the link is correct")
	assert.Equal(result["Description"], "Amazing doc.", "1. Check to see if the description is correct")
	// 2. Link without description
	v2 := "- [Effective Go](https://golang.org/doc/effective_go.html)"
	result2 := ParseLink(v2)
	assert.Equal(result2["Title"], "Effective Go", "2. Check to see if the title is correct")
	assert.Equal(result2["Link"], "https://golang.org/doc/effective_go.html", "2. Check to see if the link is correct")
	assert.Equal(result2["Description"], "", "2. Check to see if the description is correct")
	// 3. Random text
	v3 := "blargh"
	result3 := ParseLink(v3)
	assert.Equal(result3["Title"], "", "3. Check to see if the title is correct")
	assert.Equal(result3["Link"], "", "3. Check to see if the link is correct")
	assert.Equal(result3["Description"], "", "3. Check to see if the description is correct")
}

func TestParseImageLink(t *testing.T) {
	// Define assertion
	assert := assert.New(t)
	// 1. Image link
	v := "![](https://i.imgur.com/ZVPjkzh.png)"
	result := ParseImageLink(v)
	assert.Equal(result, "https://i.imgur.com/ZVPjkzh.png", "1. Check to see if the image link is correct")
	// 4. Random text
	v2 := "blargh"
	result2 := ParseImageLink(v2)
	assert.Equal(result2, "", "2. Check to see if the image link is correct")
}

func TestParseMarkdownFile(t *testing.T) {
	// result, err := ParseMarkdownFile("websites.md")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("test")
}
