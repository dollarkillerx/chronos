package lexer

import (
	"github.com/dollarkillerx/chronos/test/lexertoken"
	"strings"
)

func LexSection(lexer *Lexer) LexFn {
	for {
		if lexer.IsEOF() {
			return lexer.Errorf("LEXER_ERROR_MISSING_RIGHT_BRACKET")
		}

		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.RIGHT_BRACKET) {
			lexer.Emit(lexertoken.TOKEN_SECTION)
			return LexRightBracket
		}

		lexer.Inc()
	}
}
