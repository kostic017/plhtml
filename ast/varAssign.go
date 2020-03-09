package ast

import "fmt"

type VarAssignNode struct {
	Identifier IdentifierNode
	Value      ExpressionNode
}

func (node VarAssignNode) Print() {
	node.Identifier.Print()
	fmt.Print("=")
	node.Value.Print()
}
