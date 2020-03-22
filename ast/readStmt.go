package ast

type ReadStmtNode struct {
	Identifier IdentifierNode
}

func (node ReadStmtNode) ToString() string {
	return "Read " + node.Identifier.ToString()
}

func (node ReadStmtNode) Accept(v Visitor) {
	v.VisitReadStmt(node)
}
