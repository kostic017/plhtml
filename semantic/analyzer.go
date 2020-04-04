package semantic

import (
    "fmt"
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

    panic(fmt.Sprintf("Error on line %d: operator %s is not supported for operands of given types", node.GetLine(), node.Operator.String()))
}

func (analyzer *Analyzer) VisitBoolConst(node ast.BoolConstNode) constant.Kind {
    return constant.Bool
}

func (analyzer *Analyzer) VisitControlFlowStmt(node ast.ControlFlowStmtNode) {

    condType := node.Condition.AcceptAnalyzer(analyzer)
    if condType != constant.Bool {
        panic(fmt.Sprintf("Error on line %d: non-bool used as condition", node.GetLine()))
    }

    analyzer.currentScope = NewScope(analyzer.currentScope.id+1, analyzer.currentScope)

    for _, stmt := range node.Statements {
        stmt.AcceptAnalyzer(analyzer)
    }

    analyzer.currentScope = analyzer.currentScope.parent

}

func (analyzer *Analyzer) VisitIdentifier(node ast.IdentifierNode) constant.Kind {
    sym, ok := analyzer.currentScope.lookup(node.Name)

    if !ok {
        panic(fmt.Sprintf("Error on line %d: identifier %s undefined", node.GetLine(), node.Name))
    }

    kind := kindOfPrimitiveType(sym.typeName)

    if kind == constant.Unknown {
        panic(fmt.Sprintf("Error on line %d: only primitive types are supported", node.GetLine()))
    }

    return kind
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
    if _, ok := analyzer.currentScope.lookup(node.Identifier.Name); !ok {
        panic(fmt.Sprintf("Error on line %d: identifier %s undefined", node.GetLine(), node.Identifier.Name))
    }
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

    panic(fmt.Sprintf("Error on line %d: operator %s is not supported for expression of given type", node.GetLine(), node.Operator.String()))
}

func (analyzer *Analyzer) VisitVarAssign(node ast.VarAssignNode) {
    sym, ok := analyzer.currentScope.lookup(node.Identifier.Name)
    if !ok {
        panic(fmt.Sprintf("Error on line %d: identifier %s undefined", node.GetLine(), node.Identifier.Name))
    }

    symType := kindOfPrimitiveType(sym.typeName)

    if symType == constant.Unknown {
        panic(fmt.Sprintf("Error on line %d: only primitive types are supported", node.GetLine()))
    }

    valType := node.Value.AcceptAnalyzer(analyzer)
    ok = (symType == valType) || (isNumType(symType) && isNumType(valType))

    if !ok {
        panic(fmt.Sprintf("Error on line %d: incompatible types for assigment", node.GetLine()))
    }
}

func (analyzer *Analyzer) VisitVarDecl(node ast.VarDeclNode) {
    if _, ok := analyzer.currentScope.lookup(node.Type.Name); !ok {
        panic(fmt.Sprintf("Error on line %d: identifier %s undefined", node.GetLine(), node.Identifier.Name))
    }
    if analyzer.currentScope.declaredLocally(node.Identifier.Name) {
        panic(fmt.Sprintf("Error on line: %d: variable %s is already declared", node.GetLine(), node.Identifier.Name))
    }
    analyzer.currentScope.insert(symbol{name: node.Identifier.Name, typeName: node.Type.Name})
}

func (analyzer *Analyzer) VisitWriteStmt(node ast.WriteStmtNode) {
    valueType := node.Value.AcceptAnalyzer(analyzer)
    if valueType != constant.String {
        panic(fmt.Sprintf("Error on line %d: cannot output non-string", node.GetLine()))
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
    return constant.Unknown
}
