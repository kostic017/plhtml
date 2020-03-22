package ast

import (
	"strings"

	"../scanner"
)

type TokenType = scanner.TokenType

type AstNode interface {
	Accept(v Visitor)
}

type StatementNode interface {
	AstNode
	ToString(lvl int) string
}

type ExpressionNode interface {
	AstNode
	ToString() string
}

type Visitor interface {
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

func ident(lvl int, str string) string {
	return strings.Repeat(" ", lvl*4) + str
}
