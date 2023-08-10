
package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)



func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/join", websocket.Handler(handleJoin))

	fmt.Println("Open http://localhost:8088 to access this demo")
	// nolint: gosec
	panic(http.ListenAndServe(":8088", nil))
}
