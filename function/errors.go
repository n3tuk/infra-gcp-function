package function

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Type   string `json:"type"`
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
}

func noRoute(c *gin.Context) {
	response := ErrorResponse{
		Type:   "error",
		Status: http.StatusNotFound,
		Error:  "File Not Found",
	}

	c.Header("Content-Type", "application/json; charset=utf-8")
	c.JSON(http.StatusNotFound, response)
}
