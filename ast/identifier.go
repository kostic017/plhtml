package ast

type IdentifierNode struct {
	Name string
}

func (node IdentifierNode) ToString() string {
	return node.Name
}

func (node IdentifierNode) Accept(v Visitor) interface{} {
	return v.VisitIdentifier(node)
}
