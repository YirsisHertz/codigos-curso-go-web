package routes

import (
	"net/http"
	"vistas/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	static := http.FileServer(http.Dir("public"))
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.Home)

	router.Handle("/static/", http.StripPrefix("/static/", static))

	return router
}
