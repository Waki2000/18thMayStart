package main

import "fmt"

//finish by 8am
//array
//slice
//map
//struct - create own data type or combine different datatype
// a struct is a collection of fields or properties.

type allBooks struct {
Title string		
Author string
		ISBN   string
		Price  float32
		Pages   int
	} 

func main() {
 
var book1 allBooks
		book1.Title = "Anintroduction into Golang"
		book1.Author = "Caleb Boxey"
		book1.ISBN = "976-09888"
		book1.Price = 10.50
		book1.Pages = 100
	
fmt.Println(book1)
		fmt.Println(book1.Title)
		fmt.Println(book1.Author)
		fmt.Println(book1.ISBN)
		fmt.Println(book1.Price)
		fmt.Println(book1.Pages)
	
}
