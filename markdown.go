package markdown

import (
	"fmt"
	"log/slog"
	"strings"
)

type nodeType int

const (
	nodeRoot nodeType = iota
	nodeParagraph
	nodeHeading1
	nodeHeading2
	nodeHeading3
	nodeHeading4
	nodeHeading5
	nodeHeading6
	NodeHorizontalRule
	nodeList
	nodeCodeBlock
	nodeBlockquote
)

type node struct {
	typ nodeType
	val string
}

type Markdown struct {
	src   string
	nodes []node
}

func NewMarkdown(src string) *Markdown {
	return &Markdown{src: src}
}

func (m Markdown) ToHTML() string {

	html := ""

	blocks := strings.Split(m.src, "\n\n")
	for index, block := range blocks {
		block = strings.TrimSpace(block)

		if block == "" {
			continue
		}

		if strings.HasPrefix(block, "- ") {

			var items []string
			for item := range strings.Lines(block) {
				item := strings.TrimSpace(item)
				item = strings.TrimPrefix(item, "- ")
				items = append(items, fmt.Sprintf("<li>%s</li>", item))
			}

			m.nodes = append(m.nodes, node{typ: nodeList, val: strings.Join(items, "\n")})
			continue
		}

		if strings.HasPrefix(block, "```") {
			val := block[3:]
			val = strings.TrimSuffix(val, "```")
			m.nodes = append(m.nodes, node{typ: nodeCodeBlock, val: val})
			continue
		}

		if strings.HasPrefix(block, "# ") {
			m.nodes = append(m.nodes, node{typ: nodeHeading1, val: block[2:]})
			continue
		}
		if strings.HasPrefix(block, "## ") {
			m.nodes = append(m.nodes, node{typ: nodeHeading2, val: block[3:]})
			continue
		}
		if strings.HasPrefix(block, "### ") {
			m.nodes = append(m.nodes, node{typ: nodeHeading3, val: block[4:]})
			continue
		}
		if strings.HasPrefix(block, "#### ") {
			m.nodes = append(m.nodes, node{typ: nodeHeading4, val: block[5:]})
			continue
		}
		if strings.HasPrefix(block, "##### ") {
			m.nodes = append(m.nodes, node{typ: nodeHeading5, val: block[6:]})
			continue
		}
		if strings.HasPrefix(block, "###### ") {
			m.nodes = append(m.nodes, node{typ: nodeHeading6, val: block[7:]})
			continue
		}
		if strings.HasPrefix(block, "---") {
			m.nodes = append(m.nodes, node{typ: NodeHorizontalRule})
			continue
		}

		// default to a paragraph
		m.nodes = append(m.nodes, node{typ: nodeParagraph, val: block})

		slog.Info("block:", "index", index, "block", block)

	}

	for _, n := range m.nodes {
		switch n.typ {
		case nodeParagraph:
			html += fmt.Sprintf("<p>%s</p>\n\n", n.val)
		case nodeHeading1:
			html += fmt.Sprintf("<h1>%s</h1>\n\n", n.val)
		case nodeHeading2:
			html += fmt.Sprintf("<h2>%s</h2>\n\n", n.val)
		case nodeHeading3:
			html += fmt.Sprintf("<h3>%s</h3>\n\n", n.val)
		case nodeHeading4:
			html += fmt.Sprintf("<h4>%s</h4>\n\n", n.val)
		case nodeHeading5:
			html += fmt.Sprintf("<h5>%s</h5>\n\n", n.val)
		case nodeHeading6:
			html += fmt.Sprintf("<h6>%s</h6>\n\n", n.val)
		case NodeHorizontalRule:
			html += "<hr />\n\n"
		case nodeBlockquote:
			html += fmt.Sprintf("<blockquote>%s</blockquote>\n\n", n.val)
		case nodeCodeBlock:
			html += fmt.Sprintf("<pre class='code'><code>%s</code></pre>\n\n", n.val)
		case nodeList:
			html += fmt.Sprintf("<ul>\n%s\n</ul>\n\n", n.val)
		}

	}
	return html
}
