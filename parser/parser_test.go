package parser

import "testing"

func TestStartsWithHash(t *testing.T) {
	v := "# Intro"
	result := startsWithHash(v)
	expected := true
	if expected != result {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// func TestDownloadURL(t *testing.T)  {
// 	v := "https://raw.githubusercontent.com/nikitavoloboev/markdown-parser/master/readme.md"
// 	result := DownloadURL(v)

// }
