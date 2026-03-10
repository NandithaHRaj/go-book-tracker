package main

import (
	"net/http"
	"log"
	"go-book-tracker/handler"
)

func main(){
	http.HandleFunc("/books", handler.BookHandler)
	http.HandleFunc("/books/", handler.BookHandlerByID)

	log.Println("Server running on 8080")

	err := http.ListenAndServe(":8080",nil)
	
	if err != nil{
		log.Fatal(err)
	}

}