package parser

import (
	l "../lexer"
)

type BlockNode struct {
	Children []ParseNode
	Type     string
}

type ParseNode interface{}

type Parser struct {
	Tokens []l.LexedTokens
	Tree   BlockNode
}

func NewParser(tokensToBeParsed []l.LexedTokens) Parser {
	var p Parser
	p.Tokens = tokensToBeParsed
	return p
}

func (p Parser) StartParse() string {

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
		node, i, err := IfToken(p, i)
	case "else":
		ElseToken(p, i)
	case "while":
		WhileToken(p, i)
	case "print":
		PrintToken(p, i)
	}
}

func (p Parser) operatorToken(i int) {

}

func (p Parser) integerToken(i int) {

}

func (p Parser) identifierToken(i int) {

}
