package ast

type VarDeclNode struct {
	TypeName IdentifierNode
	VarName  IdentifierNode
}

func (node VarDeclNode) ToString(lvl int) string {
	return ident(lvl, node.TypeName.ToString()+" "+node.VarName.ToString())
}

func (node VarDeclNode) Accept(v Visitor) {
	v.VisitVarDecl(node)
}
