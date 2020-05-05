package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"have-you-ordered/cmd/orderserver/app/config"
	"have-you-ordered/internal/app/orderserver"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
}

func Run() {
	fmt.Printf("os.Args: %s\n", os.Args)
	r := gin.Default()

	r.GET("/api", orderserver.ApiHelloGo)
	r.GET("/api/ordered/:date", orderserver.ApiOrdered)
	r.GET("/api/order-types", orderserver.ApiOrderTypes)
	r.POST("/api/order", orderserver.PostOrder)
	r.POST("/api/fake-order", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	r.GET("/api/dashboard/agg-by-day", orderserver.AggHistogram)

	r.POST("/api/control/fetch", orderserver.ManuallyFetch)

	duration, e := time.ParseDuration(config.Config.FetchInterval)
	if e != nil {
		log.Fatal(e)
	}
	go orderserver.FetchInterval(duration)
	_ = r.Run(config.Config.Server)
}
