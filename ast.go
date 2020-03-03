package main

type (
	StatementNode interface {
	}

	ExpressionNode interface {
	}

	ProgramNode struct {
		Title StringConstNode
		Body  ProgramBodyNode
	}

	ProgramBodyNode struct {
		MainFunc MainFuncNode
	}

	MainFuncNode struct {
		Statements []StatementNode
	}

	VarDeclNode struct {
		Identifier IdentifierNode
		Type       TokenType
	}

	VarAssignNode struct {
		Identifier IdentifierNode
		Value      ExpressionNode
	}

	IdentifierNode struct {
		Name string
	}

	StringConstNode struct {
		Value string
	}

	IntConstNode struct {
		Value int
	}

	RealConstNode struct {
		Value float64
	}

	BoolConstNode struct {
		Value bool
	}

	WriteStmtNode struct {
		Value ExpressionNode
	}

	ReadStmtNode struct {
		Identifier IdentifierNode
	}

	WhileStmtNode struct {
		Condition  ExpressionNode
		Statements []StatementNode
	}

	BinaryOpExprNode struct {
		Value1   ExpressionNode
		Value2   ExpressionNode
		Operator TokenType
	}

/*
   AstNode interface {
   }

   UnaryExprNode struct {
   	Value ExpressionNode
   }

   NegationExprNode struct {
   	Value bool
   }
*/
)
