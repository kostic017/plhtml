package ast

type IdentifierNode struct {
	Name string
}

func (node IdentifierNode) ToString() string {
	return node.Name
}
