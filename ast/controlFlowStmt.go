package ast

type ControlFlowStmtNode struct {
	Type       TokenType
	Condition  ExpressionNode
	Statements []StatementNode
}

func (node ControlFlowStmtNode) ToString(lvl int) string {
	str := ident(lvl, string(node.Type)+" "+node.Condition.ToString())
	for _, stmt := range node.Statements {
		str += "\n" + stmt.ToString(lvl+1)
	}
	return str
}

func (node ControlFlowStmtNode) Accept(v Visitor) {
	v.VisitControlFlowStmt(node)
}
