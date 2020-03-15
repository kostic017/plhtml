package parser

import (
	"fmt"

	"../ast"
	"../scanner"
)

func (parser *Parser) parseExpr() ast.ExpressionNode {
	return parser.parseExprL(1)
}

func (parser *Parser) parseExprL(l int) ast.ExpressionNode {
	var operators []TokenType

	switch l {
	case 1:
		operators = []TokenType{scanner.TokEqOp, scanner.TokNeqOp}
	case 2:
		operators = []TokenType{scanner.TokLtOp, scanner.TokGtOp, scanner.TokLeqOp, scanner.TokGeqOp}
	case 3:
		operators = []TokenType{TokenType('+'), TokenType('-')}
	case 4:
		operators = []TokenType{TokenType('*'), TokenType('/')}
	case 5:
		return parser.parseFactor()
	}

	expr := parser.parseExprL(l + 1)
	if parser.expectOpt(operators...) {
		return ast.BinaryOpExprNode{LeftExpr: expr, Operator: parser.next().Type, RightExpr: parser.parseExprL(l + 1)}
	}
	return expr
}

func (parser *Parser) parseFactor() ast.ExpressionNode {

	switch parser.peek().Type {
	case scanner.TokIntConst:
		return parser.parseIntConst()
	case scanner.TokRealConst:
		return parser.parseRealConst()
	case scanner.TokBoolConst:
		return parser.parseBoolConst()
	case scanner.TokStringConst:
		return parser.parseStringConst()
	case scanner.TokIdentifier:
		return parser.parseIdentifier()
	case TokenType('('):
		return parser.parseParenExpr()
	case TokenType('+'), TokenType('-'), TokenType('!'):
		return parser.parseUnaryExpr()
	}

	panic(fmt.Sprintf("Invalid factor '%s'.", string(parser.peek().Type)))
}

func (parser *Parser) parseIdentifier() ast.IdentifierNode {
	parser.logger.Debug("=BEG= Identifier")
	parser.expect(scanner.TokIdentifier)
	parser.logger.Debug("=END= Identifier")
	return ast.IdentifierNode{Name: parser.current().StrVal}
}

func (parser *Parser) parseIntConst() ast.IntConstNode {
	parser.logger.Debug("=BEG= Integer Constant")
	parser.expect(scanner.TokIntConst)
	parser.logger.Debug("=END= Integer Constant")
	return ast.IntConstNode{Value: parser.current().IntVal}
}

func (parser *Parser) parseRealConst() ast.RealConstNode {
	parser.logger.Debug("=BEG= Real Constant")
	parser.expect(scanner.TokRealConst)
	parser.logger.Debug("=END= Real Constant")
	return ast.RealConstNode{Value: parser.current().RealVal}
}

func (parser *Parser) parseBoolConst() ast.BoolConstNode {
	parser.logger.Debug("=BEG= Boolean Constant")
	parser.expect(scanner.TokBoolConst)
	parser.logger.Debug("=END= Boolean Constant")
	return ast.BoolConstNode{Value: parser.current().BoolVal}
}

func (parser *Parser) parseStringConst() ast.StringConstNode {
	parser.logger.Debug("=BEG= String Constant")
	parser.expect(scanner.TokStringConst)
	parser.logger.Debug("=END= String Constant")
	return ast.StringConstNode{Value: parser.current().StrVal}
}

func (parser *Parser) parseParenExpr() ast.ExpressionNode {
	parser.logger.Debug("=BEG= ()")
	parser.expect(TokenType('('))
	expr := parser.parseExpr()
	parser.expect(TokenType(')'))
	parser.logger.Debug("=END= ()")
	return expr
}

func (parser *Parser) parseUnaryExpr() ast.UnaryExprNode {
	parser.logger.Debug("=BEG= Unary")
	op := parser.expect(TokenType('+'), TokenType('-'), TokenType('!'))
	expr := parser.parseFactor()
	parser.logger.Debug("=END= Unary")
	return ast.UnaryExprNode{Operator: op, Expr: expr}
}
