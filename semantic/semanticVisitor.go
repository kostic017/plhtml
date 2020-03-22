package semantic

import "../ast"

type SemanticVisitor struct {
	st *SymbolTable
}

func NewSemanticVisitor(st *SymbolTable) *SemanticVisitor {
	v := new(SemanticVisitor)
	v.st = st
	return v
}

func (v *SemanticVisitor) VisitBinaryOpExpr(node ast.BinaryOpExprNode) {
	node.LeftExpr.Accept(v)
	node.RightExpr.Accept(v)
}

func (v *SemanticVisitor) VisitBoolConst(node ast.BoolConstNode) {
}

func (v *SemanticVisitor) VisitControlFlowStmt(node ast.ControlFlowStmtNode) {
	node.Condition.Accept(v)
	for _, stmt := range node.Statements {
		stmt.Accept(v)
	}
}

func (v *SemanticVisitor) VisitIdentifier(node ast.IdentifierNode) {
}

func (v *SemanticVisitor) VisitIntConst(node ast.IntConstNode) {
}

func (v *SemanticVisitor) VisitMainFunc(node ast.MainFuncNode) {
	for _, stmt := range node.Statements {
		stmt.Accept(v)
	}
}

func (v *SemanticVisitor) VisitProgram(node ast.ProgramNode) {
	node.Body.Accept(v)
}

func (v *SemanticVisitor) VisitProgramBody(node ast.ProgramBodyNode) {
	node.MainFunc.Accept(v)
}

func (v *SemanticVisitor) VisitReadStmt(node ast.ReadStmtNode) {
	v.st.expect(node.Identifier.Name)
}

func (v *SemanticVisitor) VisitRealConst(node ast.RealConstNode) {
}

func (v *SemanticVisitor) VisitStringConst(node ast.StringConstNode) {
}

func (v *SemanticVisitor) VisitUnaryExpr(node ast.UnaryExprNode) {
	node.Expr.Accept(v)
}

func (v *SemanticVisitor) VisitVarAssign(node ast.VarAssignNode) {
	v.st.expect(node.Identifier.Name)
	node.Value.Accept(v)
}

func (v *SemanticVisitor) VisitVarDecl(node ast.VarDeclNode) {
	name := node.Identifier.Name
	if _, ok := v.st.lookup(name); ok {
		panic("Identifier " + name + " is already declared.")
	}
	var sym symbol
	sym.name = name
	v.st.insert(sym)
}

func (v *SemanticVisitor) VisitWriteStmt(node ast.WriteStmtNode) {
}
