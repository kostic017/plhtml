package parser

import "../scanner"

func (parser *Parser) parseStatements() []StatementNode {
	parser.expect(TokenType('<'))

	switch parser.peek().Type {
	case scanner.TokVar:
		parser.parseVarDecl()
	case scanner.TokData:
		parser.parseVarAssign()
	case scanner.TokOutput:
		parser.parseWriteStmt()
	case scanner.TokInput:
		parser.parseReadStmt()
	case scanner.TokDiv:
		parser.parseControlFlowStmt()
	}

	return nil // TODO
}

func (parser *Parser) parseVarDecl() VarDeclNode {
	parser.expect(scanner.TokVar)
	parser.expect(scanner.TokClass)
	parser.expect(TokenType('='))
	parser.expect(TokenType('"'))
	varType := parser.expect(scanner.TokIntType, scanner.TokRealType, scanner.TokBoolType, scanner.TokStringType)
	parser.expect(TokenType('"'))
	parser.expect(TokenType('>'))
	identifier := parser.parseIdentifier()
	parser.parseCloseTag(scanner.TokVar)
	return VarDeclNode{Type: varType, Identifier: identifier}
}

func (parser *Parser) parseVarAssign() VarAssignNode {
	parser.expect(scanner.TokData)
	parser.expect(scanner.TokValue)
	parser.expect(TokenType('='))
	parser.expect(TokenType('"'))
	value := parser.parseExpr()
	parser.expect(TokenType('"'))
	parser.expect(TokenType('>'))
	identifier := parser.parseIdentifier()
	parser.parseCloseTag(scanner.TokData)
	return VarAssignNode{Identifier: identifier, Value: value}
}

func (parser *Parser) parseWriteStmt() WriteStmtNode {
	parser.expect(scanner.TokOutput)
	parser.expect(TokenType('>'))
	value := parser.parseExpr()
	parser.parseCloseTag(scanner.TokOutput)
	return WriteStmtNode{Value: value}
}

func (parser *Parser) parseReadStmt() ReadStmtNode {
	parser.expect(scanner.TokInput)
	parser.expect(scanner.TokName)
	parser.expect(TokenType('='))
	parser.expect(TokenType('"'))
	identifier := parser.parseIdentifier()
	parser.expect(TokenType('"'))
	parser.expect(TokenType('>'))
	return ReadStmtNode{Identifier: identifier}
}

func (parser *Parser) parseControlFlowStmt() StatementNode {
	// TODO if, if-else
	parser.expect(scanner.TokDiv)
	parser.expect(scanner.TokData)
	parser.expect(TokenType('-'))
	parser.expect(scanner.TokWhile)
	parser.expect(TokenType('='))
	parser.expect(TokenType('"'))
	condition := parser.parseExpr()
	parser.expect(TokenType('"'))
	parser.expect(TokenType('>'))
	statements := parser.parseStatements()
	parser.parseCloseTag(scanner.TokDiv)
	return WhileStmtNode{Condition: condition, Statements: statements}
}
