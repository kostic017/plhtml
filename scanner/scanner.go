package scanner

import (
    "fmt"
    "strconv"
    "strings"
    "unicode"

    "../logger"
    "../utility"
)

var myLogger = logger.New("SCANNER")

func (scanner *Scanner) SetLogLevel(level logger.LogLevel) {
    myLogger.SetLevel(level)
}

type Scanner struct {
    line             int
    column           int
    prevColumn       int
    tokenStartLine   int
    tokenStartColumn int
    tabSize          int

    index  int
    source []rune

    keywords  map[string]TokenType
    operators map[string]TokenType
}

func New() *Scanner {
    scanner := new(Scanner)
    scanner.tabSize = 4

    scanner.keywords = map[string]TokenType{
        "doctype": TokDoctype,
        "lang":    TokLang,
        "html":    TokHTML,
        "head":    TokHead,
        "title":   TokTitle,
        "body":    TokBody,
        "main":    TokMain,
        "var":     TokVar,
        "class":   TokClass,
        "output":  TokOutput,
        "input":   TokInput,
        "name":    TokName,
        "data":    TokData,
        "value":   TokValue,
        "div":     TokDiv,
        "if":      TokIf,
        "while":   TokWhile,
        "integer": TokIntType,
        "real":    TokRealType,
        "boolean": TokBoolType,
        "string":  TokStringType,
    }

    scanner.operators = map[string]TokenType{
        "&lt;":     TokLtOp,
        "&gt;":     TokGtOp,
        "&leq;":    TokLeqOp,
        "&geq;":    TokGeqOp,
        "&equals;": TokEqOp,
        "&ne;":     TokNeqOp,
    }

    return scanner
}

func (scanner *Scanner) SetTabSize(tabSize int) {
    scanner.tabSize = tabSize
}

func (scanner *Scanner) Scan(source string) []Token {
    scanner.line = 1
    scanner.index = 0
    scanner.column = 0
    scanner.source = []rune(source)

    var tokens []Token

    for {
        tok := scanner.nextToken()
        tokens = append(tokens, tok)
        if tok.Type == TokEOF {
            break
        }
    }

    return tokens
}

func (scanner *Scanner) nextToken() Token {

    for {

        ch, ok := scanner.nextChar()

        if !ok {
            break
        }

        if unicode.IsSpace(ch) {
            continue
        }

        scanner.tokenStartLine = scanner.line
        scanner.tokenStartColumn = scanner.column

        if ch == '`' {
            return scanner.lexString(ch)
        }

        // after lexString because <!-- could be part of string literal
        if ch == '<' && scanner.lexComment() {
            continue
        }

        if unicode.IsDigit(ch) {
            return scanner.lexNumber(ch)
        }

        if ch == '&' || unicode.IsLetter(ch) {
            return scanner.lexWord(ch)
        }

        switch ch {
        case '+', '-', '*', '/', '(', ')', '!', '"', '=', '<', '>', '.':
            return scanner.newToken(TokenType(ch))
        }

        panic(fmt.Sprintf("Illegal character %c at %d:%d.", ch, scanner.tokenStartLine, scanner.tokenStartColumn))

    }

    return Token{Type: TokEOF}
}

func (scanner *Scanner) lexString(ch rune) Token {
    var ok bool
    str := string(ch)

    for {
        ch, ok = scanner.nextChar()

        if ok {

            if ch == '`' {
                break
            }

            str += string(ch)

        } else {
            panic(fmt.Sprintf("Unterminated string at %d:%d.", scanner.tokenStartLine, scanner.tokenStartColumn))
        }
    }

    token := scanner.newToken(TokStringConst)
    token.StrVal = str[1:]
    return token
}

func (scanner *Scanner) lexComment() bool {
    // <!--.*-->

    ch1, ok1 := scanner.lookahead(0)
    ch2, ok2 := scanner.lookahead(1)
    ch3, ok3 := scanner.lookahead(2)

    if ok1 && ok2 && ok3 && ch1 == '!' && ch2 == '-' && ch3 == '-' {
        for {
            ch, ok := scanner.nextChar()
            if !ok {
                panic(fmt.Sprintf("End of file inside comment at %d:%d.", scanner.tokenStartLine, scanner.tokenStartColumn))
            }
            if ch == '-' {
                ch1, ok1 := scanner.lookahead(0)
                ch2, ok2 := scanner.lookahead(1)
                if ok1 && ok2 && ch1 == '-' && ch2 == '>' {
                    scanner.nextChar()
                    scanner.nextChar()
                    return true
                }
            }
        }
    }

    return false
}

