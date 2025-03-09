package markdown

// NodeType is the type of node in the markdown AST
type NodeType int

const (
	NodeError     = iota // NodeError is a special case where the node is really and error from parsing
	NodeDocument         // NodeText is unformatted text (shouldn't be used)
	NodeEOF              // NodeEOF is the end of the document
	NodeParagraph        // NodeParagraph is a block of text
)

// Node is a node in the markdown AST
type Node struct {
	Type     NodeType // Type is the type of node
	Content  string   // Content is the content of the node (if any)
	Children []Node   // Children is a list of child nodes
	Parent   *Node    // Parent is the parent node (if any)
}

func (nt NodeType) String() string {
	switch nt {
	case NodeError:
		return "NodeError"
	case NodeDocument:
		return "NodeDocument"
	case NodeEOF:
		return "NodeEOF"
	case NodeParagraph:
		return "NodeOpenParagraph"
	default:
		return "Unknown"
	}
}
