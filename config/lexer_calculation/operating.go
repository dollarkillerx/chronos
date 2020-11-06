package lexer_calculation

import (
	"strings"
)

func IsEOF(token Token) bool {
	return token.Type == TOKEN_EOF
}

func BeginLexing(input string) *LexerCalculation {
	l := &LexerCalculation{
		Input:  input,
		State:  LexBegin,
		Tokens: make(chan Token, 3),
	}

	return l
}

func LexBegin(lexer *LexerCalculation) LexFn {
	lexer.SkipWhiteSpace()
	switch {
	case strings.HasPrefix(lexer.InputToEnd(), EVAL):
		return LexEval
	case strings.HasPrefix(lexer.InputToEnd(), LEFT_BRACKET):
		return LexLeftBracket
	case strings.HasPrefix(lexer.InputToEnd(), AND):
		return LexAND
	case strings.HasPrefix(lexer.InputToEnd(), OR):
		return LexOR
	default:
		return LexKey
	}
}

func LexKey(lexer *LexerCalculation) LexFn {
	for {
		if strings.HasPrefix(lexer.InputToEnd(), EQUAL_SIGN) {
			lexer.Emit(TOKEN_KEY)
			return LexEqualSign
		}

		lexer.Inc()

		if lexer.IsEOF() {
			return lexer.Errorf("LEXER_ERROR_UNEXPECTED_EOF")
		}
	}
}

func LexEqualSign(lexer *LexerCalculation) LexFn {
	lexer.Pos += len(EQUAL_SIGN)
	lexer.Emit(TOKEN_EQUAL_SIGN)
	return LexValue
}

func LexValue(lexer *LexerCalculation) LexFn {
	for {
		if strings.HasPrefix(lexer.InputToEnd(), "&") || strings.HasPrefix(lexer.InputToEnd(), "|") {
			lexer.Emit(TOKEN_VALUE)
			return LexBegin
		}

		lexer.Inc()

		if lexer.IsEOF() {
			lexer.Emit(TOKEN_VALUE)
			return lexer.Errorf("LEXER_ERROR_UNEXPECTED_EOF")
		}
	}
}

func LexEval(lexer *LexerCalculation) LexFn {
	lexer.Pos += len(EVAL)
	lexer.Emit(TOKEN_EVAL)
	return LexBegin
}

func LexAND(lexer *LexerCalculation) LexFn {
	lexer.Pos += len(AND)
	lexer.Emit(TOKEN_AND)
	return LexBegin
}

func LexOR(lexer *LexerCalculation) LexFn {
	lexer.Pos += len(OR)
	lexer.Emit(TOKEN_OR)
	return LexBegin
}

func LexLeftBracket(lexer *LexerCalculation) LexFn {
	lexer.Pos += len(LEFT_BRACKET)
	lexer.Emit(TOKEN_LEFT_BRACKET)

	return LexSection
}

func LexSection(lexer *LexerCalculation) LexFn {
	for {
		if lexer.IsEOF() {
			return lexer.Errorf("LEXER_ERROR_MISSING_RIGHT_BRACKET")
		}

		if strings.HasPrefix(lexer.InputToEnd(), RIGHT_BRACKET) {
			lexer.Emit(TOKEN_SECTION)
			return LexRightBracket
		}

		lexer.Inc()
	}
}

func LexRightBracket(lexer *LexerCalculation) LexFn {
	lexer.Pos += len(RIGHT_BRACKET)
	lexer.Emit(TOKEN_RIGHT_BRACKET)
	return LexBegin
}
