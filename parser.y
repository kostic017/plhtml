%{
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
%token <stringVal> C_STRING

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
    void* nodeVal;
};

%type <nodeVal> lval args stmt decl expr

%left ADD SUB
%left MUL DIV

%%

lval : IDENTIFIER { }
     ;

args : expr { }
     | args '.' expr { }
     ;
     
stmt : '<' OUTPUT '>' args '<' '/' OUTPUT '>' { }
     ;

decl : '<' VAR CLASS '=' '"' T_INT '"' '>' IDENTIFIER '<' '/' VAR '>' { }
     | '<' VAR CLASS '=' '"' T_REAL '"' '>' IDENTIFIER '<' '/' VAR '>' { }
     | '<' VAR CLASS '=' '"' T_BOOL '"' '>' IDENTIFIER '<' '/' VAR '>' { }
     | '<' VAR CLASS '=' '"' T_STRING '"' '>' IDENTIFIER '<' '/' VAR '>' { }
     ;

expr : lval
     | '(' expr ')' {  }
     | C_INT { }
     | C_REAL { }
     | C_BOOL { }
     | C_STRING { }
     | expr ADD expr { }
     | expr SUB expr { }
     | expr MUL expr { }
     | expr DIV expr { }
     | '<' INPUT NAME '=' '"' lval '"' '>' { }
     | '<' DATA VALUE '=' '"' expr '"' '>' lval '<' '/' DATA '>' { }
     ;

%%

void yyerror(char const *s) {
    fprintf(stderr, "%s\n", s);
}

int main(void) {
    yyparse();
    return 0;
}
