package lexer

import (
	"github.com/dollarkillerx/chronos/test/lexertoken"
	"strings"
)

// start
func LexBegin(lexer *Lexer) LexFn {
	lexer.SkipWhiteSpace()
	if strings.HasPrefix(lexer.InputToEnd(), lexertoken.LEFT_BRACKET) {
		return LexLeftBracket
	} else {
		return LexKey
	}
}
