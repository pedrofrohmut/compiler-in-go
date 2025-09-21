// monkey/token/token.go

package token

import "fmt"

const (
    // Special types
    Illegal    = "ILLEGAL"
    Eof        = "EOF"

    // indentifiers
    Ident      = "IDENT" // add, foobar, x, y

    // Literals
    Int        = "INT"
    String     = "STRING"

    // Operators
    Assign     = "="
    Plus       = "+"
    Minus      = "-"
    Bang       = "!"
    Asterisk   = "*"
    Slash      = "/"

    // Comparison
    Lt         = "<"
    Gt         = ">"
    Eq         = "=="
    NotEq     = "!="

    // Delimiters
    Comma      = ","
    Semicolon  = ";"

	// Others
    Dot        = "."
    Colon      = ":"

    // Grouping
    Lparen     = "("
    Rparen     = ")"
    Lbrace     = "{"
    Rbrace     = "}"
    Lbracket   = "["
    Rbracket   = "]"

    // Keywords
    Function   = "FUNCTION"
    Let        = "LET"
    True       = "TRUE"
    False      = "FALSE"
    If         = "IF"
    Else       = "ELSE"
    Return     = "RETURN"
)

type Token struct {
    Type string
    Literal string
}

func NewToken(typ string, literal string) Token {
	return Token{ Type: typ, Literal: literal }
}

func PrintToken(token Token) {
	switch token.Type {
	case String:
		fmt.Printf("{ %s, \"%s\" }\n", token.Type, token.Literal)
	default:
		fmt.Printf("{ %s, \"%s\" }\n", token.Type, token.Literal)
	}
}
