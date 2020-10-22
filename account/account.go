package account

import (
	"errors"
	"fmt"
)

//Account struct
type Account struct {
	owner   string //첫글자가 소문자면 private상태
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

//NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit x amount on your account
func (a *Account) Deposit(amount int) { //receiver  Go에서는 이게 method
	//fmt.Println("Gonna deposit", amount)
	a.balance += amount //이 메소드라는 것은 Account의 식구가 된 것이다.
	//뜬금없이 갑자기 왜 식구라는 이야기를 했냐면 이제부턴 private변수에 마음껏 접근할 수 있다.
}

//Balance of yout account
func (a Account) Balance() int {
	return a.balance
}

//Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// CHangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
