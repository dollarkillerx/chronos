package lexertoken

type TokenType string

const (
	TOKEN_ERROR TokenType = "err"
	TOKEN_EOF   TokenType = "eof"

	TOKEN_LEFT_BRACKET  TokenType = "left_bracket"
	TOKEN_RIGHT_BRACKET TokenType = "right_bracket"
	TOKEN_EQUAL_SIGN    TokenType = "eq"

	TOKEN_SECTION TokenType = "section"
	TOKEN_KEY     TokenType = "key"
	TOKEN_VALUE   TokenType = "val"
)
