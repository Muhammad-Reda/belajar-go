package strc

import "fmt"

type number int

type BankAccount struct {
	Name    string
	Balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.Balance += amount
}

func (b *BankAccount) WithDraw(amount int) error {
	if b.Balance < amount {
		return fmt.Errorf("unsufficient balance")
	}
	b.Balance -= amount
	return nil
}

func (b BankAccount) CheckBalance() {
	fmt.Printf("%s's balance: %d\n\n", b.Name, b.Balance)
}

func SaveBalance(amount int) int {
	bank := BankAccount{Name: "Redha", Balance: 1000_000}

	bank.Deposit(amount)
	return bank.Balance
}
