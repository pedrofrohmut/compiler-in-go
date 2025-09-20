// monkey/lexer/lexer.go

package lexer

import (
	"monkey/token"
)

// Some constants to make life easier
const (
	CharEof = 0
)

type Lexer struct {
	input string
	pos int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input: input,
		pos: 0,
	}
}

func (this *Lexer) hasNext() bool {
	return this.pos < len(this.input) - 1
}

func (this *Lexer) currCh() byte {
	if !this.hasNext() {
		return CharEof
	}
	return this.input[this.pos]
}

func (this *Lexer) peekCh() byte {
	if this.isCurr(CharEof) { return 0 }
	return this.input[this.pos + 1]
}

func (this *Lexer) incPos() {
	if this.hasNext() { this.pos++ }
}

func (this *Lexer) isCurr(ch byte) bool {
	return this.input[this.pos] == ch
}

func (this *Lexer) isPeek(ch byte) bool {
	if !this.hasNext() {
		return false
	}
	return this.input[this.pos + 1] == ch
}

func (this *Lexer) getWord() string {
	var start = this.pos
	for isAlphaNumeric(this.peekCh()) {
		this.incPos()
	}
	return this.input[start:this.pos + 1]
}

func (this *Lexer) tokenizeIntegerLiteral() token.Token {
	var intValue = this.getWord()
	return token.NewToken(token.Int, intValue)
}

func isNumeric(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isAlpha(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')
}

func isAlphaNumeric(ch byte) bool {
	return isNumeric(ch) || isAlpha(ch)
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func (this *Lexer) tokenizeStringLiteral() token.Token {
	this.incPos() // Jumps the '"' at the start
	var start = this.pos
	for !this.isCurr('"') && !this.isCurr(CharEof) {
		this.incPos()
	}
	var content = this.input[start:this.pos]
	this.incPos() // Jumps the '"' at the end
	return token.NewToken(token.String, content)
}

func (this *Lexer) skipWhiteSpace() {
	for isWhiteSpace(this.currCh()) {
		this.incPos()
	}
}

func (this *Lexer) Next() token.Token {
	var result token.Token

	this.skipWhiteSpace()

	switch this.currCh() {
	case CharEof:
		result = token.NewToken(token.Eof, "")
	case '"':
		result = this.tokenizeStringLiteral()
	case '=':
		if this.isPeek('=') {
			result = token.NewToken(token.Eq, "==")
			this.incPos()
		} else {
			result = token.NewToken(token.Assign, "=")
		}
	case '+':
		result = token.NewToken(token.Plus, "+")
	case '-':
		result = token.NewToken(token.Minus, "-")
	case '!':
		if this.isPeek('=') {
			result = token.NewToken(token.NotEq, "!=")
			this.incPos()
		} else {
			result = token.NewToken(token.Bang, "!")
		}
	case '*':
		result = token.NewToken(token.Asterisk, "*")
	case '/':
		result = token.NewToken(token.Slash, "/")
	case '<':
		result = token.NewToken(token.Lt, "<")
	case '>':
		result = token.NewToken(token.Gt, ">")
	case ',':
		result = token.NewToken(token.Comma, ",")
	case ';':
		result = token.NewToken(token.Semicolon, ";")
	case '.':
		result = token.NewToken(token.Dot, ".")
	case ':':
		result = token.NewToken(token.Colon, ":")
	case '(':
		result = token.NewToken(token.Lparen, "(")
	case ')':
		result = token.NewToken(token.Rparen, ")")
	case '[':
		result = token.NewToken(token.Lbracket, "[")
	case ']':
		result = token.NewToken(token.Rbracket, "]")
	case '{':
		result = token.NewToken(token.Lbrace, "{")
	case '}':
		result = token.NewToken(token.Rbrace, "}")
	default:
		// Integer Literals
		if isNumeric(this.currCh()) {
			result = this.tokenizeIntegerLiteral()
			break
		}
		// Identifiers and reserved words
		var word = this.getWord()
		switch word {
		case "fn":
			result = token.NewToken(token.Function, word)
		case "let":
			result = token.NewToken(token.Let, word)
		case "true":
			result = token.NewToken(token.True, word)
		case "false":
			result = token.NewToken(token.False, word)
		case "if":
			result = token.NewToken(token.If, word)
		case "else":
			result = token.NewToken(token.Else, word)
		case "return":
			result = token.NewToken(token.Return, word)
		default:
			result = token.NewToken(token.Ident, word)
		}
	}

	this.incPos()

	return result
}

// Print all tokens and restore position before being called
func (this *Lexer) PrintTokens() {
	var backupPos = this.pos
	this.pos = 0

	for {
		var tk = this.Next()
		token.PrintToken(tk)
		if tk.Type == token.Eof { break }
	}

	this.pos = backupPos
}
