package main

/*
prog = prog_dtype prog_html
prog_dtype = '<' '!' TokDoctype TokHtml '>'
prog_html = '<' TokHTML TokLang '=' '"' TokIdentifier '"' '>' ... '<' '/' TokHTML '>'

stmt =
     |
     | '<' TokHead '>' ... '<' '/' TokHEAD '>'
     | '<' TokTitle '>' TokStringConst '<' '/' TokTitle '>'
     | '<' TokBody '>' ... '<' '/' TokBody '>'
     | '<' TokMain '>' ... '<' '/' TokMain '>'
     | '<' TokVar TokClass '=' '"' TokIntType '"' '>' TokIdentifier '<' '/' TokVar '>'
     | '<' TokVar TokClass '=' '"' TokRealType '"' '>' TokIdentifier '<' '/' TokVar '>'
     | '<' TokVar TokClass '=' '"' TokBoolType '"' '>' TokIdentifier '<' '/' TokVar '>'
     | '<' TokVar TokClass '=' '"' TokStringType '"' '>' TokIdentifier '<' '/' TokVar '>'
     | '<' TokData TokValue '=' expression '>' TokIdentifier '<' '/' TokData '>'
     | '<' TokOutput '>' TokStringConst '<' '/' TokOutput '>'
     | '<' TokInput TokName '=' '"' TokIdentifier '"' '>'
     | '<' TokDiv TokData '-' TokWhile '=' '"' expression '"' '>' ... '<' '/' TokDiv '>'
     ;
expr = TokIntConst
     | TokRealConst
     | TokBoolConst
     | TokStringConst
     | TokIdentifier
     | expr '+' expr
     | expr '-' expr
     | expr '*' expr
     | expr '/' expr
     | expr TokLtOp expr
     | expr TokGtOp expr
     | expr TokLeqOp expr
     | expr TokGeqOp expr
     | expr TokEqOp expr
     | expr TokNeqOp expr
     | "(" expr ")"
     ;
*/
