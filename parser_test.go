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
	// 3. Image link
	v3 := "![](https://i.imgur.com/ZVPjkzh.png)"
	result3 := ParseLink(v3)
	assert.Equal(result3["Title"], "", "3. Check to see if the title is correct")
	assert.Equal(result3["Link"], "https://i.imgur.com/ZVPjkzh.png", "3. Check to see if the link is correct")
	assert.Equal(result3["Description"], "", "3. Check to see if the description is correct")
	// 4. Random text
	v4 := "blargh"
	result4 := ParseLink(v4)
	assert.Equal(result4["Title"], "", "4. Check to see if the title is correct")
	assert.Equal(result4["Link"], "", "4. Check to see if the link is correct")
	assert.Equal(result4["Description"], "", "4. Check to see if the description is correct")
}

func TestParseMarkdownFile(t *testing.T) {
	// result, err := ParseMarkdownFile("websites.md")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("test")
}
