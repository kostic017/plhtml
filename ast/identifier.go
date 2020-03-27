package ast

import "go/constant"

type IdentifierNode struct {
    Name string
}

func (node IdentifierNode) ToString() string {
    return node.Name
}

func (node IdentifierNode) AcceptAnalyzer(analyzer IAnalyzer) {
    analyzer.VisitIdentifier(node)
}

func (node IdentifierNode) AcceptInterpreter(interp IInterpreter) constant.Value {
    return interp.VisitIdentifier(node)
}
