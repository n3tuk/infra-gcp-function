package function

import (
	f "github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/gin-gonic/gin"
)

func init() {
	router := NewRouter()

	f.HTTP("server", router.Handler().ServeHTTP)
}

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(
		logger(),
		tracer(),
		gin.Recovery(),
	)

	r.NoRoute(noRoute)
	r.GET("/alive", alive)

	return r
}
