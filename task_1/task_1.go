package task_1

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	mute    sync.Mutex
}

func createBankAccount() BankAccount {
	return BankAccount{}
}

func (ba *BankAccount) deposit(amount int) {
	ba.mute.Lock()
	ba.balance += amount
	ba.mute.Unlock()
}

func (ba *BankAccount) withdraw(amount int) {
	ba.mute.Lock()
	defer ba.mute.Unlock()
	if ba.balance < amount {
		fmt.Println("Not enough money")
		return
	}
	ba.balance -= amount
}

func (ba *BankAccount) getBalance() {
	ba.mute.Lock()
	fmt.Println("Balance: ", ba.balance)
	ba.mute.Unlock()
}

func StartTask1() {
	bank := createBankAccount()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			bank.deposit(10)
			wg.Done()
		}()
		go func() {
			bank.withdraw(100)
			wg.Done()
		}()
		go func() {
			bank.getBalance()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Print("Final balance: ")
	bank.getBalance()
}
