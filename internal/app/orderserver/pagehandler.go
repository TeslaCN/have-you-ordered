package orderserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetOrderingPage(context *gin.Context) {
	context.HTML(http.StatusOK, "ordering.html", gin.H{})
}

func PostOrderPage(context *gin.Context) {
	var form OrderingForm
	if err := context.ShouldBind(&form); err != nil {

	}
	result := form.PostOrdering()
	context.HTML(http.StatusOK, "ordering.html", gin.H{
		"message": result,
	})
}

func OrderedView(context *gin.Context) {
	date := context.Query("date")
	if len(date) == 0 {
		date = time.Now().Format("20060102")
	}
	context.HTML(http.StatusOK, "ordered.html", *PostOrdered(date))
}
