package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello!",
	})
}
