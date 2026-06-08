package main

import "fmt"

type Account struct {
    Owner   string
    Balance float64
}

func Deposit(a *Account, amount float64) {  // pointer — changes balance
    a.Balance = a.Balance + amount
}

func PrintBalance(a Account) {              // value — only reads
    fmt.Println(a.Owner, "has balance:", a.Balance)
}

func main() {
    acc := Account{Owner: "Manasa", Balance: 1000000000}
    PrintBalance(acc)      
    Deposit(&acc, 500)
    PrintBalance(acc)     
    Deposit(&acc, 250)
    PrintBalance(acc)      
}