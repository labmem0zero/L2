package timeModule

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func AccurateTime(){
	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err!=nil{
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
	fmt.Println("Текущее время:",time.Now())
	fmt.Println(" Beevik время:",t.String())
}