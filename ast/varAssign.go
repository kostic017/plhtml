package ast

type VarAssignNode struct {
    Identifier IdentifierNode
    Value      ExpressionNode
}

func (node VarAssignNode) ToString(lvl int) string {
    return ident(lvl, node.Identifier.ToString()+" = "+node.Value.ToString())
}

func (node VarAssignNode) AcceptAnalyzer(analyzer IAnalyzer) {
    analyzer.VisitVarAssign(node)
}

func (node VarAssignNode) AcceptInterpreter(interp IInterpreter) {
    interp.VisitVarAssign(node)
}
