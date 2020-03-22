package ast

type WriteStmtNode struct {
	Value ExpressionNode
}

func (node WriteStmtNode) ToString() string {
	return "Write: " + node.Value.ToString()
}

func (node WriteStmtNode) Accept(v Visitor) {
	v.VisitWriteStmt(node)
}
