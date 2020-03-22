package ast

import "../scanner"

type TokenType = scanner.TokenType

type AstNode interface {
	ToString() string
	Accept(v Visitor)
}

type StatementNode interface {
	AstNode
}

type ExpressionNode interface {
	AstNode
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
