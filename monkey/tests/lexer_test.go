// monkey/tests/token_test.go

package tests

import (
	"monkey/lexer"
	"monkey/token"
	"testing"
)

/*
	Todos: Add comments
*/

type ExpectedToken struct {
    Type    string
    Literal string
}

func checkTokens(t *testing.T, lexer *lexer.Lexer, expectedTokens []ExpectedToken) {
    for i, expected := range expectedTokens {
        var token = lexer.Next()

		var hasError = false
        if token.Type != expected.Type {
            t.Errorf("Test[%d] - tokentype wrong. Expected=%q but got=%q instead",
                i, expected.Type, token.Type)
			hasError = true
        }

        if token.Literal != expected.Literal {
            t.Errorf("Test[%d] - literal wrong. Expected='%q' but got='%q' instead",
                i, expected.Literal, token.Literal)
			hasError = true
        }
		if hasError { t.Fatalf("Stoped on expected token: { %s, %s }", expected.Type, expected.Literal) }
    }
}

func Test_Lexer_Tokens(t *testing.T) {
	var input = `
		foobar
	 	123
		"foobar"
	 	=+-!*/
	 	<>==!=
	 	,;
	 	.:
	 	(){}[]
		fn let true false if else return
	`
	var lexer = lexer.NewLexer(input)

	var expectedTokens = []ExpectedToken{
		// Indentifier
		{ token.Ident, "foobar" },

		// Literals
		{ token.Int,    "123" },
		{ token.String, "foobar" },

		// // Operators
		{ token.Assign,	  "=" },
		{ token.Plus,	  "+" },
		{ token.Minus,	  "-" },
		{ token.Bang,	  "!" },
		{ token.Asterisk, "*" },
		{ token.Slash,	  "/" },

		// // Comparison
		{ token.Lt,	   "<"  },
		{ token.Gt,	   ">"  },
		{ token.Eq,	   "==" },
		{ token.NotEq, "!=" },

		// Delimiters
		{ token.Comma,     "," },
		{ token.Semicolon, ";" },

		// Others
		{ token.Dot,   "." },
		{ token.Colon, ":" },

		// Grouping
		{ token.Lparen,	  "(" },
		{ token.Rparen,	  ")" },
		{ token.Lbrace,	  "{" },
		{ token.Rbrace,	  "}" },
		{ token.Lbracket, "[" },
		{ token.Rbracket, "]" },

		// Keywords
		{ token.Function, "fn"	   },
		{ token.Let,	  "let"	   },
		{ token.True,	  "true"   },
		{ token.False,	  "false"  },
		{ token.If,		  "if"	   },
		{ token.Else,	  "else"   },
		{ token.Return,	  "return" },

		// End Of File
		{ token.Eof, "" },
	}

	// lexer.PrintTokens()
	checkTokens(t, lexer, expectedTokens)
}

func Test_Lexer_TokenizeCode(t *testing.T) {
	var input = `
		let five = 5;
		let ten = 10;

		let add = fn(x, y) {
		    x + y;
		};

		let result = add(five, ten);

		if (five < ten) {
			return true;
		} else {
			return false;
		}

		!true == false;
		five != ten;

		"foo bar";
		"foobar";
		"";
		"hello, world!";
`
		// [1, 2, 3];

		// { "foo": "bar" };

	var lexer = lexer.NewLexer(input)

    var expectedTokens = []ExpectedToken{
        // let
        {token.Let,		   "let"  },
		{ token.Ident,	   "five" },
		{ token.Assign,	   "="	  },
		{ token.Int,	   "5"	  },
		{ token.Semicolon, ";"	  },

		// let
		{ token.Let,	   "let" },
		{ token.Ident,	   "ten" },
		{ token.Assign,	   "="	 },
		{ token.Int,	   "10"	 },
		{ token.Semicolon, ";"	 },

		// let + function literal
		{ token.Let,	   "let" },
		{ token.Ident,	   "add" },
		{ token.Assign,	   "="	 },
		{ token.Function,  "fn"	 },
		{ token.Lparen,	   "("	 },
		{ token.Ident,	   "x"	 },
		{ token.Comma,	   ","	 },
		{ token.Ident,	   "y"	 },
		{ token.Rparen,	   ")"	 },
		{ token.Lbrace,	   "{"	 },
		{ token.Ident,	   "x"	 },
		{ token.Plus,	   "+"	 },
		{ token.Ident,	   "y"	 },
		{ token.Semicolon, ";"	 },
		{ token.Rbrace,	   "}"	 },
		{ token.Semicolon, ";"	 },

		// let + call with identifiers
		{ token.Let,	   "let"	},
		{ token.Ident,	   "result" },
		{ token.Assign,	   "="	    },
		{ token.Ident,	   "add"	},
		{ token.Lparen,	   "("	    },
		{ token.Ident,	   "five"	},
		{ token.Comma,	   ","	    },
		{ token.Ident,	   "ten"	},
		{ token.Rparen,	   ")"	    },
		{ token.Semicolon, ";"	    },

		// If Else
		{ token.If,		   "if"		},
		{ token.Lparen,	   "("		},
		{ token.Ident,	   "five"	},
		{ token.Lt,	       "<"		},
		{ token.Ident,	   "ten"	},
		{ token.Rparen,	   ")"		},
		{ token.Lbrace,	   "{"		},
		{ token.Return,	   "return" },
		{ token.True,	   "true"	},
		{ token.Semicolon, ";"		},
		{ token.Rbrace,	   "}"		},
		{ token.Else,	   "else"	},
		{ token.Lbrace,	   "{"		},
		{ token.Return,	   "return" },
		{ token.False,	   "false"	},
		{ token.Semicolon, ";"		},
		{ token.Rbrace,	   "}"		},

		// Comparison 1
		{ token.Bang,	   "!"	   },
		{ token.True,	   "true"  },
		{ token.Eq,		   "=="	   },
		{ token.False,	   "false" },
		{ token.Semicolon, ";"	   },

		// Comparison 2
		{ token.Ident,	   "five" },
		{ token.NotEq,	   "!="	  },
		{ token.Ident,	   "ten"  },
		{ token.Semicolon, ";"	  },

		// Strings
		{token.String,	   "foo bar"	   },
        {token.Semicolon,  ";"			   },
        {token.String,	   "foobar"		   },
        {token.Semicolon,  ";"			   },
        {token.String,	   ""			   },
        {token.Semicolon,  ";"			   },
        {token.String,	   "hello, world!" },
        {token.Semicolon,  ";"			   },

	}

	checkTokens(t, lexer, expectedTokens)
}
