package main

import "fmt"

func step8(){
	langs := [3]string{"Go", "Rust", "Python"}
	langs[1] = "Zip"
	fmt.Println(langs)
}