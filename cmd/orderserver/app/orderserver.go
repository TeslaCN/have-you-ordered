package app

import (
	"flag"
	"github.com/gin-gonic/gin"
	"have-you-ordered/internal/app/orderserver"
)

var server = flag.String("server", ":23333", "[IP]:port")

func init() {
	flag.Parse()
}

func Run() {
	r := gin.Default()

	//deprecatedWeb(r)

	r.GET("/", orderserver.OrderedView)
	r.GET("/api", orderserver.ApiHelloGo)
	r.GET("/api/ordered/:date", orderserver.ApiOrdered)
	r.POST("/api/order", orderserver.PostOrder)
	r.GET("/api/dashboard/agg-by-day", orderserver.AggHistogram)
	_ = r.Run(*server)
}

// Deprecated: Using Rest
func deprecatedWeb(r *gin.Engine) {
	r.LoadHTMLGlob("web/static/*")
	r.GET("/ordered", orderserver.OrderedView)
	r.GET("/ordering", orderserver.GetOrderingPage)
	r.POST("/order", orderserver.PostOrderPage)
}
