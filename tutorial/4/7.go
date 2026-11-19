package main

import "fmt"

type Book struct {
	Title string
	Pages int
}

func step6(){
	book := Book {Title: "All", Pages: 150}
	book.Pages = 250
	fmt.Println(book.Title, book.Pages)
}