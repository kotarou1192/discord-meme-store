package main

import (
	"discord-meme-store/src/web"
	"log"
)

func main() {
	err := web.Run()
	if err != nil {
		log.Panic(err)
		panic(err)
	}
}
