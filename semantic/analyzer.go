package semantic

import (
    "go/constant"
    "plhtml/ast"
    "plhtml/logger"
    "plhtml/token"
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
    analyzer.rootScope.symbols[typeInteger] = symbol{name: typeInteger}
    analyzer.rootScope.symbols[typeReal] = symbol{name: typeReal}
    analyzer.rootScope.symbols[typeBoolean] = symbol{name: typeBoolean}
    analyzer.rootScope.symbols[typeString] = symbol{name: typeString}
    analyzer.currentScope = NewScope(1, analyzer.rootScope)
    return analyzer
}

func (analyzer *Analyzer) VisitBinaryOpExpr(node ast.BinaryOpExprNode) constant.Kind {
    leftType := node.LeftExpr.AcceptAnalyzer(analyzer)
    rightType := node.RightExpr.AcceptAnalyzer(analyzer)

    switch node.Operator {
    case token.EqOp, token.NeqOp:
        if leftType == rightType {
            return constant.Bool
        }
    case token.LtOp, token.GtOp, token.LeqOp, token.GeqOp:
        if isNumType(leftType) && isNumType(rightType) {
            return constant.Bool
        }
    case token.Plus:
        if leftType == constant.String || rightType == constant.String {
            return constant.String
        }
        fallthrough
    case token.Minus, token.Multiply, token.Slash, token.Modulo:
        if leftType == constant.Int && rightType == constant.Int {
            return constant.Int
        }
        if isNumType(leftType) && isNumType(rightType) {
            return constant.Float
        }
    }

    panic("Operator " + node.Operator.String() + " is not supported for operands of given types.")
}

func (analyzer *Analyzer) VisitBoolConst(node ast.BoolConstNode) constant.Kind {
    return constant.Bool
}

func (analyzer *Analyzer) VisitControlFlowStmt(node ast.ControlFlowStmtNode) {

    condType := node.Condition.AcceptAnalyzer(analyzer)
    if condType != constant.Bool {
        panic("Non-bool used as condition.")
    }

    analyzer.currentScope = NewScope(analyzer.currentScope.id+1, analyzer.currentScope)

    for _, stmt := range node.Statements {
        stmt.AcceptAnalyzer(analyzer)
    }

    analyzer.currentScope = analyzer.currentScope.parent

}

func (analyzer *Analyzer) VisitIdentifier(node ast.IdentifierNode) constant.Kind {
    sym := analyzer.currentScope.expect(node.Name)
    return kindOfPrimitiveType(sym.typeName)
}

func (analyzer *Analyzer) VisitIntConst(node ast.IntConstNode) constant.Kind {
    return constant.Int
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

func (analyzer *Analyzer) VisitRealConst(node ast.RealConstNode) constant.Kind {
    return constant.Float
}

func (analyzer *Analyzer) VisitStringConst(node ast.StringConstNode) constant.Kind {
    return constant.String
}

func (analyzer *Analyzer) VisitUnaryExpr(node ast.UnaryExprNode) constant.Kind {
    exprType := node.Expr.AcceptAnalyzer(analyzer)

    switch node.Operator {
    case token.Minus:
        if isNumType(exprType) {
            return exprType
        }
    case token.Excl:
        if exprType == constant.Bool {
            return exprType
        }
    }

    panic("Unary operator " + node.Operator.String() + " is not supported for expression of given type.")
}

func (analyzer *Analyzer) VisitVarAssign(node ast.VarAssignNode) {
    sym := analyzer.currentScope.expect(node.Identifier.Name)
    symType := kindOfPrimitiveType(sym.typeName)
    valType := node.Value.AcceptAnalyzer(analyzer)
    ok := (symType == valType) || (isNumType(symType) && isNumType(valType))
    if !ok {
        panic("Incompatible types for assigment")
    }
}

func (analyzer *Analyzer) VisitVarDecl(node ast.VarDeclNode) {
    analyzer.currentScope.expect(node.Type.Name)
    if analyzer.currentScope.declaredLocally(node.Identifier.Name) {
        panic("Variable " + node.Identifier.Name + " is already declared.")
    }
    analyzer.currentScope.insert(symbol{name: node.Identifier.Name, typeName: node.Type.Name})
}

func (analyzer *Analyzer) VisitWriteStmt(node ast.WriteStmtNode) {
    valueType := node.Value.AcceptAnalyzer(analyzer)
    if valueType != constant.String {
        panic("Cannot output non-string.")
    }
}

func isNumType(kind constant.Kind) bool {
    return kind == constant.Int || kind == constant.Float
}

func kindOfPrimitiveType(typeName string) constant.Kind {
    switch typeName {
    case typeInteger:
        return constant.Int
    case typeReal:
        return constant.Float
    case typeBoolean:
        return constant.Bool
    case typeString:
        return constant.String
    }
    panic(typeName + " is not primitive type.")
}
