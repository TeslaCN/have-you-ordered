package config

import (
	"encoding/json"
	"log"
	"os"
)

type configuration struct {
	Elasticsearch struct {
		Hosts    []string `json:"hosts"`
		Username string   `json:"username"`
		Password string   `json:"password"`
	} `json:"elasticsearch"`
	Server string `json:"server"`
	Index  string `json:"index"`
}

var Config = &configuration{}

func init() {
	file, e := os.Open("configs/config.json")
	if e != nil {
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if e := decoder.Decode(Config); e != nil {
		log.Println(e)
	}
	log.Println(*Config)
}
