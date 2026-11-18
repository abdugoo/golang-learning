package main
import "fmt"

func checkScore(a, b int) string{
	if sum := a + b; sum >= 100{
		return "Win"
	}else{
		return "Less points"
	}
}

func step2(){
	fmt.Println(checkScore(45, 35))
}