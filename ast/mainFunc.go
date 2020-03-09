package ast

type MainFuncNode struct {
	Statements []StatementNode
}

func (node MainFuncNode) Print() {
	for _, s := range node.Statements {
		s.Print()
	}
}
