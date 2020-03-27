package ast

type ReadStmtNode struct {
    Identifier IdentifierNode
}

func (node ReadStmtNode) ToString(lvl int) string {
    return ident(lvl, "Read "+node.Identifier.ToString())
}

func (node ReadStmtNode) AcceptAnalyzer(analyzer IAnalyzer) {
    analyzer.VisitReadStmt(node)
}

func (node ReadStmtNode) AcceptInterpreter(interp IInterpreter) {
    interp.VisitReadStmt(node)
}
