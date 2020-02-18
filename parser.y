%{

#include "ast.h"
#include "ast_expr.h"
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

%token <intVal> C_INT
%token <realVal> C_REAL
%token <boolVal> C_BOOL

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

%type <nodeVal> expr

%left ADD SUB
%left MUL DIV

%%

expr:
    C_INT { $$ = new IntConst($1); }
    | C_REAL { $$ = new RealConst($1); }
    | C_BOOL { $$ = new BoolConst($1); }
    | IDENTIFIER { $$ = new Identifier($1); }
    | expr ADD expr { $$ = new ArithmeticExpr((Expr*) $1, PLUS, (Expr*) $3); }
    | expr SUB expr { $$ = new ArithmeticExpr((Expr*) $1, MINUS, (Expr*) $3); }
    | expr MUL expr { $$ = new ArithmeticExpr((Expr*) $1, TIMES, (Expr*) $3); }
    | expr DIV expr { $$ = new ArithmeticExpr((Expr*) $1, DIVIDE, (Expr*) $3); }
    | '(' expr ')' { $$ = $2;  }
    ;

%%

void yyerror(char const *s) {
    fprintf(stderr, "%s\n", s);
}

int main(void) {
    yyparse();
    return 0;
}
