package model

type Book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Status string `json:"status"`
	Rating float64 `json:"rating"`
}