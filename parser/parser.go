package parser

import (
	l "../lexer"
)

type ParseNode struct {
	children []ParseNode
}

type Parser struct {
	Tokens []l.LexedTokens
	Tree   ParseNode
}

func NewParser(tokensToBeParsed []l.LexedTokens) Parser {
	var p Parser
	p.Tokens = tokensToBeParsed
	return p
}

func (p Parser) StartParse() string {
	p.Tree = ParseNode{}
	for i, v := range p.Tokens {
		switch v.TokenType {
		case "RESERVED":
			p.reservedToken(i)
		case "OPERATOR":
			p.operatorToken(i)
		case "INTEGER":
			p.integerToken(i)
		case "IDENTIFIER":
			p.identifierToken(i)
		default:
			return "Not a valid token"
		}
	}
	return ""
}

func (p Parser) reservedToken(i int) {
	switch p.Tokens[i].TokenValue {
	case "if":
		IfToken()
	case "else":
		ElseToken()
	case "while":
		WhileToken()
	case "print":
		PrintToken()
	}
}

func (p Parser) operatorToken(i int) {

}

func (p Parser) integerToken(i int) {

}

func (p Parser) identifierToken(i int) {

}
