package main

import (
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64)
	Name() string
}

type UPI struct{}
type Card struct{}
type Wallet struct{}

func (u UPI) Pay(amount float64) {
	fmt.Printf("Paid ₹%.2f using %s\n", amount, u.Name())
}
func (u UPI) Name() string { return "UPI" }

// Same for Card and Wallet
func (c Card) Pay(amount float64) {
	fmt.Printf("Paid ₹%.2f using %s\n", amount, c.Name())
}
func (c Card) Name() string { return "Card" }

func (w Wallet) Pay(amount float64) {
	fmt.Printf("Paid ₹%.2f using %s\n", amount, w.Name())
}
func (w Wallet) Name() string { return "Wallet" }

func ProcessPayment(method PaymentMethod, amount float64) {
	method.Pay(amount)
}

func main() {
	payments := []PaymentMethod{UPI{}, Card{}, Wallet{}}
	amount := 100.0
	for _, payment := range payments {
		ProcessPayment(payment, amount)
	}
}
