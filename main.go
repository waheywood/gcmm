package main

import (
	"./lexer"
	"./tokenizer"
	"fmt"
	"log"
)

func main() {

	var file string
	file = "tost.cmm"
	tokens, err := tokenizer.Tokenize(file)

	if err != nil {
		log.Fatalf("Could not tokenize file: %s", err)
	}

	lexedTokens := lexer.Lex(tokens)

	fmt.Println(lexedTokens)

	//parse, err := parser.Parser.Parse(lex)
	//
	//if err != nil {
	//	log.Fatalf("Could not parse the lexed tokens: %s", err)
	//}

}
