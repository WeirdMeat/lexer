package main

type Symbol string

const (
	SymbolIf       Symbol = "if"
	SymbolThen            = "then"
	SymbolEndif           = "endif"
	SymbolWhile           = "while"
	SymbolDo              = "do"
	SymbolEndwhile        = "endwhile"
	SymbolLet             = "let"
	SymbolVar             = "var"
	SymbolConst           = "const"
	SymbolBecomes         = "="
	SymbolInput           = "input"
	SymbolPrint           = "print"
	SymbolComma           = ","
	SymbolPeriod          = "."
	SymbolProc            = "procedure"
	SymbolBegin           = "begin"
	SymbolEnd             = "end"
	SymbolCall            = "call"
	SymbolLparen          = "("
	SymbolRparen          = ")"
)

const (
	SymbolPlus  Symbol = "+"
	SymbolMinus        = "-"
	SymbolTimes        = "*"
	SymbolSlash        = "/"
	SymbolEql          = "=="
	SymbolNeq          = "!="
	SymbolLss          = "<"
	SymbolLeq          = "<="
	SymbolGtr          = ">"
	SymbolGeq          = ">="
)
