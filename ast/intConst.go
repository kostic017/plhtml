package ast

import (
    "fmt"
    "go/constant"
)

type IntConstNode struct {
    Value int
}

func (node IntConstNode) ToString() string {
    return fmt.Sprintf("%d", node.Value)
}

func (node IntConstNode) AcceptAnalyzer(analyzer IAnalyzer) constant.Kind {
    return analyzer.VisitIntConst(node)
}

func (node IntConstNode) AcceptInterpreter(interp IInterpreter) constant.Value {
    return interp.VisitIntConst(node)
}