func (scanner *Scanner) lexNumber(ch rune) Token {
    var ok bool
    isReal := false
    number := "" + string(ch)

    for ch, ok = scanner.nextChar(); ok; ch, ok = scanner.nextChar() {

        if ch == '.' {
            isReal = true
        }

        if ch == '.' || unicode.IsNumber(ch) {
            number += string(ch)
        } else {
            scanner.goBack()
            break
        }

    }

    if isReal {
        realVal, err := strconv.ParseFloat(number, 64)
        utility.Check(err)
        token := scanner.newToken(TokRealConst)
        token.RealVal = realVal
        return token
    }

    intVal, err := strconv.Atoi(number)
    utility.Check(err)
    token := scanner.newToken(TokIntConst)
    token.IntVal = intVal
    return token
}

func (scanner *Scanner) lexWord(ch rune) Token {
    // &[a-zA-Z];            operators
    // [a-zA-Z][a-zA-Z0-9]*  identifiers/keywords

    word := string(ch)

    for ch, ok := scanner.nextChar(); ok; ch, ok = scanner.nextChar() {

        valid := unicode.IsLetter(ch) ||
            (word[0] == '&' && ch == ';') ||
            (word[0] != '&' && unicode.IsNumber(ch))

        if valid {
            word += string(ch)
        } else {
            scanner.goBack()
            break
        }

    }

    if tok, ok := scanner.lexOperator(word); ok {
        return tok
    }

    word = strings.ToLower(word)

    if word == "true" || word == "false" {
        token := scanner.newToken(TokBoolConst)
        token.BoolVal = word == "true"
        return token
    }

    if tok, ok := scanner.keywords[word]; ok {
        return scanner.newToken(tok)
    }

    token := scanner.newToken(TokIdentifier)
    token.StrVal = word
    return token
}

func (scanner *Scanner) lexOperator(word string) (Token, bool) {

    firstChar := word[0]
    lastChar := word[len(word)-1:]

    if firstChar != '&' {
        return Token{Type: TokEOF}, false
    }

    if lastChar != ";" {
        panic(fmt.Sprintf("Unterminated operator %s at %d:%d.", word, scanner.tokenStartLine, scanner.tokenStartColumn))
    }

    if tok, ok := scanner.operators[word]; ok {
        return scanner.newToken(tok), true
    }

    panic(fmt.Sprintf("Operator %s is not valid at %d:%d.", word, scanner.tokenStartLine, scanner.tokenStartColumn))

}

func (scanner Scanner) newToken(tokType TokenType) Token {
    return Token{Type: tokType, Line: scanner.tokenStartLine, Column: scanner.tokenStartColumn}
}

func (scanner *Scanner) nextChar() (rune, bool) {
    if scanner.index != len(scanner.source)-1 {
        ch := scanner.source[scanner.index]
        myLogger.Debug("Read char: %s\n", strconv.Quote(string(ch)))
        scanner.incCounters(ch)
        scanner.index++
        return ch, true
    }

    return 0, false
}

func (scanner *Scanner) lookahead(i int) (rune, bool) {
    if scanner.index+i < len(scanner.source)-1 {
        return scanner.source[scanner.index+i], true
    }
    return 0, false
}

func (scanner *Scanner) goBack() {
    if scanner.index < len(scanner.source) {
        scanner.index--
        ch := scanner.source[scanner.index]
        myLogger.Debug("Unread char: %s\n", strconv.Quote(string(ch)))
        scanner.decCounters(ch)
    }
}

func (scanner *Scanner) decCounters(ch rune) {
    if ch == '\n' {
        scanner.column = scanner.prevColumn
        scanner.line--
    } else if ch == '\t' {
        scanner.column -= scanner.tabSize
    } else {
        scanner.column--
    }
}

func (scanner *Scanner) incCounters(ch rune) {
    if ch == '\n' {
        scanner.prevColumn = scanner.column
        scanner.column = 0
        scanner.line++
    } else if ch == '\t' {
        scanner.column += scanner.tabSize
    } else {
        scanner.column++
    }
}
