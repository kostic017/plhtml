package ast

type ProgramNode struct {
    Title StringConstNode
    Body  ProgramBodyNode
}

func (node ProgramNode) ToString() string {
    return node.Title.ToString() + node.Body.ToString() + "\n"
}

func (node ProgramNode) AcceptAnalyzer(analyzer IAnalyzer) {
    analyzer.VisitProgram(node)
}

func (node ProgramNode) AcceptInterpreter(interp IInterpreter) {
    interp.VisitProgram(node)
}
