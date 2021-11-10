package banking

import (
	"errors"
	"fmt"
)

// Important:
// Lowercase means the variable or function is private or not exportable.
// Uppercase means they are public or exportable.

type bankingAccount struct{
	// private fields
	owner string
	balance int
}

// A function that creates and returns the object, similar to constructors in Java or JS.
// Return is the value of the object's memory address, preventing from copy the object.
func CreateBankAccount (name string) *bankingAccount {
	bankingAccount := bankingAccount{owner: name, balance: 0}

	// Memory address is given for preventing duplicated creation of the object.
	return &bankingAccount
}

// Methods
// The things in between `func` and `deposit` is called reciever.
// `Reciever` writing convention: The left one is initial of the struct, the right is the struct name.
// This enables the object to be equipped with certain methods, like the 'Deposite' method below.
func (b *bankingAccount) Deposit(amount int) {
	b.balance += amount
}

// Error
// It is good practice to make the errors as variables for maintenance.
var errNoMoney = errors.New("Not enough money.") 
func (b *bankingAccount) Withdraw(amount int) error {
	if b.balance < amount {
		return errNoMoney // Return the error message.
	}
	b.balance -= amount
	return nil // Go's syntax requirement.
}

func (b *bankingAccount) CheckBalance() int {
	return b.balance
}

func (b bankingAccount) Owner() string {
	return b.owner
}

func (b *bankingAccount) ChangeOwner(newOwner string) {
	b.owner = newOwner
}

// Similar to 'toString()' in Java.
func (b bankingAccount) String() string {
	return fmt.Sprint(b.Owner(), " has $", b.balance)
}
