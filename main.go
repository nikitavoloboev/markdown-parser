package main

import (
	"fmt"
	"nikitavoloboev/markdown/parser"
)

func main() {
	fmt.Println(parser.DownloadUrl("https://raw.githubusercontent.com/nikitavoloboev/markdown-parser/master/readme.md"))
}
