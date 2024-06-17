package spike

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func ExportEndpoints() map[string]func(c *gin.Context) {
	return map[string]func(c *gin.Context){
		"/world": HandleRequest,
	}
}

func HandleRequest(c *gin.Context) {
	header := c.Request.Header
	log.Println("Header:", header)

	reqCtx := c.Request.Context()
	span := trace.SpanFromContext(otel.GetTextMapPropagator().Extract(reqCtx, propagation.HeaderCarrier(c.Request.Header)))
	defer span.End()

	c.JSON(http.StatusOK, gin.H{
		"message": "reply from server Spike",
	})
}