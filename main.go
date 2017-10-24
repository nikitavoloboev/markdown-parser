package main

import (
	"fmt"
	"io/ioutil"
)

//  downloadMD given URL, downloads Markdown file from the URL.
func downloadMD() {

}

//  getAllLinks returns all links and their names from a given markdown file.
func getAllLinks(fileContent string) {

}

// readFile Reads file.
func readFile(filename string) string {
	b, err := ioutil.ReadFile(filename) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func main() {
	fileContent := readFile("Summary.md")
	getAllLinks(fileContent)
}
