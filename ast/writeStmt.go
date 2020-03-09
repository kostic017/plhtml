package ast

import "fmt"

type WriteStmtNode struct {
	Value ExpressionNode
}

func (node WriteStmtNode) Print() {
	fmt.Print("Write: ")
	node.Value.Print()
}
