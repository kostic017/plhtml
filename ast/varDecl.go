package ast

type VarDeclNode struct {
	Identifier IdentifierNode
	Type       TokenType
}

func (node VarDeclNode) ToString() string {
	return string(node.Type) + " " + node.Identifier.ToString()
}

func (node VarDeclNode) Accept(v Visitor) {
	v.VisitVarDecl(node)
}
