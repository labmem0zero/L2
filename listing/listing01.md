package main

import "fmt"

func main(){
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
}

Код выводит 77,78,79, потому что результирующий срез у нас является срезом a[1:4],
где 1 - начальный элемент, 4 - конечный невключителньый элемент