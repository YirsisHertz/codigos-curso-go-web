package controllers

import (
	"html/template"
	"net/http"
)

type Tarea struct {
	Nombre     string
	Completado bool
}

type Actividades struct {
	Usuario string
	Tareas  []Tarea
}

func Home(w http.ResponseWriter, r *http.Request) {
	vista, err := template.ParseFiles("views/home.html", "views/layout.html", "views/navbar.html")

	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	actividades := Actividades{"Yirsis", []Tarea{
		{"Landing Page", true},
		{"Laravel App", false},
		{"Vue App", true},
		{"React App", true},
		{"Go API", false},
	}}

	vista.Execute(w, actividades)
}
