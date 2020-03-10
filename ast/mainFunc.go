package ast

type MainFuncNode struct {
	Statements []StatementNode
}

func (node MainFuncNode) ToString() string {
	str := ""
	for _, stmt := range node.Statements {
		str += "\n" + stmt.ToString()
	}
	return str
}
