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
	peek token.Token
	errors []string
}

func NewParser(lexer *lexer.Lexer) *Parser {
	var parser = &Parser { lexer: lexer }
	parser.curr = lexer.Next()
	parser.peek = lexer.Next()
	parser.errors = []string {}
	return parser
}

func (this *Parser) isCurr(check string) bool {
	return check == this.curr.Type
}

func (this *Parser) isPeek(check string) bool {
	return check == this.peek.Type
}

func (this *Parser) next() {
	this.curr = this.peek
	this.peek = this.lexer.Next()
}

func (this *Parser) Errors() []string {
	return this.errors
}

func (this *Parser) addError(error string) {
	this.errors = append(this.errors, error)
}

func (this *Parser) parseAtPoint() ast.Expression {
	switch this.curr.Type {
	case token.Int:
		var value, errValue = strconv.ParseInt(this.curr.Literal, 10, 64)
		if errValue != nil {
			this.addError(fmt.Sprintf("Error to parse token literal '%s' into int64", this.curr.Literal))
		}
		return &ast.IntegerLiteral { Value: value }
	default:
		this.addError("Invalid or not covered tokenType in parseAtPoint: " + this.curr.Type)
		return nil
	}
}

func (this *Parser) makeInfix(left ast.Expression) ast.Expression {
    // Start: Curr is operator
    var inf = &ast.InfixExpression {}
    inf.Left = left
    inf.Operator = this.curr.Literal
    var precedence = this.currPrecedence()
    this.next() // Curr to next value
    inf.Right = this.newInfixGroup(precedence)
    return inf
}

func (this *Parser) parseInfix(expression ast.Expression) ast.Expression {
    switch this.curr.Type {
    case token.Plus, token.Minus, token.Slash, token.Asterisk, token.Eq, token.NotEq, token.Lt, token.Gt:
        return this.makeInfix(expression)
    default:
        this.addError("Invalid or not covered symbol for infix parse: " + this.curr.Type)
        return nil
    }
}

func (this *Parser) newInfixGroup(ctxPrecedence int) ast.Expression {
	var parsedValue = this.parseAtPoint()

	var acc = parsedValue

    for !this.isPeek(token.Semicolon) && this.peekPrecedence() > ctxPrecedence {
        this.next() // Curr to operator
        acc = this.parseInfix(acc)
    }

	return acc
}

const (
    // iota gives the constants a ascending numbers
    // _ skips the 0 value
    _ int = iota
    Lowest
    Equals      // ==
    LessGreater // > or <
    Sum         // + -
    Product     // * /
    Prefix      // -X or !X
    Call        // myFunction(X)
)

var precedences = map[string] int {
    token.Eq:       Equals,
    token.NotEq:    Equals,
    token.Lt:       LessGreater,
    token.Gt:       LessGreater,
    token.Plus:     Sum,
    token.Minus:    Sum,
    token.Slash:    Product,
    token.Asterisk: Product,
    token.Lparen:   Call,
}

func (this *Parser) currPrecedence() int {
    return precedences[this.curr.Type]
}

func (this *Parser) peekPrecedence() int {
    return precedences[this.peek.Type]
}

func (this *Parser) parseExpr(precedence int) ast.Expression {
	return this.newInfixGroup(precedence)
}

// TODO: Add more types of statements here
func (this *Parser) parseStatement() ast.Statement {
	var stm = &ast.ExpressionStatement{}
	stm.Expression = this.parseExpr(Lowest)
	this.next()
	for !this.isCurr(token.Semicolon) {
		this.addError(fmt.Sprintf("Parser Error: Expected a semicolon but found %s instead", this.curr.Literal))
		this.next()
	}
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
