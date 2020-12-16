package main

import (
	"net/url"

	"github.com/flashxgo/flashx"
	"github.com/gin-gonic/gin"
)

var engine *flashx.Engine

func main() {
	engine = &flashx.Engine{}
	engine.Setup()
	r := gin.Default()
	r.GET("/hn/stories", hnEndpoint)
	r.GET("/r/stories", redditEndpoint)
	r.Run(":2020")
}

func hnEndpoint(c *gin.Context) {
	url, _ := url.Parse("http://localhost:4000")
	engine.Initiate(url, c.Writer, c.Request)
}

func redditEndpoint(c *gin.Context) {
	url, _ := url.Parse("http://localhost:4000")
	engine.Initiate(url, c.Writer, c.Request)
}
