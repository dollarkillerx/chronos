package lexer

import (
	"github.com/dollarkillerx/chronos/test/lexertoken"
	"strings"
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
