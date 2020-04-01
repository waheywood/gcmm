package parser

import (
	"fmt"
	"strconv"

	l "../lexer"
)

type Parser struct {
	tokens []l.LexedTokens
}

func NewParser(tokensToBeParsed []l.LexedTokens) Parser {
	var p Parser
	p.tokens = tokensToBeParsed
	return p
}

func (p Parser) StartParse() {
	for i, v := range p.tokens {
		fmt.Println(strconv.Itoa(i) + " " + v.TokenType + " " + v.TokenValue)
	}
}
