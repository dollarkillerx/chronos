package lexer

import (
	"strings"

	"github.com/dollarkillerx/chronos/conf/lexertoken"
)

func LexKey(lexer *Lexer) LexFn {
	for {
		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.EQUAL_SIGN) {
			lexer.Emit(lexertoken.TOKEN_KEY)
			return LexEqualSign
		}

		lexer.Inc()

		if lexer.IsEOF() {
			return lexer.Errorf("LEXER_ERROR_UNEXPECTED_EOF")
		}
	}
}

func LexEqualSign(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.EQUAL_SIGN)
	lexer.Emit(lexertoken.TOKEN_EQUAL_SIGN)
	return LexValue
}

func LexValue(lexer *Lexer) LexFn {
	for {
		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.NEWLINE) {
			lexer.Emit(lexertoken.TOKEN_VALUE)
			return LexBegin
		}

		lexer.Inc()

		if lexer.IsEOF() {
			lexer.Emit(lexertoken.TOKEN_VALUE)
			return lexer.Errorf("LEXER_ERROR_UNEXPECTED_EOF")
		}
	}
}
