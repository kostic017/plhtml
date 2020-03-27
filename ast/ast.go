package ast

import (
    "go/constant"
    "strings"

    "plhtml/token"
)

type TokenType = token.Type

type Node interface {
}

type StatementNode interface {
    Node
    ToString(lvl int) string
    AcceptAnalyzer(analyzer IAnalyzer)
    AcceptInterpreter(interp IInterpreter)
}

type ExpressionNode interface {
    Node
    ToString() string
    AcceptAnalyzer(analyzer IAnalyzer)
    AcceptInterpreter(interp IInterpreter) constant.Value
}

type IAnalyzer interface {
    VisitBinaryOpExpr(node BinaryOpExprNode)
    VisitBoolConst(node BoolConstNode)
    VisitControlFlowStmt(node ControlFlowStmtNode)
    VisitIdentifier(node IdentifierNode)
    VisitIntConst(node IntConstNode)
    VisitMainFunc(node MainFuncNode)
    VisitProgram(node ProgramNode)
    VisitProgramBody(node ProgramBodyNode)
    VisitReadStmt(node ReadStmtNode)
    VisitRealConst(node RealConstNode)
    VisitStringConst(node StringConstNode)
    VisitUnaryExpr(node UnaryExprNode)
    VisitVarAssign(node VarAssignNode)
    VisitVarDecl(node VarDeclNode)
    VisitWriteStmt(node WriteStmtNode)
}

type IInterpreter interface {
    VisitBinaryOpExpr(node BinaryOpExprNode) constant.Value
    VisitBoolConst(node BoolConstNode) constant.Value
    VisitControlFlowStmt(node ControlFlowStmtNode)
    VisitIdentifier(node IdentifierNode) constant.Value
    VisitIntConst(node IntConstNode) constant.Value
    VisitMainFunc(node MainFuncNode)
    VisitProgram(node ProgramNode)
    VisitProgramBody(node ProgramBodyNode)
    VisitReadStmt(node ReadStmtNode)
    VisitRealConst(node RealConstNode) constant.Value
    VisitStringConst(node StringConstNode) constant.Value
    VisitUnaryExpr(node UnaryExprNode) constant.Value
    VisitVarAssign(node VarAssignNode)
    VisitVarDecl(node VarDeclNode)
    VisitWriteStmt(node WriteStmtNode)
}

func ident(lvl int, str string) string {
    return strings.Repeat(" ", lvl*4) + str
}
