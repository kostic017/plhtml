package ast

type WriteStmtNode struct {
    Value ExpressionNode
}

func (node WriteStmtNode) ToString(lvl int) string {
    return ident(lvl, "Write: "+node.Value.ToString())
}

func (node WriteStmtNode) AcceptAnalyzer(analyzer IAnalyzer) {
    analyzer.VisitWriteStmt(node)
}

func (node WriteStmtNode) AcceptInterpreter(interp IInterpreter) {
    interp.VisitWriteStmt(node)
}
