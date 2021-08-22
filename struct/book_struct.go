package main

import "fmt"

type Books struct {
	title    string
	author   string
	subtitle string
	book_id  int
}

func main() {
	var book1 Books
	// var book2 Books

	book1.title = "go language"
	book1.author = "runoob"
	book1.subtitle = "Go tutorial"
	book1.book_id = 1

	fmt.Println(book1)
	fmt.Println("The title of the book1 is: ", book1.title)
}
