package semantic

import "../ast"

type SemanticAnalyzer struct {
	scope *SymbolTable
}

func NewSemanticAnalyzer() *SemanticAnalyzer {
	v := new(SemanticAnalyzer)
	v.scope = NewSymbolTable(0)
	return v
}

func (v *SemanticAnalyzer) VisitBinaryOpExpr(node ast.BinaryOpExprNode) {
	node.LeftExpr.Accept(v)
	node.RightExpr.Accept(v)
}

func (v *SemanticAnalyzer) VisitBoolConst(node ast.BoolConstNode) {
}

func (v *SemanticAnalyzer) VisitControlFlowStmt(node ast.ControlFlowStmtNode) {
	node.Condition.Accept(v)
	for _, stmt := range node.Statements {
		stmt.Accept(v)
	}
}

func (v *SemanticAnalyzer) VisitIdentifier(node ast.IdentifierNode) {
}

func (v *SemanticAnalyzer) VisitIntConst(node ast.IntConstNode) {
}

func (v *SemanticAnalyzer) VisitMainFunc(node ast.MainFuncNode) {
	for _, stmt := range node.Statements {
		stmt.Accept(v)
	}
}

func (v *SemanticAnalyzer) VisitProgram(node ast.ProgramNode) {
	node.Body.Accept(v)
}

func (v *SemanticAnalyzer) VisitProgramBody(node ast.ProgramBodyNode) {
	node.MainFunc.Accept(v)
}

func (v *SemanticAnalyzer) VisitReadStmt(node ast.ReadStmtNode) {
	v.scope.expect(node.Identifier.Name)
}

func (v *SemanticAnalyzer) VisitRealConst(node ast.RealConstNode) {
}

func (v *SemanticAnalyzer) VisitStringConst(node ast.StringConstNode) {
}

func (v *SemanticAnalyzer) VisitUnaryExpr(node ast.UnaryExprNode) {
	node.Expr.Accept(v)
}

func (v *SemanticAnalyzer) VisitVarAssign(node ast.VarAssignNode) {
	v.scope.expect(node.Identifier.Name)
	node.Value.Accept(v)
}

func (v *SemanticAnalyzer) VisitVarDecl(node ast.VarDeclNode) {
	name := node.Identifier.Name
	if _, ok := v.scope.lookup(name); ok {
		panic("Identifier " + name + " is already declared.")
	}
	v.scope.insert(symbol{name: name})
}

func (v *SemanticAnalyzer) VisitWriteStmt(node ast.WriteStmtNode) {
}
