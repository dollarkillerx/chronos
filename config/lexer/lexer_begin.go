package lexer

import (
	"github.com/dollarkillerx/chronos/config/lexertoken"
	"strings"
)

// 找到开始[
func LexBegin(lexer *Lexer) LexFn {
	lexer.SkipWhiteSpace()
	if strings.HasPrefix(lexer.InputToEnd(), lexertoken.LEFT_BRACKET) {
		return LexLeftBracket
	} else {
		return LexKey
	}
}
