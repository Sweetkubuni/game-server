package api

import (

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		header := w.Header()
		header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})
	router.GET("/player", ListPlayer)
	router.POST("/player", GeneratePlayer)

	return router
}
