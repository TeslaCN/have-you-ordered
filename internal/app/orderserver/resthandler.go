package orderserver

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"have-you-ordered/cmd/orderserver/app/dto"
	"have-you-ordered/internal/app/orderserver/elasticsearch"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// GET /api
func ApiHelloGo(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"hello": "go",
	})
}

func ApiOrdered(context *gin.Context) {
	date := context.Param("date")
	if len(date) == 0 {
		date = time.Now().Format("20060102")
	}
	orderList := PostOrdered(date)
	context.JSON(http.StatusOK, *orderList)
}

func ApiOrderTypes(context *gin.Context) {
	context.JSON(http.StatusOK, *PostAllOrderTypes())
}

func AggHistogram(c *gin.Context) {
	//language=JSON
	query := `{
		  "size": 0,
		  "aggs": {
		    "NAME": {
		      "date_histogram": {
		        "field": "time",
		        "interval": "1D",
		        "format": "yyyy-MM-dd"
		      }
		    }
		  }
		}
		`
	es := elasticsearch.Client()
	response, e := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithBody(strings.NewReader(query)),
		es.Search.WithIndex("having-meal"),
	)
	if e != nil || response.IsError() {

	}
	s, _ := ioutil.ReadAll(response.Body)
	agg := &struct {
		elasticsearch.EsBaseResponseBody
		Aggregations struct {
			NAME struct {
				Buckets []struct {
					KeyAsString string `json:"key_as_string"`
					Key         uint64 `json:"key"`
					DocCount    uint64 `json:"doc_count"`
				} `json:"buckets"`
			}
		} `json:"aggregations"`
	}{}
	_ = json.Unmarshal(s, agg)
	var results []*dto.KeyValuePairDto
	buckets := agg.Aggregations.NAME.Buckets
	for _, bucket := range buckets {
		results = append(results, &dto.KeyValuePairDto{
			Key:   bucket.KeyAsString,
			Value: bucket.DocCount,
		})
	}
	c.JSON(http.StatusOK, results)
}

func AggOrderType(c *gin.Context) {
	//language=JSON
	query := `{
  "size": 0,
  "aggs": {
    "NAME": {
      "terms": {
        "field": "orderType",
        "size": 1000
      }
    }
  }
}`
	es := elasticsearch.Client()
	response, e := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithBody(strings.NewReader(query)),
		es.Search.WithIndex("having-meal"),
	)
	if e != nil || response.IsError() {
	}
	s, _ := ioutil.ReadAll(response.Body)
	agg := &struct {
		elasticsearch.EsBaseResponseBody
		Aggregations struct {
			NAME struct {
				Buckets []struct {
					Key      string `json:"key"`
					DocCount uint64 `json:"doc_count"`
				} `json:"buckets"`
			}
		} `json:"aggregations"`
	}{}
	_ = json.Unmarshal(s, agg)
	var results []*dto.KeyValuePairDto
	buckets := agg.Aggregations.NAME.Buckets
	for _, bucket := range buckets {
		results = append(results, &dto.KeyValuePairDto{
			Key:   bucket.Key,
			Value: bucket.DocCount,
		})
	}
	c.JSON(http.StatusOK, results)
}

func AggOrderPerson(c *gin.Context) {
	//language=JSON
	query := `{
  "size": 0,
  "aggs": {
    "NAME": {
      "terms": {
        "field": "personName",
        "size": 1000
      }
    }
  }
}`
	es := elasticsearch.Client()
	response, e := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithBody(strings.NewReader(query)),
		es.Search.WithIndex("having-meal"),
	)
	if e != nil || response.IsError() {
	}
	s, _ := ioutil.ReadAll(response.Body)
	agg := &struct {
		elasticsearch.EsBaseResponseBody
		Aggregations struct {
			NAME struct {
				Buckets []struct {
					Key      string `json:"key"`
					DocCount uint64 `json:"doc_count"`
				} `json:"buckets"`
			}
		} `json:"aggregations"`
	}{}
	_ = json.Unmarshal(s, agg)
	var results []*dto.KeyValuePairDto
	buckets := agg.Aggregations.NAME.Buckets
	for _, bucket := range buckets {
		results = append(results, &dto.KeyValuePairDto{
			Key:   coverPersonName(bucket.Key),
			Value: bucket.DocCount,
		})
	}
	c.JSON(http.StatusOK, results)
}

func coverPersonName(name string) string {
	chars := []rune(name)
	if 0 < len(chars) {
		return string(chars[0]) + strings.Repeat("*", len(chars)-1)
	}
	return name
}

func PostOrder(context *gin.Context) {
	var form OrderingForm
	if err := context.ShouldBindJSON(&form); err != nil {
	}
	result := form.PostOrdering()
	context.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
