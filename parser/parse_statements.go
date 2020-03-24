package parser

import (
	"../ast"
	"../token"
)

func (parser *Parser) parseStatements() []ast.StatementNode {
	myLogger.Debug("=BEG= Statements")

	var stmts []ast.StatementNode
	for stmt := parser.parseStatement(); stmt != nil; stmt = parser.parseStatement() {
		stmts = append(stmts, stmt)
	}

	myLogger.Debug("=END= Statements")
	return stmts
}

func (parser *Parser) parseStatement() ast.StatementNode {
	parser.eat(token.LessThan)

	switch parser.peek().Type {
	case token.Var:
		return parser.parseVarDecl()
	case token.Data:
		return parser.parseVarAssign()
	case token.Output:
		return parser.parseWriteStmt()
	case token.Input:
		return parser.parseReadStmt()
	case token.Div:
		return parser.parseControlFlowStmt()
	}

	myLogger.Debug("'%s' not a statement", string(parser.peek().Type))
	parser.goBack() // '<'
	return nil
}

func (parser *Parser) parseVarDecl() ast.VarDeclNode {
	myLogger.Debug("=BEG= Var Declaration")
	parser.eat(token.Var)
	parser.eat(token.Class)
	parser.eat(token.Equal)
	parser.eat(token.DQuote)
	varType := parser.parseIdentifier()
	parser.eat(token.DQuote)
	parser.eat(token.GreaterThan)
	varName := parser.parseIdentifier()
	parser.parseCloseTag(token.Var)
	myLogger.Debug("=END= Var Declaration")
	return ast.VarDeclNode{TypeName: varType, VarName: varName}
}

func (parser *Parser) parseVarAssign() ast.VarAssignNode {
	myLogger.Debug("=BEG= Var Assignment")
	parser.eat(token.Data)
	parser.eat(token.Value)
	parser.eat(token.Equal)
	parser.eat(token.DQuote)
	value := parser.parseExpr()
	parser.eat(token.DQuote)
	parser.eat(token.GreaterThan)
	identifier := parser.parseIdentifier()
	parser.parseCloseTag(token.Data)
	myLogger.Debug("=END= Var Assignment")
	return ast.VarAssignNode{Identifier: identifier, Value: value}
}

func (parser *Parser) parseWriteStmt() ast.WriteStmtNode {
	myLogger.Debug("=BEG= Write")
	parser.eat(token.Output)
	parser.eat(token.GreaterThan)
	value := parser.parseExpr()
	parser.parseCloseTag(token.Output)
	myLogger.Debug("=END= Write")
	return ast.WriteStmtNode{Value: value}
}

func (parser *Parser) parseReadStmt() ast.ReadStmtNode {
	myLogger.Debug("=BEG= Read")
	parser.eat(token.Input)
	parser.eat(token.Name)
	parser.eat(token.Equal)
	parser.eat(token.DQuote)
	identifier := parser.parseIdentifier()
	parser.eat(token.DQuote)
	parser.eat(token.GreaterThan)
	myLogger.Debug("=END= Read")
	return ast.ReadStmtNode{Identifier: identifier}
}

func (parser *Parser) parseControlFlowStmt() ast.ControlFlowStmtNode {
	myLogger.Debug("=BEG= Control Flow")
	parser.eat(token.Div)
	parser.eat(token.Data)
	parser.eat(token.Minus)
	stmtType := parser.eat(token.If, token.While)
	parser.eat(token.Equal)
	parser.eat(token.DQuote)
	condition := parser.parseExpr()
	parser.eat(token.DQuote)
	parser.eat(token.GreaterThan)
	stmts := parser.parseStatements()
	parser.parseCloseTag(token.Div)
	myLogger.Debug("=END= Control Flow")
	return ast.ControlFlowStmtNode{Type: stmtType, Condition: condition, Statements: stmts}
}
