package handlers

import (
	"crud/models"
	"html/template"
	"log"
	"net/http"
)

func ReadUsers(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles("views/readUsers.html", "views/layout.html")

	if err != nil {
		log.Fatal("Read User View error")
	}

	usuarios := models.ReadUsers()

	view.Execute(w, usuarios)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")
		password := r.FormValue("password")

		models.CreateUser(nombre, correo, password)

		http.Redirect(w, r, "/users/", http.StatusFound)
	}

	view, err := template.ParseFiles("views/createUser.html", "views/layout.html")

	if err != nil {
		log.Fatal("Create User View error")
	}

	view.Execute(w, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	models.DeleteUser(id)

	http.Redirect(w, r, "/users/", http.StatusFound)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if r.Method == http.MethodPost {
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")
		password := r.FormValue("password")

		models.UpdateUser(id, nombre, correo, password)

		http.Redirect(w, r, "/users/", http.StatusFound)
	}

	view, err := template.ParseFiles("views/updateUser.html", "views/layout.html")

	if err != nil {
		log.Fatal("Create User View error")
	}

	usuario := models.ReadUser(id)

	view.Execute(w, usuario)
}
