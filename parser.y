%{

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
%token INTEGER
%token OUTPUT
%token INPUT
%token NAME
%token DATA
%token VALUE

%token ADD
%token SUB
%token MUL
%token DIV

%token INTEGER_CONST
%token REAL_CONST
%token BOOL_CONST

%token IDENTIFIER

%%

html:
    '<' HTML '>' { std::cout << "OPEN HTML" << std::endl; }
    | '<' '/' HTML '>' { std::cout << "CLOSE HTML" << std::endl; }

%%

void yyerror(char const *s) {
    fprintf(stderr, "%s\n", s);
}

int main(void) {
    yyparse();
    return 0;
}
