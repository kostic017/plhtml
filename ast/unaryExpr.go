package ast

type UnaryExprNode struct {
    Operator TokenType
    Expr     ExpressionNode
}

func (node UnaryExprNode) ToString() string {
    return node.Operator.String() + node.Expr.ToString()
}

func (node UnaryExprNode) Accept(v Visitor) interface{} {
    return v.VisitUnaryExpr(node)
}
