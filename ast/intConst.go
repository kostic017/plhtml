package ast

import "fmt"

type IntConstNode struct {
    Value int
}

func (node IntConstNode) ToString() string {
    return fmt.Sprintf("%d", node.Value)
}

func (node IntConstNode) Accept(v Visitor) {
    v.VisitIntConst(node)
}
