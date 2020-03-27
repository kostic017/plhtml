package ast

import "go/constant"

type BinaryOpExprNode struct {
    LeftExpr  ExpressionNode
    Operator  TokenType
    RightExpr ExpressionNode
}

func (node BinaryOpExprNode) ToString() string {
    return node.LeftExpr.ToString() + " " + node.Operator.String() + " " + node.RightExpr.ToString()
}

func (node BinaryOpExprNode) AcceptAnalyzer(analyzer IAnalyzer) {
    analyzer.VisitBinaryOpExpr(node)
}

func (node BinaryOpExprNode) AcceptInterpreter(interp IInterpreter) constant.Value {
    return interp.VisitBinaryOpExpr(node)
}
