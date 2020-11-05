package lexer_factory

import (
	"github.com/dollarkillerx/chronos/test/lexer"
	"github.com/dollarkillerx/chronos/test/lexertoken"
)

func BeginLexing(name, input string) *lexer.Lexer {
	l := &lexer.Lexer{
		Name:   name,
		Input:  input,
		State:  lexer.LexBegin,
		Tokens: make(chan lexertoken.Token, 3),
	}

	return l
}
