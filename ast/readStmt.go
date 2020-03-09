package ast

import "fmt"

type ReadStmtNode struct {
	Identifier IdentifierNode
}

func (node ReadStmtNode) Print() {
	fmt.Print("Read ")
	node.Identifier.Print()
}
