package lexer

import (
	"testing"

	"internal/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	test := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{}
}
