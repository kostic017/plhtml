package ast

type BinaryOpExprNode struct {
	LeftExpr  ExpressionNode
	Operator  TokenType
	RightExpr ExpressionNode
}

func (node BinaryOpExprNode) ToString() string {
	return node.LeftExpr.ToString() + " " + string(node.Operator) + " " + node.RightExpr.ToString()
}

func (node BinaryOpExprNode) Accept(v Visitor) {
	v.VisitBinaryOpExpr(node)
}
