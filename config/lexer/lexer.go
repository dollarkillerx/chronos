package lexer

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/dollarkillerx/chronos/config/lexertoken"
)

func BeginLexing(input string) *Lexer {
	l := &Lexer{
		Input:  input,
		State:  LexBegin,
		Tokens: make(chan lexertoken.Token, 3),
	}

	return l
}

type Lexer struct {
	Input  string                // 输入文本
	Tokens chan lexertoken.Token // 向解析器发送Token
	State  LexFn                 // 状态函数 当前是[ 下次期望是 ]的状态

	Start int
	Pos   int // Token end
	Width int
}

type LexFn func(lexer *Lexer) LexFn

// 提交当前Token
func (l *Lexer) Emit(tokenType lexertoken.TokenType) {
	l.Tokens <- lexertoken.Token{Type: tokenType, Value: l.Input[l.Start:l.Pos]}
	l.Start = l.Pos
}

// 前进
func (l *Lexer) Inc() {
	l.Pos++
	if l.Pos >= utf8.RuneCountInString(l.Input) {
		l.Emit(lexertoken.TOKEN_EOF)
	}
}

// 后退
func (l *Lexer) Dec() {
	l.Pos--
}

// 前Token 到最后面
func (l *Lexer) InputToEnd() string {
	return l.Input[l.Pos:]
}

// 是否EOF
func (l *Lexer) IsEOF() bool {
	return l.Pos >= utf8.RuneCountInString(l.Input)
}

// 向前推进一格
func (l *Lexer) Next() rune {
	// 是否EOF
	if l.IsEOF() {
		l.Width = 0
		return lexertoken.EOF
	}

	// 读一个RUNE  返回item,长度
	result, width := utf8.DecodeRuneInString(l.Input[l.Pos:])
	l.Width = width
	l.Pos += l.Width
	return result
}

// 跳过空白处
func (l *Lexer) SkipWhiteSpace() {
	for {
		ch := l.Next()
		// 为空继续  反之 退一格
		if !unicode.IsSpace(ch) {
			l.Dec()
			break
		}

		// EOF
		if ch == lexertoken.EOF {
			l.Emit(lexertoken.TOKEN_EOF)
			break
		}
	}
}

/*
返回一个包含错误信息的标记。
*/
func (l *Lexer) Errorf(format string, args ...interface{}) LexFn {
	l.Tokens <- lexertoken.Token{
		Type:  lexertoken.TOKEN_ERROR,
		Value: fmt.Sprintf(format, args...),
	}

	return nil
}

//从通道中返回下一个令牌
func (l *Lexer) NextToken() lexertoken.Token {
	for {
		select {
		case token := <-l.Tokens:
			return token
		default:
			l.State = l.State(l)
		}
	}
}
