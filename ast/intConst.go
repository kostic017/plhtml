package ast

import "fmt"

type IntConstNode struct {
	Value int
}

func (node IntConstNode) Print() {
	fmt.Printf("%d", node.Value)
}
