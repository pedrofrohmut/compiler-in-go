package main

import "monkey/lexer"

func main() {
	var input = `"bar"; "foo"; "baz";`
	var lexer = lexer.NewLexer(input)
	lexer.PrintTokens()
}
