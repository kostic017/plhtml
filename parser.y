%{

#include "ast.h"
#include "ast_type.h"
#include "ast_expr.h"
#include "ast_decl.h"
#include "ast_stmt.h"

#include <vector>
#include <iostream>

int yylex();
void yyerror(char const *s);
%}

%token DOCTYPE
%token LANG
%token HTML
%token HEAD
%token TITLE
%token BODY
%token MAIN
%token VAR
%token CLASS
%token OUTPUT
%token INPUT
%token NAME
%token DATA
%token VALUE

%token T_INT
%token T_REAL
%token T_BOOL
%token T_STRING

%token <intVal> C_INT
%token <realVal> C_REAL
%token <boolVal> C_BOOL
%token <stringVal> C_STRING;

%token ADD
%token SUB
%token MUL
%token DIV

%token <stringVal> IDENTIFIER

%union
{
    int intVal;
    float realVal;
    bool boolVal;
    char* stringVal;
    Node* nodeVal;
};

%type <nodeVal> lval args stmt decl expr

%left ADD SUB
%left MUL DIV

%%

lval : IDENTIFIER { $$ = new LVal(); }
     ;

args : expr { $$ = new std::vector<Expr*>($1); }
     | args '.' expr { ($$ = $1)->push_back($3); }
     ;
     
stmt : '<' OUTPUT '>' args '<' '/' OUTPUT '>' { $$ = new PrintStmt($4); }
     ;

decl : '<' VAR CLASS '=' '"' T_INT '"' '>' IDENTIFIER '<' '/' VAR '>' { $$ = new VarDecl(new Identifier($9, Type::intType)); }
     | '<' VAR CLASS '=' '"' T_REAL '"' '>' IDENTIFIER '<' '/' VAR '>' { $$ = new VarDecl(new Identifier($9, Type::realType)); }
     | '<' VAR CLASS '=' '"' T_BOOL '"' '>' IDENTIFIER '<' '/' VAR '>' { $$ = new VarDecl(new Identifier($9, Type::boolType)); }
     | '<' VAR CLASS '=' '"' T_STRING '"' '>' IDENTIFIER '<' '/' VAR '>' { $$ = new VarDecl(new Identifier($9, Type::stringType)); }
     ;

expr : lval
     | '(' expr ')' { $$ = $2;  }
     | C_INT { $$ = new IntConst($1); }
     | C_REAL { $$ = new RealConst($1); }
     | C_BOOL { $$ = new BoolConst($1); }
     | C_STRING { $$ = new StringConst($1); }
     | expr ADD expr { $$ = new ArithmeticExpr($1, PLUS, $3); }
     | expr SUB expr { $$ = new ArithmeticExpr($1, MINUS, $3); }
     | expr MUL expr { $$ = new ArithmeticExpr($1, TIMES, $3); }
     | expr DIV expr { $$ = new ArithmeticExpr($1, DIVIDE, $3); }
     | '<' INPUT NAME '=' '"' lval '"' '>' { $$ = new AssignExpr($6, new ReadValueExpr()); }
     | '<' DATA VALUE '=' '"' expr '"' '>' lval '<' '/' DATA '>' { $$ = new AssignExpr($9, $6); }
     ;

%%

void yyerror(char const *s) {
    fprintf(stderr, "%s\n", s);
}

int main(void) {
    yyparse();
    return 0;
}
