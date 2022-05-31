package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	gops "github.com/mitchellh/go-ps"
	"syscall"
)

type data struct {
	cd string
	cdItems []string
	pids []int
	delim string
}
var curData data

func rightPath(in string)string{
	var s string
	s=strings.ReplaceAll(in,"/",curData.delim)
	s=strings.ReplaceAll(in,"\\",curData.delim)
	if strings.HasSuffix(s, curData.delim){
		s=s[:len(s)-1]
	}
	return s
}

func showDir(){
	fmt.Println()
	items,err:=ioutil.ReadDir(curData.cd)
	pwd()
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%-50v%-15v%v\n", "Имя","Размер","Дата изменения")
	var files []string
	for _,i:=range items{
		files=append(files,i.Name())
		fmt.Printf("%-50v%-15v%v\n",i.Name(),i.Size(),i.ModTime())
	}
	curData.cdItems=files
}

func checkExists(dir string)bool{
	_,err:=os.Stat(dir)
	return !os.IsNotExist(err)
}

func dirContains(name string)bool{
	for _,f:=range curData.cdItems{
		if f==name{
			return true
		}
	}
	return false
}

func cd(dir string){
	if dir==".."{
		curData.cd=filepath.Dir(curData.cd)
		showDir()
		return
	}
	if dirContains(dir){
		cd(curData.cd+"\\"+dir)
	}
	if checkExists(dir){
		curData.cd=rightPath(dir)
	}else{
		fmt.Println("директория не существует")
	}
	showDir()
}

func pwd(){
	fmt.Println("Текущая директория:",curData.cd)
}

func echo(in string){
	fmt.Println(in)
}

func ps(){
	ps,_:=gops.Processes()
	fmt.Printf("%-10v%-10v%v\n", "Pid","PPid","Executable")
	var pids []int
	for _,p:=range ps{
		pids=append(pids,p.PPid())
		fmt.Printf("%-10v%-30v%v\n",p.Pid(),p.PPid(),p.Executable())
	}
	curData.pids=pids
}

func kill(pid string){
	procID,err:=strconv.Atoi(pid)
	if err!=nil{
		fmt.Println("некорректный PID, введено:",pid)
		return
	}
	proc,err:=os.FindProcess(procID)
	if err!=nil{
		fmt.Printf("процесса с PID=%v не существует\n",procID)
		return
	}
	err=proc.Kill()
	if err!=nil{
		fmt.Printf("невозможно завершить процесс:%v\n",err)
		return
	}
}

func fork(){
	bin,_:=os.Executable()
	args:=[]string{""}
	env:=os.Environ()
	env=append(env,"cd="+curData.cd)
	err:=syscall.Exec(bin,args,env)
	if err!=nil{
		log.Println(err)
	}
}

func Shell(){
	curData.cd,_=os.Getwd()
	if runtime.GOOS=="windows"{
		curData.delim="\\"
	}else{
		curData.delim="/"
	}
	showDir()
	scan:=bufio.NewScanner(os.Stdin)
	for scan.Scan(){
		line:=scan.Text()
		sep:=strings.Index(line," ")
		bash:=strings.Split(line," ")[0]
		var args string
		if sep>-1{
			args=line[sep+1:]
		}
		switch bash {
		case "cd":
			cd(args)
		case "echo":
			echo(args)
		case "pwd":
			pwd()
		case "ps":
			ps()
		case "kill":
			kill(args)
		case "fork":
			fork()
		}
	}
}

func main(){
	Shell()
}
