package web

import (
	"discord-meme-store/src/web/api/home"

	"github.com/gin-gonic/gin"
)

const API_PREFIX string = "/api"

func Run() error {
	eg := gin.Default()

	eg.GET(API_PREFIX+"/home", home.Show)

	err := eg.Run()
	return err
}
