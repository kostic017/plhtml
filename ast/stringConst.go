package ast

import "fmt"

type StringConstNode struct {
	Value string
}

func (node StringConstNode) Print() {
	fmt.Printf("\"%s\"", node.Value)
}
