package orderserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ManuallyFetch(context *gin.Context) {
	since := context.Query("since")
	until := context.Query("until")
	totalBulk := StartFetchingMealRecord(since, until)
	context.JSON(http.StatusOK, gin.H{
		"total_bulk": totalBulk,
	})
}
