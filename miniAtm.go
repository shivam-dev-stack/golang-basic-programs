package main

import (
	"fmt"
)

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64)
	CheckBalance() float64
	AccountType() string
}

type SavingsAccount struct {
	balance float64
}
type CurrentAccount struct {
	balance float64
}

func (s *SavingsAccount) Deposit(amount float64) {
	s.balance += amount
}

func (s *SavingsAccount) Withdraw(amount float64) {
	if amount > s.balance {
		fmt.Println("Insufficient balance!")
		return
	}
	s.balance -= amount
}

func (s *SavingsAccount) CheckBalance() float64 {
	return s.balance
}

func (s *SavingsAccount) AccountType() string {
	return "Savings Account"
}

func (c *CurrentAccount) Deposit(amount float64) {
	c.balance += amount
}

func (c *CurrentAccount) Withdraw(amount float64) {
	if amount > c.balance {
		fmt.Println("Insufficient balance!")
		return
	}
	c.balance -= amount
}

func (c *CurrentAccount) CheckBalance() float64 {
	return c.balance
}

func (c *CurrentAccount) AccountType() string {
	return "Current Account"
}

func OperateAccount(acc Account) {
	fmt.Printf("Operating on: %s\n", acc.AccountType())
	acc.Deposit(1000)
	acc.Withdraw(400)
	fmt.Printf("Balance: ₹%.2f\n\n", acc.CheckBalance())
}

func main() {
	savings := &SavingsAccount{balance: 500}
	current := &CurrentAccount{balance: 2000}

	accounts := []Account{savings, current}
	for _, acc := range accounts {
		OperateAccount(acc)
	}
}
