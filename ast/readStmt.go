package ast

type ReadStmtNode struct {
	Identifier IdentifierNode
}

func (node ReadStmtNode) ToString(lvl int) string {
	return ident(lvl, "Read "+node.Identifier.ToString())
}

func (node ReadStmtNode) Accept(v Visitor) {
	v.VisitReadStmt(node)
}
