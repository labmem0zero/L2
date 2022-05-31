package control

import (
	"l2_http/service"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type result struct {
	Result interface{}
}

func ValidateDateTime(d, t string)bool{
	dt:=d+" "+t
	_,err:=time.Parse("02.01.2006 15:04:05",dt)
	if err!=nil {
		return false
	}
	return true
}

func processError(w http.ResponseWriter,r *http.Request){
	val,ok:=service.Contexts[r.Context()]["err"].(string);if !ok{
		return
	}
	data:=strings.Split(val,"-")
	errcode,_:=strconv.Atoi(data[0])
	retErr:=strings.Replace(`{"error":"{err}"}`,"{err}",val,-1)
	w.WriteHeader(errcode)
	w.Write([]byte(retErr))

}
