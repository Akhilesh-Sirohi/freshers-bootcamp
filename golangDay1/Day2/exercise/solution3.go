package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	balance int=500
	mutex sync.Mutex
)
func deposit(amount int, wg *sync.WaitGroup){
	mutex.Lock();
	fmt.Printf("deposited %d to account with balance %d\n", amount,balance)
	balance+=amount;
	mutex.Unlock()
	wg.Done();
}
func withdraw(amount int, wg *sync.WaitGroup){
	if(balance<amount){
		fmt.Printf("cannot withdraw as balance is %d\n", balance);
		time.Sleep(time.Second)
	}
	mutex.Lock();
	fmt.Printf("withdraw %d to account with balance %d\n", amount,balance)
	balance-=amount;
	mutex.Unlock()
	wg.Done();
}
func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	go deposit(1000,&wg)
	go withdraw(600,&wg)
	wg.Wait()
}
