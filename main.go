package main

import (
	"letswork/handlers"
	"log"
	"net/http"
)

func main() {
	log.Println("Lets go!")
	http.Handle("/bumble/", http.StripPrefix("/bumble/", http.FileServer(http.Dir("./bumble"))))
	http.HandleFunc("/", handlers.RootHandler)
	log.Println(http.ListenAndServe(":9797", nil))
}
