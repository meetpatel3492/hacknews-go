package main

import (
	"fmt"
	"github.com/enescakir/emoji"
	"github.com/gin-gonic/gin"
	"github.com/meetpatel3492/go-first-module/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/", controllers.HandleMain)
	r.GET("/new/topstories", controllers.TopStories)
	r.GET("/new/topstories/concurrent", controllers.TopStoriesConcurrent)
	r.Run()
	fmt.Println("u1F602")
	fmt.Println(emoji.WavingHand)
}
