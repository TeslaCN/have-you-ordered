package test

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"testing"
)

func TestElasticsearch(t *testing.T) {
	config := elasticsearch.Config{Addresses: []string{"http://es.0:49204"}}
	es, _ := elasticsearch.NewClient(config)
	response, _ := es.Info()
	fmt.Println(response.String())
	es.Search()
}
