package ast

import (
    "fmt"
    "go/constant"
)

type BoolConstNode struct {
    Value bool
}

func (node BoolConstNode) ToString() string {
    return fmt.Sprintf("%t", node.Value)
}

func (node BoolConstNode) AcceptAnalyzer(analyzer IAnalyzer) constant.Kind {
    return analyzer.VisitBoolConst(node)
}

func (node BoolConstNode) AcceptInterpreter(interp IInterpreter) constant.Value {
    return interp.VisitBoolConst(node)
}
