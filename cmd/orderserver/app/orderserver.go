package app

import (
	"github.com/gin-gonic/gin"
	"have-you-ordered/cmd/orderserver/app/config"
	"have-you-ordered/internal/app/orderserver"
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

	_ = r.Run(config.Config.Server)
}
