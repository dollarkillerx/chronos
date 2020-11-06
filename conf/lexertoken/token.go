package lexertoken

import (
	"fmt"
)

type Token struct {
	Type  TokenType
	Value string
}

func (t *Token) String() string {
	switch t.Type {
	case TOKEN_EOF:
		return "EOF"

	case TOKEN_ERROR:
		return t.Value
	}

	return fmt.Sprintf("%q", t.Value)
}

func IsEOF(token Token) bool {
	return token.Type == TOKEN_EOF
}
