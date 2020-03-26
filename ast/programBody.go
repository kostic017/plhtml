package ast

type ProgramBodyNode struct {
	MainFunc MainFuncNode
}

func (node ProgramBodyNode) ToString() string {
	return node.MainFunc.ToString()
}

func (node ProgramBodyNode) Accept(v Visitor) {
	v.VisitProgramBody(node)
}
