package semantic

import "../ast"

type Analyzer struct {
    scope *SymbolTable
}

func NewAnalyzer() *Analyzer {
    v := new(Analyzer)
    v.scope = NewSymbolTable(0)
    return v
}

func (v *Analyzer) VisitBinaryOpExpr(node ast.BinaryOpExprNode) {
    node.LeftExpr.Accept(v)
    node.RightExpr.Accept(v)
}

func (v *Analyzer) VisitBoolConst(node ast.BoolConstNode) {
}

func (v *Analyzer) VisitControlFlowStmt(node ast.ControlFlowStmtNode) {
    node.Condition.Accept(v)
    for _, stmt := range node.Statements {
        stmt.Accept(v)
    }
}

func (v *Analyzer) VisitIdentifier(node ast.IdentifierNode) {
}

func (v *Analyzer) VisitIntConst(node ast.IntConstNode) {
}

func (v *Analyzer) VisitMainFunc(node ast.MainFuncNode) {
    for _, stmt := range node.Statements {
        stmt.Accept(v)
    }
}

func (v *Analyzer) VisitProgram(node ast.ProgramNode) {
    node.Body.Accept(v)
}

func (v *Analyzer) VisitProgramBody(node ast.ProgramBodyNode) {
    node.MainFunc.Accept(v)
}

func (v *Analyzer) VisitReadStmt(node ast.ReadStmtNode) {
    v.scope.expect(node.Identifier.Name)
}

func (v *Analyzer) VisitRealConst(node ast.RealConstNode) {
}

func (v *Analyzer) VisitStringConst(node ast.StringConstNode) {
}

func (v *Analyzer) VisitUnaryExpr(node ast.UnaryExprNode) {
    node.Expr.Accept(v)
}

func (v *Analyzer) VisitVarAssign(node ast.VarAssignNode) {
    v.scope.expect(node.Identifier.Name)
    node.Value.Accept(v)
}

func (v *Analyzer) VisitVarDecl(node ast.VarDeclNode) {
    name := node.Identifier.Name
    if _, ok := v.scope.lookup(name); ok {
        panic("Identifier " + name + " is already declared.")
    }
    v.scope.insert(symbol{name: name})
}

func (v *Analyzer) VisitWriteStmt(node ast.WriteStmtNode) {
}
