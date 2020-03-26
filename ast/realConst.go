package ast

import "fmt"

type RealConstNode struct {
    Value float64
}

func (node RealConstNode) ToString() string {
    return fmt.Sprintf("%f", node.Value)
}

func (node RealConstNode) Accept(v Visitor) interface{} {
    return v.VisitRealConst(node)
}
