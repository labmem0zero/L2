package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"syscall"
	"time"
)

func findArg(args []string, s string)string{
	reg:=regexp.MustCompile(s)
	for _,v:=range args{
		if reg.MatchString(v){
			return v
		}
	}
	return ""
}

type tcpDialer struct {
	timeout time.Duration
	host string
	port string
	conn net.Conn
	err error
	wg *sync.WaitGroup
	stopChan []chan struct{}
}
var myTelnet tcpDialer

func (t *tcpDialer)Listen(){
	t.wg.Add(1)
	fmt.Println("Слушаем ответы хоста:")
	for {
		message,err:=bufio.NewReader(t.conn).ReadString('\n')
		if err!=nil{
			fmt.Println("Хост завершил работу")
			break
		}
		fmt.Printf("%s: %s",t.host, message)
	}
	t.wg.Done()
}

func (t *tcpDialer) Write(){
	t.wg.Add(1)
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		_,err:=fmt.Fprintf(t.conn, text+"\n")
		if err != nil {
			t.Stop(1)
			break
		}
	}
	t.wg.Done()
}

func (t *tcpDialer)Stop(sig int){
	t.conn.Close()
	t.wg.Add(-1)
	os.Exit(sig)
}

func (t *tcpDialer) Dial()error{
	conn,err:=net.DialTimeout("tcp",t.host+":"+t.port,t.timeout)
	if err!=nil{
		return err
	}
	t.conn=conn
	return nil
}

func calcTimeout(arg string)time.Duration{
	if !strings.Contains(arg,"="){
		res,_:=time.ParseDuration("10s")
		return res
	}
	dur:=strings.Split(arg,"=")[1]
	res,err:=time.ParseDuration(dur)
	if err!=nil{
		log.Fatal("Некорректно указан таймаут, параметр при запуске:",dur)
	}
	return res
}

type launchArgs struct {
	timeout time.Duration
	host string
	port string
}

func MakeArgs()launchArgs{
	var la launchArgs
	args:=os.Args[1:]
	la.timeout=calcTimeout(findArg(args,`--timeout=`))
	la.host=args[len(args)-2]
	la.port=args[len(args)-1]
	return la
}

func prepTcpDialer(la launchArgs)*tcpDialer{
	tcpDialer:=&tcpDialer{
		timeout: la.timeout,
		host: la.host,
		port: la.port,
		wg: new(sync.WaitGroup),
	}
	return tcpDialer
}

func main(){
	params:=MakeArgs()
	myTelnet=*prepTcpDialer(params)
	err:=myTelnet.Dial()
	go myTelnet.Listen()
	go myTelnet.Write()
	if err!=nil{
		log.Fatal(err)
	}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGINT,
		syscall.SIGHUP,
	)
	go func(){
		<-sc
		fmt.Println("Программа завершена пользователем")
		myTelnet.Stop(0)
	}()
	myTelnet.wg.Wait()
}
