package main

import (
	"github.com/flashxgo/flashx"
	"github.com/gin-gonic/gin"
)

var engine *flashx.Engine

func main() {
	engine = &flashx.Engine{
		URLs:                  []string{"http://localhost:4000", "http://localhost:5000"},
		LoadBalancingStrategy: flashx.WeightedRoundRobin,
		RoundRobinWeights:     []int{1, 2},
	}
	engine.Setup()
	r := gin.Default()
	r.GET("/r/stories", redditEndpoint)
	r.Run(":2020")
}

func redditEndpoint(c *gin.Context) {
	engine.Initiate(c.Writer, c.Request)
}
