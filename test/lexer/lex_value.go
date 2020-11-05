package lexer

import (
	"github.com/dollarkillerx/chronos/test/lexertoken"
	"strings"
)

func LexValue(lexer *Lexer) LexFn {
	for {
		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.NEWLINE) {
			lexer.Emit(lexertoken.TOKEN_VALUE)
			return LexBegin
		}

		lexer.Inc()

		if lexer.IsEOF() {
			return lexer.Errorf("LEXER_ERROR_UNEXPECTED_EOF")
		}
	}
}
