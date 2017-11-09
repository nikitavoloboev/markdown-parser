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

// GetAllLinks returns all links and their names from a given markdown file.
func GetAllLinks(markdown string) map[string]string {
	// holds all the links and their corresponding values
	m := make(map[string]string)

	// regex to extract link and text attached to link
	re := regexp.MustCompile(`\[([^\]]*)\]\(([^)]*)\)`)

	scanner := bufio.NewScanner(strings.NewReader(markdown))
	// scans line by line
	for scanner.Scan() {
		// make regex
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)

		// only apply regex if there are links and the link does not start with #
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

// TODO: transfer whats below
// package main

// import (
// 	"bufio"
// 	"encoding/csv"
// 	"encoding/json"
// 	"fmt"
// 	"github.com/urfave/cli"
// 	"log"
// 	"os"
// 	"regexp"
// 	"strings"
// )

// type Link struct {
// 	Title        string `json:"title"`
// 	Arg          string `json:"arg"`
// 	Autocomplete string `json:"autocomplete"`
// }

// type Items struct {
// 	Items []Link `json:"items"`
// }

// func mdToJSON(fileName string) {

// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	var results Items

// 	s := bufio.NewScanner(file)

// 	re := regexp.MustCompile(`^[#-]+ \[([^\]]+)\]\(([^\)]+)`)
// 	for s.Scan() {
// 		ln := s.Text()

// 		matches := re.FindAllStringSubmatch(ln, -1)
// 		if len(matches) != 1 {
// 			continue
// 		}
// 		results.Items = append(results.Items, Link{
// 			Title:        matches[0][1],
// 			Arg:          matches[0][2],
// 			Autocomplete: matches[0][1],
// 		})
// 	}

// 	if err := s.Err(); err != nil {
// 		log.Fatalln("could not scan input", err)
// 	}

// 	out, err := json.Marshal(results)
// 	if err != nil {
// 		log.Fatalln("could not marshal result", err)
// 	}

// 	f, err := os.Create(strings.Replace(fileName, ".md", ".json", -1))
// 	if err != nil {
// 		log.Fatalln("could not make output file", err)
// 	}
// 	defer f.Close()
// 	f.WriteString(string(out))
// 	fmt.Println(strings.Replace(fileName, ".md", ".json", -1) + " was created")
// }

// // TODO: clean it up and export it to + test + view the local readme I have for instructions (document things)
// func mdToCSV(fileName string) {
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	out := csv.NewWriter(file)
// 	re := regexp.MustCompile(`^- \[([^\]]+)\]\(([^)]+)`)

// 	s := bufio.NewScanner(file)
// 	for s.Scan() {
// 		ln := s.Text()

// 		if !re.MatchString(ln) {
// 			continue
// 		}
// 		matches := re.FindAllStringSubmatch(ln, -1)

// 		row := []string{
// 			matches[0][1],
// 			"",
// 			matches[0][2],
// 		}
// 		if err := out.Write(row); err != nil {
// 			log.Print(err)
// 		}
// 	}
// 	if err := s.Err(); err != nil {
// 		log.Fatalln("scan error", err)
// 	}

// 	out.Flush()
// 	if err := out.Error(); err != nil {
// 		log.Fatalln("write error", err)
// 	}

// 	// write to file
// 	f, err := os.Create(strings.Replace(fileName, ".md", ".csv", -1))
// 	if err != nil {
// 		log.Fatalln("could not make output file", err)
// 	}
// 	defer f.Close()

// 	fmt.Println("alfred.csv file was created")
// }

// func main() {

// 	// read options
// 	app := cli.NewApp()
// 	app.Action = func(c *cli.Context) error {
// 		action := c.Args().Get(0)
// 		fileName := c.Args().Get(1)
// 		switch action {
// 		case "json":
// 			mdToJSON(fileName)
// 		case "csv":
// 			mdToCSV(fileName)
// 		}
// 		return nil
// 	}

// 	app.Run(os.Args)

// }
