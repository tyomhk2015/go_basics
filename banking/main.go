package main

import (
	"fmt"

	"github.com/tyomhk2015/go_basics/banking/banking"
)

func main() {
	// Creating an object.
	// Syntax: variable := package.ConstructorFunction()
	myBankAccount := banking.CreateBankAccount("ASH")
	myBankAccount.Deposit(20)
	fmt.Println(myBankAccount.CheckBalance())

	err := myBankAccount.Withdraw(30)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}

	fmt.Println(myBankAccount.CheckBalance())
	myBankAccount.Deposit(44)
	fmt.Println(myBankAccount.CheckBalance())

	myBankAccount.ChangeOwner("Gary")
	fmt.Println(myBankAccount.String())
}