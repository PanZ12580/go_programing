package main

import (
	"fmt"
	"go_programing/chapter9/bank"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		bank.Deposit(200)
	}()
	bank.Deposit(100)
	wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()
		bank.Withdraw(150)
	}()
	wg.Wait()
	fmt.Println(bank.Balance())
}
