package ast

import (
    "fmt"
    "go/constant"
)

type RealConstNode struct {
    Value float64
}

func (node RealConstNode) ToString() string {
    return fmt.Sprintf("%f", node.Value)
}

func (node RealConstNode) AcceptAnalyzer(analyzer IAnalyzer) {
    analyzer.VisitRealConst(node)
}

func (node RealConstNode) AcceptInterpreter(interp IInterpreter) constant.Value {
    return interp.VisitRealConst(node)
}
