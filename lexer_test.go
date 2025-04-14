package main

import (
	"bufio"
	"strings"
	"testing"
)

var test_text = `
if 1 < 2 then
	print "HI"
endif`

func TestLexer(t *testing.T) {
	want := []Token{0, 101, 2, 207, 2, 102, 0, 112, 3, 0, 103}
	r := bufio.NewReader(strings.NewReader(text))
	tokens := Lexer(r)
	for i := range want {
		if len(tokens) <= i {
			t.Errorf("Got not enough tokens, got only %d right tokens, wanted %d", len(tokens), len(want))
			break
		}
		if want[i] != tokens[i] {
			t.Errorf("Wrong token in place %d. Wanted %s, got %s", i, want[i], tokens[i])
			break
		}
	}
	if len(tokens) > len(want) {
		t.Errorf("Too much tokens. Wanted %d, got %d", len(want), len(tokens))
	}
}
