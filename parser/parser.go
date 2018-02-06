// Package parser provides methods to grab all links from markdown files.
package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// DownloadURL returns Body response from the URL.
func DownloadURL(URL string) (string, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.String(), nil
}

// DownloadURLToFile downloads Body response from URL to a file.
func DownloadURLToFile(urlStr string, fileName string) error {
	var f *os.File

	resp, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)

	return err
}

// startsWithSlash checks if string starts with a hash.
func startsWithHash(line string) bool {
	return strings.HasPrefix(line, "#")
}

// ParseLinkWithDescription parses a line and grabs the Link, Title and the Description attached to it.
// The format of the line should be as follows: `- [Title](Link) - Description.
// Description can be omitted.
// TODO: test it
func ParseLinkWithDescription(line string) {
	re := regexp.MustCompile(`\[([^]]+)\]\(([^)]+)\)(.*)`)
	match := re.FindStringSubmatch(line)
	if len(match) != 0 {
		fmt.Printf(match[0])
	}
}

// GetAllLinks returns all links and their names from a given markdown file.
func GetAllLinks(markdown string) map[string]string {
	// Holds all the links and their corresponding values
	m := make(map[string]string)

	// Regex to extract link and text attached to link
	re := regexp.MustCompile(`\[([^\]]*)\]\(([^)]*)\)`)

	scanner := bufio.NewScanner(strings.NewReader(markdown))
	// Scans line by line
	for scanner.Scan() {
		// Make regex
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)

		// Only apply regex if there are links and the link does not start with #
		if matches != nil {
			if startsWithHash(matches[0][2]) == false {
				// fmt.Println(matches[0][2])
				m[matches[0][1]] = matches[0][2]
			}
		}
	}
	return m
}

// fileToString returns string representation of a file.
func fileToString(file string) (string, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	s := string(bytes)
	return s, nil
}

// ParseMarkdownFile parses a markdown file and returns all markdown links from it.
func ParseMarkdownFile(fileName string) (map[string]string, error) {
	file, err := fileToString(fileName)
	if err != nil {
		log.Fatal()
	}
	return GetAllLinks(file), nil
}

// ParseMarkdownURL parses an URL and returns all markdown links from it.
func ParseMarkdownURL(URL string) (map[string]string, error) {
	file, err := DownloadURL(URL)
	if err != nil {
		return make(map[string]string), err
	}
	return GetAllLinks(file), nil
}

// readFile returns contents of the file.
func readFile(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}
