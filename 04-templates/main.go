package main

import (
	"fmt"
	"net/http"
	"vistas/routes"
)

func main() {

	router := routes.Router()

	fmt.Print("Server on port: 3000")
	http.ListenAndServe(":3000", router)
}
