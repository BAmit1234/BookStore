package main

import (
	"fmt"
	"log"
	"main1/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(controller.Greet())

	router := mux.NewRouter()
	router.HandleFunc("/", controller.GetBooks).Methods("GET")
	router.HandleFunc("/{id}", controller.GetBook).Methods("GET")
	router.HandleFunc("/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/{id}", controller.DeleteBook).Methods("DELETE")
	router.HandleFunc("/{id}", controller.UpdateBook).Methods("PATCH")
	fmt.Println("running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
