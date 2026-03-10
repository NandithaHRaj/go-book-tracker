package main

import (
	"log"
	"go-book-tracker/router"
	"net/http"
)

func main(){

	r := router.SetupRouter()

	log.Println("Server running on 8080")

	err := http.ListenAndServe(":8080", r)
	
	if err != nil{
		log.Fatal(err)
	}

}