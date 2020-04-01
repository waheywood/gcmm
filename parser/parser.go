package parser

import (
	l "../lexer"
)

type ParseNode struct {
}

type Parser struct {
	tokens []l.LexedTokens
}

func NewParser(tokensToBeParsed []l.LexedTokens) Parser {
	var p Parser
	p.tokens = tokensToBeParsed
	return p
}

func (p Parser) StartParse() string {
	for i, v := range p.tokens {
		switch v.TokenType {
		case "RESERVED":
			p.reservedToken(p.tokens[i+1:])
		case "OPERATOR":
			p.operatorToken(p.tokens[i+1:])
		case "INTEGER":
			p.integerToken(p.tokens[i+1:])
		case "IDENTIFIER":
			p.identifierToken(p.tokens[i+1:])
		default:
			return "Not a valid token"
		}
	}
}

func (p Parser) reservedToken(t []l.LexedTokens) {

}

func (p Parser) operatorToken(t []l.LexedTokens) {

}

func (p Parser) integerToken(t []l.LexedTokens) {

}

func (p Parser) identifierToken(t []l.LexedTokens) {

}
