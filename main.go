package main

import (
	"log"
	"go-book-tracker/router"
	"go-book-tracker/handler"
	"go-book-tracker/service"
	"go-book-tracker/repository"

	"net/http"
)

func main(){

	repo := &repository.MemoryRepository{}

	service := service.NewBookService(repo)

	handler := handler.NewBookHandler(service)

	r := router.SetupRouter(handler)

	log.Println("Server running on 8080")

	err := http.ListenAndServe(":8080", r)
	
	if err != nil{
		log.Fatal(err)
	}

}