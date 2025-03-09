package markdown

import (
	"fmt"
	"log/slog"
)

type Markdown struct{}

func New() Markdown {
	return Markdown{}
}

func (m Markdown) MarkdownToHTML(markdown string) (string, error) {

	html := ""

	l := lex(markdown)

	for {
		select {
		case n := <-l.nodes:
			switch n.Type {
			case NodeError:
				slog.Error("error parsing markdown", "content", n.Content)
				return "", fmt.Errorf("error parsing markdown: %s", n.Content)

			case NodeEOF:
				slog.Info("end of file")
				return html, nil

			case NodeDocument:
				slog.Info("emit", "type", n.Type, "content", n.Content)
				html += n.Content

			case NodeParagraph:
				slog.Info("emit", "type", n.Type, "content", n.Content)
				html += fmt.Sprintf("<p>%s</p>\n\n", n.Content)

			}
		}
	}
}
