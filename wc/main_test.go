package main

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 ")
	exp := 2
	res := Count(b, false, false)
	if res != exp {
		t.Errorf("Expected %d , got %d instead", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word 1 workd2 \n word3")
	exp := 2
	res := Count(b, true, false)
	if res != exp {
		t.Errorf("Expected %d , got %d instead", exp, res)
	}
}
