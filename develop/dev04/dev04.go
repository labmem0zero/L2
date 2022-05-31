package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func findDupes(s string, in[]string)bool{
	for _,v:=range in{
		if v==s{
			return true
		}
	}
	return false
}

func FindAnagrams(in *[]string)*map[string][]string{
	anagramIdx:=make(map[string][]int)
	for i,w:=range *in{
		word:=[]rune(w)
		sort.Slice(word, func(i, j int) bool {
			return word[i]<word[j]
		})
		tmpW:=string(word)
		anagramIdx[tmpW]=append(anagramIdx[tmpW],i)
	}
	res:=make(map[string][]string)
	var tmp []string
	for _,v:=range anagramIdx{
		if len(v)<2{
			continue
		}
		tmp=[]string{}
		for _,j:=range v{
			if findDupes((*in)[j],tmp)==false{
				tmp=append(tmp,(*in)[j])
			}
		}
		zero:=tmp[0]
		sort.Strings(tmp)
		res[zero]=tmp
	}
	return &res
}

func randSlice()[]string{
	rand.Seed(time.Now().UnixNano())
	var res []string
	for i:=0;i<10;i++{
		x:=3+rand.Intn(1)
		var word []rune
		for j:=0;j<x;j++{
			word=append(word,rune('а' + rand.Intn(2)))
		}
		res=append(res,string(word))
		word=[]rune{}
	}
	fmt.Println("Входные данные:", res)
	return res
}

func main(){
	test:=randSlice()
	fmt.Println(FindAnagrams(&test))
}
