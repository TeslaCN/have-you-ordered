package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

type configuration struct {
	Elasticsearch struct {
		Hosts    []string `json:"hosts"`
		Username string   `json:"username"`
		Password string   `json:"password"`
	} `json:"elasticsearch"`
	Server        string `json:"server"`
	Index         string `json:"index"`
	FetchInterval string `json:"fetch_interval"`
}

var (
	Config     = &configuration{}
	configPath = flag.String("config", "configs/config.json", "[-config path/to/config.json]")
)

func init() {
	flag.Parse()
	file, e := os.Open(*configPath)
	if e != nil {
		log.Fatalln("Config file not found: " + *configPath)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if e := decoder.Decode(Config); e != nil {
		log.Println(e)
	}
	log.Println(*Config)
}
