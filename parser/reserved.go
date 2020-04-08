package parser

import (
	"fmt"

	l "../lexer"
)

func IfToken(p Parser, i int) (BlockNode, int, string) {
	if p.Tokens[i+1].TokenValue != "(" {
		return BlockNode{}, i, fmt.Sprintf("Invalid token, missing ( after if at %d", i)
	}
	var node BlockNode
	node.Type = "if"
	var values []l.LexedTokens
	i++
	for {
		if p.Tokens[i+1].TokenValue != "}" {
			if p.Tokens[i].TokenValue == "(" ||
				p.Tokens[i].TokenValue == ")" ||
				p.Tokens[i].TokenValue == "{" ||
				p.Tokens[i].TokenValue == "}" {
				i++
				continue
			}
			i++
			values = append(values, p.Tokens[i])
		} else {
			break
		}

	}
	np := NewParser(values)
	return np.Tree, i, ""
}

func ElseToken(p Parser, i int) (BlockNode, int, string) {

}

func WhileToken(p Parser, i int) (BlockNode, int, string) {

}

func PrintToken(p Parser, i int) (BlockNode, int, string) {

}
