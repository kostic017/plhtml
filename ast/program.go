package ast

type ProgramNode struct {
	Title StringConstNode
	Body  ProgramBodyNode
}

func (node ProgramNode) Print() {
	node.Title.Print()
	node.Body.Print()
}
