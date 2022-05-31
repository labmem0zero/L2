package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func fileStrings(file string)[]string{
	f,err:=os.Open(file)
	if err!=nil{
		log.Fatal(err)
	}
	var res []string
	fScan:=bufio.NewScanner(f)
	for fScan.Scan(){
		res=append(res,fScan.Text())
	}
	f.Close()
	return res
}


type flags struct{
	A int
	B int
	C int
	c bool
	i bool
	v bool
	F bool
	n bool
	e bool
	regular string
	fileName string
}

var Flags flags

var Input[]string

func extractFlags(){
	flag.IntVar(&Flags.A,"A", 0,"считать строк после совпадения")
	flag.IntVar(&Flags.B,"B", 0,"считать строк перед совпадением")
	flag.IntVar(&Flags.C,"C", 0,"считать строк до и столько же после совпадения")
	flag.BoolVar(&Flags.c,"c", false,"количество результатов")
	flag.BoolVar(&Flags.i,"i", false,"игнорировать регистр")
	flag.BoolVar(&Flags.v,"v", false,"исключать находки")
	flag.BoolVar(&Flags.F,"F", false,"точное совпадение со строкой, не паттерн")
	flag.BoolVar(&Flags.F,"e", false,"заэскейпить паттерн")
	flag.BoolVar(&Flags.n,"n", false,"печать номеров строки")
	flag.Parse()
	if flag.NArg()<2{
		log.Fatal("Отсутствует име файла либо паттерн для поиска!")
	}
	Flags.regular=flag.Args()[0]
	Flags.fileName=flag.Args()[1]
	fmt.Println(Flags)
}

//проверка наличия в срезе числа
func linesContains(lines []int,target int)bool{
	for _,v:=range lines{
		if v==target{
			return true
		}
	}
	return false
}

//инверсия. Берет глобальную переменную Input и выводит в результат все строки, номера которых отсутствуют в lines
//как пример: у нас 10 строк в инпуте. 2, 3 и 5 строки соответствуют паттерну. После выполнения этой функции в вывод
//попадут все строки инпута кроме 2, 3 и 5
func invertLines(lines []int)[]int{
	var res []int
	if len(lines)==0{
		for i:=0;i<len(Input);i++{
			res=append(res,i)
		}
		return res
	}
	for i:=0;i<len(Input);i++{
		if !linesContains(lines,i){
			res=append(res,i)
		}
	}
	return res
}

//промежуточная функция. Берет область вокруг каждой найденной строки и добавляет ее в вывод
func getLinesBAC(lines[]int)[]int{
	var res []int
	//цикл по каждому номеру строки из результатов поиска
	for _,l:=range lines{
		//проверка присутствия флагов
		//флаг C имеет больший приоритет, поэтому проверяется после B(before)
		before:=l-Flags.B
		if Flags.C>0{
			before=l-Flags.C
		}
		//если диапозон выходит за нижнюю границу среза, то диапозон корректируется
		if before<0{
			before=0
		}
		//флаг C имеет больший приоритет, поэтому проверяется после A(after)
		after:=l+Flags.A
		if Flags.C>0{
			after=l+Flags.C
		}
		//если диапозон выходит за верхнюю границу среза, то диапозон корректируется
		if after>len(Input)-1{
			after=len(Input)-1
		}
		//цикл по каждому элементу диапозона. Если строка уже есть в выводе, то она НЕ повторяется в этом выводе
		//например: найдено две строки, соответствующих паттерну. Они расположены на соседник индексах(к пример, 5 и 6)
		//Если С=2, то мы должны вывести 3,4,5,6,7 строки для первой находки и 4,5,6,7,8 строки для второй находки.
		//благодаря проверке, будут выведены 3,4,5,6,7,8 строки(без повторений)
		for j:=before;j<=after;j++{
			if !linesContains(res,j){
				res=append(res,j)
			}
		}
	}
	return res
}

//вывод строк
func printLines(lines []int){
	for _,v:=range lines{
		var line string
		//с флагом -n будет выводиться номер строки перед самой строкой
		if Flags.n{
			line=strconv.Itoa(v+1)+":\t"
		}
		line=line+Input[v]
		fmt.Println(line)
	}
}

//поиск по регулярному выражению
func doRegular(pattern string)([]string,[]int){
	var results []string
	var indxs []int
	reg,err:=regexp.Compile(pattern)
	if err!=nil{
		log.Fatal("Регулярное выражение составлено некорректно. Для отключение поиска по регулярным выражением, добавьте флаг -e")
	}
	for i,s:=range Input{
		//нижний регистр с флагом -i
		if Flags.i{
			s=strings.ToLower(s)
		}
		if reg.MatchString(s){
			results=append(results,Input[i])
			indxs=append(indxs,i)
		}
	}
	return results,indxs
}

//поиск по точному соответствию строки
func doFixed(pattern string)([]string,[]int){
	var results []string
	var indxs []int
	for i,s:=range Input{
		if Flags.i{
			s=strings.ToLower(s)
		}
		if s==pattern{
			results=append(results,Input[i])
			indxs=append(indxs,i)
		}
	}
	return results,indxs
}

//запуск поиска по паттерну
func doSearch(pattern string)([]string,[]int){
	//в зависимости от наличия флага -F, поиск по полному либо неполному совпадению строки
	if Flags.F{
		return doFixed(pattern)
	}
	return doRegular(pattern)
}

//фича. Возможность заэскейпить строку перед подачей в поиск по регулярному выражению
func escapeString(s string)string{
	//если стоит флаг -F, то у нас априори нет регулярного выражения, можно возвращать как есть
	if Flags.F{
		return s
	}
	escChars:=[]byte(".^$*+?()[]{}\\|")
	var bytes []byte
	//цикл по всем байтам паттерна. Если байт является символом для эскейпа, то
	//он эскейпится и отправляется в буффер
	for _,b:=range []byte(s){
		buffer:= func(b byte)[]byte {
			for _,item:=range escChars{
				if item==b{
					return []byte{92,b}
				}
			}
			return []byte{b}
		}(b)
		//после каждого символа, буффер добавляется к результирующему срезу
		bytes=append(bytes,buffer...)
	}
	//котороый затем конвертируется в строку
	return string(bytes)
}

func main(){
	//считываю флаги
	extractFlags()
	//включить режим Escape. Экранирует спецсимволы, используемые в регулярных выражениях
	if Flags.e{
		Flags.regular=escapeString(Flags.regular)
	}
	//игнорирование регистра
	if Flags.i{
		Flags.regular=strings.ToLower(Flags.regular)
	}
	Input=fileStrings(Flags.fileName)
	_,lines:=doSearch(Flags.regular)
	lines=getLinesBAC(lines)
	//инверсия результатов. Выполняется после флагов A, B и C. Поэтому запросто может "съесть" дополнительные строки
	//в комбинации с флагами A, B или C
	if Flags.v{
		lines=invertLines(lines)
	}
	//Вывод количества строк
	if Flags.c{
		fmt.Println(len(lines))
	}else{
		//либо самих строк
		printLines(lines)
	}
}