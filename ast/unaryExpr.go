package ast

import "fmt"

type UnaryExprNode struct {
	Operator TokenType
	Expr     ExpressionNode
}

func (node UnaryExprNode) Print() {
	fmt.Print(string(node.Operator))
	node.Expr.Print()
}
