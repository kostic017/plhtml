package ast

type ProgramBodyNode struct {
	MainFunc MainFuncNode
}

func (node ProgramBodyNode) Print() {
	node.MainFunc.Print()
}
