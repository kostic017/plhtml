package parser

import (
    "../ast"
    "../scanner"
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

    myLogger.Debug("'%s' not a statement", string(parser.peek().Type))
    parser.goBack() // '<'
    return nil
}

func (parser *Parser) parseVarDecl() ast.VarDeclNode {
    myLogger.Debug("=BEG= Var Declaration")
    parser.expect(scanner.TokVar)
    parser.expect(scanner.TokClass)
    parser.expect(TokenType('='))
    parser.expect(TokenType('"'))
    varType := parser.expect(scanner.TokIntType, scanner.TokRealType, scanner.TokBoolType, scanner.TokStringType)
    parser.expect(TokenType('"'))
    parser.expect(TokenType('>'))
    identifier := parser.parseIdentifier()
    parser.parseCloseTag(scanner.TokVar)
    myLogger.Debug("=END= Var Declaration")
    return ast.VarDeclNode{Type: varType, Identifier: identifier}
}

func (parser *Parser) parseVarAssign() ast.VarAssignNode {
    myLogger.Debug("=BEG= Var Assignment")
    parser.expect(scanner.TokData)
    parser.expect(scanner.TokValue)
    parser.expect(TokenType('='))
    parser.expect(TokenType('"'))
    value := parser.parseExpr()
    parser.expect(TokenType('"'))
    parser.expect(TokenType('>'))
    identifier := parser.parseIdentifier()
    parser.parseCloseTag(scanner.TokData)
    myLogger.Debug("=END= Var Assignment")
    return ast.VarAssignNode{Identifier: identifier, Value: value}
}

func (parser *Parser) parseWriteStmt() ast.WriteStmtNode {
    myLogger.Debug("=BEG= Write")
    parser.expect(scanner.TokOutput)
    parser.expect(TokenType('>'))
    value := parser.parseExpr()
    parser.parseCloseTag(scanner.TokOutput)
    myLogger.Debug("=END= Write")
    return ast.WriteStmtNode{Value: value}
}

func (parser *Parser) parseReadStmt() ast.ReadStmtNode {
    myLogger.Debug("=BEG= Read")
    parser.expect(scanner.TokInput)
    parser.expect(scanner.TokName)
    parser.expect(TokenType('='))
    parser.expect(TokenType('"'))
    identifier := parser.parseIdentifier()
    parser.expect(TokenType('"'))
    parser.expect(TokenType('>'))
    myLogger.Debug("=END= Read")
    return ast.ReadStmtNode{Identifier: identifier}
}

func (parser *Parser) parseControlFlowStmt() ast.ControlFlowStmtNode {
    myLogger.Debug("=BEG= Control Flow")
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
    myLogger.Debug("=END= Control Flow")
    return ast.ControlFlowStmtNode{Type: stmtType, Condition: condition, Statements: stmts}
}
