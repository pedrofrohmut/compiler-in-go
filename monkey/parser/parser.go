// monkey/parser/parser.go

/*
	A parser is a software that takes text as input data, reads it and covert it
	to a data structure. The output is a structural representation of the text
	and checks for errors during the process.
*/

package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"strconv"
)

type Parser struct {
	lexer *lexer.Lexer
	curr token.Token
	errors []string
}

func NewParser(lexer *lexer.Lexer) *Parser {
	var parser = &Parser { lexer: lexer }
	parser.curr = lexer.Next()
	parser.errors = []string {}
	return parser
}

func (this *Parser) isCurr(check string) bool {
	return check == this.curr.Type
}

func (this *Parser) next() {
	this.curr = this.lexer.Next()
}

func (this *Parser) addError(error string) {
	this.errors = append(this.errors, error)
}

func (this *Parser) parseStatement() ast.Statement {
	var stm = &ast.ExpressionStatement{}

	var value, errValue = strconv.ParseInt(this.curr.Literal, 10, 64)
	if errValue != nil {
		this.addError(fmt.Sprintf("Error to parse token literal '%s' into int64", this.curr.Literal))
	}
	var integerLiteral = &ast.IntegerLiteral { Value: value }
	stm.Expression = integerLiteral

	this.next()

	return stm
}

func (this *Parser) ParseProgram() *ast.Program {
	var program = &ast.Program {}
	program.Statements = []ast.Statement {}

	for !this.isCurr(token.Eof) {
		var stm = this.parseStatement()

		if stm != nil {
			program.Statements = append(program.Statements, stm)
		}

		this.next()
	}

	return program
}
