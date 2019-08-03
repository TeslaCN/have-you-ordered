package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var server = flag.String("server", ":23333", "[IP]:port")

type OrderedForm struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	IsAM      int    `json:"isAM"`
}

type Ordered struct {
	Id             string `json:"_id"`
	Name           string `json:"name"`
	OrderType      string `json:"orderType"`
	Time           string `json:"time"`
	IsAM           int    `json:"isAM"`
	YYMMDD         string
	SuggestContent string `json:"suggestContent"`
	V              int    `json:"__v"`
}

type OrderList struct {
	ErrorCode int        `json:"error_code"`
	Msg       string     `json:"msg"`
	Data      []*Ordered `json:"data"`
}

//{
//	"name": "吴伟杰",
//	"orderType": "套餐饭",
//	"time": 1564552800000,
//	"isAM": 1,
//	"YYMMdd": "20190731",
//	"suggestContent": ""
//}
type OrderingForm struct {
	Name           string `json:"name" form:"name"`
	OrderType      string `json:"orderType" form:"orderType"`
	Time           uint64 `json:"time" form:"time"`
	IsAM           int    `json:"isAM" form:"isAM"`
	YYMMdd         string
	SuggestContent string `json:"suggestContent" form:"suggestContent"`
}

func main() {
	flag.Parse()
	fmt.Println(os.Args)
	r := gin.Default()
	r.LoadHTMLGlob("static/**")
	r.GET("/api", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"hello": "go",
		})
	})
	r.GET("/api/ordered/:date", func(context *gin.Context) {
		date := context.Param("date")
		if len(date) == 0 {
			date = time.Now().Format("20060102")
		}
		orderList := PostOrdered(date)
		context.JSON(http.StatusOK, orderList)
	})
	r.GET("/ordered", OrderedView)
	r.GET("/", OrderedView)
	r.GET("/ordering", func(context *gin.Context) {
		context.HTML(http.StatusOK, "ordering.html", gin.H{})
	})
	r.POST("/order", func(context *gin.Context) {
		var form OrderingForm
		if err := context.ShouldBind(&form); err != nil {

		}
		result := PostOrdering(&form)
		context.HTML(http.StatusOK, "ordering.html", gin.H{
			"message": result,
		})
	})
	_ = r.Run(*server)
}

func OrderedView(context *gin.Context) {
	date := context.Query("date")
	if len(date) == 0 {
		date = time.Now().Format("20060102")
	}
	context.HTML(http.StatusOK, "ordered.html", PostOrdered(date))
}

func PostOrdered(date string) OrderList {
	form := &OrderedForm{
		StartTime: date,
		EndTime:   date,
		IsAM:      1,
	}
	bytes, _ := json.Marshal(form)
	response, err := http.Post("http://www.rainholer.com:81/orderDish/getAllOrderInfo", "application/json", strings.NewReader(string(bytes)))
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	orderList := OrderList{}
	err = json.Unmarshal(body, &orderList)
	if err != nil {
		fmt.Println(err)
	}
	return orderList
}

func PostOrdering(form *OrderingForm) string {
	bytes, _ := json.Marshal(form)
	postBody := string(bytes)
	fmt.Println(postBody)
	response, err := http.Post("http://www.rainholer.com:81/orderDish/detailInfo/save", "application/json", strings.NewReader(postBody))
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	return string(body)
}
