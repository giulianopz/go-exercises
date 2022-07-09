package pointers

import (
	"errors"
	"fmt"
)

// extend an existing type
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Printf("memory address of balance in Deposit is %v \n", &w.balance)
	// there's no need to dereference a pointer
	// these types of pointers are called "strcut pointers" and they are automatically deferenced
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// to double-check that calling code check the err returned by this func, execute: errcheck .
// you can install it by: go install github.com/kisielk/errcheck@latest
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
