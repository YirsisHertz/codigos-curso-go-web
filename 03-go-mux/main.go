package main

import (
	"fmt"
	"muxpkg/routes"
	"net/http"
)

func main() {

	router := routes.MuxRoutes()

	fmt.Println("Server on port: 3000")
	http.ListenAndServe(":3000", router)
}
