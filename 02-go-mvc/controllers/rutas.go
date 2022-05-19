package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Controller")
}

func Productos(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")

	fmt.Fprint(w, "Bienvenido: "+name+" el id de usuario es: "+id)
}
