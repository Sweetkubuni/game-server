package main

import (
	"fmt"
	"net/http"

	"game-server/api"
)

func main() {
	router := api.NewRouter()

	fmt.Println("Open index.html to access this demo")
	// nolint: gosec
	panic(http.ListenAndServe(":8080", router))
}
