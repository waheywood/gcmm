package main

import (
	"./lexer"
	"./tokenizer"
	"log"
)

func main() {

	var file string
	file = "tost.cmm"
	tokens, err := tokenizer.Tokenize(file)

	if err != nil {
		log.Fatalf("Could not tokenize file: %s", err)
	}

	lex, err := lexer.Lex(tokens)

	if err != nil {
		log.Fatalf("Could not lex the tokens: %s", err)
	}

	//parse, err := parser.Parser.Parse(lex)
	//
	//if err != nil {
	//	log.Fatalf("Could not parse the lexed tokens: %s", err)
	//}

}
