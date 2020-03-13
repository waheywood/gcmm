package lexer

import (
	"reflect"
	"strconv"
)

var operators map[string]bool

type LexedTokens struct {
	tokenType  string
	tokenValue interface{}
}

func init() {
	operators = make(map[string]bool)
	operators["="] = true
	operators["-"] = true
	operators["/"] = true
	operators["+"] = true
	operators["*"] = true
	operators["^"] = true
	operators["|"] = true
	operators["&"] = true
	operators[","] = true
	operators["%"] = true
	operators["("] = true
	operators[")"] = true
	operators[">"] = true
	operators["<"] = true
}

func Lex(tokens []string) (lexedTokens []LexedTokens) {

	for _, v := range tokens {
		if isOperator(v) {
			t := LexedTokens{
				tokenType:  "OPERATOR",
				tokenValue: v,
			}
			lexedTokens = append(lexedTokens, t)
		} else if isNumber(v) {
			t := LexedTokens{
				tokenType:  "INTEGER",
				tokenValue: v,
			}
			lexedTokens = append(lexedTokens, t)
		} else if isFloat(v) {
			t := LexedTokens{
				tokenType:  "FLOAT",
				tokenValue: v,
			}
			lexedTokens = append(lexedTokens, t)
		} else if isIdentifier(v) {
			t := LexedTokens{
				tokenType:  "IDENTIFIER",
				tokenValue: v,
			}
			lexedTokens = append(lexedTokens, t)
		}
	}

	return lexedTokens
}

func isOperator(token string) bool {
	if operators[token] {
		return true
	}
	return false
}

func isNumber(token string) bool {
	if _, err := strconv.ParseInt(token, 10, 64); err == nil {
		return true
	}
	return false
}

func isFloat(token string) bool {
	if _, err := strconv.ParseFloat(token, 64); err == nil {
		return true
	}
	return false
}

func isIdentifier(token string) bool {
	if reflect.TypeOf(token).Kind() == reflect.String {
		return true
	}
	return false
}
