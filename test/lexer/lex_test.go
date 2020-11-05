package lexer

import (
	"log"
	"strings"
	"testing"
)

var TestSkipWhiteSpace1 = `
        
 [Stats]

 va = sdsd
`

func TestSkipWhiteSpace(t *testing.T) {
	lexer := Lexer{Input: TestSkipWhiteSpace1}
	lexer.SkipWhiteSpace()
	log.Println(lexer.Start)
	log.Println(lexer.Pos)
	log.Println(lexer.Width)
	log.Println(lexer.Input[:lexer.Pos])
}

func TestInputToEnd(t *testing.T) {
	lexer := Lexer{Input: TestSkipWhiteSpace1}
	lexer.SkipWhiteSpace()
	log.Println(strings.HasPrefix(lexer.InputToEnd(), "["))
}
