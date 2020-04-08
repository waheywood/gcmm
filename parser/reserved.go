package parser

import (
	l "../lexer"
)

func IfToken(p Parser, i *int) {
	var node BlockNode
	node.Type = "if"
	var values []l.LexedTokens
	*i++
	for {
		if p.Tokens[*i+1].TokenValue != "}" {
			if p.Tokens[*i].TokenValue == "(" ||
				p.Tokens[*i].TokenValue == ")" ||
				p.Tokens[*i].TokenValue == "{" {
				*i++
				continue
			}
			*i++
			values = append(values, p.Tokens[*i])
		} else {
			break
		}

	}
	np := NewParser(values)
	_ = np.StartParse()
	p.Tree.Children = append(p.Tree.Children, np.Tree)

}

func ElseToken(p Parser, i *int) {
	IfToken(p, i)
}

func WhileToken(p Parser, i *int) {

}

func PrintToken(p Parser, i *int) {
	var print BlockNode
	var tokens []string
	for _, v := range p.Tokens[*i+1:] {
		if v.TokenType == "IDENTIFIER" || v.TokenType == "INTEGER" {
			tokens = append(tokens, v.TokenValue)
		} else if v.TokenType == "RESERVED" {
			break
		}
	}
	print.Type = "print"
	print.Value = tokens
	p.Tree.Children = append(p.Tree.Children, print)
}
