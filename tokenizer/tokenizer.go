package tokenizer

import (
	"io"
	"os"
	"text/scanner"
)

//The tokenizer package reads the content of the file it is passed and returns
//a slice of tokens

func Tokenize(file string) (tokens []string, err error) {
	data, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	var s scanner.Scanner
	s.Init(io.Reader(data))

	for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
		tokens = append(tokens, s.TokenText())
	}

	return tokens, nil
}
