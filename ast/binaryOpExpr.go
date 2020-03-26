package ast

import (
    "../token"
)

type BinaryOpExprNode struct {
    LeftExpr  ExpressionNode
    Operator  TokenType
    RightExpr ExpressionNode
}

func (node BinaryOpExprNode) ToString() string {
    return node.LeftExpr.ToString() + " " + token.TypeToStr[node.Operator] + " " + node.RightExpr.ToString()
}

func (node BinaryOpExprNode) Accept(v Visitor) interface{} {
    return v.VisitBinaryOpExpr(node)
}
