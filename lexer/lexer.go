package lexer

import (
	"reflect"
	"strconv"
)

var operators map[string]bool
var reservedWords map[string]bool

type LexedTokens struct {
	tokenType  string
	tokenValue string
}

func init() {
	//Initialise the maps
	operators = make(map[string]bool)
	reservedWords = make(map[string]bool)
	//Populate operator map with all operators
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
	operators["!"] = true

	//Populate reservedWords map
	reservedWords["if"] = true
	reservedWords["else"] = true
	reservedWords["print"] = true
	reservedWords["while"] = true

}

func Lex(tokens []string) (lexedTokens []LexedTokens) {

	for i := 0; i < len(tokens); i++ {
		curToken := tokens[i]
		if isOperator(curToken) {
			t := LexedTokens{
				tokenType:  "OPERATOR",
				tokenValue: curToken,
			}
			var nextToken string
			if i < len(tokens)-1 {
				nextToken = tokens[i+1]
			} else {
				lexedTokens = append(lexedTokens, t)
				continue
			}

			if isOperator(nextToken) {
				if curToken == "!" && nextToken == "=" {
					t.tokenValue = t.tokenValue + nextToken
				} else if curToken == "=" && nextToken == "=" {
					t.tokenValue = t.tokenValue + nextToken
				} else if curToken == "|" && nextToken == "|" {
					t.tokenValue = t.tokenValue + nextToken
				} else if curToken == ">" && nextToken == "=" {
					t.tokenValue = t.tokenValue + nextToken
				} else if curToken == "<" && nextToken == "=" {
					t.tokenValue = t.tokenValue + nextToken
				} else if curToken == "&" && nextToken == "&" {
					t.tokenValue = t.tokenValue + nextToken
				}
				lexedTokens = append(lexedTokens, t)
				i++
				continue
			}
			lexedTokens = append(lexedTokens, t)
		} else if isNumber(curToken) {
			t := LexedTokens{
				tokenType:  "INTEGER",
				tokenValue: curToken,
			}
			lexedTokens = append(lexedTokens, t)
		} else if isFloat(curToken) {
			t := LexedTokens{
				tokenType:  "FLOAT",
				tokenValue: curToken,
			}
			lexedTokens = append(lexedTokens, t)
		} else if isIdentifier(curToken) {
			t := LexedTokens{
				tokenType:  "IDENTIFIER",
				tokenValue: curToken,
			}
			if isReservedWord(curToken) {
				t.tokenType = "RESERVED"
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

func isReservedWord(token string) bool {
	if reservedWords[token] {
		return true
	}
	return false
}
