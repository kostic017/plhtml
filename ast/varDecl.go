package ast

type VarDeclNode struct {
    Type       IdentifierNode
    Identifier IdentifierNode
}

func (node VarDeclNode) ToString(lvl int) string {
    return ident(lvl, node.Type.ToString()+" "+node.Identifier.ToString())
}

func (node VarDeclNode) Accept(v Visitor) {
    v.VisitVarDecl(node)
}
