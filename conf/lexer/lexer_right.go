package lexer

import "github.com/dollarkillerx/chronos/conf/lexertoken"

func LexRightBracket(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.RIGHT_BRACKET)
	lexer.Emit(lexertoken.TOKEN_RIGHT_BRACKET)
	return LexBegin
}
