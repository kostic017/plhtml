package ast

type WriteStmtNode struct {
	Value ExpressionNode
}

func (node WriteStmtNode) ToString(lvl int) string {
	return ident(lvl, "Write: "+node.Value.ToString())
}

func (node WriteStmtNode) Accept(v Visitor) interface{} {
	return v.VisitWriteStmt(node)
}
