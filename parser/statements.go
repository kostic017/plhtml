package parser

import "../scanner"

func (parser *Parser) parseStatements() []StatementNode {
    var stmts []StatementNode
    for stmt := parser.parseStmt(); stmt != nil; stmt = parser.parseStmt() {
        stmts = append(stmts, stmt)
    }
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

    parser.goBack() // '<'
    return nil
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

func (parser *Parser) parseControlFlowStmt() ControlFlowStmtNode {
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
    return ControlFlowStmtNode{Type: stmtType, Condition: condition, Statements: stmts}
}
