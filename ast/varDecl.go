package ast

import "fmt"

type VarDeclNode struct {
	Identifier IdentifierNode
	Type       TokenType
}

func (node VarDeclNode) Print() {
	fmt.Print(string(node.Type))
	node.Identifier.Print()
}
