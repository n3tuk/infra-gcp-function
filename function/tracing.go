package function

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func tracer() gin.HandlerFunc {
	// Create a static UUID value which will persist over the runtime of the
	// server, providing a persistent backend value for debugging
	sID := uuid.New().String()

	return func(c *gin.Context) {
		// Attach the static UUID and a randomly generated UUID value for each
		// request to the response
		c.Header("X-Server-ID", sID)
		c.Header("X-Request-ID", uuid.New().String())
		c.Next()
	}
}
