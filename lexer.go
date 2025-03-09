package markdown

import (
	"log/slog"
	"unicode/utf8"
)

// lexer is a markdown lexer
type lexer struct {
	input    string    // markdown input
	index    int       // start position of the current token
	position int       // current position in the input
	width    int       // width of the last rune read
	nodes    chan Node // channel to send nodes to
}

// run starts the lexer at the root state
func (l *lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.nodes) // close the channel when done
}

// lex creates a new lexer
func lex(input string) *lexer {
	l := &lexer{
		input: input,
		nodes: make(chan Node),
	}
	go l.run() // start the lexer
	return l
}

// emit sends a node to the channel
func (l *lexer) emit(t NodeType) {

	//slog.Info("emit", "type", t, "content", l.input[l.index:l.position])

	l.nodes <- Node{
		Type:    t,
		Content: l.input[l.index:l.position], // content for the current substring of the input (index .. position)
	}

	// move the index to the current position
	l.index = l.position

}

func (l *lexer) next() rune {
	if l.position >= len(l.input) {
		l.width = 0
		return 0
	}

	r, w := rune(l.input[l.position]), 1

	if r >= utf8.RuneSelf {
		r, w = utf8.DecodeRuneInString(l.input[l.position:])
	}

	l.width = w
	l.position += l.width

	return r
}

// stateFn is a state function for the lexer
type stateFn func(*lexer) stateFn

// lexText is the root state function
func lexText(l *lexer) stateFn {

	slog.Info("lexText", "pos", l.position, "input", l.input)

	for {
		// position is beyond the input length
		if l.position >= len(l.input) {
			l.emit(NodeDocument)
		}

		if l.next() == 0 {
			return lexEOF
		}

		if l.next() == '\n' {
			return lexParagraph
		}
	}
}

// lexHeading is the heading state function
func lexEOF(l *lexer) stateFn {

	slog.Info("lexEOF", "pos", l.position, "input", l.input)

	l.emit(NodeEOF)

	return nil
}

func lexParagraph(l *lexer) stateFn {

	if l.position >= len(l.input) {
		l.emit(NodeParagraph)
		return lexEOF
	}

	return nil
}
