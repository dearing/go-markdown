package markdown

import (
	"fmt"
	"strings"
	"testing"
)

// see: https://github.github.com/gfm/#atx-headings
func TestMarkdownATXHeadings(t *testing.T) {
	for i := range 6 {
		test := strings.Repeat("#", i+1) + " Hello World\n"
		expected := fmt.Sprintf("<h%d>Hello World</h%d>\n", i+1, i+1)
		md := New()

		result, err := md.MarkdownToHTML(test)
		if err != nil {
			t.Errorf("Error: %s", err)
		}

		if result != expected {
			t.Errorf("Expected: %s, Got: %s", expected, result)
		}
	}

	test := `# foo *bar* \*baz\*`
	expected := `<h1>foo <em>bar</em> *baz*</h1>`

	md := New()
	result, err := md.MarkdownToHTML(test)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}

}

// TestMarkdownThematicBreaks test the horizontal rule syntax
// see: https://github.github.com/gfm/#thematic-breaks
func TestMarkdownThematicBreaks(t *testing.T) {

	syntax := []string{"---", "***", "___"}
	for _, s := range syntax {
		test := s
		expected := "<hr />"
		md := New()

		result, err := md.MarkdownToHTML(test)
		if err != nil {
			t.Errorf("Error: %s", err)
		}

		if result != expected {
			t.Errorf("Expected: %s, Got: %s", expected, result)
		}
	}
}

func TestMarkdownIndentedCodeBlock(t *testing.T) {
	test := `a simple
  indented code block`

	expected := "<pre><code>a simple\n  indented code block\n</code></pre>"
	md := New()

	result, err := md.MarkdownToHTML(test)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}
