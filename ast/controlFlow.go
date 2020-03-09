package ast

import "fmt"

type ControlFlowStmtNode struct {
	Type       TokenType
	Condition  ExpressionNode
	Statements []StatementNode
}

func (node ControlFlowStmtNode) Print() {
	fmt.Print(string(node.Type))
	fmt.Print("(")
	node.Condition.Print()
	fmt.Print(")")
	for _, stmt := range node.Statements {
		stmt.Print()
	}
}
