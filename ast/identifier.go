package ast

import "fmt"

type IdentifierNode struct {
	Name string
}

func (node IdentifierNode) Print() {
	fmt.Printf(node.Name)
}
