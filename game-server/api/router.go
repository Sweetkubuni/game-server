"github.com/julienschmidt/httprouter"

func newRouter() {
	router := httprouter.New()
	router.GET("/player", GetPlayer)
	router.POST("/player", PostPlayer)
}