package parser

import (
	"log"

	l "../lexer"
)

type BlockNode struct {
	Children []ParseNode
	Type     string
	Value    interface{}
}

type Variables struct {
}

type ParseNode interface{}

type Parser struct {
	Tokens []l.LexedTokens
	Tree   BlockNode
	Vars   map[string]interface{}
}

func NewParser(tokensToBeParsed []l.LexedTokens) Parser {
	var p Parser
	p.Tokens = tokensToBeParsed
	p.Vars = make(map[string]interface{})
	return p
}

func (p Parser) StartParse() error {

	for i, v := range p.Tokens {
		switch v.TokenType {
		case "RESERVED":
			p.reservedToken(&i)
		case "OPERATOR":
			p.operatorToken(&i)
		case "INTEGER":
			p.integerToken(&i)
		case "IDENTIFIER":
			p.identifierToken(&i)
		default:
			return "Not a valid token"
		}
	}
	return error{}
}

func (p Parser) reservedToken(i *int) {
	switch p.Tokens[*i].TokenValue {
	case "if":
		node, err := IfToken(p, i)
	case "else":
		ElseToken(p, i)
	case "while":
		WhileToken(p, i)
	case "print":
		PrintToken(p, i)
	}
}

func (p Parser) operatorToken(i *int) {
	switch p.Tokens[*i].TokenValue {
	case "&&":
		AndToken(p, i)
	case "||":
		OrToken(p, i)
	case "/":
		Div(p, i)
	case "%":
		Mod(p, i)
	case "!=":
		NotEqToken(p, i)
	case "==":
		EqToken(p, i)
	case ">=":
		Gte(p, i)
	case "<=":
		Lte(p, i)
	case "<":
		Lt(p, i)
	case ">":
		Gt(p, i)
	case "+":
		Add(p, i)
	case "*":
		Mul(p, i)
	case "^":
		Exp(p, i)
	default:
		log.Fatal("Unrecognised token")
	}
}

func (p Parser) integerToken(i *int) {

}

func (p Parser) identifierToken(i *int) {
	IdentToken(p, i)
}
