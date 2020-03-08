package parser

import "../scanner"

func (parser *Parser) parseStatements() []StatementNode {
	parser.logger.Debug("=BEG= Statements")

	var stmts []StatementNode
	for stmt := parser.parseStatement(); stmt != nil; stmt = parser.parseStatement() {
		stmts = append(stmts, stmt)
	}

	parser.logger.Debug("=END= Statements")
	return stmts
}

func (parser *Parser) parseStatement() StatementNode {
	parser.expect(TokenType('<'))

	switch parser.peek().Type {
	case scanner.TokVar:
		return parser.parseVarDecl()
	case scanner.TokData:
		return parser.parseVarAssign()
	case scanner.TokOutput:
		return parser.parseWriteStmt()
	case scanner.TokInput:
		return parser.parseReadStmt()
	case scanner.TokDiv:
		return parser.parseControlFlowStmt()
	}

	parser.logger.Debug("'%s' not a statement", string(parser.peek().Type))
	parser.goBack() // '<'
	return nil
}

func (parser *Parser) parseVarDecl() VarDeclNode {
	parser.logger.Debug("=BEG= Var Declaration")
	parser.expect(scanner.TokVar)
	parser.expect(scanner.TokClass)
	parser.expect(TokenType('='))
	parser.expect(TokenType('"'))
	varType := parser.expect(scanner.TokIntType, scanner.TokRealType, scanner.TokBoolType, scanner.TokStringType)
	parser.expect(TokenType('"'))
	parser.expect(TokenType('>'))
	identifier := parser.parseIdentifier()
	parser.parseCloseTag(scanner.TokVar)
	parser.logger.Debug("=END= Var Declaration")
	return VarDeclNode{Type: varType, Identifier: identifier}
}

func (parser *Parser) parseVarAssign() VarAssignNode {
	parser.logger.Debug("=BEG= Var Assignment")
	parser.expect(scanner.TokData)
	parser.expect(scanner.TokValue)
	parser.expect(TokenType('='))
	parser.expect(TokenType('"'))
	value := parser.parseExpr()
	parser.expect(TokenType('"'))
	parser.expect(TokenType('>'))
	identifier := parser.parseIdentifier()
	parser.parseCloseTag(scanner.TokData)
	parser.logger.Debug("=END= Var Assignment")
	return VarAssignNode{Identifier: identifier, Value: value}
}

func (parser *Parser) parseWriteStmt() WriteStmtNode {
	parser.logger.Debug("=BEG= Write")
	parser.expect(scanner.TokOutput)
	parser.expect(TokenType('>'))
	value := parser.parseExpr()
	parser.parseCloseTag(scanner.TokOutput)
	parser.logger.Debug("=END= Write")
	return WriteStmtNode{Value: value}
}

func (parser *Parser) parseReadStmt() ReadStmtNode {
	parser.logger.Debug("=BEG= Read")
	parser.expect(scanner.TokInput)
	parser.expect(scanner.TokName)
	parser.expect(TokenType('='))
	parser.expect(TokenType('"'))
	identifier := parser.parseIdentifier()
	parser.expect(TokenType('"'))
	parser.expect(TokenType('>'))
	parser.logger.Debug("=END= Read")
	return ReadStmtNode{Identifier: identifier}
}

func (parser *Parser) parseControlFlowStmt() ControlFlowStmtNode {
	parser.logger.Debug("=BEG= Control Flow")
	parser.expect(scanner.TokDiv)
	parser.expect(scanner.TokData)
	parser.expect(TokenType('-'))
	stmtType := parser.expect(scanner.TokIf, scanner.TokWhile)
	parser.expect(TokenType('='))
	parser.expect(TokenType('"'))
	condition := parser.parseExpr()
	parser.expect(TokenType('"'))
	parser.expect(TokenType('>'))
	stmts := parser.parseStatements()
	parser.parseCloseTag(scanner.TokDiv)
	parser.logger.Debug("=END= Control Flow")
	return ControlFlowStmtNode{Type: stmtType, Condition: condition, Statements: stmts}
}
