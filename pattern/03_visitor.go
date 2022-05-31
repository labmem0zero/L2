package main

import (
	"fmt"
	"main/03_visitor_module"
)

func visitorCheck(){
	lamp:=&visitor_module.SmartLamp{}
	teapot:=&visitor_module.SmartTeapot{}
	powerCalc:=&visitor_module.PowerCalculator{}
	lamp.SetParams(1,1,1)
	lamp.Accept(powerCalc)
	fmt.Println(lamp.GetType()," - мощность -",powerCalc.GetPower())
	teapot.SetParams(1,1,1,1)
	teapot.Accept(powerCalc)
	fmt.Println(teapot.GetType()," - мощность -",powerCalc.GetPower())
}


