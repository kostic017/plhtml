package ast

type MainFuncNode struct {
	Statements []StatementNode
}

func (node MainFuncNode) ToString() string {
	str := "\nMain:"
	for _, stmt := range node.Statements {
		str += "\n" + stmt.ToString(1)
	}
	return str
}

func (node MainFuncNode) Accept(v Visitor) {
	v.VisitMainFunc(node)
}
