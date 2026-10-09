package main

import (
	"fmt"
)

type Counter struct{ n int }

func (c Counter) IncValue()    { c.n++ } // works on a COPY
func (c *Counter) IncPointer() { c.n++ } // works on the ORIGINAL

func (c Counter) Get() int { return c.n }
func (c *Counter) Inc()    { c.n++ }
func NewCounter() Counter  { return Counter{} }


type Account struct {
	Owner string
	Balance float64
}
func (a *Account) Deposit(amount float64) { a.Balance += amount }


// tasks
// 1. Create BankAccount with an exported Owner and an unexported balance.
// 2. Add Deposit(amount float64) , Withdraw(amount float64) error (return an
// error if there's not enough money) and Balance() float64 .
// 3. Choose the receiver type for each method

type BankAccount struct{
	Owner string
	balance float64
}
func (b *BankAccount) Deposit(amount float64){ b.balance += amount }
func (b *BankAccount) Withdraw(amount float64) error {
	if(b.balance < amount){
		return fmt.Errorf("insufficient funds")
	}
	b.balance -= amount
	return nil
}
func (b *BankAccount) Balance() float64 {
	return b.balance
}


func main() {
	c := Counter{} //&Counter - addressable
	c.IncValue()
	fmt.Println(c.n) // 0
	c.IncPointer()   //(*c)
	fmt.Println(c.n) // 1

	//task: compile or not?
	// var v Counter
	// v.Inc()            //compile
	// (&v).Inc()         //no	| already does for us
	// Counter{}.Get()    //compile
	// Counter{}.Inc()    //compile
	// NewCounter().Inc() //no	| NewCounter() itself doesn't have that method

	// m := map[string]Counter{"a": {}}
	// m["a"].Get() //compile
	// m["a"].Inc()	//no	| bc map doesn't give a permission to do so, bc of rehashing

	// s := []Counter{{}}
	// s[0].Inc()

	// pm := map[string]*Counter{"a": {}}
	// pm["a"].Inc()

	accounts := []Account{{"Ali", 100}, {"Bota", 200}}
	for _, acc := range accounts {
		acc.Deposit(50)
	}
	fmt.Println(accounts) // 100, 200
	for i := range accounts {
		accounts[i].Deposit(50)
	}
	fmt.Println(accounts) // 150, 250


	// tasks
	bank_accounts := []*BankAccount{
		&BankAccount{"A", 10},
		&BankAccount{"B", 30},
		&BankAccount{"C", 100},
	}
	for _, v := range bank_accounts{
		v.Deposit(v.balance * 0.05)
	}
	for _, v := range bank_accounts{
		fmt.Println(v.Balance())
	}
}
