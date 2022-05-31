package main

import (
	"fmt"
)

type subsAlg interface {
	set(s *subsequence)
	getName()string
}

type algIter struct {}
func (a *algIter) set(s *subsequence){
	s.curNumb=s.curNumb+1
}
func (a *algIter) getName()string{
	return "Итерация +1"
}

type algMulti struct {}
func (a *algMulti) set(s *subsequence){
	s.curNumb=s.curNumb*2
}
func (a *algMulti) getName()string{
	return "Умножить на 2"
}

type algSubtract struct {}
func (a *algSubtract) set(s *subsequence){
	s.curNumb=s.curNumb-2
}
func (a *algSubtract) getName()string{
	return "Вычесть 2"
}

type subsequence struct {
	storage []int
	subsAlg subsAlg
	curNumb int
}

func InitSubs(startAlgo subsAlg)*subsequence{
	return &subsequence{
		storage: []int{1},
		subsAlg: startAlgo,
		curNumb: 1,
	}
}

func (s *subsequence)SetAlgo(a subsAlg){
	s.subsAlg=a
}

func (s *subsequence)setCur(){
	s.subsAlg.set(s)
}

func (s *subsequence)algoName()string{
	return s.subsAlg.getName()
}

func(s *subsequence)Next(){
	fmt.Printf("Выбранный алгоритм:%s\n",s.algoName())
	fmt.Printf("Прошлое число:%v ",s.curNumb)
	s.setCur()
	fmt.Printf("новое число:%v\n",s.curNumb)
	s.storage=append(s.storage,s.curNumb)
}

func(s *subsequence)ReturnSubsequence()[]int{
	return s.storage
}
