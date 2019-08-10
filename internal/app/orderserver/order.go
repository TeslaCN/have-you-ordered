package orderserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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

type OrderingForm struct {
	Name           string `json:"name" form:"name"`
	OrderType      string `json:"orderType" form:"orderType"`
	Time           uint64 `json:"time" form:"time"`
	IsAM           int    `json:"isAM" form:"isAM"`
	YYMMdd         string
	SuggestContent string `json:"suggestContent" form:"suggestContent"`
}

func PostOrdered(date string) *OrderList {
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
	return &orderList
}

func (o *OrderingForm) PostOrdering() string {
	bytes, _ := json.Marshal(o)
	postBody := string(bytes)
	fmt.Println(postBody)
	response, err := http.Post("http://www.rainholer.com:81/orderDish/detailInfo/save", "application/json", strings.NewReader(postBody))
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	return string(body)
}
