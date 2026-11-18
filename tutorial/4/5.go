package main

import "fmt"

func step4(){
	defer fmt.Println("third")
	defer fmt.Println("second")
	fmt.Println("first")
}