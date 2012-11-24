package token

type Token int

// Tokens for a very simple first version of the language:
//
// var ident.
// ident = 3 + 8.
// ident = ident - 2.
// ident = ident * 9 / 3.
//
// Optional parenthesis, forced cast to integer, just integer and variables.
// Clever use of _bet and _end tags stolen from Go's tokens.
const (
	ILLEGAL Token = iota
	EOF
	COMMENT

	literal_beg
	IDENT
	INT
	literal_end

	operator_beg
	ADD
	SUB
	MUL
	DIV

	ASSIGN
	LPAREN
	RPAREN
	PERIOD
	operator_end

	keyword_beg
	VAR
	keyword_end
)

type TokenInfo struct {
	T        Token
	L        string
	StartPos int
	Line     int
	Col      int
	Len      int
}

var Tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	COMMENT: "COMMENT",

	IDENT: "IDENT",
	INT:   "INT",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",

	ASSIGN: "=",
	LPAREN: "(",
	RPAREN: ")",
	PERIOD: ".",

	VAR: "var",
}

var Keywords map[string]Token

func init() {
	Keywords = make(map[string]Token, keyword_end-keyword_beg-1)
	for i := keyword_beg + 1; i < keyword_end; i++ {
		Keywords[Tokens[i]] = i
	}
}
