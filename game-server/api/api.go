package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetPlayer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func PostPlayer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
