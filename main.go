package main

import (
	"fmt"
	"log"

	"./lexer"
	"./parser"
	"./tokenizer"
)

func main() {

	var file string
	file = "tost.cmm"
	tokens, err := tokenizer.Tokenize(file)

	if err != nil {
		log.Fatalf("Could not tokenize file: %s", err)
	}
	fmt.Println(tokens)
	lexedTokens := lexer.Lex(tokens)

	fmt.Println(lexedTokens)

	p := parser.NewParser(lexedTokens)
	p.StartParse()

	//if err != nil {
	//	log.Fatalf("Could not parse the lexed tokens: %s", err)
	//}

}
