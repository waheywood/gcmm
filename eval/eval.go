package eval

import (
	 "../parser"
)

func Eval(p *parser.Parser) {
	var total int64
	for i, v range p.Tree.Children {
		switch v.Type{
		case "Add":
			total += v.Value
		case "Div":
			total /= v.Value
		case "Mul":
			total *= v.Value
		case "Mod":
			total %= v.Value
		case "Exp": 
			total ^= v.Value
		case "print":
			fmt.Printf(v.Value + "\n")
		}
	}
}
