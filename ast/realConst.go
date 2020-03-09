package ast

import "fmt"

type RealConstNode struct {
	Value float64
}

func (node RealConstNode) Print() {
	fmt.Printf("%f", node.Value)
}
