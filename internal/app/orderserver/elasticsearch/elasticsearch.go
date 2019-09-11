package elasticsearch

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"have-you-ordered/cmd/orderserver/app/config"
	"log"
)

var c *elasticsearch.Client

var IndexName = config.Config.Index

type EsBaseResponseBody struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
}

func connect() (*elasticsearch.Client, error) {
	var client *elasticsearch.Client
	cfg := elasticsearch.Config{
		Addresses: config.Config.Elasticsearch.Hosts,
		Username:  config.Config.Elasticsearch.Username,
		Password:  config.Config.Elasticsearch.Password,
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	var r map[string]interface{}
	res, err := client.Info()
	if err != nil {
		log.Printf("Error getting response: %s", err)
	}
	// Check response status
	if res.IsError() {
		//log.Fatalf("Error: %s", res.String())
	}
	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		//log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	return client, nil
}

func Client() *elasticsearch.Client {
	if c == nil {
		newClient, err := connect()
		if err != nil {
			log.Printf("")
			return c
		}
		c = newClient
	}
	return c
}
