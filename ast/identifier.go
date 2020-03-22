package ast

type IdentifierNode struct {
	Name string
}

func (node IdentifierNode) ToString() string {
	return node.Name
}

func (node IdentifierNode) Accept(v Visitor) {
	v.VisitIdentifier(node)
}
