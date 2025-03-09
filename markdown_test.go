package markdown

import (
	"os"
	"testing"
)

func TestMarkdownFile(t *testing.T) {
	src, err := os.ReadFile("test.md")
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	md := NewMarkdown(string(src))
	html := md.ToHTML()

	err = os.WriteFile("test.html", []byte(html), 0644)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

}
