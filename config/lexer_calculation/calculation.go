package lexer_calculation

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/dollarkillerx/chronos/config/lexertoken"
)

type LexerCalculation struct {
	Input  string
	Tokens chan Token
	State  LexFn

	Start int
	Pos   int // Token end
}

type LexFn func(lexer *LexerCalculation) LexFn

// 提交当前Token
func (l *LexerCalculation) Emit(tokenType TokenType) {
	l.Tokens <- Token{Type: tokenType, Value: l.Input[l.Start:l.Pos]}
	l.Start = l.Pos
}

// 前进
func (l *LexerCalculation) Inc() {
	l.Pos++
	if l.Pos >= utf8.RuneCountInString(l.Input) {
		l.Emit(TOKEN_EOF)
	}
}

// 后退
func (l *LexerCalculation) Dec() {
	l.Pos--
}

// 前Token 到最后面
func (l *LexerCalculation) InputToEnd() string {
	return l.Input[l.Pos:]
}

// 是否EOF
func (l *LexerCalculation) IsEOF() bool {
	return l.Pos >= utf8.RuneCountInString(l.Input)
}

// 向前推进一格
func (l *LexerCalculation) Next() rune {
	// 是否EOF
	if l.IsEOF() {
		return lexertoken.EOF
	}

	// 读一个RUNE  返回item,长度
	result, width := utf8.DecodeRuneInString(l.Input[l.Pos:])
	l.Pos += width
	return result
}

// 跳过空白处
func (l *LexerCalculation) SkipWhiteSpace() {
	for {
		ch := l.Next()
		// 为空继续  反之 退一格
		if !unicode.IsSpace(ch) {
			l.Dec()
			break
		}

		// EOF
		if ch == lexertoken.EOF {
			l.Emit(TOKEN_EOF)
			break
		}
	}
}

/*
返回一个包含错误信息的标记。
*/
func (l *LexerCalculation) Errorf(format string, args ...interface{}) LexFn {
	l.Tokens <- Token{
		Type:  TOKEN_ERROR,
		Value: fmt.Sprintf(format, args...),
	}

	return nil
}

//从通道中返回下一个令牌
func (l *LexerCalculation) NextToken() Token {
	for {
		select {
		case token := <-l.Tokens:
			return token
		default:
			l.State = l.State(l)
		}
	}
}
