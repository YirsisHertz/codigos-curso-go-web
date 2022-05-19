package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetLibro(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	nombre_libro := params["nombre_libro"]

	fmt.Fprint(w, nombre_libro)
}

func CrearLibro(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	nombre_libro := params["nombre_libro"]

	fmt.Fprint(w, " CREAR: "+nombre_libro)
}
