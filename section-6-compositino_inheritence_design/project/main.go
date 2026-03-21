// Bank Management System
package main

import (
	"errors"
	"fmt"
)

// Normal Account
type Account struct {
	AccountNumber     int
	Balance           float64
	AccountHolderName string
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit must be greater than 0")
	}
	a.Balance += amount
	fmt.Printf("Deposited $%.2f to %d. New Balance: $%.2f\n", amount, a.AccountNumber, a.Balance)
	return nil
}
func (acc *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Withdrawl amount must be positive")
	}
	if acc.Balance < amount {
		return fmt.Errorf("insufficient funds in %d. Balance: $%.2f, Tried to withdraw: $%.2f",
			acc.AccountNumber, acc.Balance, amount)
	}
	acc.Balance -= amount
	fmt.Printf("Withdrew $%.2f from %d. New Balance: $%.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil
}
func (acc *Account) GetBalance() float64 {
	return acc.Balance
}
func (acc *Account) String() string {
	return fmt.Sprintf("Account [%d] Owner: %s, Balance: $%.2f",
		acc.AccountNumber, acc.AccountHolderName, acc.Balance)
}

type SavingAccount struct {
	Account
	InterestRate float64
}

func (sa *SavingAccount) AddInterest() {
	interest := sa.Balance * sa.InterestRate
	fmt.Printf("Adding interest amount %.2f to saving account %d", interest, sa.AccountNumber)
	err := sa.Deposit(interest)
	if err != nil {
		fmt.Printf("AddInterst: Error depositing $%.2f to savings account. %v\n", interest, err)
	}
}

type OverdraftAccount struct {
	Account
	OverdraftLimit float64
}

func (oa *OverdraftAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	// Allow withdrawal up to Balance + OverdraftLimit
	if (oa.Balance + oa.OverdraftLimit) < amount {
		return fmt.Errorf("withdrawal of $%.2f exceeds overdraft limit for %d. Available including overdraft: $%.2f",
			amount, oa.AccountNumber, oa.Balance+oa.OverdraftLimit)
	}
	oa.Balance -= amount // Balance can go negative
	fmt.Printf("Withdrew $%.2f from overdraft account %d. New Balance: $%.2f\n", amount, oa.AccountNumber, oa.Balance)
	return nil
}

func main() {
	fmt.Println("---- Bank Account System ---")
	savAcc := SavingAccount{
		Account: Account{
			AccountNumber:     1,
			Balance:           1000.00,
			AccountHolderName: "Saver_Man",
		},
		InterestRate: 0.70,
	}
	fmt.Println("--- Saving Account Operation ---")
	fmt.Println(savAcc.Account.String())
	depositAmt := 120.00
	err := savAcc.Deposit(depositAmt)
	if err != nil {
		fmt.Printf("Error depositing $%.2f to savings account. %v\n", depositAmt, err)

	}
	savAcc.AddInterest()
	fmt.Println("Final Savings Details : ", savAcc.Account.String())

	fmt.Println("--- Overdraft Account Operation ---")
	ovdAcc := OverdraftAccount{
		Account: Account{
			AccountNumber:     2,
			Balance:           1200.00,
			AccountHolderName: "Overdraft_Man",
		},
		OverdraftLimit: 1200.00,
	}
	err = ovdAcc.Deposit(50.00)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = ovdAcc.Withdraw(200.00)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = ovdAcc.Withdraw(100.00)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = ovdAcc.Withdraw(100.00)
	if err != nil {
		fmt.Println("Error (expected for overdraft limit):", err)
	}

	fmt.Println("Final Overdraft Details:", ovdAcc.Account.String())

}
