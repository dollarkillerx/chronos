package lexer

import "github.com/dollarkillerx/chronos/test/lexertoken"

func LexEqualSign(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.EQUAL_SIGN)
	lexer.Emit(lexertoken.TOKEN_EQUAL_SIGN)
	return LexValue
}
