package ast

type ProgramNode struct {
    Title StringConstNode
    Body  ProgramBodyNode
}

func (node ProgramNode) ToString() string {
    return node.Title.ToString() + node.Body.ToString() + "\n"
}

func (node ProgramNode) Accept(v Visitor) {
    v.VisitProgram(node)
}
