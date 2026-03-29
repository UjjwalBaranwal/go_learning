package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	mtx     sync.Mutex
}

func (b *BankAccount) Deposit(val int) {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	b.balance += val
	fmt.Println("Deposit ", val)
	fmt.Println("Current Balance ", b.balance)
}
func (b *BankAccount) Withdraw(val int) {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	if b.balance < val {
		fmt.Println("cannot withdraw that val:", val)
		return
	}

	b.balance -= val
	fmt.Println("Withdraw", val)
}
func (b *BankAccount) Balance() {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	fmt.Println("Balance ", b.balance)
}
func main() {
	// now doing the mutex solution in the bank system
	var wg sync.WaitGroup
	var account = &BankAccount{
		balance: 100,
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			time.Sleep(time.Duration(amount) * time.Millisecond)
			account.Deposit(amount)
		}(i + 1)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			time.Sleep(time.Duration(amount) * time.Millisecond)
			account.Withdraw(amount * 10)
		}(i + 1)
	}

	wg.Wait()
	account.Balance()

	// demostrateRaceCondition()
}
func demostrateRaceCondition() {
	var mtx sync.Mutex
	// this is a classic race condition problem
	// can be solve via mutex
	cnt := 0 // critical section
	var wg sync.WaitGroup
	// without using mutex
	// increaseCnt := func() {
	// 	defer wg.Done()
	// 	cnt++
	// 	fmt.Println(cnt)
	// }
	//with using mutex
	increaseCnt := func() {
		defer wg.Done()
		mtx.Lock()
		defer mtx.Unlock()
		cnt++

	}
	for range 1000 {
		// fmt.Println("crnt ", i)
		wg.Add(1)
		go increaseCnt()
	}

	wg.Wait()
	fmt.Println("cnt : ", cnt)
}
