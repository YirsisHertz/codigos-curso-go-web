package main

import (
	"fmt"
	"net/http"
	"paquetes/rutas"
)

func main() {
	rutas.Rutas()

	fmt.Print("Server on port: 3000")
	http.ListenAndServe(":3000", nil)
}
