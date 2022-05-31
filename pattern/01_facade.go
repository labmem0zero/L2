package main

import (
	"fmt"
	"main/hardcore_module"
	"time"
)

func FacadeRegAccount(name string, phone string)(accId int){
	newAcc:=hardcore_module.NewAccount(name,phone)
	id:=hardcore_module.Accounts.AccountCount()
	newAcc.SetId(id)
	hardcore_module.Accounts.AddAccount(newAcc)
	return id
}

func FacadeSendMoney(owner int,aim int, sum float64)string{
	hardcore_module.Accounts.AddTransaction(owner,hardcore_module.TransactionOut,sum,aim,time.Now().String())
	money:=hardcore_module.Accounts.GetMoneyLeft(owner)
	hardcore_module.Accounts.SetMoney(owner,money-sum)
	hardcore_module.Accounts.AddTransaction(aim,hardcore_module.TransactionIn,sum,owner,time.Now().String())
	money=hardcore_module.Accounts.GetMoneyLeft(aim)
	hardcore_module.Accounts.SetMoney(aim,money+sum)
	return fmt.Sprintf("Отправлено %.2f денежек с акканта %v на аккаунт %v",sum,owner,aim)
}

func FacadeAccInfo(id int)string{
	acc:=hardcore_module.Accounts.GetAccountInfo(id)
	accTransactions:=hardcore_module.Accounts.GetTransactions(id)
	return fmt.Sprintf("Основные данные:\n%s\nТранзакции:\n%s\n",acc,accTransactions)
}

func FacadeAllAccs()string{
	accs:=hardcore_module.Accounts.GetAccounts()
	return accs
}