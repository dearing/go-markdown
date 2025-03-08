package markdown

import "testing"

func TestMarkdownH1(t *testing.T) {

	test := "# Hello World"
	expected := "<h1>Hello World</h1>"

	md := New()

	result, err := md.MarkdownToHTML(test)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}
