package ast

type BinaryOpExprNode struct {
	Expr1    ExpressionNode
	Expr2    ExpressionNode
	Operator TokenType
}

func (node BinaryOpExprNode) ToString() string {
	return node.Expr1.ToString() + " " + string(node.Operator) + " " + node.Expr2.ToString()
}
