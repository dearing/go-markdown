package markdown

import "strings"

type Markdown struct{}

const (
	parseError = iota
	parseHeading
	parseHorizontalRule
	parseUnorderedList
)

func New() Markdown {
	return Markdown{}
}

func (m Markdown) MarkdownToHTML(markdown string) (string, error) {

	html := ""

	for line := range strings.Lines(markdown) {

		// headings <h1> to <h6>
		if strings.HasPrefix(line, "# ") {
			html += "<h1>" + strings.TrimSpace(line[2:]) + "</h1>"
			continue
		}
		if strings.HasPrefix(line, "## ") {
			html += "<h2>" + strings.TrimSpace(line[3:]) + "</h2>"
			continue
		}
		if strings.HasPrefix(line, "### ") {
			html += "<h3>" + strings.TrimSpace(line[4:]) + "</h3>"
			continue
		}
		if strings.HasPrefix(line, "#### ") {
			html += "<h4>" + strings.TrimSpace(line[5:]) + "</h4>"
			continue
		}
		if strings.HasPrefix(line, "##### ") {
			html += "<h5>" + strings.TrimSpace(line[6:]) + "</h5>"
			continue
		}
		if strings.HasPrefix(line, "###### ") {
			html += "<h6>" + strings.TrimSpace(line[7:]) + "</h6>"
			continue
		}

		// horizontal rule <hr/>
		if strings.HasPrefix(line, "---") {
			html += "<hr/>"
			continue
		}

	}

	return html, nil
}
