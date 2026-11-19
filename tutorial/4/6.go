package main

import "fmt"

func step5(){
	score := 100
	ptr := &score
	*ptr = 500
	fmt.Println(score)
}