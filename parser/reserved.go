package parser

import (
	"fmt"

	l "../lexer"
)

func IfToken(p Parser, i *int) (BlockNode, string) {
	if p.Tokens[*i+1].TokenValue != "(" {
		return BlockNode{}, fmt.Sprintf("Invalid token, missing ( after if at %d", i)
	}
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
	err := np.StartParse()
	if err != "" {
		return BlockNode{}, "Error in parsing"
	}
	return np.Tree, ""
}

func ElseToken(p Parser, i *int) (BlockNode, string) {
	return IfToken(p, i)
}

func WhileToken(p Parser, i *int) (BlockNode, string) {

}

func PrintToken(p Parser, i *int) (BlockNode, string) {
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
	return print, ""
}
