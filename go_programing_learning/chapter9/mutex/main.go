package main

import (
	"fmt"
	"go_programing/chapter9/mutex/bank"
)

func main() {
	go func() {
		bank.Deposit(200)
		fmt.Println(bank.Balance())
	}()

	bank.Deposit(200)
	bank.Withdraw(100)
	fmt.Println(bank.Balance())
}
