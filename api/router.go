package api

import "github.com/julienschmidt/httprouter"

func newRouter() {
	router := httprouter.New()
	router.GET("/player", ListPlayer)
	router.POST("/player", GeneratePlayer)
}
