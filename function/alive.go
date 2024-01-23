package function

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AliveResponse struct {
	Type  string `json:"type"`
	Alive bool   `json:"alive"`
	Date  string `json:"date,omitempty"`
}

// alive provides a simple response to verify that the endpoint is alive and
// service responses back to the client, including a date/time to validate the
// request is not cached
func alive(c *gin.Context) {
	a := AliveResponse{
		Type:  "alive",
		Alive: true,
		Date:  time.Now().Format(time.RFC3339),
	}

	c.Header("Content-Type", "application/json; charset=utf-8")
	c.JSON(http.StatusOK, a)
}
