package main

type Token int

const (
	TokenEof Token = iota - 1 //-1
	TokenNewline
	TokenIdent
	TokenNumber
	TokenString
)

const (
	TokenIf Token = iota + 101 //101
	TokenThen
	TokenEndif
	TokenWhile
	TokenDo
	TokenEndwhile
	TokenLet
	TokenVar
	TokenConst
	TokenBecomes
	TokenInput
	TokenPrint
	TokenComma
	TokenPeriod
	TokenProc
	TokenBegin
	TokenEnd
	TokenCall
	TokenLparen
	TokenRparen
)

const (
	TokenPlus Token = iota + 201 //201
	TokenMinus
	TokenTimes
	TokenSlash
	TokenEql
	TokenNeq
	TokenLss
	TokenLeq
	TokenGtr
	TokenGeq
)
