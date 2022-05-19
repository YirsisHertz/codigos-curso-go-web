package routes

import (
	"muxpkg/controllers"

	"github.com/gorilla/mux"
)

func MuxRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/libro/{nombre_libro}", controllers.CrearLibro).Methods("POST")
	router.HandleFunc("/libro/{nombre_libro}", controllers.GetLibro).Methods("GET")

	return router
}
