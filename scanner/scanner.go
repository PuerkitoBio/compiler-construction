package scanner

import (
	"bufio"
	"bytes"
	"github.com/PuerkitoBio/compiler-construction/token"
	"io"
	"os"
	"unicode"
)

type Scanner struct {
	r              *bufio.Reader // Input source code
	ch             rune          // Current character
	pos            int           // Current position
	line           int           // Current line
	col            int           // Current column within the line
	eof            bool
	err            error
	pendingIncLine bool
}

func NewScanner(f *os.File) *Scanner {
	s := &Scanner{r: bufio.NewReader(f), line: 1}
	s.read()
	return s
}

func (this *Scanner) read() {
	var sz int

	this.ch, sz, this.err = this.r.ReadRune()
	if this.err == nil {
		// Position in bytes increments with the size of the rune
		this.pos += sz

		if this.pendingIncLine {
			this.pendingIncLine = false
			this.line++
			this.col = 0
		}
		// Column increments by 1 (one unicode character read)
		this.col++
		// If this is a newline, will increment line number on next read
		if this.ch == '\n' {
			this.pendingIncLine = true
		}
	} else if this.err == io.EOF {
		this.eof = true
	}
}

func (this *Scanner) skipWhitespace() {
	for {
		switch this.ch {
		case ' ', '\n', '\t', '\r':
			this.read()
		default:
			return
		}
	}
}

func (this *Scanner) scanIdentifier() string {
	var b bytes.Buffer

	for isLetter(this.ch) || isDigit(this.ch) {
		b.WriteRune(this.ch)
		this.read()
	}
	return b.String()
}

func (this *Scanner) scanInteger() string {
	var b bytes.Buffer

	for isDigit(this.ch) {
		b.WriteRune(this.ch)
		this.read()
	}
	return b.String()
}

func (this *Scanner) consumeRestOfLine() string {
	var b bytes.Buffer

	for this.ch != '\n' && !this.eof {
		b.WriteRune(this.ch)
		this.read()
	}
	return b.String()
}

func lookupKeyword(lit string) token.Token {
	t, kw := token.Keywords[lit]
	if kw {
		return t
	}
	return token.IDENT
}

func (this *Scanner) scan() token.TokenInfo {
	var ti token.TokenInfo

	// Skip blanks
	this.skipWhitespace()
	ti.StartPos = this.pos
	ti.Line = this.line
	ti.Col = this.col

	// Find token
	switch {
	case isLetter(this.ch):
		// Scan an identifier, then check if this is a keyword. The output of this case
		// can be an identifier or a keyword.
		ti.L = this.scanIdentifier()
		ti.T = lookupKeyword(ti.L)

	case isDigit(this.ch):
		ti.T = token.INT
		ti.L = this.scanInteger()

	default:
		// Check for EOF, error
		if this.eof {
			ti.T = token.EOF
		} else if this.err != nil {
			ti.T = token.ILLEGAL
			ti.L = this.err.Error()
		} else {
			// Case for operators, always read another char for next call
			ch := this.ch
			this.read()
			ti.L = string(ch)
			switch ch {
			case '.':
				ti.T = token.PERIOD
			case '+':
				ti.T = token.ADD
			case '-':
				ti.T = token.SUB
			case '*':
				ti.T = token.MUL
			case '/':
				if this.ch == '/' {
					ti.T = token.COMMENT
					// Get ready for next char
					this.read()
					ti.L = this.consumeRestOfLine()
				} else {
					ti.T = token.DIV
				}
			case '=':
				ti.T = token.ASSIGN
			}
		}
	}

	ti.Len = len(ti.L)
	return ti
}

func (this *Scanner) GetToken() token.TokenInfo {
	return this.scan()
}

func isLetter(ch rune) bool {
	// Copied from go's scanner
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= 0x80 && unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	// Copied from go's scanner, limited to 0-9
	return '0' <= ch && ch <= '9'
}
