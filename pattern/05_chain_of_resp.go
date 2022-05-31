package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"
)

var goodsPrices =map[string]float64{"0001":15.99}

type orderedGood struct{
	GoodId string
	GoodPrice float64
	GoodCount float64
	FinalSum float64
	OrderId string
}

func NewGood(id string, count float64)*orderedGood{
	good:=new(orderedGood)
	good.GoodId=id
	good.GoodCount=count
	return good
}

type orderHandler interface {
	execute(good *orderedGood)
	setNext(handler orderHandler)
}

type priceHandler struct {
	next orderHandler
}

func (h *priceHandler) execute(g *orderedGood){
	if g.GoodPrice!=0{
		fmt.Println("Цена на товар уже извлечена")
		h.next.execute(g)
		return
	}
	fmt.Println("Цена на товар извлечена")
	g.GoodPrice=goodsPrices[g.GoodId]
	h.next.execute(g)
}

func (h *priceHandler) setNext(next orderHandler){
	h.next=next
}

type finalSumHandler struct {
	next orderHandler
}

func (h *finalSumHandler) execute(g *orderedGood){
	if g.FinalSum!=0{
		fmt.Println("Сумма заказа уже расчитана")
		h.next.execute(g)
		return
	}
	if (g.GoodPrice==0)||(g.GoodCount==0){
		fmt.Printf("Ошибка! Цена товара=%.2f, количество товара для заказа=%.2f!\n")
		return
	}
	fmt.Println("Сумма заказа расчитана")
	g.FinalSum=g.GoodPrice*g.GoodCount
	h.next.execute(g)
}

func (h *finalSumHandler) setNext(next orderHandler){
	h.next=next
}

type orderIdHandler struct {
	next orderHandler
}

func (h *orderIdHandler) execute(g *orderedGood){
	if g.OrderId!=""{
		fmt.Println("ID заказа уже установлен")
		fmt.Println("Финальный заказ:")
		res,_:=json.Marshal(g)
		fmt.Println(string(res))
		return
	}
	rgxp:=regexp.MustCompile("[A-z]|[[:punct:]]|[\x20]")
	tmp:=g.GoodId+fmt.Sprintf("%.2f",g.FinalSum)+time.Now().String()
	g.OrderId=rgxp.ReplaceAllString(tmp,"")
	fmt.Println("ID заказа установлен")
	fmt.Println("Финальный заказ:")
	res,_:=json.Marshal(g)
	fmt.Println(string(res))
}

func (h *orderIdHandler) setNext(next orderHandler){
	h.next=next
}

