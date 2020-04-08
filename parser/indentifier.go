package parser

func IdentToken(p Parser, i *int) {
	if p.Tokens[*i+1].TokenValue == "=" {
		p.Vars[p.Tokens[*i].TokenValue] = p.Tokens[*i+2].TokenValue
	}
}
