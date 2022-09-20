package main

import (
	"testing"
)

func TestNewAccountPrintStatement(t *testing.T) {
	a := newAccount()

	have := a.printStatement()
	want := "Date\tAmount\tBalance\n"
	if have != want {
		t.Fatalf("want = %s, have = %s", want, have)
	}

}
