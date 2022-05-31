package hardcore_module

import "encoding/json"

const (
	TransactionIn=true
	TransactionOut=false
)
var	transacType=map[bool]string{true:"receive",false:"send"}

var Accounts Accs

type Accs struct {
	accounts map[int]Account
}

type Account struct {
	OwnerId int
	OwnerName string
	OwnerPhone string
	Money float64
	transactions []Transaction
}

type Transaction struct {
	TransactionDate string
	TransactionType string
	TransactionSum float64
	TransactionAim int
}

func NewAccount(name string, phone string)Account{
	acc:=Account{}
	acc.OwnerName=name
	acc.OwnerPhone=phone
	return acc
}

func (a *Account) SetId(id int){
	a.OwnerId=id
}

func (a Accs) SetMoney(accId int,money float64){
	acc:=a.accounts[accId]
	acc.Money=money
	a.accounts[accId]=acc
}

func (a Accs) AddTransaction(accId int,transactionType bool, transactionSum float64, transactionAim int, transactionDate string){
	acc:=a.accounts[accId]
	acc.transactions=append(acc.transactions, Transaction{
		transactionDate,
		transacType[transactionType],
		transactionSum,
		transactionAim,
	})
	a.accounts[accId]=acc
}

func (a Accs) AddAccount(acc Account){
	a.accounts[acc.OwnerId]=acc
}

func (a Accs) AccountCount() int{
	return len(a.accounts)
}

func (a Accs) GetMoneyLeft(accId int) float64{
	return a.accounts[accId].Money
}

func (a Accs) GetAccounts() string{
	accs,_:=json.MarshalIndent(a.accounts," ","\t")
	return string(accs)
}

func (a Accs) GetAccountInfo(accId int) string{
	acc,_:=json.MarshalIndent(a.accounts[accId]," ","\t")
	return string(acc)
}

func (a Accs) GetTransactions(accId int) string{
	transactions,_:=json.MarshalIndent(a.accounts[accId].transactions," ","\t")
	return string(transactions)
}

func init(){
	Accounts=Accs{make(map[int]Account)}
}