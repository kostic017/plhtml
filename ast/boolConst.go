package ast

import "fmt"

type BoolConstNode struct {
	Value bool
}

func (node BoolConstNode) ToString() string {
	return fmt.Sprintf("%t", node.Value)
}
