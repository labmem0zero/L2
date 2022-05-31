package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
)

var goods=map[string]string{
	"101101":"{\n\t\t\"name\":\"Умная розетка\",\n\t\t\"price\":300,\n\t\t\"power\":3.14,\n\t\t\"description\":\"Самая умная розетка, обладает красивой выходной мощностью числа Пи, что равняется 3.14Вт! Подходит для всех маломощных устройств.\"\n\t}",
	"501101":"{\n\t\t\"name\":\"Умная лампа\",\n\t\t\"price\":400.0,\n\t\t\"power\":60.0,\n\t\t\"description\":\"Отличная лампочка мощностью 60Вт\"\n\t}",
}

var stocks=map[string]int{
	"101101":50,
	"501101":20,
}

type importedGood struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Power       float64 `json:"power"`
	Description string  `json:"description"`
}

func getGoodFromServer(goodID string)importedGood{
	var good importedGood
	err:=json.Unmarshal([]byte(goods[goodID]),&good)
	if err!=nil{
		log.Println(err)
	}
	return good
}

func getStockFromServer(goodID string)int{
	return stocks[goodID]
}

type ISmartDevice interface {
	setInfo(x importedGood)
	setSupply()
	setType(goodType string)
	setID(id string)
	getDeviceInfo()string
}

type SmartDevice struct {
	goodID,name,description,goodType string
	power float64
	stock int
	price float64
}

func (s *SmartDevice) setInfo(x importedGood){
	s.name=x.Name
	s.power=x.Power
	s.price=x.Price
	s.description=x.Description
}

func (s *SmartDevice) setID(id string){
	s.goodID=id
}

func (s *SmartDevice) setSupply(){
	s.stock=getStockFromServer(s.goodID)
}

func (s *SmartDevice) setType(goodType string){
	s.goodType=goodType
}

func (s *SmartDevice)getDeviceInfo()string{
	info:=fmt.Sprintf("Название устройства: %s\nТип устройства: %s\nМощность устройства: %.2f\nОписание устройства: %s\nЦена устройства: %.2f\nОстаток устройств на складе: %v\nКод устройства для заказа: %s\n",s.name,s.goodType,s.power,s.description,s.price,s.stock,s.goodID)
	return info
}

func PrintDeviceInfo(i ISmartDevice){
	fmt.Println(i.getDeviceInfo())
}

type smartLamp struct{
	SmartDevice
}

func newSmartLamp() ISmartDevice{
	return &smartLamp{
		SmartDevice{goodType: "Лампа"},
	}
}

type smartSwitch struct{
	SmartDevice
}

func newSmartSwitch() ISmartDevice{
	return &smartSwitch{
		SmartDevice{goodType: "Лампа"},
	}
}

func GetGood(goodID string)(ISmartDevice,error){
	good:=getGoodFromServer(goodID)
	if strings.Contains(good.Name,"лампа"){
		device:=newSmartLamp()
		device.setID(goodID)
		device.setType("Лампочка")
		device.setInfo(good)
		device.setSupply()
		return device,nil
	}
	if strings.Contains(good.Name,"розетка"){
		device:=newSmartSwitch()
		device.setID(goodID)
		device.setType("Розетка")
		device.setInfo(good)
		device.setSupply()
		return device,nil
	}
	return nil,errors.New("необработанная ошибка")
}





