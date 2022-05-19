package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	msg := "Metodo Invalido"

	if method == "GET" {
		msg = "Bienvenido"
	}

	fmt.Fprint(w, msg)
}

func main() {
	http.HandleFunc("/", index)

	fmt.Println("App Server in a port: 3000")
	http.ListenAndServe(":3000", nil)
}
