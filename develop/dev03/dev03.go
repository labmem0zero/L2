package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadFile(fileName string)[]string{
	var str []string
	f,err:=os.Open(fileName)
	if err!=nil{
		log.Fatal(err)
	}
	scanner:=bufio.NewScanner(f)
	for scanner.Scan(){
		str=append(str,scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	f.Close()
	return str
}

func WriteFile(fileName string, in []string){
	f,err:=os.OpenFile(fileName,os.O_TRUNC|os.O_RDWR,0664)
	if err!=nil{
		log.Fatal("writerr:",err)
	}
	for _,s:=range in{
		f.WriteString(s+"\n")
	}
	f.Close()
}

//удаление повторяющихся элементов
func deleteDupes(in []string)[]string{
	var res []string
	tmp:=make(map[string]struct{})
	for _,s:=range in{
		if _,ok:=tmp[s];ok==false{
			tmp[s]= struct{}{}
			res=append(res,s)
		}
	}
	return res
}

//переворачивание среза. Нужно для ключа 'c'
func reverseSlice(in []string)[]string{
	if len(in)<2{
		return in
	}
	for i:=0;i<len(in)/2;i++{
		in[i],in[len(in)-1-i]=in[len(in)-1-i],in[i]
	}
	return in
}

//функция проверки существования ключа. При отсутствии ключа возвращает -1, false
func keysContains(keys []string,key string)(int,bool){
	for i,k:=range keys{
		if k==key{
			return i,true
		}
	}
	return -1,false
}

//сортировка по возрастанию чисел. Параметры
//in - строки, отвечающие параметрам сортировки
//left - 0
//right - len(in)-1
//col - срез строк, относительно которых происходит сравнение(нужно для сортировки по колонкам),
//если не нужно, то передается in
func sortQuickNum(in []string,left int, right int, col []string)[]string{
	l:=left
	r:=right
	center,_:=strconv.Atoi(col[(left+right)/2])
	for l<=r{
		colL,_:=strconv.Atoi(col[l])
		for colL<center{
			l++
			colL,_=strconv.Atoi(col[l])
		}
		colR,_:=strconv.Atoi(col[r])
		for colR>center{
			r--
			colR,_=strconv.Atoi(col[r])
		}
		if l<=r{
			in[r],in[l]=in[l],in[r]
			col[l],col[r]=col[r],col[l]
			l++
			r--
		}
	}
	if r>left{
		sortQuickNum(in,left,r,col)
	}
	if l<right{
		sortQuickNum(in,l,right,col)
	}
	return in
}

//сравнение двух строк. Используется при совпадении символов

func compareTwo(in,col []string,l,r,idx int){
	if (len(col[l])<idx+1)||(len(col[r])<idx+1){
		return
	}
	if col[l][idx]==col[r][idx]{
		compareTwo(in,col,l,r,idx+1)
	}
	if col[l][idx]>col[r][idx]{
		col[l],col[r]=col[r],col[l]
		in[l],in[r]=in[r],in[l]
		return
	}
}

func sortQuickLex(in []string,left int, right int, col []string)[]string{
	l:=left
	r:=right
	center:=col[(left+right)/2][0]
	for l<=r{
		for col[l][0]<center{
			l++
		}
		for col[r][0]>center{
			r--
		}
		if l<=r{
			in[r],in[l]=in[l],in[r]
			col[l],col[r]=col[r],col[l]
			if col[r][0]==col[l][0]{
				compareTwo(in,col,l,r,0)
			}
			l++
			r--
		}
	}
	if r>left{
		sortQuickLex(in,left,r,col)
	}
	if l<right{
		sortQuickLex(in,l,right,col)
	}
	return in
}

//промежуточная стадия сортировки. Считывает ключи и подготавливает строки,
//которые подходят дял сравнения. Так же подготавливает col, в котором хранится срез с выбранной колонкой
func prepareSort(in []string, key string)[]string{
	var col []string
	var prepared []string
	var wrong []string
	keys:=strings.Split(key, " ")
	if n,ok:=keysContains(keys,"-k");ok==true{
		//проверяем, есть ли параметр после ключа 'k'. Нужно, что бы не паниковать при IndexOutOfRange
		if n+1>len(keys)-1{
			return append([]string{"Ошибка ввода, недостаточно параметров"},in...)
		}
		var tmpIn []string
		var tmpCol []string
		//считываем параметр с номером столбца, проверяем, что бы он был числом
		k,err:=strconv.Atoi(keys[n+1])
		if err!=nil{
			return append([]string{"Ошибка ввода, неверный параметр для колонки"},in...)
		}
		for _,s:=range in {
			//разбиваем строки на столбцы. Если столбец существует, то базовую строку отправляем в основной срез,
			//а подстроку(столбец) отправляем в дополнительный срез, по которому будет происходить сортировка.
			//если в строке нет такого столбца, то она добавляется в срез wrong для последующего добавления в результат
			tmpStrings:=strings.Split(s," ")
			if len(tmpStrings)>=k{
				tmpIn=append(tmpIn,s)
				tmpCol=append(tmpCol,tmpStrings[k-1])
			}else{
				wrong=append(wrong,s)
			}
		}
		//если вышло 0 строк, то возвращаем результат
		if len(tmpIn)==0{
			return append([]string{"Такого столбца не существует"},in...)
		}
		prepared=tmpIn
		col=tmpCol
	}else{
		prepared=in
		col=in
	}
	//подготовка срезов для сортировки по числам
	if _,ok:=keysContains(keys,"-n");ok==true {
		fmt.Println("есть ключ -n")
		reg := regexp.MustCompile("^[0-9]+")
		var tmpIn []string
		var tmpCol []string
		//ищем строки, начинающиеся с чисел в срезе с колонкой, который мы заранее подготовили
		//опять же, строки, начинающиеся не с числа запомниаются, для последующего дописывания
		for i, s:=range col {
			numSubs:=reg.FindString(s)
			if numSubs!= "" {
				tmpIn=append(tmpIn, prepared[i])
				tmpCol=append(tmpCol, numSubs)
			} else {
				wrong=append(wrong, prepared[i])
			}
		}
		if len(tmpIn)==0{
			return append([]string{"подходящих по параметрам строк не найдено"},in...)
		}
		prepared=tmpIn
		col=tmpCol
		in=sortQuickNum(prepared,0,len(prepared)-1,col)
		if _,ok:=keysContains(keys,"c");ok==true{
			in=reverseSlice(in)
		}
		in=append(in,wrong...)
		if _,ok:=keysContains(keys,"u");ok==true{
			in=deleteDupes(in)
		}
		return in
	}
	in=sortQuickLex(prepared,0,len(prepared)-1,col)
	if _,ok:=keysContains(keys,"c");ok==true{
		in=reverseSlice(in)
	}
	in=append(in,wrong...)
	if _,ok:=keysContains(keys,"u");ok==true{
		in=deleteDupes(in)
	}
	return in
}

func SortStrings(in []string,key string)[]string{
	return prepareSort(in,key)
}

func PrintArray(in []string){
	fmt.Println("-----------------------")
	for _,s:=range in{
		fmt.Println(s)
	}
	fmt.Println("-----------------------")
}

func main()  {
	testStrings:=ReadFile("dev03test.txt")
	res:=SortStrings(testStrings,"-k 2 -n")
	for _,s:=range res{
		fmt.Println(s)
	}
	WriteFile("dev03test.txt",res)
}