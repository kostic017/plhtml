package ast

import "fmt"

type BoolConstNode struct {
	Value bool
}

func (node BoolConstNode) Print() {
	fmt.Printf("%t", node.Value)
}
