package ast

import "fmt"

type BoolConstNode struct {
    Value bool
}

func (node BoolConstNode) ToString() string {
    return fmt.Sprintf("%t", node.Value)
}

func (node BoolConstNode) Accept(v Visitor) {
    v.VisitBoolConst(node)
}
