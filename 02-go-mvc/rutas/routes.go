package rutas

import (
	"net/http"
	"paquetes/controllers"
)

func Rutas() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/productos", controllers.Productos)
}
