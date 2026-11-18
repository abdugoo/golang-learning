package main
import "fmt"

func calc(x, y int)(sum, dif int){
	sum = x + y
	dif = x - y
	return
}

func main(){
	fmt.Println(calc(30, 15))
}