package main

import (
	"errors"
	"fmt"
)

type Server struct {
	busyState state
	idleState state
	offState state

	currState state

	jobs []IJob

	data map[int]string
}

func InitServer()*Server{
	data:=make(map[int]string)
	s:=&Server{data: data}
	busyState:=busyState{Server: s}
	idleState:=idleState{Server: s}
	offState:=offState{Server: s}
	s.busyState=busyState
	s.idleState=idleState
	s.offState=offState
	s.currState=s.offState
	return s
}

func (s *Server)Start(){
	s.currState=s.idleState
}

func(s *Server)DoJob(job IJob)error{
	err:=s.currState.doJob(job)
	return err
}

func (s *Server)appendJob(job IJob){
	s.jobs=append(s.jobs,job)
}

func (s *Server)getJob()(IJob,error){
	if len(s.jobs)==0{
		return nil,errors.New("на сервере нет заданий")
	}
	job:=s.jobs[len(s.jobs)-1]
	return job,nil
}

type state interface {
	doJob(job IJob)error
	leftJobs()(int,error)
	stateInfo()string
}

type busyState struct {
	Server *Server
}
func (s busyState)doJob(job IJob)error{
	s.Server.jobs=append(s.Server.jobs,job)
	return errors.New("сервер занят, задание добавлено в очередь")
}
func (s busyState)leftJobs()(int,error){
	jobs:=len(s.Server.jobs)
	return jobs,nil
}
func (s busyState)stateInfo()string{
	return "busyState"
}

type idleState struct {
	Server *Server
}
func (s idleState)doJob(job IJob)error{
	err:=job.doJob(s.Server)
	if err!=nil{
		return err
	}
	s.Server.currState=s.Server.busyState
	return nil
}
func (s idleState)leftJobs()(int,error){
	jobs:=len(s.Server.jobs)
	return jobs,nil
}
func (s idleState)stateInfo()string{
	return "idleState"
}

type offState struct {
	Server *Server
}
func (s offState)doJob(job IJob)error{
	return errors.New("сервер выключен, выполнение заданий невозможно")
}
func (s offState)leftJobs()(int,error){
	return 0,errors.New("сервер выключен, инфомрация о заданиях недоступна")
}
func (s offState)stateInfo()string{
	return "offState"
}

type IJob interface {
	doJob(s *Server)error
}

type DeleteDataJob struct {
	inputData int
}

func (j *DeleteDataJob) doJob(s *Server)error{
	_,ok:=s.data[j.inputData]
	if ok==false{
		return errors.New("невозможно удалить данные карты, такого ключа не существует")
	}
	delete(s.data,j.inputData)
	fmt.Println("Ключ удален")
	return nil
}

type AddDataJob struct {
	inputData1 int
	inputData2 string
}

func (j *AddDataJob) doJob(s *Server)error{
	_,ok:=s.data[j.inputData1]
	if ok==true{
		return errors.New("невозможно добавить данные в карту, ключ уже занят")
	}
	s.data[j.inputData1]=j.inputData2
	fmt.Println("----имитация потока заданий----")
	return nil
}

type ReplaceDataJob struct {
	inputData1 int
	inputData2 string
}

func (j *ReplaceDataJob) doJob(s *Server)error{
	_,ok:=s.data[j.inputData1]
	if ok==false{
		return errors.New("невозможно изменить данные карты, такого ключа не существует")
	}
	s.data[j.inputData1]=j.inputData2
	fmt.Println("Значение ключа изменено")
	return nil
}
