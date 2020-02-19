%{

#include "ast.h"
#include "y.tab.h"
#include <string>

void yyerror(char const *s);

%}

digit  [0-9]
letter [a-zA-Z]

%%

(?i:doctype)                      return DOCTYPE;
(?i:lang)                         return LANG;
(?i:html)                         return HTML;
(?i:head)                         return HEAD;
(?i:title)                        return TITLE;
(?i:body)                         return BODY;
(?i:main)                         return MAIN;
(?i:var)                          return VAR;
(?i:class)                        return CLASS;
(?i:output)                       return OUTPUT;
(?i:input)                        return INPUT;
(?i:name)                         return NAME;
(?i:data)                         return DATA;
(?i:value)                        return VALUE;

(?i:int)                          return T_INT;
(?i:real)                         return T_REAL;
(?i:bool)                         return T_BOOL;
(?i:string)                       return T_STRING;

{digit}+                          { yylval.intVal = std::stoi(yytext); return C_INT; }
(true|false)                      { yylval.boolVal = yytext == "true"; return C_BOOL; }
{digit}+\.{digit}+                { yylval.realVal = std::stof(yytext); return C_REAL; }
`.*`                              { yylval.stringVal = yytext; return C_STRING; } // TODO

&plus;                            return ADD;
&minus;                           return SUB;
&times;                           return MUL;
&divide;                          return DIV;

(_|{letter})(_|{letter}|{digit})* { yylval.stringVal = yytext; return IDENTIFIER; }

[!"/=<>()]                        return yytext[0];

"<!--".*"-->"                     ;
[ \t\n]                           ;

.                                 yyerror(yytext);

%%

int yywrap() {
    return 1;
}
