// Package parser provides methods to grab all links from markdown files.
package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// DownloadURL returns Body response from the URL.
func DownloadURL(URL string) string {
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		return buf.String()
	}
	return "test"
}

// GetAllLinks returns all links and their names from a given markdown file.
func GetAllLinks(markdown string) map[string]string {
	// holds all the links and their corresponding values
	m := make(map[string]string)

	// regex to extract link and text attached to link
	re := regexp.MustCompile(`\[([^\]]*)\]\(([^)]*)\)`)

	scanner := bufio.NewScanner(strings.NewReader(markdown))
	// scans line by line
	for scanner.Scan() {
		// apply regex
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)

		// only apply regex if there are links
		if matches != nil {
			m[matches[0][1]] = matches[0][2]
		}
	}
	return m
}

// readFile returns contents of the file.
func readFile(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}
