package ast

import "go/constant"

type StringConstNode struct {
    Value string
}

func (node StringConstNode) ToString() string {
    return "\"" + node.Value + "\""
}

func (node StringConstNode) AcceptAnalyzer(analyzer IAnalyzer) constant.Kind {
    return analyzer.VisitStringConst(node)
}

func (node StringConstNode) AcceptInterpreter(v IInterpreter) constant.Value {
    return v.VisitStringConst(node)
}
