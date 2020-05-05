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

// {"error_code":0,"msg":"ok 😄","data":[{"_id":"5c3c2067795d1e7c0a8dc2e7","value":"套餐饭","openStatus":true,"__v":0},{"_id":"5c665072795d1e7c0a8dc81a","value":"中餐","openStatus":false,"__v":0},{"_id":"5c66507a795d1e7c0a8dc81b","value":"晚餐","openStatus":false,"__v":0},{"_id":"5c665083795d1e7c0a8dc81c","value":"中餐和晚餐","openStatus":false,"__v":0},{"_id":"5e9662ec795d1e7c0a8ded0b","value":"牛肉汤面","openStatus":false,"__v":0},{"_id":"5e966309795d1e7c0a8ded0c","value":"干面套餐+猪杂汤","openStatus":true,"__v":0},{"_id":"5e9be1ee795d1e7c0a8ded5f","value":"干炒粿条","openStatus":true,"__v":0},{"_id":"5e9e6526795d1e7c0a8ded94","value":"牛肉炒饭","openStatus":false,"__v":0},{"_id":"5ea12dd8795d1e7c0a8deddb","value":"牛肉炒面","openStatus":false,"__v":0}]}
type OrderTypes struct {
	ErrorCode int          `json:"error_code"`
	Msg       string       `json:"msg"`
	Data      []*OrderType `json:"data"`
}

type OrderType struct {
	Id         string `json:"_id"`
	Value      string `json:"value"`
	OpenStatus bool   `json:"openStatus"`
	V          int    `json:"__v"`
}

type OrderingForm struct {
	Name           string `json:"name" form:"name"`
	OrderType      string `json:"orderType" form:"orderType"`
	Time           uint64 `json:"time" form:"time"`
	IsAM           int    `json:"isAM" form:"isAM"`
	YYMMdd         string
	SuggestContent string `json:"suggestContent" form:"suggestContent"`
}

func PostAllOrderTypes() *OrderTypes {
	response, err := http.Post("http://www.rainholer.com:81/orderDish/getAllOrderType", "application/json", strings.NewReader("{}"))
	if err != nil {
		fmt.Println(err)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	orderTypes := OrderTypes{}
	err = json.Unmarshal(bytes, &orderTypes)
	if err != nil {
		fmt.Println(err)
	}
	return &orderTypes
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
