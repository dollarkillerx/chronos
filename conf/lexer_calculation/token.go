package lexer_calculation

type Token struct {
	Type  TokenType
	Value string
}

type TokenType string

const (
	TOKEN_EOF   TokenType = "eof"
	TOKEN_ERROR TokenType = "error"

	TOKEN_EVAL TokenType = "eval"
	TOKEN_AND  TokenType = "and"
	TOKEN_OR   TokenType = "or"

	TOKEN_LEFT_BRACKET  TokenType = "left_bracket"
	TOKEN_RIGHT_BRACKET TokenType = "right_bracket"

	TOKEN_SECTION    TokenType = "section"
	TOKEN_EQUAL_SIGN TokenType = "eq"
	TOKEN_KEY        TokenType = "key"
	TOKEN_VALUE      TokenType = "val"
)

const EOF rune = 0

const EVAL string = "eval"
const AND string = "&&"
const OR string = "or"
const LEFT_BRACKET string = "("
const RIGHT_BRACKET string = ")"
const EQUAL_SIGN string = "=="
