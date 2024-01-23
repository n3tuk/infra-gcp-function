package function

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func logger() gin.HandlerFunc {
	logger := &log.Logger

	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		var logEvent *zerolog.Event
		if c.Writer.Status() >= http.StatusInternalServerError {
			logEvent = logger.Error()
		} else {
			logEvent = logger.Info()
		}

		logEvent.
			Str("id", c.Writer.Header().Get("X-Request-ID")).
			Str("client_ip", c.ClientIP()).
			Str("host", c.Request.Host).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status_code", c.Writer.Status()).
			Int("body_size", c.Writer.Size()).
			Float64("latency", time.Since(start).Seconds()).
			Str("user_agent", c.Request.UserAgent()).
			Msg(c.Errors.ByType(gin.ErrorTypePrivate).String())
	}
}
