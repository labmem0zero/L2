package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func Facade01(){
	fmt.Println("Зарегистрировал аккаунт:",FacadeRegAccount("Vasya","3495394592"))
	fmt.Println("Зарегистрировал аккаунт:",FacadeRegAccount("Kolya","1109978696"))
	fmt.Println("Вся информация аккаунта 0:")
	fmt.Println(FacadeAccInfo(0))
	fmt.Println("Вся информация аккаунта 1:")
	fmt.Println(FacadeAccInfo(1))
	fmt.Println(FacadeSendMoney(0,1,500))
	fmt.Println(FacadeSendMoney(1,0,250))
	fmt.Println("Вся информация аккаунта 0:")
	fmt.Println(FacadeAccInfo(0))
	fmt.Println("Вся информация аккаунта 1:")
	fmt.Println(FacadeAccInfo(1))
}

func Builder02(){
	name:="Вася Пупкин"
	text:="необходимо подготовить 2 тонны картофеля для инвестиций. " +
		"Если верить интернету, то всего через год его станет не менее 10 тонн. В итоге окупаемость составляет 400%!"
	fmt.Println("Оффициальный документ:")
	data,_:=json.MarshalIndent(BuildLetter(name,text,OfficialDocument)," ","\t")
	fmt.Println(string(data))
	fmt.Println("Неформальное письмо:")
	data,_=json.MarshalIndent(BuildLetter(name,text,InformalLetter)," ","\t")
	fmt.Println(string(data))
}

func Visitor03(){
	visitorCheck()
}

func Command04(){
	clock:=&Clock{}
	timeSetter:=&SetTimeCommand{clock}
	timeGetter:=&GetTimeCommand{clock}
	commandSender:=CommandSender{timeSetter}
	commandSender.Send()
	commandSender=CommandSender{timeGetter}
	commandSender.Send()
}

func ChainOfResp05(){
	orderIdHandler:=&orderIdHandler{}
	finalSumHandler:=&finalSumHandler{}
	finalSumHandler.setNext(orderIdHandler)
	priceHandler:=&priceHandler{}
	priceHandler.setNext(finalSumHandler)
	orderedGood:=NewGood("0001",5)
	priceHandler.execute(orderedGood)
}

func FactoryMethod06(){
	goodOne,err:=GetGood("101101")
	if err==nil{
		PrintDeviceInfo(goodOne)
	}else{
		log.Println(err)
	}
	goodTwo,err:=GetGood("501101")
	if err==nil{
		PrintDeviceInfo(goodTwo)
	}else{
		log.Println(err)
	}
	fmt.Printf("Товар 1 тип=%T, товар 2 тип=%T\n",goodOne,goodTwo)
}

func Strategy07(){

	alg1:=&algIter{}
	alg2:=&algMulti{}
	alg3:=&algSubtract{}
	subsequence:=InitSubs(alg1)
	for i:=0;i<10;i++{
		subsequence.Next()
		subsequence.SetAlgo(alg2)
		subsequence.Next()
		subsequence.SetAlgo(alg3)
		subsequence.Next()
		subsequence.SetAlgo(alg1)
	}
	fmt.Println("Вся последовательность:",subsequence.ReturnSubsequence())
}

func State08(){
	server:=InitServer()
	fmt.Println("Текущее состояние сервера:",server.currState.stateInfo())
	jobAdd:=&AddDataJob{
		10,
		"десять",
	}
	err:=server.DoJob(jobAdd)
	if err!=nil {
		fmt.Println("Не получилось выполнить задание:",err)
	}
	server.Start()
	fmt.Println("Текущее состояние сервера:",server.currState.stateInfo())
	err=server.DoJob(jobAdd)
	if err!=nil {
		fmt.Println("Не получилось выполнить задание:",err)
	}else{
		fmt.Println("Сервер доступен, задание выполняется...")
	}
	fmt.Println("Текущее состояние сервера:",server.currState.stateInfo())
	err=server.DoJob(jobAdd)
	if err!=nil {
		fmt.Println("Не получилось выполнить задание:",err)
	}
}
func main() {
	//Facade01()
	//Builder02()
	//Visitor03()
	//Command04()
	//ChainOfResp05()
	//FactoryMethod06()
	//Strategy07()
	State08()
}