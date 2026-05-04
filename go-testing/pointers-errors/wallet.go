package main

import "fmt"

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(bitcoin Bitcoin) {
	// fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
