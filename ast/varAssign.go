package ast

type VarAssignNode struct {
	Identifier IdentifierNode
	Value      ExpressionNode
}

func (node VarAssignNode) ToString() string {
	return node.Identifier.ToString() + " = " + node.Value.ToString()
}

func (node VarAssignNode) Accept(v Visitor) {
	v.VisitVarAssign(node)
}
