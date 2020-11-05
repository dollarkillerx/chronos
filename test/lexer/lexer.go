package lexer

import (
	"fmt"
	"github.com/dollarkillerx/chronos/test/lexertoken"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	Name   string
	Input  string                // 输入文本
	Tokens chan lexertoken.Token // 用于向词法分析发送Token 的Channel
	State  LexFn                 // 状态函数

	Start int // Token开始位置，结束位置可以通过start + len(token)获得
	Pos   int // 词法器处理文本位置 当确认Token结尾时,即相当于知道Token的end position
	Width int
}

type LexFn func(lexer *Lexer) LexFn // 词法器状态函数的定义，返回下一个期望 Token 的分析函数。

/**
将一个Token放入令牌Channel。这个标记的值是
根据当前词典的位置从输入端读取。
*/
func (l *Lexer) Emit(tokenType lexertoken.TokenType) {
	l.Tokens <- lexertoken.Token{Type: tokenType, Value: l.Input[l.Start:l.Pos]}
	l.Start = l.Pos
}

// 推进
func (l *Lexer) Inc() {
	l.Pos++
	if l.Pos >= utf8.RuneCountInString(l.Input) {
		l.Emit(lexertoken.TOKEN_EOF)
	}
}

// 倒退
func (l *Lexer) Dec() {
	l.Pos--
}

/**
返回当前词典位置的输入片段
到输入字符串的末尾。
*/
func (l *Lexer) InputToEnd() string {
	return l.Input[l.Pos:]
}

// 跳过空白处，直到我们得到有意义的东西。
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

func (l *Lexer) Next() rune {
	// 是否EOF
	if l.Pos >= utf8.RuneCountInString(l.Input) {
		l.Width = 0
		return lexertoken.EOF
	}

	// 读一个RUNE  返回item,长度
	result, width := utf8.DecodeRuneInString(l.Input[l.Pos:])
	l.Width = width
	l.Pos += l.Width
	return result
}

// 是否EOF
func (l *Lexer) IsEOF() bool {
	return l.Pos >= len(l.Input)
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
