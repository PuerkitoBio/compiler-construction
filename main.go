package main

import (
	"fmt"
	"github.com/PuerkitoBio/compiler-construction/scanner"
	"github.com/PuerkitoBio/compiler-construction/token"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Expected a file name as argument.")
	}
	f, e := os.Open(os.Args[1])
	if e != nil {
		panic(e)
	}
	s := scanner.NewScanner(f)
	for t := s.GetToken(); t.T != token.EOF; t = s.GetToken() {
		fmt.Printf("%+v\n", t)
	}
}
