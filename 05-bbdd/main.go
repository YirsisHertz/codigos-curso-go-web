package main

import (
	"bbdd/routes"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Print("Variables de entorno no cargadas")
	}
}

func main() {
	router := routes.Router()

	http.ListenAndServe(":3000", router)
}
