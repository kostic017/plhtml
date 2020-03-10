package ast

type WriteStmtNode struct {
	Value ExpressionNode
}

func (node WriteStmtNode) ToString() string {
	return "Write: " + node.Value.ToString()
}
