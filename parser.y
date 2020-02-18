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

%token C_INT
%token C_REAL
%token C_BOOL

%token ADD
%token SUB
%token MUL
%token DIV

%token IDENTIFIER

%%

stmt:
    

expr:
    C_INT { $$ = new IntConst($1); }
    | C_REAL { $$ = new RealConst($1); }
    | C_BOOL { $$ = new BoolConst($1); }
    | IDENTIFIER { $$ = new Identifier($1); }
    | expr ADD expr { $$ = new ArithmeticExpr($1, PLUS, $3);  }
    | expr SUB expr { $$ = new ArithmeticExpr($1, MINUS, $3); }
    | expr MUL expr { $$ = new ArithmeticExpr($1, TIMES, $3); }
    | expr DIV expr { $$ = new ArithmeticExpr($1, DIVIDE, $3); }
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
