package ast

import "fmt"

type BinaryOpExprNode struct {
	Expr1    ExpressionNode
	Expr2    ExpressionNode
	Operator TokenType
}

func (node BinaryOpExprNode) Print() {
	node.Expr1.Print()
	fmt.Print(string(node.Operator))
	node.Expr2.Print()
}
