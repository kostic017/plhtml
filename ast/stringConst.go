package ast

type StringConstNode struct {
	Value string
}

func (node StringConstNode) ToString() string {
	return "\"" + node.Value + "\""
}

func (node StringConstNode) Accept(v Visitor) {
	v.VisitStringConst(node)
}
