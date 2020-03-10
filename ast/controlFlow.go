package ast

import "fmt"

type ControlFlowStmtNode struct {
	Type       TokenType
	Condition  ExpressionNode
	Statements []StatementNode
}

func (node ControlFlowStmtNode) ToString() string {
	str := fmt.Sprintf("%s (%s)", string(node.Type), node.Condition.ToString())
	for _, stmt := range node.Statements {
		str += "\n    " + stmt.ToString()
	}
	return str
}
