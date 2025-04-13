package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var text = `
if 1 < 2 then
	print "HI"
endif`

func main() {
	r := bufio.NewReader(strings.NewReader(text))
	fmt.Println("Hi, Daddy")
	Lexer(r)
	fmt.Println()
	f1, err := os.OpenFile("test1", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	rd1 := bufio.NewReader(f1)
	Lexer(rd1)
	fmt.Println()
	f2, err := os.OpenFile("test2", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	rd2 := bufio.NewReader(f2)
	Lexer(rd2)
}
