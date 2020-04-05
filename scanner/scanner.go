package scanner

import (
    "fmt"
    "strconv"
    "strings"
    "unicode"

    "plhtml/logger"
    "plhtml/token"
    "plhtml/util"
)

var myLogger = logger.New("SCANNER")

func (scanner *Scanner) SetLogLevel(level logger.LogLevel) {
    myLogger.SetLevel(level)
}

type Scanner struct {
    line       int
    column     int
    prevColumn int
    tabSize    int
    index      int
    source     []rune
}

type Token struct {
    Type    token.Type
    Line    int
    Column  int
    IntVal  int
    BoolVal bool
    RealVal float64
    StrVal  string
}

func New() *Scanner {
    scanner := new(Scanner)
    scanner.tabSize = 4
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
        if tok.Type == token.EOF {
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

        tokenStartLine := scanner.line
        tokenStartColumn := scanner.column

        if ch == '`' {
            return scanner.lexString(ch, tokenStartLine, tokenStartColumn)
        }

        // after lexString because <!-- could be part of string literal
        if ch == '<' && scanner.lexComment(tokenStartLine, tokenStartColumn) {
            continue
        }

        if unicode.IsDigit(ch) {
            return scanner.lexNumber(ch, tokenStartLine, tokenStartColumn)
        }

        if ch == '&' || unicode.IsLetter(ch) {
            return scanner.lexWord(ch, tokenStartLine, tokenStartColumn)
        }

        charType := token.Illegal

        switch ch {
        case '+':
            charType = token.Plus
            break
        case '-':
            charType = token.Minus
            break
        case '*':
            charType = token.Multiply
            break
        case '/':
            charType = token.Slash
            break
        case '%':
            charType = token.Modulo
            break
        case '(':
            charType = token.LParen
            break
        case ')':
            charType = token.RParen
            break
        case '!':
            charType = token.Excl
            break
        case '"':
            charType = token.DQuote
            break
        case '=':
            charType = token.Equal
            break
        case '<':
            charType = token.LessThan
            break
        case '>':
            charType = token.GreaterThan
            break
        }

        if charType != token.Illegal {
            return newTok(charType, tokenStartLine, tokenStartColumn)
        }

        panic(fmt.Sprintf("Illegal character %c at %d:%d.", ch, tokenStartLine, tokenStartColumn))

    }

    return Token{Type: token.EOF}
}

func (scanner *Scanner) lexString(ch rune, line int, column int) Token {
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
            panic(fmt.Sprintf("Unterminated string at %d:%d.", line, column))
        }
    }

    tok := newTok(token.StringConst, line, column)
    tok.StrVal = str[1:]
    return tok
}

func (scanner *Scanner) lexComment(line int, column int) bool {
    // <!--.*-->

    ch1, ok1 := scanner.lookahead(0)
    ch2, ok2 := scanner.lookahead(1)
    ch3, ok3 := scanner.lookahead(2)

    if ok1 && ok2 && ok3 && ch1 == '!' && ch2 == '-' && ch3 == '-' {
        for {
            ch, ok := scanner.nextChar()
            if !ok {
                panic(fmt.Sprintf("End of file inside comment at %d:%d.", line, column))
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

func (scanner *Scanner) lexNumber(ch rune, line int, column int) Token {
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
        util.Check(err)
        tok := newTok(token.RealConst, line, column)
        tok.RealVal = realVal
        return tok
    }

    intVal, err := strconv.Atoi(number)
    util.Check(err)
    tok := newTok(token.IntConst, line, column)
    tok.IntVal = intVal
    return tok
}

func (scanner *Scanner) lexWord(ch rune, line int, column int) Token {
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

    if tok, ok := scanner.lexBoolOperator(word, line, column); ok {
        return tok
    }

    word = strings.ToLower(word)

    if word == "true" || word == "false" {
        tok := newTok(token.BoolConst, line, column)
        tok.BoolVal = word == "true"
        return tok
    }

    if tok, ok := token.KeywordLexemes[word]; ok {
        return newTok(tok, line, column)
    }

    tok := newTok(token.Identifier, line, column)
    tok.StrVal = word
    return tok
}

func (scanner *Scanner) lexBoolOperator(word string, line int, column int) (Token, bool) {

    firstChar := word[0]
    lastChar := word[len(word)-1:]

    if firstChar != '&' {
        return Token{Type: token.Illegal}, false
    }

    if lastChar != ";" {
        panic(fmt.Sprintf("Unterminated operator %s at %d:%d.", word, line, column))
    }

    if tok, ok := token.BoolOpLexemes[word]; ok {
        return newTok(tok, line, column), true
    }

    panic(fmt.Sprintf("Operator %s is not valid at %d:%d.", word, line, column))

}

func newTok(tokType token.Type, line int, column int) Token {
    return Token{Type: tokType, Line: line, Column: column}
}

func (scanner *Scanner) lookahead(i int) (rune, bool) {
    if scanner.index+i < len(scanner.source)-1 {
        return scanner.source[scanner.index+i], true
    }
    return 0, false
}

func (scanner *Scanner) nextChar() (rune, bool) {

    if scanner.index >= len(scanner.source)-1 {
        return 0, false
    }

    ch := scanner.source[scanner.index]
    scanner.index++

    myLogger.Debug("Read char: %s\n", strconv.Quote(string(ch)))

    if ch == '\n' {
        scanner.prevColumn = scanner.column
        scanner.column = 0
        scanner.line++
    } else if ch == '\t' {
        scanner.column += scanner.tabSize
    } else {
        scanner.column++
    }

    return ch, true

}

func (scanner *Scanner) goBack() {
    if scanner.index < len(scanner.source) {

        scanner.index--
        ch := scanner.source[scanner.index]

        myLogger.Debug("Unread char: %s\n", strconv.Quote(string(ch)))

        if ch == '\n' {
            scanner.column = scanner.prevColumn
            scanner.line--
        } else if ch == '\t' {
            scanner.column -= scanner.tabSize
        } else {
            scanner.column--
        }

    }
}
