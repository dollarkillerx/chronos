package lexer

import (
	"log"
	"testing"
)

var TestSkipWhiteSpace1 = `
                    sd
我的天
这是测试代码
我雷德去
`

func TestSkipWhiteSpace(t *testing.T) {
	lexer := Lexer{Input: TestSkipWhiteSpace1}
	lexer.SkipWhiteSpace()
	log.Println(lexer.Start)
	log.Println(lexer.Pos)
	log.Println(lexer.Width)
	log.Println(lexer.Input[:lexer.Pos])
}
