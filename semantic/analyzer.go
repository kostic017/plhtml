package semantic

import (
    "plhtml/ast"
    "plhtml/logger"
)

var myLogger = logger.New("ANALYZER")

func SetLogLevel(level logger.LogLevel) {
    myLogger.SetLevel(level)
}

type Analyzer struct {
    rootScope    *Scope
    currentScope *Scope
}

func NewAnalyzer() *Analyzer {
    analyzer := new(Analyzer)
    analyzer.rootScope = NewScope(0, nil)
    analyzer.rootScope.symbols["integer"] = symbol{"integer"}
    analyzer.rootScope.symbols["real"] = symbol{"real"}
    analyzer.rootScope.symbols["boolean"] = symbol{"boolean"}
    analyzer.rootScope.symbols["string"] = symbol{"string"}
    analyzer.currentScope = NewScope(1, analyzer.rootScope)
    return analyzer
}

func (analyzer *Analyzer) VisitBinaryOpExpr(node ast.BinaryOpExprNode) {
    node.LeftExpr.AcceptAnalyzer(analyzer)
    node.RightExpr.AcceptAnalyzer(analyzer)
}

func (analyzer *Analyzer) VisitBoolConst(node ast.BoolConstNode) {
}

func (analyzer *Analyzer) VisitControlFlowStmt(node ast.ControlFlowStmtNode) {
    node.Condition.AcceptAnalyzer(analyzer)

    analyzer.currentScope = NewScope(analyzer.currentScope.id+1, analyzer.currentScope)
    for _, stmt := range node.Statements {
        stmt.AcceptAnalyzer(analyzer)
    }
    analyzer.currentScope = analyzer.currentScope.parent
}

func (analyzer *Analyzer) VisitIdentifier(node ast.IdentifierNode) {
    analyzer.currentScope.expect(node.Name)
}

func (analyzer *Analyzer) VisitIntConst(node ast.IntConstNode) {
}

func (analyzer *Analyzer) VisitMainFunc(node ast.MainFuncNode) {
    analyzer.currentScope = NewScope(analyzer.currentScope.id+1, analyzer.currentScope)
    for _, stmt := range node.Statements {
        stmt.AcceptAnalyzer(analyzer)
    }
    analyzer.currentScope = analyzer.currentScope.parent
}

func (analyzer *Analyzer) VisitProgram(node ast.ProgramNode) {
    node.Body.AcceptAnalyzer(analyzer)
}

func (analyzer *Analyzer) VisitProgramBody(node ast.ProgramBodyNode) {
    node.MainFunc.AcceptAnalyzer(analyzer)
}

func (analyzer *Analyzer) VisitReadStmt(node ast.ReadStmtNode) {
    analyzer.currentScope.expect(node.Identifier.Name)
}

func (analyzer *Analyzer) VisitRealConst(node ast.RealConstNode) {
}

func (analyzer *Analyzer) VisitStringConst(node ast.StringConstNode) {
}

func (analyzer *Analyzer) VisitUnaryExpr(node ast.UnaryExprNode) {
    node.Expr.AcceptAnalyzer(analyzer)
}

func (analyzer *Analyzer) VisitVarAssign(node ast.VarAssignNode) {
    analyzer.currentScope.expect(node.Identifier.Name)
    node.Value.AcceptAnalyzer(analyzer)
}

func (analyzer *Analyzer) VisitVarDecl(node ast.VarDeclNode) {
    analyzer.currentScope.expect(node.Type.Name)
    name := node.Identifier.Name
    if analyzer.currentScope.declaredLocally(name) {
        panic("Variable " + name + " is already declared.")
    }
    analyzer.currentScope.insert(symbol{name: name})
}

func (analyzer *Analyzer) VisitWriteStmt(node ast.WriteStmtNode) {
    node.Value.AcceptAnalyzer(analyzer)
}
