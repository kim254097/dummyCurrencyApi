package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	handler "github.com/kim254097/dummyApiCurrency/handlers"
)

func main() {
	var port int

	fmt.Print("Enter port number for the server: ")
	fmt.Scan(&port)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler.HomeLink)
	router.HandleFunc("/currency", handler.CreateCurrency).Methods("POST")
	router.HandleFunc("/currencies", handler.GetAllCurrency).Methods("GET")
	router.HandleFunc("/currency/{id}", handler.GetOneCurrency).Methods("GET")
	router.HandleFunc("/currency/{id}", handler.UpdateCurrency).Methods("PATCH")
	router.HandleFunc("/currency/{id}", handler.DeleteCurrency).Methods("DELETE")

	fmt.Println("")
	fmt.Println("http://localhost:" + strconv.Itoa(port) + "/  =>  Home")
	fmt.Println("http://localhost:" + strconv.Itoa(port) + "/currency  =>  POST")
	fmt.Println("http://localhost:" + strconv.Itoa(port) + "/currencies  =>  GET")
	fmt.Println("http://localhost:" + strconv.Itoa(port) + "/currency/{id}  =>  GET")
	fmt.Println("http://localhost:" + strconv.Itoa(port) + "/currency/{id}  =>  PATCH")
	fmt.Println("http://localhost:" + strconv.Itoa(port) + "/currency/{id}  =>  DELETE")
	fmt.Println("")
	fmt.Println("* request body example => { 'ID': 'ARG', 'Name': 'Peso', 'Country': 'Argentina', }")
	fmt.Println("")
	fmt.Print("Server connected successfully...")
	panic(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
