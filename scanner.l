%{

#include "y.tab.h"

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

{digit}+                          return C_INT;
(true|false)                      return C_BOOL;
{digit}+\.{digit}+                return C_REAL;

&plus;                            return ADD;
&minus;                           return SUB;
&times;                           return MUL;
&divide;                          return DIV;

(_|{letter})(_|{letter}|{digit})* return IDENTIFIER;

[!"/=<>()]                        return yytext[0];

"<!--".*"-->"                     ;
[ \t\n]                           ;

.                                 yyerror(yytext);

%%

int yywrap() {
    return 1;
}
