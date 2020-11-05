package lexertoken

type Token struct {
	Type  TokenType
	Value string
}

type TokenType int

const (
	TOKEN_ERROR TokenType = iota
	TOKEN_EOF

	TOKEN_LEFT_BRACKET
	TOKEN_RIGHT_BRACKET
	TOKEN_EQUAL_SIGN
	TOKEN_NEWLINE

	TOKEN_SECTION
	TOKEN_KEY
	TOKEN_VALUE
)

const EOF rune = 0

const LEFT_BRACKET string = "["
const RIGHT_BRACKET string = "]"
const EQUAL_SIGN string = "="
const NEWLINE string = "\n"
