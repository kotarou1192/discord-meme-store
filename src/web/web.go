package web

import (
	"discord-meme-store/src/web/api/home"

	"github.com/gin-gonic/gin"
)

type Engine struct {
	webEngine *gin.Engine
}

func Run() error {
	eg := gin.Default()

	eg.GET("/home", home.Show)

	err := eg.Run()
	return err
}
