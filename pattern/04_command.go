package main

import (
	"fmt"
	"time"
)

type Clock struct{
	currentTime time.Time
}

func (c *Clock) NowTime(){
	c.currentTime=time.Now()
	fmt.Println("Установил время:",c.currentTime)
}

func (c Clock) GetTime(){
	fmt.Println("Хранимое время:",c.currentTime)
}

type CommandSender struct {
	command command
}

func (c *CommandSender) Send(){
	c.command.execute()
}

type command interface {
	execute()
}

type SetTimeCommand struct {
	target target
}

func (c *SetTimeCommand) execute(){
	c.target.NowTime()
}

type GetTimeCommand struct {
	target target
}

func (c *GetTimeCommand) execute(){
	c.target.GetTime()
}

type target interface {
	NowTime()
	GetTime()
}