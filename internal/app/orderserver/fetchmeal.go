package orderserver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"have-you-ordered/internal/app/orderserver/elasticsearch"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func StartFetchingMealRecord(since string, until string) int {

	var dates []string
	layout := "20060102"
	untilDate, _ := time.Parse(layout, until)
	for moment, _ := time.Parse(layout, since); moment.Before(untilDate); moment = moment.Add(time.Hour * 24) {
		formattedDate := moment.Format(layout)
		dates = append(dates, formattedDate)
	}
	totalSize := 0
	for _, date := range dates {
		orders := Fetch(date)
		//for k, v := range orders {
		//	fmt.Println(k, *v)
		//	IndexOrdered(v)
		//}
		log.Printf("Fetching: %s", date)
		totalSize += IndexBulk(orders)
	}
	//log.Printf("Total Bulk: %d\n", totalSize)
	return totalSize
}

func Fetch(date string) []*Ordered {
	orderList := PostOrdered(date)
	return (*orderList).Data
}

type OrderedEsReq struct {
	Id             string `json:"id"`
	PersonName     string `json:"personName"`
	OrderType      string `json:"orderType"`
	Time           int64  `json:"time"`
	IsAM           bool   `json:"isAm"`
	YYMMDD         string `json:"yymmdd"`
	SuggestContent string `json:"suggestContent"`
	V              int    `json:"v"`
}

func (o *OrderedEsReq) ReadOrdered(ordered *Ordered) {
	parsedTime, err := time.Parse(layout, ordered.Time)
	if err != nil {
		log.Println(err)
	}
	epochmillis := parsedTime.Unix()
	o.Id = ordered.Id
	o.PersonName = ordered.Name
	o.OrderType = ordered.OrderType
	o.Time = epochmillis
	o.IsAM = ordered.IsAM == 1
	o.YYMMDD = ordered.YYMMDD
	o.SuggestContent = ordered.SuggestContent
	o.V = ordered.V
}

// Thu Aug 08 2019 14:26:30 GMT+0800 (CST)
const layout = "Mon Jan 02 2006 15:04:05 GMT-0700 (MST)"

func IndexBulk(ordereds []*Ordered) int {
	var buf bytes.Buffer
	size := len(ordereds)
	if size < 1 {
		return size
	}
	for _, ordered := range ordereds {
		orderedEsReq := &OrderedEsReq{}
		orderedEsReq.ReadOrdered(ordered)
		dataBytes, _ := json.Marshal(orderedEsReq)
		dataBytes = append(dataBytes, "\n"...)
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%s" } }%s`, Md5FromByte(dataBytes), "\n"))
		buf.Grow(len(meta) + len(dataBytes))
		buf.Write(meta)
		buf.Write(dataBytes)
	}
	response, e := elasticsearch.Client().Bulk(bytes.NewReader(buf.Bytes()), elasticsearch.Client().Bulk.WithIndex("having-meal"))
	if e != nil {
		log.Println(e)
	}
	log.Printf("Bulk: %d ", size)
	b, _ := ioutil.ReadAll(response.Body)
	log.Println(string(b))
	return size
}

func IndexOrdered(ordered *Ordered) {

	orderedEsReq := &OrderedEsReq{}
	orderedEsReq.ReadOrdered(ordered)

	bytes, _ := json.Marshal(orderedEsReq)

	request := esapi.IndexRequest{
		Index:      elasticsearch.IndexName,
		DocumentID: ordered.Id,
		Body:       strings.NewReader(string(bytes)),
		Refresh:    "true",
	}
	res, err := request.Do(context.Background(), elasticsearch.Client())
	defer res.Body.Close()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		b, _ := ioutil.ReadAll(res.Body)
		log.Printf("[%s] Error indexing document ID=%s %s", res.Status(), ordered.Id, string(b))
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
}

func FetchInterval(duration time.Duration) {
	log.Printf("Start Fetch Interval: %d sec", duration/time.Second)
	ticker := time.NewTicker(duration)
	f := func() {
		defer func() {
			if e := recover(); e != nil {
				log.Printf("Fetch interval error: %s", e)
			}
		}()
		now := time.Now()
		log.Printf("%d", IndexBulk(Fetch(now.Format("20060102"))))
	}
	f()
	for {
		select {
		case <-ticker.C:
			f()
		}
	}
}
