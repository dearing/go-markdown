package markdown

import (
	"fmt"
	"strings"
	"testing"
)

func TestMarkdownHeading(t *testing.T) {
	for i := range 5 {
		test := strings.Repeat("#", i+1) + " Hello World"
		expected := fmt.Sprintf("<h%d>Hello World</h%d>", i+1, i+1)
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

func TestMarkdownHorizontalRule(t *testing.T) {
	test := "---"
	expected := "<hr/>"
	md := New()

	result, err := md.MarkdownToHTML(test)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}
