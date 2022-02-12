package main

import (
	"fmt"
	"net/http"
)
	
func main() {
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Seja bem vindo ou bem vinda ao meu código!")
})

	http.ListenAndServe(":1337", nil) //DefaultServerMux
}