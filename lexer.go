package main

import (
	"bufio"
	"fmt"
	"io"
)

var curChar rune
var r *bufio.Reader

func charNext() Token {
	var err error
	if err == io.EOF {
		return TokenEof
	}
	return tokenGet()
}

func tokenGet() Token {
	if curChar == '\n' {
		return TokenNewline
	}

	if curChar >= 'a' && curChar <= 'z' {
		word := []rune{curChar}
		for peek() >= 'a' && peek() <= 'z' {
			goNextChar()
			word = append(word, curChar)
		}
		if keyword, ok := checkKeyword(string(word)); ok {
			return keyword
		}
		return TokenIdent
	}
	if curChar == '"' {
		for peek() != '"' {
			goNextChar()
		}
		goNextChar()
		return TokenString
	}
	if curChar >= '0' && curChar <= '9' {
		for peek() >= '0' && peek() <= '9' {
			goNextChar()
		}
		return TokenNumber
	}

	if curChar == '(' {
		return TokenLparen
	}
	if curChar == ')' {
		return TokenLparen
	}

	if curChar == '+' {
		return TokenPlus
	}
	if curChar == '-' {
		return TokenMinus
	}
	if curChar == '*' {
		return TokenTimes
	}
	if curChar == '/' {
		return TokenSlash
	}
	if curChar == '=' {
		switch peek() {
		case '=':
			goNextChar()
			return TokenEql
		default:
			return TokenBecomes
		}
	}
	if curChar == '<' {
		switch peek() {
		case '=':
			goNextChar()
			return TokenLeq
		default:
			return TokenLss
		}
	}
	if curChar == '>' {
		switch peek() {
		case '=':
			goNextChar()
			return TokenGeq
		default:
			return TokenGtr
		}
	}
	panic(fmt.Errorf("invalid symbol: %c", curChar))
}

func checkKeyword(w string) (Token, bool) {
	if w == "if" {
		return TokenIf, true
	}
	if w == "then" {
		return TokenThen, true
	}
	if w == "endif" {
		return TokenEndif, true
	}
	if w == "while" {
		return TokenWhile, true
	}
	if w == "do" {
		return TokenDo, true
	}
	if w == "endwhile" {
		return TokenEndwhile, true
	}
	if w == "let" {
		return TokenLet, true
	}
	if w == "input" {
		return TokenInput, true
	}
	if w == "print" {
		return TokenPrint, true
	}
	return -1, false
}

func peek() rune {
	c, _, err := r.ReadRune()
	if err != nil {
		if err == io.EOF {
			return 0
		} else {
			panic(err)
		}
	}
	if err = r.UnreadRune(); err != nil {
		panic(err)
	}
	return c
}

func goNextChar() bool {
	c, _, err := r.ReadRune()
	if err != nil {
		if err == io.EOF {
			return false
		} else {
			panic(err)
		}
	}
	curChar = c
	return true
}

func Lexer(rr *bufio.Reader) {
	r = rr
	for {
		if goNextChar() == false {
			break
		}
		if curChar == ' ' || curChar == '\r' || curChar == '\t' {
			continue
		}
		token := tokenGet()
		fmt.Println(token)
	}
}
