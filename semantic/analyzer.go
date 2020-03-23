package semantic

import (
	"../ast"
	"../logger"
)

var myLogger = logger.New("ANALYZER")

func SetLogLevel(level logger.LogLevel) {
	myLogger.SetLevel(level)
}

type Analyzer struct {
	scope *Scope
}

func NewAnalyzer() *Analyzer {
	analyzer := new(Analyzer)
	analyzer.scope = NewScope(0, nil)
	return analyzer
}

func (analyzer *Analyzer) VisitBinaryOpExpr(node ast.BinaryOpExprNode) {
	node.LeftExpr.Accept(analyzer)
	node.RightExpr.Accept(analyzer)
}

func (analyzer *Analyzer) VisitBoolConst(node ast.BoolConstNode) {
}

func (analyzer *Analyzer) VisitControlFlowStmt(node ast.ControlFlowStmtNode) {
	node.Condition.Accept(analyzer)

	analyzer.scope = NewScope(analyzer.scope.id+1, analyzer.scope)
	for _, stmt := range node.Statements {
		stmt.Accept(analyzer)
	}
	analyzer.scope = analyzer.scope.parent
}

func (analyzer *Analyzer) VisitIdentifier(node ast.IdentifierNode) {
	analyzer.scope.expect(node.Name)
}

func (analyzer *Analyzer) VisitIntConst(node ast.IntConstNode) {
}

func (analyzer *Analyzer) VisitMainFunc(node ast.MainFuncNode) {
	analyzer.scope = NewScope(analyzer.scope.id+1, analyzer.scope)
	for _, stmt := range node.Statements {
		stmt.Accept(analyzer)
	}
	analyzer.scope = analyzer.scope.parent
}

func (analyzer *Analyzer) VisitProgram(node ast.ProgramNode) {
	node.Body.Accept(analyzer)
}

func (analyzer *Analyzer) VisitProgramBody(node ast.ProgramBodyNode) {
	node.MainFunc.Accept(analyzer)
}

func (analyzer *Analyzer) VisitReadStmt(node ast.ReadStmtNode) {
	analyzer.scope.expect(node.Identifier.Name)
}

func (analyzer *Analyzer) VisitRealConst(node ast.RealConstNode) {
}

func (analyzer *Analyzer) VisitStringConst(node ast.StringConstNode) {
}

func (analyzer *Analyzer) VisitUnaryExpr(node ast.UnaryExprNode) {
	node.Expr.Accept(analyzer)
}

func (analyzer *Analyzer) VisitVarAssign(node ast.VarAssignNode) {
	analyzer.scope.expect(node.Identifier.Name)
	node.Value.Accept(analyzer)
}

func (analyzer *Analyzer) VisitVarDecl(node ast.VarDeclNode) {
	analyzer.scope.expect(node.Type.Name)
	name := node.Identifier.Name
	if analyzer.scope.declaredLocally(name) {
		panic("Variable " + name + " is already declared.")
	}
	analyzer.scope.insert(symbol{name: name})
}

func (analyzer *Analyzer) VisitWriteStmt(node ast.WriteStmtNode) {
	node.Value.Accept(analyzer)
}
