package main

import (
	"goweb/config"
	"goweb/controllers/categories"
	"goweb/controllers/home"
	"log"
	"net/http"
)

func main() {
	config.Database()

	//home
	http.HandleFunc("/", home.Welcome)

	//products
	http.HandleFunc("/products", home.Welcome)

	//categories
	http.HandleFunc("/categories", categories.Index)
	http.HandleFunc("/categories/add", categories.Add)
	http.HandleFunc("/categories/edit", categories.Edit)
	http.HandleFunc("/categories/delete", categories.Delete)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
	log.Println("Listening on localhost:8080")
}
