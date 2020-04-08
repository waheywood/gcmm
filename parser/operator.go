package parser

func AndToken(p Parser, i *int) {
	if p.Tokens[*i-1].TokenValue && p.Tokens[*i+1].TokenValue {
		p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "And", Value: true})
	}
	p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "And", Value: false})
}

func OrToken(p Parser, i *int) {
	if p.Tokens[*i-1].TokenValue || p.Tokens[*i+1].TokenValue {
		p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "And", Value: true})
	}
	p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "And", Value: false})

}

func Assign(p Parser, i *int) {

}

func Comma(p Parser, i *int) {

}

func Div(p Parser, i *int) {
	l := p.Tokens[*i-1].TokenValue
	r := p.Tokens[*i+1].TokenValue
	if p.Tokens[*i-1].TokenValue && p.Tokens[*i+1].TokenValue {
		p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "Div", Value: l / r})
	}

}

func Mod(p Parser, i *int) {
	l := p.Tokens[*i-1].TokenValue
	r := p.Tokens[*i+1].TokenValue
	if p.Tokens[*i-1].TokenValue && p.Tokens[*i+1].TokenValue {
		p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "Mod", Value: l % r})
	}
}

func NotEqToken(p Parser, i *int) {
	if p.Tokens[*i-1].TokenValue != p.Tokens[*i+1].TokenValue {
		p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "NotEq", Value: true})
	}
	p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "NotEq", Value: false})

}

func EqToken(p Parser, i *int) {
	if p.Tokens[*i-1].TokenValue == p.Tokens[*i+1].TokenValue {
		p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "Eq", Value: true})
	}
	p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "Eq", Value: false})

}

func Gte(p Parser, i *int) {
	if p.Tokens[*i-1].TokenValue >= p.Tokens[*i+1].TokenValue {
		p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "Gte", Value: true})
	}
	p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "Lte", Value: false})

}

func Lte(p Parser, i *int) {
	if p.Tokens[*i-1].TokenValue <= p.Tokens[*i+1].TokenValue {
		p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "Lte", Value: true})
	}
	p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "Lte", Value: false})

}

func Lt(p Parser, i *int) {
	if p.Tokens[*i-1].TokenValue < p.Tokens[*i+1].TokenValue {
		p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "Lt", Value: true})
	}
	p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "Lt", Value: false})

}

func Gt(p Parser, i *int) {
	if p.Tokens[*i-1].TokenValue > p.Tokens[*i+1].TokenValue {
		p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "Gt", Value: true})
	}
	p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "Gt", Value: false})

}

func Lb(p Parser, i *int) {

}

func Rb(p Parser, i *int) {

}

func Add(p Parser, i *int) {
	l := p.Tokens[*i-1].TokenValue
	r := p.Tokens[*i+1].TokenValue
	if p.Tokens[*i-1].TokenValue && p.Tokens[*i+1].TokenValue {
		p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "Add", Value: l + r})
	}
}

func Mul(p Parser, i *int) {
	l := p.Tokens[*i-1].TokenValue
	r := p.Tokens[*i+1].TokenValue
	if p.Tokens[*i-1].TokenValue && p.Tokens[*i+1].TokenValue {
		p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "Mul", Value: l * r})
	}
}

func Exp(p Parser, i *int) {
	l := p.Tokens[*i-1].TokenValue
	r := p.Tokens[*i+1].TokenValue
	if p.Tokens[*i-1].TokenValue && p.Tokens[*i+1].TokenValue {
		p.Tree.Children = append(p.Tree.Children, BlockNode{Type: "Exp", Value: l ^ r})
	}
}

func Lcb(p Parser, i *int) {

}

func Rcb(p Parser, i *int) {

}
