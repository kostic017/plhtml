package parser

import (
	"fmt"

	"../ast"
	"../token"
)

func (parser *Parser) parseExpr() ast.ExpressionNode {
	return parser.parseExprL(1)
}

func (parser *Parser) parseExprL(l int) ast.ExpressionNode {
	var operators []TokenType

	switch l {
	case 1:
		operators = []TokenType{token.EqOp, token.NeqOp}
	case 2:
		operators = []TokenType{token.LtOp, token.GtOp, token.LeqOp, token.GeqOp}
	case 3:
		operators = []TokenType{token.Plus, token.Minus}
	case 4:
		operators = []TokenType{token.Asterisk, token.Slash}
	case 5:
		return parser.parseFactor()
	}

	expr := parser.parseExprL(l + 1)
	if parser.eatOpt(operators...) {
		return ast.BinaryOpExprNode{LeftExpr: expr, Operator: parser.next().Type, RightExpr: parser.parseExprL(l + 1)}
	}
	return expr
}

func (parser *Parser) parseFactor() ast.ExpressionNode {
	nextToken := parser.peek()
	switch nextToken.Type {
	case token.IntConst:
		return parser.parseIntConst()
	case token.RealConst:
		return parser.parseRealConst()
	case token.BoolConst:
		return parser.parseBoolConst()
	case token.StringConst:
		return parser.parseStringConst()
	case token.Identifier:
		return parser.parseIdentifier()
	case token.LParen:
		return parser.parseParenExpr()
	case token.Plus, token.Minus, token.Exclamation:
		return parser.parseUnaryExpr()
	default:
		panic(fmt.Sprintf("Invalid factor '%s' at %d:%d.", string(nextToken.Type), nextToken.Line, nextToken.Column))
	}
}

func (parser *Parser) parseIdentifier() ast.IdentifierNode {
	myLogger.Debug("=BEG= Identifier")
	parser.eat(token.Identifier)
	myLogger.Debug("=END= Identifier")
	return ast.IdentifierNode{Name: parser.current().StrVal}
}

func (parser *Parser) parseIntConst() ast.IntConstNode {
	myLogger.Debug("=BEG= Integer Constant")
	parser.eat(token.IntConst)
	myLogger.Debug("=END= Integer Constant")
	return ast.IntConstNode{Value: parser.current().IntVal}
}

func (parser *Parser) parseRealConst() ast.RealConstNode {
	myLogger.Debug("=BEG= Real Constant")
	parser.eat(token.RealConst)
	myLogger.Debug("=END= Real Constant")
	return ast.RealConstNode{Value: parser.current().RealVal}
}

func (parser *Parser) parseBoolConst() ast.BoolConstNode {
	myLogger.Debug("=BEG= Boolean Constant")
	parser.eat(token.BoolConst)
	myLogger.Debug("=END= Boolean Constant")
	return ast.BoolConstNode{Value: parser.current().BoolVal}
}

func (parser *Parser) parseStringConst() ast.StringConstNode {
	myLogger.Debug("=BEG= String Constant")
	parser.eat(token.StringConst)
	myLogger.Debug("=END= String Constant")
	return ast.StringConstNode{Value: parser.current().StrVal}
}

func (parser *Parser) parseParenExpr() ast.ExpressionNode {
	myLogger.Debug("=BEG= ()")
	parser.eat(token.LParen)
	expr := parser.parseExpr()
	parser.eat(token.RParen)
	myLogger.Debug("=END= ()")
	return expr
}

func (parser *Parser) parseUnaryExpr() ast.UnaryExprNode {
	myLogger.Debug("=BEG= Unary")
	op := parser.eat(token.Plus, token.Minus, token.Exclamation)
	expr := parser.parseFactor()
	myLogger.Debug("=END= Unary")
	return ast.UnaryExprNode{Operator: op, Expr: expr}
}
