package ast

import (
	"../token"
)

type UnaryExprNode struct {
	Operator TokenType
	Expr     ExpressionNode
}

func (node UnaryExprNode) ToString() string {
	return token.TypeToStr[node.Operator] + node.Expr.ToString()
}

func (node UnaryExprNode) Accept(v Visitor) interface{} {
	return v.VisitUnaryExpr(node)
}
