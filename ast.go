package main

type (
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

	StatementNode interface {
	}

	IdentifierNode struct {
		Name string
	}

	StringConstNode struct {
		Value string
	}

/*
   AstNode interface {
   }

   VarDeclNode struct {
   	Identifier IdentifierNode
   	Type       TokenType
   }

   AssignStmtNode struct {
   	Identifier IdentifierNode
   	Value      ExpressionNode
   }

   WriteStmtNode struct {
   	Value ExpressionNode
   }

   ReadStmtNode struct {
   	Identifier IdentifierNode
   }

   WhileStmtNode struct {
   	Condition ExpressionNode
   	// statements
   }

   ExpressionNode interface {
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

   AddExprNode struct {
   	Value1 ExpressionNode
   	Value2 ExpressionNode
   }

   SubExprNode struct {
   	Value1 ExpressionNode
   	Value2 ExpressionNode
   }

   MulExprNode struct {
   	Value1 ExpressionNode
   	Value2 ExpressionNode
   }

   DivExprNode struct {
   	Value1 ExpressionNode
   	Value2 ExpressionNode
   }

   LtExprNode struct {
   	Value1 ExpressionNode
   	Value2 ExpressionNode
   }

   GtExprNode struct {
   	Value1 ExpressionNode
   	Value2 ExpressionNode
   }

   LeqExprNode struct {
   	Value1 ExpressionNode
   	Value2 ExpressionNode
   }

   GeqExprNode struct {
   	Value1 ExpressionNode
   	Value2 ExpressionNode
   }

   EqExprNode struct {
   	Value1 ExpressionNode
   	Value2 ExpressionNode
   }

   NeqExprNode struct {
   	Value1 ExpressionNode
   	Value2 ExpressionNode
   }

   UnaryExprNode struct {
   	Value ExpressionNode
   }

   NegationExprNode struct {
   	Value bool
   }
*/
)
