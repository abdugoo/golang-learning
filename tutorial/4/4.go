package main

import "fmt"


func checkScore2(a, b int){
	switch score := a + b;{
	case score >= 90:
		fmt.Println("A")
	case score >= 70:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}
}

func step3(){
	var a, b int
	fmt.Scan(&a, &b)
	checkScore2(a, b)

}