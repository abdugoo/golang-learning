package main

import "fmt"

func step7(){
	book1 := Book {Title: "Friends", Pages: 250}
	p := book1
	p.Pages = 220
	fmt.Println(p.Title, p.Pages)

}