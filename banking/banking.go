package main

type Account struct {
}

func newAccount() Account {
	a := Account{}
	return a
}

func (a *Account) printStatement() string {
	ret := "Date\tAmount\tBalance\n"
	return ret
}
