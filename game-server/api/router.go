"github.com/julienschmidt/httprouter"

func newRouter() {
	router := httprouter.New()
	router.GET("/user", Index)
	router.POST("/user", Index)
}