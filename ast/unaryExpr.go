package ast

import "go/constant"

type UnaryExprNode struct {
    Operator TokenType
    Expr     ExpressionNode
}

func (node UnaryExprNode) ToString() string {
    return node.Operator.String() + node.Expr.ToString()
}

func (node UnaryExprNode) AcceptAnalyzer(analyzer IAnalyzer) constant.Kind {
    return analyzer.VisitUnaryExpr(node)
}

func (node UnaryExprNode) AcceptInterpreter(interp IInterpreter) constant.Value {
    return interp.VisitUnaryExpr(node)
}
