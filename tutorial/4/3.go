package main
import "fmt"

func step1(){
	for i := 5; i > 0; i--{
		fmt.Println(i)
	}
	i := 5
	for i > 0{
		fmt.Println(i)
		i--
	}
}