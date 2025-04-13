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

func (t Token) String() string {
	if t < 100 {
		return []string{"Token.newline", "Token.ident", "Token.number", "Token.string"}[t]
	} else if t < 200 {
		return []string{"Token.if", "Token.then", "Token.endif", "Token.while", "Token.do", "Token.endwhile",
			"Token.let", "Token.var", "Token.const", "Token.becomes", "Token.input", "Token.print", "Token.comma",
			"Token.period", "Token.proc", "Token.begin", "Token.end", "Token.call", "Token.lparen", "Token.rparen"}[t-101]
	} else {
		return []string{"Token.plus", "Token.minus", "Token.times", "Token.slash", "Token.eql", "Token.neq", "Token.lss",
			"Token.leq", "Token.gtr", "Token.geq"}[t-201]
	}
}
