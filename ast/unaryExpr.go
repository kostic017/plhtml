package ast

type UnaryExprNode struct {
	Operator TokenType
	Expr     ExpressionNode
}

func (node UnaryExprNode) ToString() string {
	return string(node.Operator) + node.Expr.ToString()
}
