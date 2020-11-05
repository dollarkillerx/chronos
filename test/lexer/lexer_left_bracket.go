package lexer

import "github.com/dollarkillerx/chronos/test/lexertoken"

func LexLeftBracket(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.LEFT_BRACKET)
	lexer.Emit(lexertoken.TOKEN_LEFT_BRACKET)
	return LexSection
}
