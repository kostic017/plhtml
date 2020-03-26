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

func (analyzer *Analyzer) VisitBinaryOpExpr(node ast.BinaryOpExprNode) interface{} {
    node.LeftExpr.Accept(analyzer)
    node.RightExpr.Accept(analyzer)
    return nil
}

func (analyzer *Analyzer) VisitBoolConst(node ast.BoolConstNode) interface{} {
    return nil
}

func (analyzer *Analyzer) VisitControlFlowStmt(node ast.ControlFlowStmtNode) {
    node.Condition.Accept(analyzer)

    analyzer.currentScope = NewScope(analyzer.currentScope.id+1, analyzer.currentScope)
    for _, stmt := range node.Statements {
        stmt.Accept(analyzer)
    }
    analyzer.currentScope = analyzer.currentScope.parent
}

func (analyzer *Analyzer) VisitIdentifier(node ast.IdentifierNode) interface{} {
    analyzer.currentScope.expect(node.Name)
    return nil
}

func (analyzer *Analyzer) VisitIntConst(node ast.IntConstNode) interface{} {
    return nil
}

func (analyzer *Analyzer) VisitMainFunc(node ast.MainFuncNode) {
    analyzer.currentScope = NewScope(analyzer.currentScope.id+1, analyzer.currentScope)
    for _, stmt := range node.Statements {
        stmt.Accept(analyzer)
    }
    analyzer.currentScope = analyzer.currentScope.parent
}

func (analyzer *Analyzer) VisitProgram(node ast.ProgramNode) {
    node.Body.Accept(analyzer)
}

func (analyzer *Analyzer) VisitProgramBody(node ast.ProgramBodyNode) {
    node.MainFunc.Accept(analyzer)
}

func (analyzer *Analyzer) VisitReadStmt(node ast.ReadStmtNode) {
    analyzer.currentScope.expect(node.Identifier.Name)
}

func (analyzer *Analyzer) VisitRealConst(node ast.RealConstNode) interface{} {
    return nil
}

func (analyzer *Analyzer) VisitStringConst(node ast.StringConstNode) interface{} {
    return nil
}

func (analyzer *Analyzer) VisitUnaryExpr(node ast.UnaryExprNode) interface{} {
    node.Expr.Accept(analyzer)
    return nil
}

func (analyzer *Analyzer) VisitVarAssign(node ast.VarAssignNode) {
    analyzer.currentScope.expect(node.Identifier.Name)
    node.Value.Accept(analyzer)
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
    node.Value.Accept(analyzer)
}
