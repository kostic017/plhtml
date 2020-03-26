package token

type Type int

const (
    Illegal Type = iota
    EOF

    Doctype
    Lang
    HTML
    Head
    Title
    Body
    Main
    Var
    Class
    Output
    Input
    Name
    Data
    Value
    Div
    If
    While

    LtOp
    GtOp
    LeqOp
    GeqOp
    EqOp
    NeqOp

    IntConst
    RealConst
    BoolConst
    StringConst

    Identifier

    Plus
    Minus
    Asterisk
    Slash
    LParen
    RParen
    Exclamation
    DQuote
    Equal
    LessThan
    GreaterThan
    Period
)

func (tokenType Type) String() string {
    return [...]string{
        Illegal:     "Illegal",
        EOF:         "EOF",
        Doctype:     "Doctype",
        Lang:        "Lang",
        HTML:        "Html",
        Head:        "Head",
        Title:       "Title",
        Body:        "Body",
        Main:        "Main",
        Var:         "Var",
        Class:       "Class",
        Output:      "Output",
        Input:       "Input",
        Name:        "Name",
        Data:        "Data",
        Value:       "Value",
        Div:         "Div",
        If:          "If",
        While:       "While",
        LtOp:        "LtOp",
        GtOp:        "GtOp",
        LeqOp:       "LeqOp",
        GeqOp:       "GeqOp",
        EqOp:        "EqOp",
        NeqOp:       "NeqOp",
        IntConst:    "IntConst",
        RealConst:   "RealConst",
        BoolConst:   "BoolConst",
        StringConst: "StringConst",
        Identifier:  "Identifier",
        Plus:        "+",
        Minus:       "-",
        Asterisk:    "*",
        Slash:       "/",
        LParen:      "(",
        RParen:      ")",
        Exclamation: "!",
        DQuote:      "\"",
        Equal:       "=",
        LessThan:    "<",
        GreaterThan: ">",
        Period:      ".",
    }[tokenType]
}

var KeywordLexemes = map[string]Type{
    "doctype": Doctype,
    "lang":    Lang,
    "html":    HTML,
    "head":    Head,
    "title":   Title,
    "body":    Body,
    "main":    Main,
    "var":     Var,
    "class":   Class,
    "output":  Output,
    "input":   Input,
    "name":    Name,
    "data":    Data,
    "value":   Value,
    "div":     Div,
    "if":      If,
    "while":   While,
}

var BoolOpLexemes = map[string]Type{
    "&lt;":     LtOp,
    "&gt;":     GtOp,
    "&leq;":    LeqOp,
    "&geq;":    GeqOp,
    "&equals;": EqOp,
    "&ne;":     NeqOp,
}
