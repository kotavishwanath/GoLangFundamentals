package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit Hello world method")
	fmt.Fprintf(w, "Just checking the Go installation")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", HelloWorld).Methods("GET")
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/users/{name}/{emain}", NewUser).Methods("POST")
	myRouter.HandleFunc("/users/{name}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/users/{name}/{email}", UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

func main() {
	handleRequests()
	//**** 2 ******
	//https://www.youtube.com/watch?v=VAGodyl84OY

	//**** 1 ******
	//https://www.youtube.com/watch?v=YMQUQ6XQgz8
	//https://stackoverflow.com/questions/41539909/cannot-find-package-github-com-gorilla-mux-in-any-of/51175238

}
