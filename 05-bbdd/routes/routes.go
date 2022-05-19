package routes

import (
	"bbdd/db"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Producto struct {
	id     int
	nombre string
}

func Router() *mux.Router {

	db := db.Conexion()

	router := mux.NewRouter()

	router.HandleFunc("/productos", func(w http.ResponseWriter, r *http.Request) {

		query := "SELECT id, nombre FROM productos"

		rows, err := db.Query(query)

		if err != nil {
			fmt.Print("Registros no creados")
		} else {
			var productos []Producto

			for rows.Next() {
				var producto Producto

				rows.Scan(&producto.id, &producto.nombre)
				productos = append(productos, producto)

				fmt.Fprintln(w, productos)
			}

			rows.Err()
		}
	})

	return router
}
