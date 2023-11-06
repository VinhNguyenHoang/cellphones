package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRequest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "reply from server Jerry",
	})
}
