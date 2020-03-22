package ast

type VarAssignNode struct {
	Identifier IdentifierNode
	Value      ExpressionNode
}

func (node VarAssignNode) ToString(lvl int) string {
	return ident(lvl, node.Identifier.ToString()+" = "+node.Value.ToString())
}

func (node VarAssignNode) Accept(v Visitor) {
	v.VisitVarAssign(node)
}
