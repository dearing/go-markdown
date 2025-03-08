package markdown

type Markdown struct{}

func New() Markdown {
	return Markdown{}
}

func (m Markdown) MarkdownToHTML(markdown string) (string, error) {
	// This is a placeholder function
	return markdown, nil
}
