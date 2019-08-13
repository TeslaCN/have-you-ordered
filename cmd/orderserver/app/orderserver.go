package app

import (
	"github.com/gin-gonic/gin"
	"have-you-ordered/cmd/orderserver/app/config"
	"have-you-ordered/internal/app/orderserver"
	"log"
	"time"
)

func init() {
}

func Run() {
	r := gin.Default()

	r.GET("/api", orderserver.ApiHelloGo)
	r.GET("/api/ordered/:date", orderserver.ApiOrdered)
	r.POST("/api/order", orderserver.PostOrder)
	r.GET("/api/dashboard/agg-by-day", orderserver.AggHistogram)

	r.POST("/api/control/fetch", orderserver.ManuallyFetch)

	duration, e := time.ParseDuration(config.Config.FetchInterval)
	if e != nil {
		log.Fatal(e)
	}
	go orderserver.FetchInterval(duration)
	_ = r.Run(config.Config.Server)
}
