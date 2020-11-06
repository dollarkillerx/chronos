package lexer

import (
	"strings"

	"github.com/dollarkillerx/chronos/config/lexertoken"
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
