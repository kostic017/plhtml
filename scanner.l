%{

#include "y.tab.h"

void yyerror(char const *s);

%}

digit  [0-9]
letter [a-zA-Z]

%%

(?i:doctype) return DOCTYPE;
(?i:lang)    return LANG;
(?i:html)    return HTML;
(?i:head)    return HEAD;
(?i:title)   return TITLE;
(?i:body)    return BODY;
(?i:main)    return MAIN;
(?i:var)     return VAR;
(?i:class)   return CLASS;
(?i:integer) return INTEGER;
(?i:output)  return OUTPUT;
(?i:input)   return INPUT;
(?i:name)    return NAME;
(?i:data)    return DATA;
(?i:value)   return VALUE;

&plus;       return ADD;

`            return STR_DELIM;
[!"/=<>]     return yytext[0];

{digit}+     return INTEGER_CONST;

(_|{letter})(_|{letter}|{digit})* return IDENTIFIER;

"<!--".*"-->"    ;
[ \t\n]          ;

.            yyerror(yytext);

%%

int yywrap() {
    return 1;
}
