package ast

import "fmt"

type RealConstNode struct {
	Value float64
}

func (node RealConstNode) ToString() string {
	return fmt.Sprintf("%f", node.Value)
}
